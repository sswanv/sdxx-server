package helper

import (
	"math/rand/v2"
	cfg "sdxx/server/internal/config/gen"
	"sync/atomic"
)

func NewTbNickname(buf []map[string]any) *TbNickname {
	t := &TbNickname{}
	t.Reload(buf)
	return t
}

type TbNickname struct {
	p atomic.Pointer[cfg.TbNickname]
}

func (t *TbNickname) Load() *cfg.TbNickname {
	return t.p.Load()
}

func (t *TbNickname) Reload(buf []map[string]any) error {
	tbl, err := cfg.NewTbNickname(buf)
	if err != nil {
		return err
	}
	t.p.Store(tbl)
	return nil
}

func (t *TbNickname) Name() string {
	return "tbnickname"
}

func (t *TbNickname) TakeOne() string {
	tbNickname := t.Load()
	n := len(tbNickname.GetDataList())
	if n == 0 {
		return ""
	}
	index := rand.IntN(n)
	return tbNickname.Get(index).Name
}
