/*
Copyright 2022 The Kubernetes Authors.

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

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	clusterv1 "sigs.k8s.io/cluster-api/api/core/v1beta2"
	runtimehooksv1 "sigs.k8s.io/cluster-api/api/runtime/hooks/v1alpha1"
	runtimecatalog "sigs.k8s.io/cluster-api/exp/runtime/catalog"
)

// FakeRequest is a response for testing
// +kubebuilder:object:root=true
type FakeRequest struct {
	metav1.TypeMeta `json:",inline"`

	// CommonRequest contains Settings field common to all request types.
	runtimehooksv1.CommonRequest `json:",inline"`

	Cluster clusterv1.Cluster

	Second string
	First  int
}

var _ runtimehooksv1.ResponseObject = &FakeResponse{}

// FakeResponse is a response for testing.
// +kubebuilder:object:root=true
type FakeResponse struct {
	metav1.TypeMeta `json:",inline"`

	runtimehooksv1.CommonResponse `json:",inline"`

	Second string
	First  int
}

func FakeHook(*FakeRequest, *FakeResponse) {}

func init() {
	catalogBuilder.RegisterHook(FakeHook, &runtimecatalog.HookMeta{
		Tags:        []string{"fake-tag"},
		Summary:     "FakeHook summary",
		Description: "FakeHook description",
		Deprecated:  true,
	})
}
