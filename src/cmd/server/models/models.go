package models

import "shoeguard-main-backend/cmd/server/utils"

func GetModels() (models []interface{}) {
	models = []interface{}{User{}, Report{}}
	return
}

func MigrateModels() {
	db := utils.GetDB()
	models := GetModels()
	for _, model := range models {
		db.AutoMigrate(model)
	}
}
