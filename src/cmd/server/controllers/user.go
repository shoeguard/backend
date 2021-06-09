package controllers

import (
	"shoeguard-main-backend/cmd/server/customErrors"
	"shoeguard-main-backend/cmd/server/forms"
	"shoeguard-main-backend/cmd/server/models"
	"shoeguard-main-backend/cmd/server/utils"

	"github.com/gofiber/fiber/v2"
)

// Register godoc
// @Summary Register an user with given information
// @Description Register user
// @Accept json
// @Produce json
// @Param body body forms.UserRegistrationForm true "body"
// @Success 201 {object} forms.UserRegistrationForm
// @Success 400 {object} customErrors.ErrorResponse
// @Router /users/register [post]
func Register(c *fiber.Ctx) error {
	form := forms.UserRegistrationForm{}

	if err := utils.ParseAndValidateForm(c, &form); err != nil {
		return customErrors.Response400WithError(c, customErrors.FormError, err.Error())
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
			return customErrors.Response400WithError(
				c,
				customErrors.PhoneNumberDuplicate,
				"The user with the same phone number already exists.",
			)
		}
		return c.SendStatus(500)
	}

	return c.Status(201).JSON(form)
}
