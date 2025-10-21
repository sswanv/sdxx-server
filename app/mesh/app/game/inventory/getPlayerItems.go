package inventory

import (
	"context"
	commonv1 "sdxx/server/protobuf/gen/common/v1"
	inventoryv1 "sdxx/server/protobuf/gen/mesh/game/inventory/v1"

	"github.com/dobyte/due/v2/errors"
	"github.com/dobyte/due/v2/log"
)

func (s *Server) GetPlayerItems(ctx context.Context, req *inventoryv1.GetPlayerItemsReq) (*inventoryv1.GetPlayerItemsResp, error) {
	playerItems, err := s.playerItemModel.FindByPlayerId(ctx, req.PlayerId)
	if err != nil {
		log.Errorf("get player items: find by player id failed, err: %v", err)
		return nil, errors.NewError(commonv1.InternalError)
	}

	var items []*commonv1.Item
	for _, playerItem := range playerItems {
		items = append(items, &commonv1.Item{
			ItemId: playerItem.ItemId,
			Count:  playerItem.Count,
		})
	}

	return &inventoryv1.GetPlayerItemsResp{
		Items: items,
	}, nil
}
