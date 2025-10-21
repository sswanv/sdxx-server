package api

import (
	"sdxx/server/internal/component/http"
	"sdxx/server/internal/service"

	platformv1 "sdxx/server/protobuf/gen/mesh/platform/v1"
)

func NewApi(proxy *http.Proxy) *Api {
	return &Api{
		proxy:                 proxy,
		platformServiceClient: service.NewPlatformServiceClient(proxy.NewMeshClient),
	}
}

type Api struct {
	proxy                 *http.Proxy
	platformServiceClient platformv1.PlatformServiceClient
}
