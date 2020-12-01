package user

import (
	"gf-app/app/api"
	user2 "gf-app/app/service/user"
	"gf-app/boot"
	"gf-app/grpc/ofc_user"
	"gf-app/inter"
	"gf-app/library"

	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/net/ghttp"
)

type User struct {
	api.Control
	conn inter.RpcInter
	RpcClient ofc_user.UserClient
}

// 初始化rpc客户端链接
func (u *User) Init(r *ghttp.Request) {
	u.conn, _ = boot.RpcPoolMap.Conn(boot.OFC_COM_SVR)
	u.RpcClient = ofc_user.NewUserClient(u.conn.GetRpc())
}

// 释放资源
func (u *User) Shut(r *ghttp.Request) {
	u.conn.Close()
}
/**
 * @Author lifeifei
 * @Date 1:14 下午 2020/11/24
 * @Summary 用户模块
 * @Tags 登出
 * @Produce json
 * @Router  /api/v1/user/logout
 * @Return {object} response.JsonResponse "执行结果"
 **/
func (u *User) Logout(r *ghttp.Request) {
	if guid, b := library.Auth(r); b {
		_ = r.Session.Remove(guid)
	}
	library.JsonExit(r, 0, "请求成功")
}

/**
 * @Author lff
 * @Date 5:41 下午 2020/11/17
 * @Summary 获取登录用户信息
 * @Tags 用户服务
 * @Produce json
 * @Router /api/v1/user [GET]
 * @Return   {object} response.JsonResponse "执行结果"
 **/
func (u *User) Index(r *ghttp.Request) {
	if guid, b := library.Auth(r); b {
		if ok := r.Session.GetBool(guid); !ok {
			u.Fail(r, gerror.New("登录已失效"), nil)
		}
		u.Success(r, r.Session.Get(guid))
	} else {
		u.Fail(r, gerror.New("登录已失效"), nil)
	}
}

/**
 * @Author lifeifei
 * @Date 5:42 下午 2020/11/17
 * @Summary 用户注册接口
 * @tags 用户服务
 * @produce json
 * @router /api/v1/user/register [POST]
 * @param  mobile  formData string  true "用户账号"
 * @param  password  formData string  true "用户密码"
 * @param  Code  formData string true "验证码"
 * @return 0 {object} response.JsonResponse "执行结果"
 **/

func (u *User) Register(r *ghttp.Request) {
	var request *ofc_user.ResgitserRequest
	if err := u.Parse(r, &request); err != nil {
		u.Fail(r, err, nil, request)
	}

	result, err := user2.SignUp(r, request, u.RpcClient)
	if err != nil {
		u.Fail(r, err, nil, request)
	}
	u.Success(r, result, request)
}

/**
 * @Author lifeifei
 * @Date 6:50 下午 2020/11/17
 * @Summary 用户登录接口
 * @Produce json
 * @Tags 用户服务
 * @Router  /api/v1/user/login [POST]
 * @Param mobile  formData string  true "用户账号"
 * @Param password  formData string  true "用户密码"
 * @Return 0 {object} response.JsonResponse "执行结果"
 **/

func (u *User) Login(r *ghttp.Request) {
	var request *ofc_user.LoginRequest

	if err := u.Parse(r, &request); err != nil {
		u.Fail(r, err, nil, request)
	}
	result, err := user2.SignIn(r, request, u.RpcClient)
	if err != nil {
		u.Fail(r, err, nil, request)
	}
	u.Success(r, result, request)
}

/**
 * @Author lifeifei
 * @Date 6:51 下午 2020/11/17
 * @Summary 用户短信接口
 * @Tags 用户服务
 * @Produce json
 * @Router /api/v1/user/sms [POST]
 * @Param  mobile  formData string  true "用户账号"
 * @Return  0 {object} response.JsonResponse "执行结果"
 **/
func (u *User) Sms(r *ghttp.Request) {
	var request *ofc_user.MessageRequest
	if err := u.Parse(r, &request); err != nil {
		u.Fail(r, err, nil, request)
	}
	if err := user2.Sms(request, u.RpcClient); err != nil {
		u.Fail(r, err, nil, request)
	}
	u.Success(r, nil, request)
}

func NewUser() *User {
	return &User{}
}
