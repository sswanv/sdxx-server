package core

import (
	"sdxx/server/app/gm/app/commands"
	"sdxx/server/internal/core/nodex"
	"sdxx/server/internal/middleware"
	"sdxx/server/internal/service"
	gmv1 "sdxx/server/protobuf/gen/gm/v1"
	inventoryv1 "sdxx/server/protobuf/gen/mesh/game/inventory/v1"

	"github.com/dobyte/due/v2/cluster/node"
)

type Core struct {
	proxy                  *node.Proxy
	commandManager         *commands.CommandManager
	inventoryServiceClient inventoryv1.InventoryServiceClient
}

func New(proxy *node.Proxy) *Core {
	inventoryServiceClient := service.NewInventoryServiceClient(proxy.NewMeshClient)
	return &Core{
		proxy:                  proxy,
		commandManager:         commands.NewCommandManager(inventoryServiceClient),
		inventoryServiceClient: inventoryServiceClient,
	}
}

func (c *Core) Init() {
	c.proxy.Router().Group(func(group *node.RouterGroup) {
		group.Middleware(middleware.Auth)
		group.AddRouteHandler(int32(gmv1.Route_ROUTE_EXECUTE_COMMAND), nodex.W(c.handleExecuteCommand))
		group.AddRouteHandler(int32(gmv1.Route_ROUTE_GET_COMMANDS), nodex.W(c.handleGetCommands))
	})
}
