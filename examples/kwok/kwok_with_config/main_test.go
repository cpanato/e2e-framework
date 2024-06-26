/*
Copyright 2023 The Kubernetes Authors.

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

package kwok

import (
	"os"
	"testing"

	"sigs.k8s.io/e2e-framework/pkg/env"
	"sigs.k8s.io/e2e-framework/pkg/envconf"
	"sigs.k8s.io/e2e-framework/pkg/envfuncs"
	"sigs.k8s.io/e2e-framework/support/kwok"
)

var testenv env.Environment

func TestMain(m *testing.M) {
	testenv = env.New()
	kwokClusterName := envconf.RandomName("kwok-cluster", 16)
	namespace := envconf.RandomName("kwok-ns", 16)

	testenv.Setup(
		envfuncs.CreateClusterWithConfig(kwok.NewProvider(), kwokClusterName, "kwok-config.yaml"),
		envfuncs.CreateNamespace(namespace),
	)

	testenv.Finish(
		envfuncs.ExportClusterLogs(kwokClusterName, "./logs"),
		envfuncs.DeleteNamespace(namespace),
		envfuncs.DestroyCluster(kwokClusterName),
	)
	os.Exit(testenv.Run(m))
}
