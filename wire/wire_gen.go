// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package wire

import (
	"tfr-game/bootstrap"
	"tfr-game/controller"
	"tfr-game/dao/impl"
)

// Injectors from wire.go:

// 正式实现的依赖注入初始化
func InitializeApp() *bootstrap.App {
	userRoomDao := impl.NewUserRoomDao()
	userHandler := controller.NewUserHandler(userRoomDao)
	handlers := controller.NewHandlers(userHandler)
	app := bootstrap.NewAPP(handlers)
	return app
}