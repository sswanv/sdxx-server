package inventory

import (
	"sdxx/server/app/mesh/app/game/inventory/config"
	"sdxx/server/internal/component/mysql"
	rediscomp "sdxx/server/internal/component/redis"
	"sdxx/server/internal/dao/mysql/game/model"
	"sdxx/server/internal/service"
	inventoryv1 "sdxx/server/protobuf/gen/mesh/game/inventory/v1"

	"github.com/dobyte/due/v2/cluster/mesh"
	"gorm.io/gorm"
)

func NewServer(proxy *mesh.Proxy) *Server {
	db := mysql.NewInstance("etc.mysql.game")
	return &Server{
		proxy:             proxy,
		table:             config.NewTables(),
		redis:             rediscomp.Instance(),
		db:                db,
		playerAssetsModel: model.NewPlayerAssetsModel(db),
		playerItemModel:   model.NewPlayerItemModel(db),
	}
}

type Server struct {
	inventoryv1.UnimplementedInventoryServiceServer
	proxy             *mesh.Proxy
	table             *config.Tables
	redis             rediscomp.Redis
	db                *gorm.DB
	playerAssetsModel model.PlayerAssetsModel
	playerItemModel   model.PlayerItemModel
}

func (s *Server) Init() {
	s.proxy.AddServiceProvider(service.GameInventory, &inventoryv1.InventoryService_ServiceDesc, s)
}
