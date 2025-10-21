package main

import (
	"sdxx/server/app/web/app"
	"sdxx/server/app/web/codec"
	"sdxx/server/internal/component/http"

	"github.com/dobyte/due/registry/nacos/v2"
	"github.com/dobyte/due/transport/grpc/v2"
	"github.com/dobyte/due/v2"
)

func main() {
	container := due.NewContainer()
	registry := nacos.NewRegistry()
	transporter := grpc.NewTransporter()
	component := http.NewServer(
		http.WithRegistry(registry),
		http.WithTransporter(transporter),
		http.WithJsonEncoder(codec.JsonEncoder),
		http.WithJsonDecoder(codec.JsonDecoder),
	)
	app.Init(component.Proxy())
	container.Add(component)
	container.Serve()
}
