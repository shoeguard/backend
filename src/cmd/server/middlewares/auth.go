package middlewares

import (
	"shoeguard-main-backend/cmd/server/models"

	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

var BasicAuth = basicauth.New(basicauth.Config{
	Authorizer: func(phoneNumber string, password string) bool {
		user := models.User{}
		user.SetUser(phoneNumber)
		return user.IsPasswordCorrect(password)
	},
})
