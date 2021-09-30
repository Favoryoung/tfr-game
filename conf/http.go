package conf

var (
	//ListenAddr http 监听端口
	ListenAddr string
)

//setConfHttp 动态配置方法
func setConfHttp() {
	ListenAddr = StringOr("web_listen_addr", "80")
}
