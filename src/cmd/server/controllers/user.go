package controllers

import (
	"shoeguard-main-backend/cmd/server/customErrors"
	"shoeguard-main-backend/cmd/server/forms"
	"shoeguard-main-backend/cmd/server/models"
	"shoeguard-main-backend/cmd/server/utils"

	"github.com/gofiber/fiber/v2"
)

func getUser(phoneNumber string) (user models.User, err error) {
	user = models.User{}
	err = user.SetUser(phoneNumber)
	return
}

// Register godoc
// @Summary Register an user with given information
// @Description Register user
// @Accept json
// @Produce json
// @Param body body forms.UserRegistrationForm true "body"
// @Success 201 {object} forms.UserReadOnlyInfo
// @Failure 400 {object} customErrors.ErrorResponse
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

	resp := forms.UserReadOnlyInfo{
		PhoneNumber:        form.PhoneNumber,
		IsStudent:          form.IsStudent,
		PartnerPhoneNumber: form.PartnerPhoneNumber,
		Nickname:           form.Nickname,
	}

	return c.Status(201).JSON(resp)
}

// GetMyInfo godoc
// @Summary Get my info
// @Description Get my profile
// @Accept json
// @Produce json
// @Security BasicAuth
// @Success 200 {object} forms.UserReadOnlyInfo
// @Failure 401 string Unauthorized
// @Router /users [get]
func GetMyInfo(c *fiber.Ctx) error {
	user, _ := getUser(c.Locals("usermame").(string))
	resp := forms.UserReadOnlyInfo{
		PhoneNumber:        user.PhoneNumber,
		IsStudent:          user.IsStudent,
		PartnerPhoneNumber: user.PartnerPhoneNumber,
		Nickname:           user.Nickname,
	}
	return c.JSON(resp)
}

// UpdateMyInfo godoc
// @Summary Update my info
// @Description Update my profile
// @Accept json
// @Produce json
// @Security BasicAuth
// @Param body body forms.UserModifiableInfo true "body"
// @Success 200 {object} forms.UserReadOnlyInfo
// @Failure 401 string Unauthorized
// @Router /users [patch]
func UpdateMyInfo(c *fiber.Ctx) error {
	// validate
	form := forms.UserModifiableInfo{}
	if err := utils.ParseAndValidateForm(c, &form); err != nil {
		return customErrors.Response400WithError(c, customErrors.FormError, err.Error())
	}

	user, _ := getUser(c.Locals("usermame").(string))

	if form.Password != "" {
		user.Password = form.Password
		user.HashPassword()
	}
	if form.IsStudent != nil {
		user.IsStudent = *form.IsStudent
	}
	if form.PartnerPhoneNumber != "" {
		user.PartnerPhoneNumber = form.PartnerPhoneNumber
	}
	if form.Nickname != "" {
		user.Nickname = form.Nickname
	}

	user.Save()

	return GetMyInfo(c)
}
