// app/gm/app/commands/add_item_command.go
package commands

import (
	"sdxx/server/internal/core/nodex"
	commonv1 "sdxx/server/protobuf/gen/common/v1"
	inventoryv1 "sdxx/server/protobuf/gen/mesh/game/inventory/v1"
	"strconv"
	"strings"

	"github.com/dobyte/due/v2/errors"
)

type AddItemCommand struct {
	BaseCommand
	inventoryServiceClient inventoryv1.InventoryServiceClient
}

func NewAddItemCommand(inventoryServiceClient inventoryv1.InventoryServiceClient) *AddItemCommand {
	return &AddItemCommand{
		BaseCommand: BaseCommand{
			Name:        "add_item",
			Description: "添加道具到玩家背包",
			Usage:       "add_item <item_id> <count>",
		},
		inventoryServiceClient: inventoryServiceClient,
	}
}

func (c *AddItemCommand) Execute(ctx nodex.Context, args string) error {
	parts := strings.Fields(args)
	if len(parts) != 2 {
		return errors.NewError(commonv1.InvalidArgument.WithMessage(c.Usage))
	}

	itemId, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return errors.NewError(commonv1.InvalidArgument.WithMessage("道具ID格式错误"))
	}

	count, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return errors.NewError(commonv1.InvalidArgument.WithMessage("数量格式错误"))
	}

	_, err = c.inventoryServiceClient.AddItems(ctx.Ctx(), &inventoryv1.AddItemsReq{
		PlayerId: ctx.UID(),
		Items: []*commonv1.ItemStack{
			{
				ItemId: itemId,
				Count:  count,
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}
