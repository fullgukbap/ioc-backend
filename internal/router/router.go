package router

import (
	"context"
	"ioc-backend/internal/application/handler"
	"ioc-backend/internal/infra/config"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
)

type Router struct {
	userHandler *handler.UserHandler
	config      *config.Config
	*fiber.App
}

func NewRouter(
	lc fx.Lifecycle,
	config *config.Config,
	userHandler *handler.UserHandler,
	errorHandler *handler.ErrorHandler,
) *Router {
	app := &Router{
		config:      config,
		userHandler: userHandler,
		App: fiber.New(
			fiber.Config{
				ErrorHandler: errorHandler.HandleError(),
			},
		),
	}

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					app.Route()
					if err := app.Start(); err != nil {
						panic(err)
					}
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				return app.Shutdown()
			},
		},
	)

	return nil
}

func (r *Router) Route() {
	api := r.Group("/api")
	{
		user := api.Group("/user")
		{
			r.userHandler.Route(user)
		}
	}
}

func (r *Router) Start() error {
	return r.Listen(r.config.HTTP.Port)
}
