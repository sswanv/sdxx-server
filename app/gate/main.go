package main

import (
	"net/http"
	"sdxx/server/internal/component/jwt"

	"github.com/dobyte/due/locate/redis/v2"
	"github.com/dobyte/due/network/ws/v2"
	"github.com/dobyte/due/registry/nacos/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/cluster/gate"
)

func main() {
	container := due.NewContainer()
	server := ws.NewServer()
	server.OnUpgrade(func(w http.ResponseWriter, r *http.Request) bool {
		if _, err := jwt.Instance().Http().ExtractPayload(r, true); err != nil {
			return false
		}
		return true
	})
	locator := redis.NewLocator()
	registry := nacos.NewRegistry()
	component := gate.NewGate(
		gate.WithServer(server),
		gate.WithLocator(locator),
		gate.WithRegistry(registry),
	)
	container.Add(component)
	container.Serve()
}
