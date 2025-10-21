package helper

import (
	cfg "sdxx/server/internal/config/gen"
	"sync/atomic"
)

func NewTbProperty(buf []map[string]any) *TbProperty {
	t := &TbProperty{}
	t.Reload(buf)
	return t
}

type TbProperty struct {
	p atomic.Pointer[cfg.TbProperty]
}

func (t *TbProperty) Load() *cfg.TbProperty {
	return t.p.Load()
}

func (t *TbProperty) Reload(buf []map[string]any) error {
	tbl, err := cfg.NewTbProperty(buf)
	if err != nil {
		return err
	}
	t.p.Store(tbl)
	return nil
}

func (t *TbProperty) Name() string {
	return "tbproperty"
}
