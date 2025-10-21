package helper

import (
	cfg "sdxx/server/internal/config/gen"
	"sync/atomic"
)

func NewTbLevel(buf []map[string]any) *TbLevel {
	t := &TbLevel{}
	t.Reload(buf)
	return t
}

type TbLevel struct {
	p atomic.Pointer[cfg.TbLevel]
}

func (t *TbLevel) Load() *cfg.TbLevel {
	return t.p.Load()
}

func (t *TbLevel) Reload(buf []map[string]any) error {
	tbl, err := cfg.NewTbLevel(buf)
	if err != nil {
		return err
	}
	t.p.Store(tbl)
	return nil
}

func (t *TbLevel) Name() string {
	return "tblevel"
}
