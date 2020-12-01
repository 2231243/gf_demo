package middleware

import (
	"gf-app/boot"
	"gf-app/library"
	"net/http"
	"strings"

	"github.com/gogf/gf/net/ghttp"
)

// jwt验证
func MiddlewareAuth(r *ghttp.Request) {

	if boot.WhiteRoute.Contains(r.URL.Path) || strings.Contains(r.URL.Path,"debug") {
		r.Middleware.Next()
		return
	}
	_, b := library.Auth(r)
	if !b {
		library.JsonExit(r, http.StatusForbidden, "auth token not found or illegal")
	}
	//if ok := r.Session.GetBool(guid); !ok {
	//	library.JsonExit(r, http.StatusForbidden, "当前会话已失效,请重新登录")
	//}
	r.Middleware.Next()
}
