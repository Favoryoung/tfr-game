package bootstrap

import (
	"net/http"
)

//默认路由
func baseRoute(app *App) {
	app.Engine.StaticFile("/favicon.ico", "./web/img/favicon.ico")

	// 前端web
	app.Engine.StaticFS("/v2/game/front", http.Dir("web"))
}
