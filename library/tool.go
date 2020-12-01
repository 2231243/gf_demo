package library

import (
	"fmt"
	"github.com/gogf/gf/util/guid"
	"strings"
	"time"

	"gf-app/grpc/ofc_user"

	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

// HTTP请求参数验证是否合法
func Parse(r *ghttp.Request, data interface{}) error {
	if err := r.Parse(data); err != nil {
		if v, ok := err.(*gvalid.Error); ok {
			return gerror.New(v.FirstString())
		}
		return err
	}
	return nil
}

// 检测日期格式是否正确 format 2020-11/2020-11-11
func TimeLegal(start, end string) error {
	dayPat := `^\d{4}-[01]{1}\d{1}$`
	monPat := `(((^((1[8-9]\d{2})|([2-9]\d{3}))([-\/\._])(10|12|0?[13578])([-\/\._])(3[01]|[12][0-9]|0?[1-9]))|(^((1[8-9]\d{2})|([2-9]\d{3}))([-\/\._])(11|0?[469])([-\/\._])(30|[12][0-9]|0?[1-9]))|(^((1[8-9]\d{2})|([2-9]\d{3}))([-\/\._])(0?2)([-\/\._])(2[0-8]|1[0-9]|0?[1-9]))|(^([2468][048]00)([-\/\._])(0?2)([-\/\._])(29))|(^([3579][26]00)([-\/\._])(0?2)([-\/\._])(29))|(^([1][89][0][48])([-\/\._])(0?2)([-\/\._])(29))|(^([2-9][0-9][0][48])([-\/\._])(0?2)([-\/\._])(29))|(^([1][89][2468][048])([-\/\._])(0?2)([-\/\._])(29))|(^([2-9][0-9][2468][048])([-\/\._])(0?2)([-\/\._])(29))|(^([1][89][13579][26])([-\/\._])(0?2)([-\/\._])(29))|(^([2-9][0-9][13579][26])([-\/\._])(0?2)([-\/\._])(29)))((\s+(0?[1-9]|1[012])(:[0-5]\d){0,2}(\s[AP]M))?$|(\s+([01]\d|2[0-3])(:[0-5]\d){0,2})?$))`
	if gregex.IsMatchString(dayPat, start) && gregex.IsMatchString(dayPat, end) {
		if err := TimeBefore(start+"-01", end+"-01"); err != nil {
			return err
		}
		return nil
	}
	if gregex.IsMatchString(monPat, start) && gregex.IsMatchString(monPat, end) {
		if err := TimeBefore(start, end); err != nil {
			return err
		}
		return nil
	}
	return gerror.New("时间格式不合法")
}

// 判断两个时间大小 2020-11-11
func TimeBefore(start, end string) error {
	start_time := gtime.New(start)
	end_time := gtime.New(end)

	if start_time.After(gtime.New(time.Now())) || start_time.After(end_time) {
		return gerror.New("开始时间大于当前时间或结束时间")
	}
	return nil
}



// 获取两个天数之间的相差集合
func GetBetweenDates(sdate, edate string) []string {
	if sdate == edate {
		return []string{sdate}
	}
	d := []string{}
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	if len(sdate) == 7 && len(sdate) == 7 {
		var t []string
		for _, v := range d {
			t = append(t, v[0:7])

		}
		return RemoveRepByLoop(t)
	}
	return RemoveRepByLoop(d)
}

// 去除重复的集合
func RemoveRepByLoop(slc []string) []string {
	result := []string{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}

func DataReass(r *gmap.ListMap, result interface{}, key string) []interface{} {
	socre := gmap.NewListMap(true)
	gconv.Maps(result)
	for _, v := range gconv.Maps(result) {
		socre.Set(v[key], v)
	}
	r.Merge(socre)
	var m []interface{}
	r.Iterator(func(k interface{}, v interface{}) bool {
		m = append(m, v)
		return true
	})
	return m
}
func CreateSessionId(r *ghttp.Request) string {
	var (
		address = r.RemoteAddr
		header  = fmt.Sprintf("%v", r.Header)
	)
	return guid.S([]byte(address), []byte(header))
}

// 登录注册返回token
func LoginAndRegister(r *ghttp.Request, user *ofc_user.LoginAndRegisterReply) (*gmap.Map, error) {
	m := gmap.New()
	gid := CreateSessionId(r)
	token, err := Jwt(gid)
	if err != nil {
		return m, err
	}

	m.Sets(g.MapAnyAny{
		"token": token,
	})
	r.Session.Id()
	_ = r.Session.Set(gid, map[string]interface{}{
		"cid" : user.Cid,
		"mobile" : user.Mobile,
		"create_time" : user.CreateTime,
	})
	return m, nil
}

// 根据token获取登录用户信息
func Auth(r *ghttp.Request) ( string, bool) {
	auth := r.Header.Get("Authorization")
	if auth == "" || !strings.HasPrefix(auth, "Bearer") {
		return "", false
	}
	auth = auth[7:]
	token, c, err := ParseToken(auth)
	if err != nil || !token.Valid {
		return "", false
	}

	return c.Guid, true
}
