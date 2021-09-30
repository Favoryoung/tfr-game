package conf

//HotUpdate 是否支持配置热更新(监听配置文件)
const HotUpdate = false

//热更新到配置 需要添加到 setConfFunc
var setConfFunc = []func(){
	setConfApp,
	setConfHttp,
}
