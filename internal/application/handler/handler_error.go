package handler

import (
	"errors"
	"ioc-backend/internal/application/presenter"
	"ioc-backend/internal/infra/exception"

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
		var excep *exception.Exception
		if errors.As(err, &excep) {
			// internal server error
			if excep.Status == fiber.StatusInternalServerError {
				h.logger.Error("error", zap.Error(excep.Err))
				excep.Err = errors.New("hides information because an internal server error occurred.")
			}

			// it's myqsl error
			if exception.IsMySQLError(excep.Type) {
				excep.Err = errors.New("mysql errors are hidden for information security reasons.")
			}

			return c.Status(excep.Status).JSON(&presenter.ErrorResponse{
				Type:    excep.Type,
				Status:  excep.Status,
				Message: excep.Message,
				Data:    excep.Data,
				Detail:  excep.Err.Error(),
			})
		}

		// it's fiber error
		code := fiber.StatusInternalServerError
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			code = fiberErr.Code
		}

		return c.Status(code).JSON(&presenter.ErrorResponse{
			Type:    exception.ErrWebServerInternal,
			Message: fiberErr.Message,
			Status:  code,
			// Detail:  err,
		})
	}
}
