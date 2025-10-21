package player

import (
	"context"
	"sdxx/server/internal/component/snowflake"
	"sdxx/server/internal/dao/mysql/game/model"
	"sdxx/server/internal/utils/hashx"
	commonv1 "sdxx/server/protobuf/gen/common/v1"
	playerv1 "sdxx/server/protobuf/gen/mesh/game/player/v1"

	"github.com/dobyte/due/v2/errors"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/utils/xconv"
	"github.com/dobyte/due/v2/utils/xrand"
)

func (s *Server) takeOneNickname(ctx context.Context) string {
	nickname := s.table.TbNickname.TakeOne()
	if nickname == "" {
		nickname = xrand.Letters(8)
	}
	for {
		nicknameHash := hashx.Sum64String(nickname)
		_, err := s.playerModel.FindOneByNicknameHash(ctx, nicknameHash)
		var isOnlyOne bool
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				isOnlyOne = true
			} else {
				log.Errorf("find one by nickname hash err: %v", err)
			}
		}
		if isOnlyOne {
			break
		} else {
			nickname += xrand.Letters(6)
		}
	}

	return nickname
}

func (s *Server) EnsurePlayer(ctx context.Context, req *playerv1.EnsurePlayerReq) (*playerv1.EnsurePlayerResp, error) {
	player, err := s.playerModel.FindOneByAccountIdServerId(ctx, req.AccountId, req.ServerId)
	var shouldNew bool
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			shouldNew = true
		} else {
			log.Errorf("find one by account id server id err: %v", err)
			return nil, errors.NewError(commonv1.InternalError)
		}
	}
	if shouldNew {
		tbRole := s.table.TbRole.Load()
		tbLevel := s.table.TbLevel.Load()
		global := s.table.TbGlobal.Load()
		role := tbRole.Get(global.InitialRole)
		level := tbLevel.Get(role.InitializeLv)

		nickname := s.takeOneNickname(ctx)
		nicknameHash := hashx.Sum64String(nickname)
		player = &model.Player{
			AccountId:    req.AccountId,
			ServerId:     req.ServerId,
			PlayerId:     snowflake.Generate(),
			Level:        uint64(role.InitializeLv),
			Nickname:     nickname,
			NicknameHash: nicknameHash,
			Avatar:       xconv.String(role.InitializeAvatar),
			Exp:          uint64(level.Exp),
		}
		err = s.playerModel.Insert(ctx, nil, player)
		if err != nil {
			log.Errorf("insert player err: %v", err)
			return nil, errors.NewError(commonv1.InternalError)
		}
	}
	return &playerv1.EnsurePlayerResp{
		PlayerId: player.PlayerId,
	}, nil
}
