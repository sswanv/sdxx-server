package service

import (
	inventoryv1 "sdxx/server/protobuf/gen/mesh/game/inventory/v1"
	playerv1 "sdxx/server/protobuf/gen/mesh/game/player/v1"
	platformv1 "sdxx/server/protobuf/gen/mesh/platform/v1"

	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/transport"
	"google.golang.org/grpc"
)

func NewPlatformServiceClient(fn transport.NewMeshClient) platformv1.PlatformServiceClient {
	client, err := fn(target(Platform))
	if err != nil {
		log.Fatalf("new platform service err: %v", err)
	}
	return platformv1.NewPlatformServiceClient(client.Client().(grpc.ClientConnInterface))
}

func NewPlayerServiceClient(fn transport.NewMeshClient) playerv1.PlayerServiceClient {
	client, err := fn(target(GamePlayer))
	if err != nil {
		log.Fatalf("new player service err: %v", err)
	}
	return playerv1.NewPlayerServiceClient(client.Client().(grpc.ClientConnInterface))
}

func NewInventoryServiceClient(fn transport.NewMeshClient) inventoryv1.InventoryServiceClient {
	client, err := fn(target(GameInventory))
	if err != nil {
		log.Fatalf("new inventory service err: %v", err)
	}
	return inventoryv1.NewInventoryServiceClient(client.Client().(grpc.ClientConnInterface))
}
