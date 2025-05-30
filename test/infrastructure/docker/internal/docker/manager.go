/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package docker

import (
	"context"
	"fmt"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/kind/pkg/apis/config/v1alpha4"
	"sigs.k8s.io/kind/pkg/cluster/constants"

	"sigs.k8s.io/cluster-api/test/infrastructure/container"
	"sigs.k8s.io/cluster-api/test/infrastructure/docker/internal/docker/types"
	"sigs.k8s.io/cluster-api/test/infrastructure/kind"
)

// KubeadmContainerPort is the port that kubeadm listens on in the container.
const KubeadmContainerPort = 6443

// ControlPlanePort is the port for accessing the control plane API in the container.
const ControlPlanePort = 6443

// HAProxyPort is the port for accessing HA proxy stats.
const HAProxyPort = 8404

// DefaultNetwork is the default network name to use in kind.
const DefaultNetwork = "kind"

// haproxyEntrypoint is the entrypoint used to start the haproxy load balancer container.
var haproxyEntrypoint = []string{"haproxy", "-W", "-db", "-f", "/usr/local/etc/haproxy/haproxy.cfg"}

// Manager is the kind manager type.
type Manager struct{}

type nodeCreateOpts struct {
	Name         string
	ClusterName  string
	Role         string
	EntryPoint   []string
	Mounts       []v1alpha4.Mount
	PortMappings []v1alpha4.PortMapping
	Labels       map[string]string
	IPFamily     container.ClusterIPFamily
	KindMapping  kind.Mapping
}

// CreateControlPlaneNode will create a new control plane container.
// NOTE: If port is 0 picking a host port for the control plane is delegated to the container runtime and is not stable across container restarts.
// This means that connection to a control plane node may take some time to recover if the underlying container is restarted.
func (m *Manager) CreateControlPlaneNode(ctx context.Context, name, clusterName, listenAddress string, port int32, mounts []v1alpha4.Mount, portMappings []v1alpha4.PortMapping, labels map[string]string, ipFamily container.ClusterIPFamily, kindMapping kind.Mapping) (*types.Node, error) {
	// add api server port mapping
	portMappingsWithAPIServer := append(portMappings, v1alpha4.PortMapping{
		ListenAddress: listenAddress,
		HostPort:      port,
		ContainerPort: KubeadmContainerPort,
		Protocol:      v1alpha4.PortMappingProtocolTCP,
	})
	createOpts := &nodeCreateOpts{
		Name:         name,
		ClusterName:  clusterName,
		Role:         constants.ControlPlaneNodeRoleValue,
		PortMappings: portMappingsWithAPIServer,
		Mounts:       mounts,
		Labels:       labels,
		IPFamily:     ipFamily,
		KindMapping:  kindMapping,
	}
	node, err := createNode(ctx, createOpts)
	if err != nil {
		return nil, err
	}

	return node, nil
}

// CreateWorkerNode will create a new worker container.
func (m *Manager) CreateWorkerNode(ctx context.Context, name, clusterName string, mounts []v1alpha4.Mount, portMappings []v1alpha4.PortMapping, labels map[string]string, ipFamily container.ClusterIPFamily, kindMapping kind.Mapping) (*types.Node, error) {
	createOpts := &nodeCreateOpts{
		Name:         name,
		ClusterName:  clusterName,
		Role:         constants.WorkerNodeRoleValue,
		PortMappings: portMappings,
		Mounts:       mounts,
		Labels:       labels,
		IPFamily:     ipFamily,
		KindMapping:  kindMapping,
	}
	return createNode(ctx, createOpts)
}

// CreateExternalLoadBalancerNode will create a new container to act as the load balancer for external access.
// NOTE: If port is 0 picking a host port for the load balancer is delegated to the container runtime and is not stable across container restarts.
// This can break the Kubeconfig in kind, i.e. the file resulting from `kind get kubeconfig -n $CLUSTER_NAME' if the load balancer container is restarted.
func (m *Manager) CreateExternalLoadBalancerNode(ctx context.Context, name, image, clusterName, listenAddress string, port int32, _ container.ClusterIPFamily) (*types.Node, error) {
	// load balancer port mapping
	portMappings := []v1alpha4.PortMapping{
		{
			ListenAddress: listenAddress,
			HostPort:      port,
			ContainerPort: ControlPlanePort,
			Protocol:      v1alpha4.PortMappingProtocolTCP,
		},
		{
			ListenAddress: listenAddress,
			HostPort:      0,
			ContainerPort: HAProxyPort,
			Protocol:      v1alpha4.PortMappingProtocolTCP,
		},
	}
	createOpts := &nodeCreateOpts{
		Name:         name,
		ClusterName:  clusterName,
		Role:         constants.ExternalLoadBalancerNodeRoleValue,
		PortMappings: portMappings,
		EntryPoint:   haproxyEntrypoint,
		// Load balancer doesn't have an equivalent in kind, but we use a kind.Mapping to
		// forward the image name to create node.
		KindMapping: kind.Mapping{
			Image: image,
			Mode:  kind.ModeNone,
		},
	}
	node, err := createNode(ctx, createOpts)
	if err != nil {
		return nil, err
	}

	return node, nil
}

func createNode(ctx context.Context, opts *nodeCreateOpts) (*types.Node, error) {
	log := ctrl.LoggerFrom(ctx)

	// Collect the labels to apply to the container
	containerLabels := map[string]string{
		clusterLabelKey:  opts.ClusterName,
		nodeRoleLabelKey: opts.Role,
	}
	for name, value := range opts.Labels {
		containerLabels[name] = value
	}

	runOptions := &container.RunContainerInput{
		Name:   opts.Name, // make hostname match container name
		Image:  opts.KindMapping.Image,
		Labels: containerLabels,
		// runtime persistent storage
		// this ensures that E.G. pods, logs etc. are not on the container
		// filesystem, which is not only better for performance, but allows
		// running kind in kind for "party tricks"
		// (please don't depend on doing this though!)
		Entrypoint:   opts.EntryPoint,
		Volumes:      map[string]string{"/var": ""},
		Mounts:       generateMountInfo(opts.Mounts),
		PortMappings: generatePortMappings(opts.PortMappings),
		Network:      DefaultNetwork,
		Tmpfs: map[string]string{
			"/tmp": "", // various things depend on working /tmp
			"/run": "", // systemd wants a writable /run
		},
		IPFamily: opts.IPFamily,
		KindMode: opts.KindMapping.Mode,
	}
	if opts.Role == constants.ControlPlaneNodeRoleValue {
		runOptions.EnvironmentVars = map[string]string{
			"KUBECONFIG": "/etc/kubernetes/admin.conf",
		}
	}

	log.V(6).Info(fmt.Sprintf("Container run options: %+v", runOptions))

	containerRuntime, err := container.RuntimeFrom(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to container runtime: %v", err)
	}

	err = containerRuntime.RunContainer(ctx, runOptions, nil)
	if err != nil {
		return nil, err
	}

	return types.NewNode(opts.Name, opts.KindMapping.Image, opts.Role), nil
}

func generateMountInfo(mounts []v1alpha4.Mount) []container.Mount {
	mountInfo := []container.Mount{}
	for _, mount := range mounts {
		mountInfo = append(mountInfo, container.Mount{
			Source:   mount.HostPath,
			Target:   mount.ContainerPath,
			ReadOnly: mount.Readonly,
		})
	}
	// some k8s things want to read /lib/modules
	mountInfo = append(mountInfo, container.Mount{
		Source:   "/lib/modules",
		Target:   "/lib/modules",
		ReadOnly: true,
	})
	return mountInfo
}

func generatePortMappings(portMappings []v1alpha4.PortMapping) []container.PortMapping {
	result := make([]container.PortMapping, 0, len(portMappings))
	for _, pm := range portMappings {
		portMapping := container.PortMapping{
			ContainerPort: pm.ContainerPort,
			HostPort:      pm.HostPort,
			ListenAddress: pm.ListenAddress,
			Protocol:      capiProtocolToCommonProtocol(pm.Protocol),
		}
		result = append(result, portMapping)
	}
	return result
}

func capiProtocolToCommonProtocol(protocol v1alpha4.PortMappingProtocol) string {
	switch protocol {
	case v1alpha4.PortMappingProtocolUDP:
		return "udp"
	case v1alpha4.PortMappingProtocolSCTP:
		return "sctp"
	default:
		return "tcp"
	}
}
