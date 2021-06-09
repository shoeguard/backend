package routes

import (
	"shoeguard-main-backend/cmd/server/controllers"
	"shoeguard-main-backend/cmd/server/middlewares"
	"shoeguard-main-backend/configs"

	_ "shoeguard-main-backend/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	if configs.ENABLE_SWAGGER == "true" {
		app.Get("/swagger/*", swagger.Handler)
		app.Get("/swagger", func(c *fiber.Ctx) error {
			return c.Redirect("./swagger/")
		})
	}

	usersGroup := app.Group("/users")
	usersGroup.Post("/register", controllers.Register)
	reportGroup := app.Group("/report")
	reportGroup.Use(middlewares.BasicAuthMiddleware())
	reportGroup.Post("", controllers.Report)
	reportGroup.Get("", controllers.GetReports)
	reportGroup.Patch("/:id", controllers.AddRecordedAudioURL)
}
