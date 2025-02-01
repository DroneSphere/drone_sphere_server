package web

import (
	"drone_sphere_server/internal/adapter"
	"drone_sphere_server/internal/adapter/web/middleware"
	user_app "drone_sphere_server/internal/domain/user/app"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// registerUserRoutes 注册用户相关的路由。
func RegisterUserRoutes(router fiber.Router, a *user_app.Application) {
	validate := validator.New()

	router.Post("/register", func(c *fiber.Ctx) error {
		command := new(user_app.RegisterCommand)
		if err := c.BodyParser(&command); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(adapter.Failed(-1, err.Error()))
		}

		if err := validate.Struct(command); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(adapter.Failed(-1, err.Error()))
		}

		result, err := a.Register(c.Context(), *command)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(adapter.Failed(-1, err.Error()))
		}
		return c.JSON(adapter.Success(result))
	})

	router.Post("/login", func(c *fiber.Ctx) error {
		command := new(user_app.LoginCommand)
		if err := c.BodyParser(&command); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(adapter.Failed(-1, err.Error()))
		}

		if err := validate.Struct(command); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(adapter.Failed(-1, err.Error()))
		}

		result, err := a.Login(c.Context(), *command)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(adapter.Failed(-1, err.Error()))
		}
		return c.JSON(adapter.Success(result))
	})

	authed := router.Use(middleware.JWTWare)

	authed.Get("/status", func(c *fiber.Ctx) error {
		result, err := a.GetUserStatus(c.Context(), 0)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(adapter.Failed(-1, err.Error()))
		}
		return c.JSON(adapter.Success(result))
	})
}
