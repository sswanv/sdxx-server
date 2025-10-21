package inventory

import (
	"context"
	inventoryv1 "sdxx/server/protobuf/gen/mesh/game/inventory/v1"
)

func (s *Server) GetPlayerEquipped(context.Context, *inventoryv1.GetPlayerEquippedReq) (*inventoryv1.GetPlayerEquippedResp, error) {
	return &inventoryv1.GetPlayerEquippedResp{}, nil
}
