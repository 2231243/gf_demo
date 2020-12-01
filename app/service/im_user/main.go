/**
* @Authore: lifeifei
* @Date: 2020/11/20 1:47 下午
 */
package im_user

import (
	"context"
	"time"

	"gf-app/grpc/ofc_im_user"
	"gf-app/library"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/frame/g"
)

func StatisUser(req *ofc_im_user.ImUserRegisterStatistics,
	client ofc_im_user.ImUserStatisticsClient) ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := library.TimeLegal(req.StartTime, req.EndTime); err != nil {
		return nil, err
	}
	result, err := client.StatisUser(ctx, req)
	if err != nil {
		return nil, err
	}
	betweenDays := library.GetBetweenDates(req.StartTime, req.EndTime)
	r := gmap.NewListMap(true)
	for _, v := range betweenDays {
		r.Set(v, g.MapAnyAny{
			"date":  v,
			"count": 0,
		})
	}
	return library.DataReass(r, result.List, "date"), nil
}
