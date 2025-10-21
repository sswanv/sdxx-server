package inventory

import (
	"context"
	inventoryv1 "sdxx/server/protobuf/gen/mesh/game/inventory/v1"
)

func (s *Server) GetPlayerEquipments(context.Context, *inventoryv1.GetPlayerEquipmentsReq) (*inventoryv1.GetPlayerEquipmentsResp, error) {
	return &inventoryv1.GetPlayerEquipmentsResp{}, nil
}
