package api

import (
	"sdxx/server/internal/component/http"
	platformv1 "sdxx/server/protobuf/gen/mesh/platform/v1"
	webv1 "sdxx/server/protobuf/gen/web/v1"
)

func (a *Api) GetServerList(ctx http.Context) error {
	getServerListResp, err := a.platformServiceClient.GetServerList(ctx.Context(), &platformv1.GetServerListReq{})
	if err != nil {
		return ctx.Failure(err)
	}

	var servers []*webv1.GetServerListResp_Server
	for _, server := range getServerListResp.Servers {
		servers = append(servers, &webv1.GetServerListResp_Server{
			Id:     server.Id,
			Name:   server.Name,
			Addr:   server.Addr,
			Status: webv1.ServerStatus(server.Status),
		})
	}
	return ctx.Success(&webv1.GetServerListResp{Servers: servers})
}
