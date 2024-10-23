package handler

import (
	"ioc-backend/internal/application/port"
	"ioc-backend/internal/application/presenter"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type UserHandler struct {
	logger      *zap.Logger
	userService port.UserService
}

func NewUserHandler(
	logger *zap.Logger,
	userService port.UserService,
) *UserHandler {
	return &UserHandler{
		logger:      logger,
		userService: userService,
	}
}

func (h UserHandler) Route(r fiber.Router) {
	r.Post("/", h.Create)
	r.Get("/:id", h.Get)
	r.Put("/:id", h.Update)
	r.Delete("/:id", h.Delete)
}

func (h UserHandler) Create(c fiber.Ctx) error {
	h.logger.Info("processing user creation request", zap.String("ip", c.IP()))

	req := new(presenter.UserCreateRequest)
	if err := c.Bind().JSON(req); err != nil {
		return err
	}

	createdUser, err := h.userService.Create(c.Context(), req.ToDomain())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewUserCreateResponse(createdUser))
}

func (h UserHandler) Get(c fiber.Ctx) error {
	h.logger.Info("processing user get request", zap.String("ip", c.IP()))

	req := new(presenter.UserGetReqeust)
	if err := c.Bind().Query(req); err != nil {
		return err
	}

	user, err := h.userService.Get(c.Context(), req.ID)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewuserGetResponse(user))
}

func (h UserHandler) Update(c fiber.Ctx) error {
	h.logger.Info("processing user update request", zap.String("ip", c.IP()))

	req := new(presenter.UserUpdateRequest)
	if err := c.Bind().JSON(req); err != nil {
		return err
	}

	if err := c.Bind().Query(req); err != nil {
		return err
	}

	updatedUser, err := h.userService.Update(c.Context(), req.NewUserUpdateRequest())
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(presenter.NewUserUpdateResponse(updatedUser))
}

func (h UserHandler) Delete(c fiber.Ctx) error {
	h.logger.Info("processing user delete request", zap.String("ip", c.IP()))

	req := new(presenter.UserDeleteRequest)
	if err := c.Bind().Query(req); err != nil {
		return err
	}

	if err := h.userService.Delete(c.Context(), req.ID); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
