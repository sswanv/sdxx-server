package core

import (
	"sdxx/server/internal/core/nodex"
	gmv1 "sdxx/server/protobuf/gen/gm/v1"
)

func (c *Core) handleGetCommands(ctx nodex.Context) {
	ctx.Task(func(ctx nodex.Context) {
		commands := c.commandManager.GetAvailableCommands()
		var results []*gmv1.GetCommandsResp_Command
		for _, command := range commands {
			results = append(results, &gmv1.GetCommandsResp_Command{
				Name:        command.GetName(),
				Description: command.GetDescription(),
				Usage:       command.GetUsage(),
			})
		}
		ctx.Response(&gmv1.GetCommandsResp{
			Commands: results,
		})
	})
}
