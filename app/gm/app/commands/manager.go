// app/gm/app/commands/manager.go
package commands

import (
	"fmt"
	"sdxx/server/internal/core/nodex"
	commonv1 "sdxx/server/protobuf/gen/common/v1"
	inventoryv1 "sdxx/server/protobuf/gen/mesh/game/inventory/v1"

	"github.com/dobyte/due/v2/errors"
)

type CommandManager struct {
	mapOfCommands map[string]Command
	commands      []Command
}

func NewCommandManager(
	inventoryServiceClient inventoryv1.InventoryServiceClient,
) *CommandManager {
	manager := &CommandManager{
		mapOfCommands: make(map[string]Command),
	}

	manager.registerCommands(inventoryServiceClient)

	return manager
}

func (m *CommandManager) registerCommands(
	inventoryServiceClient inventoryv1.InventoryServiceClient,
) {
	command := NewAddItemCommand(inventoryServiceClient)
	m.mapOfCommands["add_item"] = command
	m.commands = append(m.commands, command)
}

func (m *CommandManager) ExecuteCommand(ctx nodex.Context, commandName string, args string) error {
	cmd, exists := m.mapOfCommands[commandName]
	if !exists {
		return errors.NewError(commonv1.NotFound.WithMessage(fmt.Sprintf("command '%s' not found", commandName)))
	}

	return cmd.Execute(ctx, args)
}

func (m *CommandManager) GetAvailableCommands() []Command {
	return m.commands
}
