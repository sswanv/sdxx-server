package helper

import (
	cfg "sdxx/server/internal/config/gen"
	"sync/atomic"
)

func NewTbItem(buf []map[string]any) *TbItem {
	t := &TbItem{}
	t.Reload(buf)
	return t
}

type TbItem struct {
	p atomic.Pointer[cfg.TbItem]
}

func (t *TbItem) Load() *cfg.TbItem {
	return t.p.Load()
}

func (t *TbItem) Reload(buf []map[string]any) error {
	tbl, err := cfg.NewTbItem(buf)
	if err != nil {
		return err
	}
	t.p.Store(tbl)
	return nil
}

func (t *TbItem) Name() string {
	return "tbitem"
}
