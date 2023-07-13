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
	"context"
	"log"

	"github.com/bufbuild/connect-go"

	apiv1 "github.com/koor-tech/version-service/api/v1"
)

type VersionServer struct{}

func (s *VersionServer) Operator(
	ctx context.Context,
	req *connect.Request[apiv1.OperatorRequest],
) (*connect.Response[apiv1.OperatorResponse], error) {
	log.Println("Request headers: ", req.Header())
	vm, err := getVersionMatrix(req.Msg.Versions.KoorOperator)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	res := connect.NewResponse(&apiv1.OperatorResponse{
		Versions: &apiv1.DetailedVersions{
			KoorOperator: vm.KoorOperator["foo"],
			Ksd:          vm.Ksd["foo"],
			Ceph:         vm.Ceph["foo"],
		},
	})
	res.Header().Set("Version-Service-Version", "v1")
	return res, nil
}
