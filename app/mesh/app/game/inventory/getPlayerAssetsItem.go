package inventory

import (
	"context"
	commonv1 "sdxx/server/protobuf/gen/common/v1"
	inventoryv1 "sdxx/server/protobuf/gen/mesh/game/inventory/v1"
)

func (s *Server) GetPlayerAssets(ctx context.Context, req *inventoryv1.GetPlayerAssetsReq) (*inventoryv1.GetPlayerAssetsResp, error) {
	global := s.table.TbGlobal.Load()
	mapOfItemStack := make(map[uint64]*commonv1.ItemStack)
	var itemIds []uint64
	for _, itemId := range global.PlayerAssets {
		itemIds = append(itemIds, uint64(itemId))
		mapOfItemStack[uint64(itemId)] = &commonv1.ItemStack{
			ItemId: uint64(itemId),
			Count:  0,
		}
	}

	playerAssets, err := s.playerAssetsModel.FindByPlayerIdAndItemId(ctx, req.PlayerId, itemIds)
	if err != nil {
		return nil, err
	}
	for _, asset := range playerAssets {
		mapOfItemStack[asset.ItemId].Count += asset.Count
	}

	var assets []*commonv1.ItemStack
	for _, itemStack := range mapOfItemStack {
		assets = append(assets, itemStack)
	}

	return &inventoryv1.GetPlayerAssetsResp{
		Assets: assets,
	}, nil
}
