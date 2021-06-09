package configs

import (
	"syscall"

	_ "github.com/joho/godotenv/autoload"
)

func getEnv(key string, defaultValue string) string {
	value, found := syscall.Getenv(key)
	if found {
		return value
	} else {
		return defaultValue
	}
}

var (
	PSQL_HOST      = getEnv("PSQL_HOST", "localhost")
	PSQL_USER      = getEnv("PSQL_USER", "postgres")
	PSQL_PASSWORD  = getEnv("PSQL_PASSWORD", "example")
	PSQL_DBNAME    = getEnv("PSQL_DBNAME", "postgres")
	PSQL_PORT      = getEnv("PSQL_PORT", "5432")
	PSQL_SSLMODE   = getEnv("PSQL_SSLMODE", "disable")
	ENABLE_SWAGGER = getEnv("ENABLE_SWAGGER", "true")
	PORT           = getEnv("PORT", "8080")
)
