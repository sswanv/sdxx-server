package config

import (
	"sdxx/server/internal/config/helper"

	"github.com/dobyte/due/v2/log"
)

func NewTables() *Tables {
	t := &Tables{
		tables:     helper.NewBaseTables(helper.DefaultJsonLoader),
		TbRole:     &helper.TbRole{},
		TbLevel:    &helper.TbLevel{},
		TbGlobal:   &helper.TbGlobal{},
		TbNickname: &helper.TbNickname{},
		TbProperty: &helper.TbProperty{},
	}
	if err := t.tables.Init(t); err != nil {
		log.Fatalf("init tables err: %v", err)
	}
	return t
}

type Tables struct {
	tables     *helper.BaseTables
	TbRole     *helper.TbRole
	TbLevel    *helper.TbLevel
	TbGlobal   *helper.TbGlobal
	TbNickname *helper.TbNickname
	TbProperty *helper.TbProperty
}
