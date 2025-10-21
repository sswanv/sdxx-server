package core

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

func New(proxy *node.Proxy) *Core {
	return &Core{
		proxy:                  proxy,
		platformServiceClient:  service.NewPlatformServiceClient(proxy.NewMeshClient),
		playerServiceClient:    service.NewPlayerServiceClient(proxy.NewMeshClient),
		inventoryServiceClient: service.NewInventoryServiceClient(proxy.NewMeshClient),
	}
}

type Core struct {
	proxy                  *node.Proxy
	platformServiceClient  platformv1.PlatformServiceClient
	playerServiceClient    playerv1.PlayerServiceClient
	inventoryServiceClient inventoryv1.InventoryServiceClient
}

func (c *Core) Init() {
	c.proxy.Router().Group(func(group *node.RouterGroup) {
		group.AddRouteHandler(int32(hallv1.Route_ROUTE_LOGIN), nodex.W(c.handleLogin))
		group.Middleware(middleware.Auth)
		group.AddRouteHandler(int32(hallv1.Route_ROUTE_FETCH_PLAYER_PROFILE), nodex.W(c.handleFetchPlayerProfile))
		group.AddRouteHandler(int32(hallv1.Route_ROUTE_FETCH_PLAYER_ATTRS), nodex.W(c.handleFetchPlayerAttrs))
	})
}
