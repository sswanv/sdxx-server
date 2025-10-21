package core

import (
	"sdxx/server/internal/core/nodex"
	commonv1 "sdxx/server/protobuf/gen/common/v1"
	hallv1 "sdxx/server/protobuf/gen/hall/v1"
	playerv1 "sdxx/server/protobuf/gen/mesh/game/player/v1"

	"github.com/dobyte/due/v2/log"
)

func (c *Core) handleFetchPlayerProfile(ctx nodex.Context) {
	ctx.Task(func(ctx nodex.Context) {
		req := new(hallv1.FetchPlayerProfileReq)
		if err := ctx.Parse(req); err != nil {
			log.Errorf("parse request message failed: %v", err)
			ctx.Error(commonv1.MessageDecodingFailed)
			return
		}

		playerProfile, err := c.playerServiceClient.GetProfile(ctx.Ctx(), &playerv1.GetProfileReq{
			PlayerId: ctx.UID(),
		})
		if err != nil {
			log.Errorf("get player profile failed: %v", err)
			ctx.Error(commonv1.InternalError)
			return
		}

		ctx.Response(&hallv1.FetchPlayerProfileResp{
			Nickname: playerProfile.Nickname,
			Avatar:   playerProfile.Avatar,
			Level:    playerProfile.Level,
			Assets:   playerProfile.Assets,
		})
	})
}
