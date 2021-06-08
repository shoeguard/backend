package utils

import (
	"fmt"
	"log"
	"shoeguard-main-backend/configs"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once

func GetDB() *gorm.DB {
	once.Do(func() {
		var err error
		envs := configs.GetEnvs()
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			envs.PSQL_HOST,
			envs.PSQL_USER,
			envs.PSQL_PASSWORD,
			envs.PSQL_DBNAME,
			envs.PSQL_PORT,
			envs.PSQL_SSLMODE,
		)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Connected to the database.")
		}
	})
	return DB
}
