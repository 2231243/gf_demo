package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
)

func MiddlewareCORS(r *ghttp.Request) {
	// 设置开始请求时间
	r.SetParam("requestTime", gtime.Now().Format("Y-m-d H:i:s.u"))
	r.Response.CORSDefault()
	r.Middleware.Next()
}
