package middleware

import (
	commonv1 "sdxx/server/protobuf/gen/common/v1"
	corev1 "sdxx/server/protobuf/gen/core/v1"

	"github.com/dobyte/due/v2/cluster/node"
	"github.com/dobyte/due/v2/log"
)

func Auth(middleware *node.Middleware, ctx node.Context) {
	if ctx.UID() == 0 {
		err := ctx.Response(&corev1.Response{
			Code: int32(commonv1.Unauthorized.Code()),
			Msg:  commonv1.Unauthorized.Message(),
		})
		if err != nil {
			log.Errorf("response message failed, err: %v", err)
		}
	} else {
		middleware.Next(ctx)
	}
}
