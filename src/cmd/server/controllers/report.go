package controllers

import (
	"shoeguard-main-backend/cmd/server/forms"
	"shoeguard-main-backend/cmd/server/models"
	"shoeguard-main-backend/cmd/server/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getUser(phoneNumber string) (user models.User, err error) {
	user = models.User{}
	err = user.SetUser(phoneNumber)
	return
}

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
	// validate
	form := forms.ReportRequestForm{}
	if err := utils.ParseAndValidateForm(c, &form); err != nil {
		return c.Status(400).
			JSON(map[string]interface{}{"error": "form error", "detail": err.Error()})
	}

	// create report
	user, _ := getUser(c.Locals("username").(string))
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

	// response report
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
// @Param report_id path int true "Report ID"
// @Param body body forms.AddRecordedAudioURLForm true "body"
// @Success 200 {object} forms.ReportResponse
// @Router /report/{report_id} [patch]
func AddRecordedAudioURL(c *fiber.Ctx) error {
	// validate
	form := forms.AddRecordedAudioURLForm{}
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return c.Status(400).
			JSON(map[string]interface{}{"error": "parameter error", "detail": "the paramter 'id' should be numeric"})
	}
	if err := utils.ParseAndValidateForm(c, &form); err != nil {
		return c.Status(400).
			JSON(map[string]interface{}{"error": "form error", "detail": err.Error()})
	}

	// if no such report
	report := models.Report{}
	if err := report.SetReportByID(id); err != nil {
		return c.SendStatus(404)
	}

	// if the reporter is not owning the report
	user, _ := getUser(c.Locals("username").(string))
	if report.ReporterID != user.ID {
		return c.SendStatus(403)
	}

	// update and save the report
	report.RecordedAudioURL = form.AudioURL
	if err := report.Save(); err != nil {
		return c.SendStatus(500)
	}

	// response
	response := forms.ReportResponse{
		ID:               report.ID,
		CreatedAt:        report.CreatedAt,
		UpdatedAt:        report.UpdatedAt,
		RecordedAudioURL: report.RecordedAudioURL,
		DeviceInfo:       report.DeviceInfo,
		Latitude:         report.Latitude,
		Longitude:        report.Longitude,
	}
	return c.JSON(response)
}
