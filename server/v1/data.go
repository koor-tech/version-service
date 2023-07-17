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
	"log"
	"os"
	"path"

	"google.golang.org/protobuf/encoding/protojson"

	apiv1 "github.com/koor-tech/version-service/api/v1"
)

var data = map[string][]byte{}

func init() {
	files, err := os.ReadDir("./data")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		name := file.Name()
		content, err := os.ReadFile(path.Join("./data", name))
		if err != nil {
			log.Fatalf("failed to read data file: %v", err)
		}
		data[name] = content
	}
}

func getVersionMatrix(koorOperatorVersion string) (*apiv1.VersionMatrix, error) {
	filename := fmt.Sprintf("koor-operator-v%s.json", koorOperatorVersion)
	contents, ok := data[filename]
	if !ok {
		return nil, fmt.Errorf("data file not found: %s", filename)
	}

	vm := &apiv1.VersionMatrix{}
	err := protojson.Unmarshal(contents, vm)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal data file: %s , err: %v", filename, err)
	}
	return vm, nil
}
