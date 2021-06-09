package controllers

import (
	"shoeguard-main-backend/cmd/server/customErrors"
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

func getReportResponses(user models.User) ([]forms.ReportResponse, error) {
	var reports []models.Report
	reportDAO := models.ReportDAO{}

	if err := reportDAO.GetReportsByReporterID(&reports, user.ID); err != nil {
		return nil, err
	}

	// convert []models.Report to []forms.ReportResponse

	reportResponses := make([]forms.ReportResponse, len(reports))
	for i, report := range reports {
		reportResponses[i].ID = report.ID
		reportResponses[i].CreatedAt = report.CreatedAt
		reportResponses[i].UpdatedAt = report.UpdatedAt
		reportResponses[i].RecordedAudioURL = report.RecordedAudioURL
		reportResponses[i].DeviceInfo = report.DeviceInfo
		reportResponses[i].Latitude = report.Latitude
		reportResponses[i].Longitude = report.Longitude
	}

	return reportResponses, nil
}

func getReportResponsesOfPartner(user models.User) ([]forms.ReportResponse, error) {
	partnerUser, err := getUser(user.PartnerPhoneNumber)
	if err != nil {
		return nil, err
	}

	return getReportResponses(partnerUser)
}

// Report godoc
// @Summary Report Request
// @Description Report Request with given information
// @Security BasicAuth
// @Accept json
// @Produce json
// @Param body body forms.ReportRequestForm true "body"
// @Success 201 {object} forms.ReportResponse
// @Success 400 {object} customErrors.ErrorResponse
// @Router /report [post]
func Report(c *fiber.Ctx) error {
	// validate
	form := forms.ReportRequestForm{}
	if err := utils.ParseAndValidateForm(c, &form); err != nil {
		return customErrors.Response400WithError(c, customErrors.FormError, err.Error())
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
		return c.SendStatus(500)
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

// GetReports godoc
// @Summary Get reports
// @Description Get reports from me or from my partners
// @Security BasicAuth
// @Accept json
// @Produce json
// @Success 200 {object} forms.ReportsResponse
// @Success 400 {object} customErrors.ErrorResponse
// @Router /report [get]
func GetReports(c *fiber.Ctx) error {
	user, _ := getUser(c.Locals("username").(string))

	response := forms.ReportsResponse{IsReportOfStudent: !user.IsStudent}
	var err error

	if user.IsStudent {
		var reports []forms.ReportResponse
		reports, err = getReportResponses(user)
		response.Reports = reports
	} else {
		if user.PartnerPhoneNumber == "" {
			return c.Status(400).
				JSON(map[string]interface{}{"error": "no child", "detail": "you haven't registered your child."})
		}
		var reports []forms.ReportResponse
		reports, err = getReportResponsesOfPartner(user)
		response.Reports = reports
	}

	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(response)
}

// AddRecordedAudioURL godoc
// @Summary Add recorded audio url
// @Description Add recorded audio url to given ID's report
// @Security BasicAuth
// @Accept json
// @Produce json
// @Param report_id path int true "Report ID"
// @Param body body forms.AddRecordedAudioURLForm true "body"
// @Success 200 {object} forms.ReportResponse
// @Success 400 {object} customErrors.ErrorResponse
// @Router /report/{report_id} [patch]
func AddRecordedAudioURL(c *fiber.Ctx) error {
	// validate
	form := forms.AddRecordedAudioURLForm{}
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return customErrors.Response400WithError(
			c,
			customErrors.ParamterError,
			"the paramter 'id' should be numeric.",
		)
	}
	if err := utils.ParseAndValidateForm(c, &form); err != nil {
		return customErrors.Response400WithError(c, customErrors.FormError, err.Error())
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
