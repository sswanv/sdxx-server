package core

import (
	"sdxx/server/internal/core/nodex"
	commonv1 "sdxx/server/protobuf/gen/common/v1"
	hallv1 "sdxx/server/protobuf/gen/hall/v1"
	playerv1 "sdxx/server/protobuf/gen/mesh/game/player/v1"
	platformv1 "sdxx/server/protobuf/gen/mesh/platform/v1"

	"github.com/dobyte/due/v2/log"
)

func (c *Core) handleLogin(ctx nodex.Context) {
	ctx.Task(func(ctx nodex.Context) {
		req := new(hallv1.LoginReq)
		if err := ctx.Parse(req); err != nil {
			log.Errorf("parse request message failed: %v", err)
			ctx.Error(commonv1.MessageDecodingFailed)
			return
		}
		if req.AccessToken == "" {
			ctx.Error(commonv1.InvalidArgument)
			return
		}
		if req.ServerId == 0 {
			ctx.Error(commonv1.InvalidArgument)
			return
		}

		validateTokenResp, err := c.platformServiceClient.ValidateToken(ctx.Ctx(), &platformv1.ValidateTokenReq{AccessToken: req.AccessToken})
		if err != nil {
			ctx.Error(err)
			return
		}
		ensurePlayerResp, err := c.playerServiceClient.EnsurePlayer(ctx.Ctx(), &playerv1.EnsurePlayerReq{
			AccountId: validateTokenResp.AccountId,
			ServerId:  req.ServerId,
		})
		if err != nil {
			ctx.Error(err)
			return
		}

		if err = ctx.BindGate(int64(ensurePlayerResp.PlayerId)); err != nil {
			log.Errorf("bind gate failed, play_id = %v err = %v", validateTokenResp.AccountId, err)
			ctx.Error(commonv1.InternalError)
			return
		}

		ctx.Response(&hallv1.LoginResp{})
	})
}
