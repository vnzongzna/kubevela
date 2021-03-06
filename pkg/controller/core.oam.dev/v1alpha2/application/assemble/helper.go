/*
Copyright 2021 The KubeVela Authors.

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

package assemble

import (
	"encoding/json"

	"k8s.io/apimachinery/pkg/runtime"

	"github.com/oam-dev/kubevela/apis/core.oam.dev/v1alpha2"
)

// TODO duplicate with PR 1708, should remove latter
func convertRawExtention2AppConfig(raw runtime.RawExtension) (*v1alpha2.ApplicationConfiguration, error) {
	obj := &v1alpha2.ApplicationConfiguration{}
	b, err := raw.MarshalJSON()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

func convertRawExtention2Component(raw runtime.RawExtension) (*v1alpha2.Component, error) {
	obj := &v1alpha2.Component{}
	b, err := raw.MarshalJSON()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, obj); err != nil {
		return nil, err
	}
	return obj, nil
}
