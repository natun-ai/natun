/*
Copyright 2022 Natun.

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

package streaming

import (
	"github.com/natun-ai/natun/pkg/plugin"
	"github.com/natun-ai/natun/pkg/runner"
)

// These variables are being overwritten by the build process
var (
	Image      = "ghcr.io/natun-ai/streaming-runner:2f72613"
	runtimeVer = "ea1cc51"
)

func init() {
	baseRunner := runner.BaseRunner{
		Image:          Image,
		RuntimeVersion: runtimeVer,
		Command:        []string{"runner"},
	}
	reconciler, err := baseRunner.Reconciler()
	if err != nil {
		panic(err)
	}

	// Register the plugin
	plugin.DataConnectorReconciler.Register("streaming", reconciler)
}
