package common

import "github.com/gofiber/fiber/v2"

func HttpErrorResponse(app *fiber.Ctx, status int, message string) error {
	return app.Status(status).JSON(fiber.Map{
		"errors": message,
	})
}
