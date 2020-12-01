package middleware

import (
	"net/http"

	"gf-app/library"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 全局捕获程序异常
func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		// 记录到自定义错误日志文件
		g.Log("exception").Error(err)
		//返回固定的友好信息
		library.JsonExit(r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}
