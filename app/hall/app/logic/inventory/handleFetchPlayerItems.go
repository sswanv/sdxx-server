package inventory

import (
	"sdxx/server/internal/core/nodex"
	commonv1 "sdxx/server/protobuf/gen/common/v1"
	hallv1 "sdxx/server/protobuf/gen/hall/v1"
	inventoryv1 "sdxx/server/protobuf/gen/mesh/game/inventory/v1"

	"github.com/dobyte/due/v2/log"
)

func (c *Inventory) handleFetchPlayerItems(ctx nodex.Context) {
	ctx.Task(func(ctx nodex.Context) {
		req := new(hallv1.FetchPlayerItemsReq)
		if err := ctx.Parse(req); err != nil {
			log.Errorf("parse request message failed: %v", err)
			ctx.Error(commonv1.MessageDecodingFailed)
			return
		}

		resp, err := c.inventoryServiceClient.GetPlayerItems(ctx.Ctx(), &inventoryv1.GetPlayerItemsReq{
			PlayerId: ctx.UID(),
		})
		if err != nil {
			log.Errorf("get player items failed: %v", err)
			ctx.Error(commonv1.InternalError)
			return
		}
		ctx.Response(resp)
	})
}
