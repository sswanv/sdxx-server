package inventory

import (
	"sdxx/server/internal/core/nodex"
	"sdxx/server/internal/middleware"
	"sdxx/server/internal/service"
	hallv1 "sdxx/server/protobuf/gen/hall/v1"
	inventoryv1 "sdxx/server/protobuf/gen/mesh/game/inventory/v1"
	playerv1 "sdxx/server/protobuf/gen/mesh/game/player/v1"
	platformv1 "sdxx/server/protobuf/gen/mesh/platform/v1"

	"github.com/dobyte/due/v2/cluster/node"
)

func New(proxy *node.Proxy) *Inventory {
	return &Inventory{
		proxy:                  proxy,
		inventoryServiceClient: service.NewInventoryServiceClient(proxy.NewMeshClient),
	}
}

type Inventory struct {
	proxy                  *node.Proxy
	platformServiceClient  platformv1.PlatformServiceClient
	playerServiceClient    playerv1.PlayerServiceClient
	inventoryServiceClient inventoryv1.InventoryServiceClient
}

func (c *Inventory) Init() {
	c.proxy.Router().Group(func(group *node.RouterGroup) {
		group.Middleware(middleware.Auth)
		group.AddRouteHandler(int32(hallv1.Route_ROUTE_FETCH_PLAYER_ITEMS), nodex.W(c.handleFetchPlayerItems))
		group.AddRouteHandler(int32(hallv1.Route_ROUTE_FETCH_PLAYER_EQUIPMENTS), nodex.W(c.handleFetchPlayerEquipments))
	})
}
