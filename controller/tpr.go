/*
Copyright 2016 The Fission Authors.

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

package controller

import (
	"github.com/fission/fission/tpr"
)

func makeTPRBackedAPI() (*API, error) {
	config, clientset, err := tpr.GetKubernetesClient()
	if err != nil {
		return nil, err
	}

	tprClient, err := tpr.GetTprClient(config)
	if err != nil {
		return nil, err
	}

	api := &API{
		function:         tpr.MakeFunctionInterface(tprClient),
		environment:      tpr.MakeEnvironmentInterface(tprClient),
		httpTrigger:      tpr.MakeHttptriggerInterface(tprClient),
		kubeWatchTrigger: tpr.MakeKuberneteswatchtriggerInterface(tprClient),
	}
	return api, nil
}