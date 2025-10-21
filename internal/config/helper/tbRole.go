package helper

import (
	cfg "sdxx/server/internal/config/gen"
	"sync/atomic"
)

func NewTbRole(buf []map[string]any) *TbRole {
	t := &TbRole{}
	t.Reload(buf)
	return t
}

type TbRole struct {
	p atomic.Pointer[cfg.TbRole]
}

func (t *TbRole) Load() *cfg.TbRole {
	return t.p.Load()
}

func (t *TbRole) Reload(buf []map[string]any) error {
	tbl, err := cfg.NewTbRole(buf)
	if err != nil {
		return err
	}
	t.p.Store(tbl)
	return nil
}

func (t *TbRole) Name() string {
	return "tbrole"
}
