package player

import (
	"context"
	inventoryv1 "sdxx/server/protobuf/gen/mesh/game/inventory/v1"
	playerv1 "sdxx/server/protobuf/gen/mesh/game/player/v1"
)

func (s *Server) GetProfile(ctx context.Context, req *playerv1.GetProfileReq) (*playerv1.GetProfileResp, error) {
	player, err := s.playerModel.FindOneByPlayerId(ctx, req.PlayerId)
	if err != nil {
		return nil, err
	}

	getPlayerAssetsResp, err := s.inventoryServiceClient.GetPlayerAssets(ctx, &inventoryv1.GetPlayerAssetsReq{
		PlayerId: req.PlayerId,
	})
	if err != nil {
		return nil, err
	}
	return &playerv1.GetProfileResp{
		Nickname: player.Nickname,
		Avatar:   player.Avatar,
		Level:    player.Level,
		Assets:   getPlayerAssetsResp.Assets,
	}, nil
}
