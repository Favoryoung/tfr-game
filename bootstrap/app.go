package bootstrap

import (
	"github.com/gin-gonic/gin"
	"log"
	"tfr-game/conf"
	"tfr-game/controller"
	"tfr-game/helper/er"
)

type Configurator func(bootstrapper *App)

type App struct {
	*gin.Engine
	AppName  string
	AppOwner string

	Handlers *controller.Handlers
}

func NewAPP(handlers *controller.Handlers) *App {
	app := &App{
		Engine:   newRouterClient(),
		AppName:  conf.AppName,
		AppOwner: conf.AppOwner,
		Handlers: handlers,
	}

	app.boot()
	return app
}

func newRouterClient() *gin.Engine {
	if !conf.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()
	return engine
}

func (app *App) Configure(cfgList ...Configurator) {
	for _, cfg := range cfgList {
		cfg(app)
	}
}

func (app *App) boot() {
	// Force log's color
	gin.ForceConsoleColor()
	app.Engine.Use(gin.Logger())
	//log.SetFlags(log.Llongfile | log.LstdFlags)
	app.Engine.Use(er.Recovery())
	app.Configure(baseRoute)
}

func (app *App) Run() {
	if !conf.Debug {
		log.Println(" Listening and serving HTTP on :" + conf.ListenAddr)
	}
	err := app.Engine.Run(":" + conf.ListenAddr)
	if err != nil {
		log.Fatalln("server start failed")
	}
}
