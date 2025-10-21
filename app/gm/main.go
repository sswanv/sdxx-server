package main

import (
	"sdxx/server/app/gm/app"

	"github.com/dobyte/due/locate/redis/v2"
	"github.com/dobyte/due/registry/nacos/v2"
	"github.com/dobyte/due/transport/grpc/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/node"
)

func main() {
	container := due.NewContainer()
	locator := redis.NewLocator()
	registry := nacos.NewRegistry()
	transporter := grpc.NewTransporter()
	transporter.SetDefaultDiscovery(registry)
	component := node.NewNode(
		node.WithLocator(locator),
		node.WithRegistry(registry),
		node.WithTransporter(transporter),
	)
	app.Init(component.Proxy())
	container.Add(component)
	container.Serve()
}
