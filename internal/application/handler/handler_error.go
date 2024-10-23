package handler

import "github.com/gofiber/fiber/v3"

type ErrorHandler struct {
}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}

func (h ErrorHandler) HandleError() func(c fiber.Ctx, err error) error {
	return func(c fiber.Ctx, err error) error {
		return fiber.DefaultErrorHandler(c, err)
	}
}
