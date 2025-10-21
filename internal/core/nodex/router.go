package nodex

import "github.com/dobyte/due/v2/cluster/node"

type RouteHandler func(ctx Context)

func W(fn RouteHandler) node.RouteHandler {
	return func(ctx node.Context) {
		fn(Context{Context: ctx})
	}
}
