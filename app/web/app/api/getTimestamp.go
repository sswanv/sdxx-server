package api

import (
	"sdxx/server/internal/component/http"
	webv1 "sdxx/server/protobuf/gen/web/v1"

	"github.com/dobyte/due/v2/utils/xtime"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (a *Api) GetTimestamp(ctx http.Context) error {
	return ctx.Success(&webv1.GetTimestampResp{
		Timestamp: timestamppb.New(xtime.Now()),
	})
}
