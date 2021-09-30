package helper

import (
	"errors"
	"fmt"
	"strconv"
	"tfr-game/conf"
	"time"
)

var timeTemplates = []string{
	"2006-01-02 15:04:05", //常规类型
	"2006/01/02 15:04:05",
	"2006-01-02",
	"2006/01/02",
	"15:04:05",
	time.RFC3339,
	"2006-01-02 15:04:05Z07:00",
	"2006-01-02 15:04:05 Z07:00",
	"2006-01-02 15:04:05 -0700 MST",
}

// 当前时间的时间戳
func NowUnix() int {
	return int(time.Now().Unix())
}

// 当前时间
func NowTime() time.Time {
	return time.Now().In(conf.SysTimeLocation)
}

// FormatFromUnixTime 将 unix 时间戳格式化为 yyyymmdd HH:II:SS 格式字符串
func FormatFromUnixTime(t int64) string {
	if t > 0 {
		return time.Unix(t, 0).Format(conf.SysTimeForm)
	} else {
		return NowTime().Format(conf.SysTimeForm)
	}
}

// 将字符串转成日期 2006-01-02 格式
func StrShortToTime(s string) time.Time {
	t, _ := time.Parse(conf.SysTimeFormShort, s)
	return t.In(conf.SysTimeLocation)
}

// ParseTime 将字符串转成时间 2006-01-02 15:04:05 格式
func ParseTime(s string) (time.Time, error) {
	for _, timeTemplate := range timeTemplates {
		t, err := time.ParseInLocation(timeTemplate, s, conf.SysTimeLocation)
		if nil == err {
			return t, nil
		}
	}

	return time.Time{}, errors.New("parseTime string to time.Time error, string = " + s)
}

// FriendTime 时间转化为字符串
func FriendTime(t time.Time) string {
	return t.In(conf.SysTimeLocation).Format(conf.SysTimeForm)
}

func GtTime(t1 time.Time, t2 time.Time) bool {
	return t1.UnixNano() > t2.UnixNano()
}

func TodayInt() int {
	y, m, d := time.Now().Date()
	strDay := fmt.Sprintf("%d%02d%02d", y, m, d)
	day, _ := strconv.Atoi(strDay)
	return day
}
