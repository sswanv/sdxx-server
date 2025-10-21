package main

import (
	"sdxx/server/app/mesh/app"

	rediscache "github.com/dobyte/due/cache/redis/v2"
	"github.com/dobyte/due/locate/redis/v2"
	"github.com/dobyte/due/registry/nacos/v2"
	"github.com/dobyte/due/transport/grpc/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cache"
	"github.com/dobyte/due/v2/cluster/mesh"
)

func main() {
	cache.SetCache(rediscache.NewCache())
	container := due.NewContainer()
	locator := redis.NewLocator()
	registry := nacos.NewRegistry()
	transporter := grpc.NewTransporter()
	transporter.SetDefaultDiscovery(registry)
	component := mesh.NewMesh(
		mesh.WithLocator(locator),
		mesh.WithRegistry(registry),
		mesh.WithTransporter(transporter),
	)
	app.Init(component.Proxy())
	container.Add(component)
	container.Serve()
}
