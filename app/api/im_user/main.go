/**
* @Authore: lifeifei
* @Date: 2020/11/20 9:58 上午
 */
package im_user

import (
	"gf-app/app/api"
	"gf-app/app/service/im_user"
	"gf-app/boot"
	"gf-app/grpc/ofc_im_user"
	"gf-app/inter"
	"github.com/gogf/gf/os/gtime"

	"github.com/gogf/gf/net/ghttp"
)

type ImUser struct {
	api.Control
	RpcClient ofc_im_user.ImUserStatisticsClient
	conn inter.RpcInter
}

func (im *ImUser) Init(r *ghttp.Request) {
	im.conn, _ = boot.RpcPoolMap.Conn(boot.OFC_COM_SVR)
	im.RpcClient = ofc_im_user.NewImUserStatisticsClient(im.conn.GetRpc())
}


func (im *ImUser) Shut(r *ghttp.Request) {
	im.conn.Close()
}
/**
 * @Author lifeifei
 * @Date 11:58 上午 2020/11/20
 * @Summary 用户统计模块
 * @Tags 按照天/月统计IM用户注册数
 * @Produce json
 * @Router  /api/v1/imUser/getRegisterUsers
 * @Param start_time  formData string  false "开始天数"
 * @Param end_time formDate string false "结束天数"
 * @Param appkey formData string false "应用key"
 * @Return  0 {object} response.JsonResponse "执行结果"
 **/
func (im *ImUser) GetRegisterUsers(r *ghttp.Request) {
	var request *ofc_im_user.ImUserRegisterStatistics
	if err := im.Parse(r, &request); err != nil {
		im.Fail(r, err, nil, request)
	}

	request.StartTime = r.GetQuery("start_time", gtime.Now().AddDate(0, 0, -7).Format("Y-m-d")).(string)
	request.EndTime = r.GetQuery("end_time", gtime.Now().Format("Y-m-d")).(string)
	result, err := im_user.StatisUser(request, im.RpcClient)
	if err != nil {
		im.Fail(r, err, nil, request)
	}
	im.Success(r, result)
}

func NewImUser() *ImUser {
	imUser := &ImUser{}
	return imUser
}
