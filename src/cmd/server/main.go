package main

import (
	"shoeguard-main-backend/cmd/server/middlewares"
	"shoeguard-main-backend/cmd/server/models"
	"shoeguard-main-backend/cmd/server/routes"

	"github.com/gofiber/fiber/v2"
)

// @title ShoeGuard Main Backend API
// @version 1.0
// @securityDefinitions.basic BasicAuth
func main() {
	models.MigrateModels()
	app := fiber.New()
	middlewares.SetupMiddlewares(app)
	routes.SetupRoutes(app)
	app.Listen(":8080")
}
