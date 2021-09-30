package conf

import (
	"time"
)

const (
	SysTimeForm      = "2006-01-02 15:04:05" // 时间转换格式
	SysTimeFormShort = "2006-01-02"          // 日期转换格式
)

var SysTimeLocation, _ = time.LoadLocation("Asia/Shanghai")

var (
	//Debug 是否开启 debug
	Debug bool
	//AppName 应用名称
	AppName string
	//AppOwner 程序作者
	AppOwner string
)

//setConfApp  动态配置方法 setConf + "文件名"
func setConfApp() {
	Debug = BoolOrFalse("app.debug")
	AppName = StringOr("app.name", "Having Fun")
	AppOwner = StringOr("app.owner", "乐正逸")
}
