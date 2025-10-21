package helper

import (
	cfg "sdxx/server/internal/config/gen"
	"sync/atomic"
)

func NewTbGlobal(buf []map[string]any) *TbGlobal {
	t := &TbGlobal{}
	t.Reload(buf)
	return t
}

type TbGlobal struct {
	p atomic.Pointer[cfg.TbGlobal]
}

func (t *TbGlobal) Load() *cfg.Global {
	return t.p.Load().Get(0)
}

func (t *TbGlobal) Reload(buf []map[string]any) error {
	tbl, err := cfg.NewTbGlobal(buf)
	if err != nil {
		return err
	}
	t.p.Store(tbl)
	return nil
}

func (t *TbGlobal) Name() string {
	return "tbglobal"
}
