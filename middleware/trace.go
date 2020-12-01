package middleware

import (
	"time"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
)

func TraceMiddleware(r *ghttp.Request) {
	glog.SetAsync(true)
	logger := glog.New()
	logger.SetLevel(glog.LEVEL_INFO)
	s := time.Now()
	r.Middleware.Next()
	logger.Info(g.Map{
		"请求时间":  gtime.Now().Format("Y-m-d H:i:s.u"),
		"请求URL": r.URL.Path,
		"响应时间":  time.Since(s).String(),
	})
}
