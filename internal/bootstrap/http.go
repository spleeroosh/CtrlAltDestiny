package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spleeroosh/CtrlAltDestiny/internal/config"
	"github.com/spleeroosh/CtrlAltDestiny/internal/pkg/application"
	"github.com/spleeroosh/CtrlAltDestiny/internal/pkg/routerfx"
	serverfx "github.com/spleeroosh/CtrlAltDestiny/internal/pkg/serverfx"
	"go.uber.org/fx"
)

// Создание нового HTTP-сервера
func newHTTPServer(lc fx.Lifecycle, sh fx.Shutdowner, engine *gin.Engine, conf config.Config, router *routerfx.AppRoute) *serverfx2.ServerFX {
	// Установите режим Gin (например, ReleaseMode или DebugMode)
	gin.SetMode(gin.ReleaseMode)
	fmt.Println("HTTP SERVER START")
	// Настройка роутера
	router.SetupRouter(engine)

	// Создание HTTP сервера
	srv := serverfx2.New(
		conf.App.Name,
		serverfx.Handler(engine.Handler()),
		serverfx.Port(conf.App.Port),
	)

	lc.Append(application.ServerHooks(sh, srv))
	fmt.Println("HTTP SERVER END")
	return srv
}
