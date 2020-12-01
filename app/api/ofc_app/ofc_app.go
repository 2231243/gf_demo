package ofc_app

import (
	"gf-app/app/api"
	service "gf-app/app/service/ofc_app"
	"gf-app/boot"
	"gf-app/grpc/ofc_app"
	"gf-app/inter"

	"github.com/gogf/gf/net/ghttp"
)

type OfcApp struct {
	api.Control
	appClient ofc_app.OfcAppServiceClient
	conn inter.RpcInter
}

func (u *OfcApp) Init(r *ghttp.Request) {
	u.conn, _ = boot.RpcPoolMap.Conn(boot.OFC_COM_SVR)
	u.appClient = ofc_app.NewOfcAppServiceClient(u.conn.GetRpc())
}

func (u *OfcApp) Shut(r *ghttp.Request) {
	u.conn.Close()
}

//添加应用
func (u *OfcApp) Create(r *ghttp.Request) {
	var request *ofc_app.AppRequest
	if err := u.Parse(r, &request); err != nil {
		u.Fail(r, err, nil, request)
	}
	//request := boot.M.Get(r.URL.Path).(*ofc_app.AppRequest)
	reply, err := service.AppCreate(request, u.appClient)
	if err != nil {
		u.Fail(r, err, nil, request)
	}
	u.Success(r, reply, nil)
}

//修改应用
func (u *OfcApp) Update(r *ghttp.Request) {

	var request *ofc_app.AppRequest
	if err := u.Parse(r, &request); err != nil {
		u.Fail(r, err, nil, request)
	}
	//request := boot.M.Get(r.URL.Path).(*ofc_app.AppRequest)
	if err := service.AppUpdate(request, u.appClient); err != nil {
		u.Fail(r, err, nil, request)
	}
	u.Success(r, nil)
}

//获取单个
func (u *OfcApp) One(r *ghttp.Request) {
	var request *ofc_app.AppGetRequest
	if err := u.Parse(r, &request); err != nil {
		u.Fail(r, err, nil, request)
	}
	//request := boot.M.Get(r.URL.Path).(*ofc_app.AppGetRequest)
	reply, err := service.AppOne(request, u.appClient)
	if err != nil {
		u.Fail(r, err, nil, request)
	}
	u.Success(r, reply, nil)
}

//获取列表
func (u *OfcApp) List(r *ghttp.Request) {
	var request *ofc_app.AppGetRequest
	if err := u.Parse(r, &request); err != nil {
		u.Fail(r, err, nil, request)
	}
	// boot.M.RLockFunc(func(m map[interface{}]interface{}) {
	// 	request = m[r.URL.Path].(*ofc_app.AppGetRequest)
	// })
	//request := boot.M.Get(r.URL.Path).(*ofc_app.AppGetRequest)
	data, err := service.AppList(request, u.appClient)
	if err != nil {
		u.Fail(r, err, nil, request)
	}
	u.Success(r, data, nil)
}

//刷新密钥
func (u *OfcApp) Secret(r *ghttp.Request) {
	var request *ofc_app.AppKeyRequest
	if err := u.Parse(r, &request); err != nil {
		u.Fail(r, err, nil, request)
	}
	//request := boot.M.Get(r.URL.Path).(*ofc_app.AppKeyRequest)
	data, err := service.AppSecret(request, u.appClient)
	if err != nil {
		u.Fail(r, err, nil, request)
	}
	u.Success(r, data, nil)
}

//消息抄送
func (u *OfcApp) Message(r *ghttp.Request) {
	var request *ofc_app.MessageUrlRequest
	if err := u.Parse(r, &request); err != nil {
		u.Fail(r, err, nil, request)
	}
	//request := boot.M.Get(r.URL.Path).(*ofc_app.MessageUrlRequest)
	if err := service.AppMessage(request, u.appClient); err != nil {
		u.Fail(r, err, nil, request)
	}
	u.Success(r, nil)
}

//标识管理
func (u *OfcApp) Ident(r *ghttp.Request) {
	var request *ofc_app.IdentifierRequest
	if err := u.Parse(r, &request); err != nil {
		u.Fail(r, err, nil, request)
	}
	//request := boot.M.Get(r.URL.Path).(*ofc_app.IdentifierRequest)
	if err := service.AppIdentifier(request, u.appClient); err != nil {
		u.Fail(r, err, request)
	}
	u.Success(r, nil)
}

//文案添加
func (u *OfcApp) Copywriting(r *ghttp.Request) {
	var request *ofc_app.CopywritingRequest
	if err := u.Parse(r, &request); err != nil {
		u.Fail(r, err, nil, request)
	}
	//request := boot.M.Get(r.URL.Path).(*ofc_app.CopywritingRequest)
	if err := service.CopyWritingCreate(request, u.appClient); err != nil {
		u.Fail(r, err, nil, request)
	}
	u.Success(r, nil)
}

//文案修改
func (u *OfcApp) CopywritingSave(r *ghttp.Request) {
	var request *ofc_app.CopywritingRequest
	if err := u.Parse(r, &request); err != nil {
		u.Fail(r, err, nil, request)
	}
	//request := boot.M.Get(r.URL.Path).(*ofc_app.CopywritingRequest)
	if err := service.CopyWritingUpdate(request, u.appClient); err != nil {
		u.Fail(r, err, request)
	}
	u.Success(r, nil)
}

//获取单个文案
func (u *OfcApp) CopywritingOne(r *ghttp.Request) {
	var request *ofc_app.CopywritingGetRequest
	if err := u.Parse(r, &request); err != nil {
		u.Fail(r, err, nil, request)
	}
	//request := boot.M.Get(r.URL.Path).(*ofc_app.CopywritingGetRequest)
	res, err := service.CopywritingGetOne(request, u.appClient)
	if err != nil {
		u.Fail(r, err, nil, request)
	}
	u.Success(r, res, nil)
}

//获取文案列表
func (u *OfcApp) CopywritingList(r *ghttp.Request) {
	var request *ofc_app.CopywritingGetRequest
	if err := u.Parse(r, &request); err != nil {
		u.Fail(r, err, nil, request)
	}
	//request := boot.M.Get(r.URL.Path).(*ofc_app.CopywritingGetRequest)
	data, err := service.CopywritingGetList(request, u.appClient)
	if err != nil {
		u.Fail(r, err, nil, request)
	}
	u.Success(r, data, nil)
}

func NewOfcApp() *OfcApp {
	app := &OfcApp{}
	return app
}
