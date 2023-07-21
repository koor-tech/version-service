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

package main

import (
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/koor-tech/version-service/api/v1/apiv1connect"
	serverv1 "github.com/koor-tech/version-service/server/v1"
)

func main() {
	vs := &serverv1.VersionServer{}
	mux := http.NewServeMux()
	path, handler := apiv1connect.NewVersionServiceHandler(vs)
	mux.Handle(path, handler)
	http.ListenAndServe(
		":8082",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
