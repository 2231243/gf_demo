package api

import (
	"net/http"

	"gf-app/library"

	"github.com/gogf/gf/net/ghttp"
	"google.golang.org/grpc/status"
)

type Control struct {}

// 参数验证
func (c *Control) Parse(r *ghttp.Request, data interface{}) error {
	if err := library.Parse(r, data); err != nil {
		return err
	}
	return nil
}

// 响应成功
func (c *Control) Success(r *ghttp.Request, data ...interface{}) {
	library.JsonExit(r, 0, "请求成功", data...)
}

// 响应失败
func (c *Control) Fail(r *ghttp.Request, err error, data ...interface{}) {
	if statusErr, ok := status.FromError(err); ok {
		if int(statusErr.Code()) == 4 {
			library.JsonExit(r, http.StatusGatewayTimeout, "远程服务器请求超时", data...)
			return
		}
		library.JsonExit(r, int(statusErr.Code()), statusErr.Message(), data...)
	} else {
		library.JsonExit(r, http.StatusBadRequest, err.Error(), data...)
	}
}
//
//// 构造函数
//func (c *Control) Init(r *ghttp.Request) {
//	boot.RpcClients.Iterator(func(v interface{}) bool {
//		value, _ := boot.RpcPoolMap.Conn(v.(string))
//		c.Conn.Set(v, value)
//		return true
//	})
//}
//
//// 析构函数
//func (c *Control) Shut(r *ghttp.Request) {
//
//	c.Conn.Iterator(func(k interface{}, v interface{}) bool {
//		v.(inter.RpcInter).Close()
//		return true
//	})
//
//}
