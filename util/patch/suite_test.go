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

package patch

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"

	"sigs.k8s.io/cluster-api/internal/test/envtest"
	"sigs.k8s.io/cluster-api/util/test/builder"
)

const (
	timeout = time.Second * 10
)

var (
	env *envtest.Environment
	ctx = ctrl.SetupSignalHandler()
)

func TestMain(m *testing.M) {
	os.Exit(envtest.Run(ctx, envtest.RunInput{
		M:        m,
		SetupEnv: func(e *envtest.Environment) { env = e },
		AdditionalCRDDirectoryPaths: []string{
			filepath.Join("util", "test", "builder", "crd"),
		},
		AdditionalSchemeBuilder: runtime.NewSchemeBuilder(
			builder.AddTransitionV1Beta2ToScheme,
		),
	}))
}
