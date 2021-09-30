package main

import (
	"tfr-game/route"
	"tfr-game/wire"
)

func main() {
	app := wire.InitializeApp()
	//配置路由
	app.Configure(route.Configure)

	app.Run()
}
