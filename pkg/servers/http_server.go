package servers

import (
	"deluze/pkg/routes"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func StartupApi(port string) {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	app.Use(logger.New(logger.Config{
		Format:     "🆕 ${yellow}Latency:${magenta}${latency} ${yellow}Time:${magenta}${time} ${yellow}Status:${status} ${yellow}Path:${magenta}${path} \n",
		TimeFormat: "15:04:05 02/01/2006",
		TimeZone:   "America/Sao_Paulo",
	}))
	healthRoutes(app)
	routes.SetupAuthRoutes(app)
	app.Listen(port)
}

func healthRoutes(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "success",
		})
	})
	app.Get("/status", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "success",
		})
	})
}
