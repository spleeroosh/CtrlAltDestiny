package bootstrap

import (
	"CtrlAltDestiny/internal/config"
	"CtrlAltDestiny/internal/pkg/application"
	"CtrlAltDestiny/internal/pkg/logger"
)

func newLogger(conf config.Config, buildVersion application.BuildVersion) log.Logger {
	return log.NewLogger(
		conf.App.Name,
		log.WithEnv(conf.App.Environment),
		log.WithLevel(log.Level(conf.App.LogLevel)),
		log.WithBuildCommit(buildVersion.Commit),
		log.WithBuildTime(buildVersion.Time),
		log.WithPrettify(conf.App.PrettyLogs),
		log.WithOverrideStdLogOut(true),
	)
}
