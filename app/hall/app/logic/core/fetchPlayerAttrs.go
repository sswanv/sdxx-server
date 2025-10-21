package core

import (
	"sdxx/server/internal/core/nodex"
	hallv1 "sdxx/server/protobuf/gen/hall/v1"
	playerv1 "sdxx/server/protobuf/gen/mesh/game/player/v1"
)

func (c *Core) handleFetchPlayerAttrs(ctx nodex.Context) {
	ctx.Task(func(ctx nodex.Context) {
		getAttrsResp, err := c.playerServiceClient.GetAttrs(ctx.Ctx(), &playerv1.GetAttrsReq{
			PlayerId: ctx.UID(),
		})
		if err != nil {
			ctx.Error(err)
			return
		}
		resp := &hallv1.FetchPlayerAttrsResp{
			Attrs:      getAttrsResp.Attrs,
			Equipments: getAttrsResp.Equipments,
		}
		ctx.Response(resp)
	})
}
