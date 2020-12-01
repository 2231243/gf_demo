package message

import (
	"gf-app/app/api"
	msg "gf-app/app/service/message"
	"github.com/gogf/gf/net/ghttp"
)

type MessageAPI struct {
	api.Control
}



// @summary 获取应用消息的数据统计(天)
// @tags 	消息服务
// @param 	appkey     query string true  "应用APPKEY"
// @param   start_date query string false "开始日期, 日期格式:YYYY-MM-DD"
// @param   end_date   query string false "结束日期, 日期格式:YYYY-MM-DD"
// @param   accid      query string false "指定的用户ID"
// @router  /api/v1/message/getAppDailyStats [GET]
// @success 0 {object} response.JsonResponse "执行结果"
func (this *MessageAPI) GetAppDailyStats(r *ghttp.Request) {
	appkey := r.URL.Query().Get("appkey")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")
	accid := r.URL.Query().Get("accid")
	msgSvc := msg.NewMessageService()
	results, _ := msgSvc.GetAppDailyMessageStats(r.Context(), appkey, startDate, endDate, accid)
	this.Success(r, results)

}
func NewMessage() *MessageAPI {
	mess := &MessageAPI{}
	return mess
}
