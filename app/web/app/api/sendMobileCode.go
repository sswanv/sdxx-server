package api

import (
	"sdxx/server/internal/component/http"
	webv1 "sdxx/server/protobuf/gen/web/v1"
)

func (a *Api) SendMobileCode(ctx http.Context) error {
	return ctx.Success(&webv1.SendMobileCodeResp{})
}
