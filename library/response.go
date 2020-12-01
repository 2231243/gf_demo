package library

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
)

// 数据返回通用JSON数据结构
type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

// 标准返回结果数据结构封装。
// @param r *gttp.Request
// @param   code int  true "响应状态码"
// @param   message string  true "响应信息"
// @param  data ...interface{} 0 响应结构 1 HTTP请求参数
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	response := JsonResponse{
		Code:    code,
		Message: message,
		Data:    responseData,
	}
	// 打印请求日志
	glog.SetAsync(true)
	l := glog.New()
	l.SetLevel(glog.LEVEL_INFO)
	if len(data) > 1 {
		gtime.Now().Format("Y-m-d H:i:s.u")
		l.Info(g.Map{"url": r.URL.Path, "header": r.Header, "request": data[1],
			"response": response, "requestTime": r.GetParamVar("requestTime").String()})
	} else {
		l.Info(g.Map{"url": r.URL.Path, "header": r.Header,
			"response": response, "requestTime" : r.GetParamVar("requestTime").String()})
	}
	_ = r.Response.WriteJsonExit(response)
}

// 返回JSON数据并退出当前HTTP执行函数。
// @param r *gttp.Request
// @param   code int  true "响应状态码"
// @param   message string  true "响应信息"
// @param  data ...interface{} 0 响应结构 1 HTTP请求参数
func JsonExit(r *ghttp.Request, err int, msg string, data ...interface{}) {
	r.Response.ClearBuffer()
	Json(r, err, msg, data...)
}
