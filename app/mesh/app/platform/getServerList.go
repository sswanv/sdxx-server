package platform

import (
	"context"
	platformv1 "sdxx/server/protobuf/gen/mesh/platform/v1"
)

func (s *Server) GetServerList(ctx context.Context, req *platformv1.GetServerListReq) (*platformv1.GetServerListResp, error) {
	tbServerList := s.table.TbServerList.Load()
	var servers []*platformv1.GetServerListResp_Server
	for _, server := range tbServerList.GetDataList() {
		servers = append(servers, &platformv1.GetServerListResp_Server{
			Id:     server.Id,
			Name:   server.Name,
			Addr:   server.Addr,
			Status: platformv1.ServerStatus_SERVER_STATUS_OPEN,
		})
	}
	return &platformv1.GetServerListResp{Servers: servers}, nil
}
