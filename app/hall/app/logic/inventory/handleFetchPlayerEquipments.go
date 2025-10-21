package inventory

import (
	"sdxx/server/internal/core/nodex"
	commonv1 "sdxx/server/protobuf/gen/common/v1"
	hallv1 "sdxx/server/protobuf/gen/hall/v1"

	"github.com/dobyte/due/v2/log"
)

func (c *Inventory) handleFetchPlayerEquipments(ctx nodex.Context) {
	ctx.Task(func(ctx nodex.Context) {
		req := new(hallv1.FetchPlayerEquipmentsReq)
		if err := ctx.Parse(req); err != nil {
			log.Errorf("parse request message failed: %v", err)
			ctx.Error(commonv1.MessageDecodingFailed)
			return
		}
	})
}
