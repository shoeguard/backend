package routes

import (
	"shoeguard-main-backend/cmd/server/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	usersGroup := app.Group("/users")
	usersGroup.Post("/register", controllers.Register)
}
