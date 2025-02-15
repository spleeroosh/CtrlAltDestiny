package bootstrap

import (
	"CtrlAltDestiny/internal/config"
	"CtrlAltDestiny/internal/pkg/application"
	"CtrlAltDestiny/internal/pkg/routerfx"
	"CtrlAltDestiny/internal/pkg/serverfx"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"

	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(providers...),
		fx.Invoke(registerHooks),
	)
}

var providers = []any{
	newHTTPRouter,
	config.GetConfig,
	application.GetBuildVersion,
	newLogger,
	newHTTPServer,
	newPostgresClient,
	//// Регистрация роутов
	//fx.Annotate(wsapi.NewRoutes, fx.As(new(routerfx.Provider)), fx.ResultTags(`group:"providers"`)),
	// Коллектор роутов
	fx.Annotate(routerfx.NewRouter, fx.ParamTags(`group:"providers"`)),
}

func registerHooks(pool *pgxpool.Pool, server *serverfx.ServerFX) {
	fmt.Println("Postgres pool successfully initialized")
}
