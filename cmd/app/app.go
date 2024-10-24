package main

import (
	"ioc-backend/internal/application/handler"
	"ioc-backend/internal/application/port"
	"ioc-backend/internal/application/repository"
	"ioc-backend/internal/application/service"
	"ioc-backend/internal/infra/config"
	"ioc-backend/internal/infra/storage/mysql"
	"ioc-backend/internal/router"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Supply(
			".env.toml",
		),
		fx.Provide(

			router.NewRouter,
			config.NewConfig,
			zap.NewProduction,
			mysql.NewMysql,

			// handlers
			handler.NewErrorHandler,
			handler.NewUserHandler,

			// services
			fx.Annotate(
				service.NewUserService,
				fx.As(new(port.UserService)),
			),

			// repositories
			fx.Annotate(
				repository.NewUserRepository,
				fx.As(new(port.UserRepository)),
			),
		),
		fx.Invoke(func(r *router.Router) {}),
	).Run()
}
