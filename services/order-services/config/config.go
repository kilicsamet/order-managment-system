package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HttpListenPort string
	PostgresHost   string
	PostgresPort   string
	PostgresDbUser string
	PostgresDbPass string
	PostgresDbName string
	SqLiteDbPath   string
	RabbitMQUser   string
	RabbitMQPass   string
	RabbitMQHost   string
	RabbitMQPort   string
}

func New() *Config {
	checkDotEnvFile()

	config := &Config{
		HttpListenPort: getEnvVariable("HTTP_LISTEN_PORT"),
		PostgresHost:   getEnvVariable("POSTGRES_HOST"),
		PostgresPort:   getEnvVariable("POSTGRES_PORT"),
		PostgresDbUser: getEnvVariable("POSTGRES_DB_USER"),
		PostgresDbPass: getEnvVariable("POSTGRES_DB_PASS"),
		PostgresDbName: getEnvVariable("POSTGRES_DB_NAME"),
		SqLiteDbPath:   getEnvVariable("SQLITE_DB_PATH"),
		RabbitMQUser:   os.Getenv("RABBITMQ_USER"),
		RabbitMQPass:   os.Getenv("RABBITMQ_PASS"),
		RabbitMQHost:   os.Getenv("RABBITMQ_HOST"),
		RabbitMQPort:   os.Getenv("RABBITMQ_PORT"),
	}
	return getDefaultConfig(config)
}

func getEnvVariable(key string) string {
	godotenv.Load(".env")
	return os.Getenv(key)
}

func checkDotEnvFile() {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		os.Create(".env")
	}
}
