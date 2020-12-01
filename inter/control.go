package inter

import "github.com/gogf/gf/net/ghttp"

type Control interface {
	// 参数验证
	Parse(r *ghttp.Request, data interface{}) error
	// 成功请求响应
	Success(r *ghttp.Request, data ...interface{})
	// 失败响应
	Fail(r *ghttp.Request, err error, data ...interface{})
	// 构造函数
	Init(r *ghttp.Request)
	// 析构函数
	Shut(r *ghttp.Request)
}
