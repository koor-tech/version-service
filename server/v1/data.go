/*
Copyright 2023 Koor Technologies, Inc. All rights reserved.

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

package serverv1

import (
	"fmt"

	semver "github.com/Masterminds/semver/v3"
	"google.golang.org/protobuf/encoding/protojson"

	apiv1 "github.com/koor-tech/version-service/api/v1"
	"github.com/koor-tech/version-service/data"
)

func getVersionMatrix(koorOperatorVersion string) (*apiv1.VersionMatrix, error) {
	sv, err := semver.NewVersion(koorOperatorVersion)
	if err != nil {
		return nil, fmt.Errorf("invalid koor operator version: %s", koorOperatorVersion)
	}

	filename := fmt.Sprintf("koor-operator-v%d.%d.json", sv.Major(), sv.Minor())
	contents, err := data.Data.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("data file not found: %s", filename)
	}

	vm := &apiv1.VersionMatrix{}
	err = protojson.Unmarshal(contents, vm)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal data file: %s , err: %v", filename, err)
	}
	return vm, nil
}
