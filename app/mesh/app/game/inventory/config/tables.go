package config

import (
	"sdxx/server/internal/config/helper"

	"github.com/dobyte/due/v2/log"
)

func NewTables() *Tables {
	t := &Tables{
		tables:   helper.NewBaseTables(helper.DefaultJsonLoader),
		TbItem:   &helper.TbItem{},
		TbGlobal: &helper.TbGlobal{},
	}
	if err := t.tables.Init(t); err != nil {
		log.Fatalf("init tables err: %v", err)
	}

	return t
}

type Tables struct {
	tables *helper.BaseTables

	TbItem   *helper.TbItem
	TbGlobal *helper.TbGlobal
}
