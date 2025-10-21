package api

import (
	"sdxx/server/internal/code"
	"sdxx/server/internal/component/http"
	platformv1 "sdxx/server/protobuf/gen/mesh/platform/v1"
	webv1 "sdxx/server/protobuf/gen/web/v1"
)

func (a *Api) RefreshToken(ctx http.Context) error {
	req := new(webv1.RefreshTokenReq)
	if err := ctx.Bind().JSON(req); err != nil {
		return ctx.Failure(code.InvalidArgument)
	}

	// if req.AccessToken == "" {
	// 	return ctx.Failure(code.InvalidArgument)
	// }

	refreshTokenResp, err := a.platformServiceClient.RefreshToken(ctx.Context(), &platformv1.RefreshTokenReq{
		AccessToken: req.AccessToken,
	})
	if err != nil {
		return ctx.Failure(err)
	}

	return ctx.Success(&platformv1.RefreshTokenResp{
		AccessToken:  refreshTokenResp.AccessToken,
		AccessExpire: refreshTokenResp.AccessExpire,
		RefreshAfter: refreshTokenResp.RefreshAfter,
	})
}
