//+build wireinject

package wire

import (
	"github.com/google/wire"
	"tfr-game/bootstrap"
	"tfr-game/controller"
	"tfr-game/dao/impl"
)

// 正式实现的依赖注入初始化
func InitializeApp() *bootstrap.App {
	wire.Build(
		impl.NewUserRoomDao,

		controller.NewUserHandler,
		controller.NewHandlers,

		bootstrap.NewAPP,
	)
	return &bootstrap.App{}
}

// mock数据的依赖注入初始化
//func InitializeMockApp() *bootstrap.App {
//	wire.Build(
//		mock.NewGirlDao,
//
//		controller.NewDemoHandler,
//		controller.NewHandlers,
//
//		bootstrap.NewAPP,
//	)
//	return &bootstrap.App{}
//}
