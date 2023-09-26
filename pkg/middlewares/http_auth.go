package middlewares

import (
	AuthService "deluze/internal/core/services/auth"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
	token := c.Cookies("_auth")
	if token == "" {
		return c.Status(401).JSON(fiber.Map{
			"message": "invalid token or not present in header",
		})
	}
	token = strings.Replace(token, "Bearer ", "", 1)
	userId, err := AuthService.Validate(token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "invalid token or already expired",
		})
	}
	c.Locals("user", userId)
	return c.Next()
}
