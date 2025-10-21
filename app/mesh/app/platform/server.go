package platform

import (
	"sdxx/server/app/mesh/app/platform/config"
	jwtcomp "sdxx/server/internal/component/jwt"
	"sdxx/server/internal/component/mysql"
	rediscomp "sdxx/server/internal/component/redis"
	"sdxx/server/internal/dao/mysql/platform/model"
	"sdxx/server/internal/service"
	platformv1 "sdxx/server/protobuf/gen/mesh/platform/v1"

	"github.com/dobyte/due/v2/cluster/mesh"
	"gorm.io/gorm"
)

func NewServer(proxy *mesh.Proxy) *Server {
	db := mysql.NewInstance("etc.mysql.platform")
	return &Server{
		proxy:              proxy,
		table:              config.NewTables(),
		redis:              rediscomp.Instance(),
		jwt:                jwtcomp.Instance(),
		db:                 db,
		accountModel:       model.NewAccountModel(db),
		accountAuthModel:   model.NewAccountAuthModel(db),
		accountPlayerModel: model.NewAccountPlayerModel(db),
	}
}

type Server struct {
	platformv1.UnimplementedPlatformServiceServer
	proxy              *mesh.Proxy
	table              *config.Tables
	jwt                *jwtcomp.JWT
	redis              rediscomp.Redis
	db                 *gorm.DB
	accountModel       model.AccountModel
	accountAuthModel   model.AccountAuthModel
	accountPlayerModel model.AccountPlayerModel
}

func (s *Server) Init() {
	s.proxy.AddServiceProvider(service.Platform, &platformv1.PlatformService_ServiceDesc, s)
}
