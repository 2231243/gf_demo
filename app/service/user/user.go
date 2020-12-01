package user

import (
	"context"
	"github.com/gogf/gf/net/ghttp"
	"time"

	"gf-app/grpc/ofc_user"
	"gf-app/library"

	"github.com/gogf/gf/container/gmap"
)

// 用户注册
func SignUp(r *ghttp.Request, req *ofc_user.ResgitserRequest, client ofc_user.UserClient) (*gmap.Map, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user, err := client.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return library.LoginAndRegister(r, user)
}

// 用户登录
func SignIn(r *ghttp.Request, req *ofc_user.LoginRequest, client ofc_user.UserClient) (*gmap.Map, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	user, err := client.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return library.LoginAndRegister(r, user)
}

//发送验证码
func Sms(req *ofc_user.MessageRequest, client ofc_user.UserClient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := client.SendMessage(ctx, req)
	return err
}
