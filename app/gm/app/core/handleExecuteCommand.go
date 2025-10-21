package core

import (
	"sdxx/server/internal/core/nodex"
	gmv1 "sdxx/server/protobuf/gen/gm/v1"
)

func (c *Core) handleExecuteCommand(ctx nodex.Context) {
	ctx.Task(func(ctx nodex.Context) {
		req := new(gmv1.ExecuteCommandReq)
		if err := ctx.Parse(req); err != nil {
			ctx.Error(err)
			return
		}

		err := c.commandManager.ExecuteCommand(ctx, req.Command, req.Args)
		if err != nil {
			ctx.Error(err)
			return
		}

		ctx.Response(&gmv1.ExecuteCommandResp{})
	})
}
