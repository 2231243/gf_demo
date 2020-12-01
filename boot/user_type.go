/**
* @Authore: lifeifei
* @Date: 2020/11/19 9:15 上午
 */
package boot

import (
	"gf-app/grpc/ofc_user"
	"github.com/gogf/gf/frame/g"
)

func UserRequest() g.MapAnyAny {
	return g.MapAnyAny{
		// 登录
		"/api/v1/user/login": &ofc_user.LoginRequest{},
		// 注册
		"/api/v1/user/register": &ofc_user.ResgitserRequest{},
		// 获取验证码
		"/api/v1/user/sms": &ofc_user.MessageRequest{},
	}
}
