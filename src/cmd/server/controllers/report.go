package controllers

import (
	"shoeguard-main-backend/cmd/server/forms"
	"shoeguard-main-backend/cmd/server/models"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

// Report godoc
// @Summary Report Request
// @Description Report Request with given information
// @security BasicAuth
// @Accept json
// @Produce json
// @Param body body forms.ReportRequestForm true "body"
// @Success 201 {object} forms.ReportResponse
// @Router /report [post]
func Report(c *fiber.Ctx) error {
	form := forms.ReportRequestForm{}

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

	response := forms.ReportResponse{
		ID:               report.ID,
		CreatedAt:        report.CreatedAt,
		UpdatedAt:        report.UpdatedAt,
		RecordedAudioURL: report.RecordedAudioURL,
		DeviceInfo:       report.DeviceInfo,
		Latitude:         report.Latitude,
		Longitude:        report.Longitude,
	}

	return c.Status(201).JSON(response)
}

// AddRecordedAudioURL godoc
// @Summary Add recorded audio url
// @Description Add recorded audio url to given ID's report
// @security BasicAuth
// @Accept json
// @Produce json
// @Success 200 {object} forms.AddRecordedAudioURLForm
// @Router /report/{id} [patch]
func AddRecordedAudioURL(c *fiber.Ctx) error {
	form := forms.AddRecordedAudioURLForm{}
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	if err != nil {
		return c.Status(400).
			JSON(map[string]interface{}{"error": "parameter error", "detail": "the paramter 'id' should be numeric"})
	}

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

	report := models.Report{}
	if err := report.SetReportByID(id); err != nil {
		return c.SendStatus(404)
	}

	report.RecordedAudioURL = form.AudioURL
	if err := report.Save(); err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(report)
}
