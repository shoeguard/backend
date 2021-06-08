package main

import (
	"shoeguard-main-backend/cmd/server/middlewares"
	"shoeguard-main-backend/cmd/server/models"
	"shoeguard-main-backend/cmd/server/routes"
	"shoeguard-main-backend/cmd/server/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := utils.GetDB()
	db.AutoMigrate(models.User{})
	app := fiber.New()
	middlewares.SetupMiddlewares(app)
	routes.SetupRoutes(app)
	app.Listen(":8080")
}
