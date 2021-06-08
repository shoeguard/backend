package configs

import (
	"syscall"

	_ "github.com/joho/godotenv/autoload"
)

type envs struct {
	PSQL_HOST     string
	PSQL_USER     string
	PSQL_PASSWORD string
	PSQL_DBNAME   string
	PSQL_PORT     string
	PSQL_SSLMODE  string
}

func getEnv(key string, defaultValue string) string {
	value, found := syscall.Getenv(key)
	if found {
		return value
	} else {
		return defaultValue
	}
}

func GetEnvs() (e envs) {
	e = envs{
		PSQL_HOST:     getEnv("PSQL_HOST", "localhost"),
		PSQL_USER:     getEnv("PSQL_USER", "postgres"),
		PSQL_PASSWORD: getEnv("PSQL_PASSWORD", "example"),
		PSQL_DBNAME:   getEnv("PSQL_DBNAME", "postgres"),
		PSQL_PORT:     getEnv("PSQL_PORT", "5432"),
		PSQL_SSLMODE:  getEnv("PSQL_SSLMODE", "disable"),
	}
	return
}
