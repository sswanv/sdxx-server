package service

import "fmt"

const (
	// platform
	Platform = "platform" // 平台服务
	// game
	GamePlayer    = "game-player"    // 玩家服务
	GameInventory = "game-inventory" // 背包服务
)

func target(service string) string {
	return fmt.Sprintf("discovery://%s", service)
}
