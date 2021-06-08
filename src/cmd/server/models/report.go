package models

import (
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
