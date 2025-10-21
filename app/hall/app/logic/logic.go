package logic

import (
	"sdxx/server/app/hall/app/logic/core"
	"sdxx/server/app/hall/app/logic/inventory"

	"github.com/dobyte/due/v2/cluster/node"
)

func Init(proxy *node.Proxy) {
	core.New(proxy).Init()
	inventory.New(proxy).Init()
}
