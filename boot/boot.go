package boot

import (
	"net/http"
	"sync"
	"time"

	"gf-app/inter"
	"gf-app/library"
	"gf-app/packed"

	"github.com/gogf/gf/container/gset"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/shimingyah/pool"
	"github.com/gogf/gf/os/gsession"
)

// 用于应用初始化。

var (
	once                   sync.Once
	RpcPoolMap             inter.GetRpcConn
	WhiteRoute, RpcClients *gset.Set
)

const (
	OFC_COM_SVR = "127.0.0.1:50051"
	IM_MSG_SVR  = "im-msg-svr:50051"
)

// 路由设置初始化
func RouteInit() {
	s := g.Server()
	// 设置路由path模式 驼峰类型
	s.SetNameToUriType(ghttp.URI_TYPE_CAMEL)

	// 服务器状态404 500异常捕获响应
	s.BindStatusHandlerByMap(map[int]ghttp.HandlerFunc{
		http.StatusNotFound: func(r *ghttp.Request) {
			library.JsonExit(r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		},
		http.StatusInternalServerError: func(r *ghttp.Request) {
			library.JsonExit(r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		},
		http.StatusBadGateway: func(r *ghttp.Request) {
			library.JsonExit(r, http.StatusBadGateway, http.StatusText(http.StatusBadGateway))
		},
		http.StatusGatewayTimeout: func(r *ghttp.Request) {
			library.JsonExit(r, http.StatusGatewayTimeout, http.StatusText(http.StatusGatewayTimeout))
		},
		http.StatusServiceUnavailable: func(r *ghttp.Request) {
			library.JsonExit(r, http.StatusServiceUnavailable, http.StatusText(http.StatusServiceUnavailable))
		},
	})

	// 绑定登录用户ID
	s.BindHookHandler("/*", ghttp.HOOK_BEFORE_SERVE, func(r *ghttp.Request) {
		if guid, b := library.Auth(r); b {
			if ok := r.Session.GetBool(guid); ok {
				user := r.Session.Get(guid).(map[string]interface{})
				r.SetForm("cid", user["cid"])
			}
		}
	})
}

// rpc客户端初始化
func RpcInit() {
	m := make(map[string]pool.Options)
	RpcClients = gset.New(true)
	RpcClients.Add(OFC_COM_SVR)
	RpcClients.Add(IM_MSG_SVR)

	RpcClients.Iterator(func(v interface{}) bool {
		m[v.(string)] = pool.DefaultOptions
		return true
	})
	once.Do(func() {
		RpcPoolMap = packed.NewRpcPool(m)
	})
}

// session初始化
func SessionInit() {
	s := g.Server()
	_ = s.SetConfigWithMap(g.Map{
		"SessionMaxAge":  time.Minute * 60,
		"SessionStorage": gsession.NewStorageRedis(g.Redis()),
	})
}

// 设置路由白名单
func WhiteParseRoute() {
	WhiteRoute = gset.New(true)
	for k, _ := range UserRequest() {
		WhiteRoute.Add(k.(string))
	}
}

func init() {
	WhiteParseRoute()
	SessionInit()
	RpcInit()
	RouteInit()
}
