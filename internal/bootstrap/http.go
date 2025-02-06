package bootstrap

import (
	"CtrlAltDestiny/internal/config"
	"CtrlAltDestiny/internal/pkg/application"
	"CtrlAltDestiny/internal/pkg/routerfx"
	serverfx "CtrlAltDestiny/internal/pkg/serverfx"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// Создание нового HTTP-сервера
func newHTTPServer(lc fx.Lifecycle, sh fx.Shutdowner, engine *gin.Engine, conf config.Config, router *routerfx.AppRoute) *serverfx.ServerFX {
	// Установите режим Gin (например, ReleaseMode или DebugMode)
	gin.SetMode(gin.ReleaseMode)
	fmt.Println("HTTP SERVER START")
	// Настройка роутера
	router.SetupRouter(engine)

	// Создание HTTP сервера
	srv := serverfx.New(
		conf.App.Name,
		serverfx.Handler(engine.Handler()),
		serverfx.Port(conf.App.Port),
	)

	lc.Append(application.ServerHooks(sh, srv))
	fmt.Println("HTTP SERVER END")
	return srv
}
