package route

import (
	"github.com/gin-gonic/gin"
	"tfr-game/bootstrap"
	"tfr-game/helper"
	"tfr-game/middleware"
)

func Configure(app *bootstrap.App) {
	user := app.Engine.Group("v2/game/api/").Use(middleware.User(app.Handlers.User.UserRoomDao))
	{
		user.GET("ping", func(c *gin.Context) {
			helper.Msg(c, "pong")
		})

		// 查询所有用户
		user.GET("users", app.Handlers.User.Users)
		// 个人用户和房间信息
		user.GET("homepage", app.Handlers.User.Mine)
	}
}
