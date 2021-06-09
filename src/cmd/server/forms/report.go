package forms

import "time"

type ReportRequestForm struct {
	DeviceInfo string  `json:"device_info" example:"00:11:22:33:FF:EE"  valid:"required"`
	Latitude   float64 `                   example:"37.5428147089301"   valid:"required"`
	Longitude  float64 `                   example:"126.96631451849314" valid:"required"`
}

type ReportResponse struct {
	ID               uint      `json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	RecordedAudioURL string    `json:"recorded_audio_url"`
	DeviceInfo       string    `json:"device_info"        example:"00:11:22:33:FF:EE"  valid:"required"`
	Latitude         float64   `json:"latitude"           example:"37.5428147089301"   valid:"required"`
	Longitude        float64   `json:"longitude"          example:"126.96631451849314" valid:"required"`
}

type AddRecordedAudioURLForm struct {
	AudioURL string `json:"audio_url"`
}

type ReportsResponse struct {
	IsReportOfStudent bool             `json:"is_report_of_student"`
	Reports           []ReportResponse `json:"reports"`
}
