package models

import (
	"shoeguard-main-backend/cmd/server/utils"

	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	ReporterID       uint
	Reporter         User
	DeviceInfo       string `gorm:"not null"`
	RecordedAudioURL string
	Latitude         float64 `gorm:"not null"`
	Longitude        float64 `gorm:"not null"`
}

func (report *Report) Create() error {
	db := utils.GetDB()
	return db.Create(&report).Error
}

func (report *Report) SetReportByID(id uint64) error {
	db := utils.GetDB()
	return db.First(&report, id).Error
}

func (report *Report) Save() error {
	db := utils.GetDB()
	return db.Save(&report).Error
}

type ReportDAO struct{}

func (reportDAO ReportDAO) GetReportsByReporterID(reports *[]Report, id uint) error {
	db := utils.GetDB()
	return db.Find(&reports, "reporter_id = ?", id).Error
}
