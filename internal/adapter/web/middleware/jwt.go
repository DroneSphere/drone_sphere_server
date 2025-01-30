package middleware

import (
	"drone_sphere_server/pkg/token"
	"github.com/gofiber/fiber/v2"
)

func JWTWare(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if len(authHeader) <= 0 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
			"code": 401,
			"msg":  "unauthorized",
		})
	}

	jwt := authHeader[7:]
	_, err := token.ValidateJWT(jwt, []byte("secret"))
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{
			"code": 401,
			"msg":  "unauthorized",
		})
	}

	return ctx.Next()
}
