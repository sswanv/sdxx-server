package app

import (
	"sdxx/server/app/web/app/api"
	"sdxx/server/internal/component/http"
)

func Init(proxy *http.Proxy) {
	a := api.NewApi(proxy)
	v1 := proxy.Router().Group("/api/v1")
	v1.Post("/get-timestamp", a.GetTimestamp)
	v1.Post("/get-server-list", a.GetServerList)
	v1.Post("/send-mobile-code", a.SendMobileCode)
	v1.Post("/mobile-login", a.MobileLogin)
	v1.Post("/refresh-token", a.RefreshToken)
}
