package config

import "os"

const (
	DevEnvironment  = "dev"
	ProdEnvironment = "prod"
)

type Config struct {
	App         string
	Environment string
	LogLevel    string
	HttpPort    string
	Postgres    struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
	}
}

func New() (*Config, error) {
	cfg := &Config{}

	cfg.App = getEnv("APP", "recipe-app")
	cfg.Environment = getEnv("ENVIRONMENT", "dev")
	cfg.LogLevel = getEnv("LOG_LEVEL", "debug")
	cfg.HttpPort = getEnv("HTTP_PORT", ":8000")

	cfg.Postgres.Host = getEnv("POSTGRES_HOST", "localhost")
	cfg.Postgres.Port = getEnv("POSTGRES_PORT", "5432")
	cfg.Postgres.User = getEnv("POSTGRES_USER", "sam")
	cfg.Postgres.Password = getEnv("POSTGRES_PASSWORD", "")
	cfg.Postgres.Database = getEnv("POSTGRES_DATABASE", "db")

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}
