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
	"fmt"
	"log"
	"sort"
	"strings"

	semver "github.com/Masterminds/semver/v3"
	"github.com/bufbuild/connect-go"

	apiv1 "github.com/koor-tech/version-service/api/v1"
)

type VersionServer struct{}

func (s *VersionServer) Operator(
	ctx context.Context,
	req *connect.Request[apiv1.OperatorRequest],
) (*connect.Response[apiv1.OperatorResponse], error) {
	log.Println("Request headers: ", req.Header())
	versions, err := getLatestVersions(req.Msg)
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&apiv1.OperatorResponse{
		Versions: versions,
	})
	res.Header().Set("Version-Service-Version", "v1")
	return res, nil
}

func getLatestVersions(opreq *apiv1.OperatorRequest) (*apiv1.DetailedProductVersions, error) {
	kov := strings.TrimPrefix(opreq.Versions.KoorOperator, "v")
	vm, err := getVersionMatrix(kov)
	if err != nil {
		return nil, connect.NewError(connect.CodeNotFound, err)
	}
	latestKoorOperator, err := findLatestVersion(vm.KoorOperator, opreq.Versions.KoorOperator)
	if err != nil {
		return nil, err
	}

	latestKsd, err := findLatestVersion(vm.Ksd, opreq.Versions.Ksd)
	if err != nil {
		return nil, err
	}

	latestCeph, err := findLatestVersion(vm.Ceph, opreq.Versions.Ceph)
	if err != nil {
		return nil, err
	}

	latest := &apiv1.DetailedProductVersions{
		KoorOperator: latestKoorOperator,
		Ksd:          latestKsd,
		Ceph:         latestCeph,
	}

	return latest, nil
}

func findLatestVersion(versions map[string]*apiv1.DetailedVersion, current string) (*apiv1.DetailedVersion, error) {
	if len(versions) == 0 {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("could not find latest versions, current: %s", current))
	}

	currentSemver, err := semver.NewVersion(current)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("invalid version: %s", current))
	}

	semvers := make([]*semver.Version, 0, len(versions))
	for k := range versions {
		sv, err := semver.NewVersion(k)
		if err != nil {
			return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("could not parse version: %s", k))
		}
		semvers = append(semvers, sv)
	}

	sort.Sort(sort.Reverse(semver.Collection(semvers)))

	if currentSemver.GreaterThan(semvers[0]) {
		return nil, connect.NewError(connect.CodeNotFound,
			fmt.Errorf("current version (%s) is bigger than latest known version (%s)", current, semvers[0]))
	}

	latest := semvers[0].String()
	return versions[latest], nil
}
