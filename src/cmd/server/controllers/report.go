package controllers

import (
	"shoeguard-main-backend/cmd/server/forms"
	"shoeguard-main-backend/cmd/server/models"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

// Report godoc
// @Summary Report Request
// @Description Report Request with given information
// @security BasicAuth
// @Accept json
// @Produce json
// @Param body body forms.ReportForm true "body"
// @Success 201 {object} models.ReportDocument
// @Router /report [post]
func Report(c *fiber.Ctx) error {
	form := forms.ReportForm{}

	// parsing body
	if err := c.BodyParser(&form); err != nil {
		return c.Status(400).
			JSON(map[string]interface{}{"error": "form error", "detail": err.Error()})
	}

	// validating body
	if succeed, result := govalidator.ValidateStruct(form); !succeed {
		return c.Status(400).
			JSON(map[string]interface{}{"error": "form error", "detail": result.Error()})
	}

	phoneNumber := c.Locals("username").(string)
	user := models.User{}
	user.SetUser(phoneNumber)

	report := models.Report{
		Reporter:   user,
		DeviceInfo: form.DeviceInfo,
		Latitude:   form.Latitude,
		Longitude:  form.Longitude,
	}

	if err := report.Create(); err != nil {
		return c.Status(500).
			JSON(map[string]interface{}{"error": "unknown error", "detail": err.Error()})
	}

	return c.Status(201).JSON(report)
}
