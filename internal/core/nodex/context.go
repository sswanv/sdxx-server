package nodex

import (
	"context"
	"sdxx/server/internal/code"
	corev1 "sdxx/server/protobuf/gen/core/v1"

	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/codes"
	"github.com/dobyte/due/v2/log"
	"github.com/zeromicro/go-zero/core/logx"
	proto "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type Context struct {
	node.Context
}

func (c Context) Response(message proto.Message) {
	data, err := anypb.New(message)
	if err != nil {
		logx.Errorf(" marshals message into a new Any instance err: %v", err)
		return
	}
	c.Context.Response(&corev1.Response{
		Code: int32(code.OK.Code()),
		Msg:  code.OK.Message(),
		Data: data,
	})
}

func (c Context) Error(err any) {
	resp := &corev1.Response{}
	switch v := err.(type) {
	case error:
		code := codes.Convert(v)
		resp.Code = int32(code.Code())
		resp.Msg = code.Message()
	case *codes.Code:
		resp.Code = int32(v.Code())
		resp.Msg = v.Message()
	default:
		resp.Code = int32(code.Unknown.Code())
		resp.Msg = code.Unknown.Message()
	}
	if err := c.Context.Response(resp); err != nil {
		log.Errorf("response message failed: %v", err)
	}
}

func (c Context) Task(fn func(ctx Context)) {
	c.Context.Task(func(ctx node.Context) {
		fn(Context{Context: ctx})
	})
}

func (c Context) Ctx() context.Context {
	return c.Context.Context()
}

func (c Context) UID() uint64 {
	return uint64(c.Context.UID())
}
