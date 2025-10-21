// internal/gm/command/command.go
package commands

import "sdxx/server/internal/core/nodex"

type Command interface {
	Execute(ctx nodex.Context, args string) error
	GetName() string
	GetDescription() string
	GetUsage() string
}

type BaseCommand struct {
	Name        string
	Description string
	Usage       string
}

func (b *BaseCommand) GetName() string {
	return b.Name
}

func (b *BaseCommand) GetDescription() string {
	return b.Description
}

func (b *BaseCommand) GetUsage() string {
	return b.Usage
}
