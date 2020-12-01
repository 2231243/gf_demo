package message

import (
	"context"
	"gf-app/boot"
	"gf-app/grpc/im_msg"
)

type ServiceInterface interface {
	GetAppDailyMessageStats(
		ctx context.Context, appkey string, startDate string, endDate string, accid string) (interface{}, error)
}

type service struct {
}

func (this *service) GetAppDailyMessageStats(
	ctx context.Context,
	appkey string,
	startDate string,
	endDate string,
	accid string) (interface{}, error) {
	conn, _ := boot.RpcPoolMap.Conn(boot.IM_MSG_SVR)
	client := im_msg.NewIMMessageServiceClient(conn.GetRpc())
	res, err := client.GetAppDailyMessageStats(ctx, &im_msg.GetAppDailyMessageStatsRequest{
		AppKey:    appkey,
		StartTime: startDate,
		EndTime:   endDate,
		Accid:     accid,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewMessageService() ServiceInterface {
	return &service{}
}
