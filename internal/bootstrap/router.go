package bootstrap

import (
	"CtrlAltDestiny/internal/config"
	"CtrlAltDestiny/internal/pkg/routerfx"
	"github.com/gin-gonic/gin"
)

func newHTTPRouter(conf config.Config) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	return routerfx.New(
		conf.App.Name,
	)
}
