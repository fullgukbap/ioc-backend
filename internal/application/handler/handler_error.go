package handler

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type ErrorHandler struct {
	logger *zap.Logger
}

func NewErrorHandler(
	logger *zap.Logger,
) *ErrorHandler {
	return &ErrorHandler{
		logger: logger,
	}
}

func (h ErrorHandler) HandleError() func(c fiber.Ctx, err error) error {
	return func(c fiber.Ctx, err error) error {
		h.logger.Error("error", zap.Error(err))
		return fiber.DefaultErrorHandler(c, err)
	}
}
