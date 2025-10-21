package api

import (
	"sdxx/server/internal/code"
	"sdxx/server/internal/component/http"
	"sdxx/server/internal/types"
	platformv1 "sdxx/server/protobuf/gen/mesh/platform/v1"
	webv1 "sdxx/server/protobuf/gen/web/v1"
)

func (a *Api) MobileLogin(ctx http.Context) error {
	req := new(webv1.MobileLoginReq)
	if err := ctx.Bind().JSON(req); err != nil {
		return ctx.Failure(code.InvalidArgument)
	}

	if req.DeviceId == "" {
		return ctx.Failure(code.InvalidArgument)
	}

	_, err := a.platformServiceClient.Register(ctx.Context(), &platformv1.RegisterReq{
		Channel:    int32(types.ChannelMobile),
		Identifier: req.Mobile,
		Credential: "",
		DeviceType: req.DeviceType,
		DeviceId:   req.DeviceId,
		ClientIp:   ctx.IP(),
	})
	if err != nil {
		return ctx.Failure(err)
	}

	loginResp, err := a.platformServiceClient.Login(ctx.Context(), &platformv1.LoginReq{
		Channel:    int32(types.ChannelMobile),
		Identifier: req.Mobile,
		Credential: "",
		DeviceType: req.DeviceType,
		DeviceId:   req.DeviceId,
		ClientIp:   ctx.IP(),
	})
	if err != nil {
		return ctx.Failure(err)
	}

	return ctx.Success(&webv1.MobileLoginResp{
		AccessToken:  loginResp.AccessToken,
		AccessExpire: loginResp.AccessExpire,
		RefreshAfter: loginResp.RefreshAfter,
	})
}
