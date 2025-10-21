package app

import (
	"sdxx/server/app/gm/app/core"

	"github.com/dobyte/due/v2/cluster/node"
)

func Init(proxy *node.Proxy) {
	core.New(proxy).Init()
}
