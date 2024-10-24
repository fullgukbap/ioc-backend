package router

import (
	"context"
	"ioc-backend/internal/application/handler"
	"ioc-backend/internal/infra/config"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Router struct {
	userHandler *handler.UserHandler
	config      *config.Config
	*fiber.App
}

func NewRouter(
	lc fx.Lifecycle,
	config *config.Config,
	logger *zap.Logger,
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
				logger.Info("starting server", zap.String("port", config.HTTP.Port))
				go func() {
					app.Route()
					if err := app.Start(); err != nil {
						panic(err)
					}
				}()

				return nil
			},
			OnStop: func(ctx context.Context) error {
				logger.Info("stopping server")
				// TODO: graceful shutdown
				// router.Shutdown()
				return nil
			},
		},
	)

	return app
}

func (r *Router) Route() {
	v1 := r.Group("/v1")
	{
		api := v1.Group("/api")
		{
			users := api.Group("/users")
			{
				r.userHandler.Route(users)
			}
		}
	}
}

func (r *Router) Start() error {
	return r.Listen(r.config.HTTP.Port)
}
