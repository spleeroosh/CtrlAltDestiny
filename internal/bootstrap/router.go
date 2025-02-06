package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/spleeroosh/CtrlAltDestiny/internal/config"
	"github.com/spleeroosh/CtrlAltDestiny/internal/pkg/routerfx"
)

func newHTTPRouter(conf config.Config) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	return routerfx.New(
		conf.App.Name,
	)
}
