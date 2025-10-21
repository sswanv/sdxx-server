package player

import (
	"sdxx/server/app/mesh/app/game/player/config"
	"sdxx/server/internal/component/mysql"
	rediscomp "sdxx/server/internal/component/redis"
	"sdxx/server/internal/dao/mysql/game/model"
	"sdxx/server/internal/service"
	inventoryv1 "sdxx/server/protobuf/gen/mesh/game/inventory/v1"
	playerv1 "sdxx/server/protobuf/gen/mesh/game/player/v1"

	"github.com/dobyte/due/v2/cluster/mesh"
	"gorm.io/gorm"
)

func NewServer(proxy *mesh.Proxy) *Server {
	db := mysql.NewInstance("etc.mysql.game")
	return &Server{
		proxy:                  proxy,
		table:                  config.NewTables(),
		redis:                  rediscomp.Instance(),
		db:                     db,
		inventoryServiceClient: service.NewInventoryServiceClient(proxy.NewMeshClient),
		playerModel:            model.NewPlayerModel(db),
	}
}

type Server struct {
	playerv1.UnimplementedPlayerServiceServer
	proxy                  *mesh.Proxy
	table                  *config.Tables
	redis                  rediscomp.Redis
	db                     *gorm.DB
	inventoryServiceClient inventoryv1.InventoryServiceClient
	playerModel            model.PlayerModel
}

func (s *Server) Init() {
	s.proxy.AddServiceProvider(service.GamePlayer, &playerv1.PlayerService_ServiceDesc, s)
}
