package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupMiddlewares(app *fiber.App) {
	app.Use(logger.New())
}
