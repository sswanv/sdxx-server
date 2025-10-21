package app

import (
	"sdxx/server/app/mesh/app/game/inventory"
	"sdxx/server/app/mesh/app/game/player"
	"sdxx/server/app/mesh/app/platform"
	"sdxx/server/internal/component/snowflake"

	"github.com/dobyte/due/config/nacos/v2"
	"github.com/dobyte/due/v2/cluster/mesh"
	"github.com/dobyte/due/v2/config"
	"github.com/dobyte/due/v2/config/file"
	"github.com/dobyte/due/v2/mode"
	"github.com/zeromicro/go-zero/core/logx"
)

func Init(proxy *mesh.Proxy) {
	// 禁用go-zero状态日志
	logx.DisableStat()
	// 初始化雪花id
	snowflake.Instance()
	// 配置初始化
	{
		var (
			source config.Source
		)
		if mode.IsDebugMode() {
			source = file.NewSource()
		} else {
			source = nacos.NewSource()
		}
		config.SetConfigurator(config.NewConfigurator(config.WithSources(source)))
	}
	{
		// platofrm
		platform.NewServer(proxy).Init()
		// game
		player.NewServer(proxy).Init()
		inventory.NewServer(proxy).Init()
	}
}
