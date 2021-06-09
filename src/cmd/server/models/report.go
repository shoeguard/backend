package models

import (
	"shoeguard-main-backend/cmd/server/utils"

	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	ReporterID       int
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
