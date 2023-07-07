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
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
