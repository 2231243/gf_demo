package router

import (
	"gf-app/app/api/certificate"
	"gf-app/app/api/im_user"
	"gf-app/app/api/message"
	app "gf-app/app/api/ofc_app"
	"gf-app/app/api/user"
	"gf-app/middleware"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	s.Use(middleware.TraceMiddleware)
	// 跨域
	s.Use(middleware.MiddlewareCORS)
	// 捕获程序崩溃异常
	s.Use(middleware.MiddlewareErrorHandler)
	// 请求验权
	s.Use(middleware.MiddlewareAuth)
	// 请求参数验证
	//s.Use(middleware.Parse)
	//gzip压缩response
	//s.Use(middleware.Compress(middleware.BestCompression,nil))
	//gzip压缩请求
	//s.Use(middleware.Decompress)

	ofcApp := app.NewOfcApp()
	ofcUser := user.NewUser()
	cert := certificate.NewCert()
	msgAPI := message.NewMessage()
	imUser := im_user.NewImUser()

	s.Group(`/api/v1`, func(group *ghttp.RouterGroup) {
		// 用户模块
		group.ALL(`/user`, ofcUser, "Index, Register, Login, Sms, Logout")
		// 应用模块
		group.ALL(`/app`, ofcApp)
		// 证书模块
		group.ALL(`/cert`, cert)
		// 消息统计
		group.ALL(`/message`, msgAPI)
		// 用户统计
		group.ALL(`/imUser`, imUser)
	})

}
