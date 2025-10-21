package app

import (
	"sdxx/server/app/hall/app/logic"

	"github.com/dobyte/due/v2/cluster/node"
)

func Init(proxy *node.Proxy) {
	logic.Init(proxy)
}
