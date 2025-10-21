package helper

import (
	cfg "sdxx/server/internal/config/gen"
	"sync/atomic"
)

func NewTbServerList(buf []map[string]any) *TbServerList {
	t := &TbServerList{}
	t.Reload(buf)
	return t
}

type TbServerList struct {
	p atomic.Pointer[cfg.TbServerList]
}

func (t *TbServerList) Load() *cfg.TbServerList {
	return t.p.Load()
}

func (t *TbServerList) Reload(buf []map[string]any) error {
	tbl, err := cfg.NewTbServerList(buf)
	if err != nil {
		return err
	}
	t.p.Store(tbl)
	return nil
}

func (t *TbServerList) Name() string {
	return "tbserverlist"
}
