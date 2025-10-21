package helper

import (
	"reflect"

	"github.com/dobyte/due/v2/config"
	"github.com/dobyte/due/v2/log"
	"github.com/pkg/errors"
)

type Table interface {
	Name() string
	Reload([]map[string]any) error
}

func NewBaseTables(loader JsonLoader) *BaseTables {
	return &BaseTables{
		loader: loader,
	}
}

type BaseTables struct {
	loader JsonLoader
	tables []Table
}

func (bt *BaseTables) autoRegister(tables any) {
	val := reflect.ValueOf(tables).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)

		if !fieldType.IsExported() {
			continue
		}
		if field.Kind() != reflect.Ptr || field.IsNil() {
			continue
		}
		if table, ok := field.Interface().(Table); ok {
			bt.tables = append(bt.tables, table)
			log.Infof("table registered: %s (%s)", table.Name(), fieldType.Name)
		}
	}
}

func (bt *BaseTables) All() []Table {
	return bt.tables
}

func (bt *BaseTables) Init(tables any) error {
	bt.autoRegister(tables)

	for _, t := range bt.tables {
		buf, err := bt.loader(t.Name())
		if err != nil {
			return errors.Errorf("load %s failed: %v", t.Name(), err)
		}
		if err := t.Reload(buf); err != nil {
			return errors.Errorf("parse %s failed: %v", t.Name(), err)
		}
		log.Infof("init %s success", t.Name())
	}

	bt.watch()

	return nil
}

func (bt *BaseTables) watch() {
	var names []string
	for _, t := range bt.tables {
		names = append(names, t.Name())
	}
	config.Watch(func(updated ...string) {
		for _, name := range updated {
			for _, t := range bt.tables {
				if t.Name() == name {
					buf, err := bt.loader(name)
					if err != nil {
						log.Errorf("tables load %s failed: %v\n", name, err)
						continue
					}
					if err := t.Reload(buf); err != nil {
						log.Errorf("tables reload %s failed: %v\n", name, err)
					} else {
						log.Infof("tables reload %s success\n", name)
					}
				}
			}
		}
	}, names...)
}
