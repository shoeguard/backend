package controllers

import (
	"shoeguard-main-backend/cmd/server/forms"
	"shoeguard-main-backend/cmd/server/models"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	form := forms.UserRegistrationForm{}

	// parsing body
	if err := c.BodyParser(&form); err != nil {
		return c.Status(400).
			JSON(map[string]interface{}{"error": "form error", "detail": err.Error()})
	}

	// validating the body
	if succeed, result := govalidator.ValidateStruct(form); !succeed {
		return c.Status(400).
			JSON(map[string]interface{}{"error": "form error", "detail": result.Error()})
	}

	user := models.User{
		PhoneNumber:        form.PhoneNumber,
		Password:           form.Password,
		IsStudent:          form.IsStudent,
		PartnerPhoneNumber: form.PartnerPhoneNumber,
		Nickname:           form.Nickname,
	}

	if err := user.Create(); err != nil {
		if err.Error() == "phone number duplicates" {
			return c.Status(400).
				JSON(map[string]interface{}{"error": "phone number duplicates", "detail": "the user with the same phone number already exists."})
		}
	}

	return c.JSON(user)
}
