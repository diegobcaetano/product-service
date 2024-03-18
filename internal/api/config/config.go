package config

import (
	"github.com/diegobcaetano/product-service/internal/logging"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBHost     string `envconfig:"DB_HOST"`
	DBKeyspace string `envconfig:"DB_KEYSPACE"`
	DBUser     string `envconfig:"DB_USER"`
	DBPassword string `envconfig:"DB_PASSWORD"`
}

type EnvLoader interface {
	Load(path string) EnvLoader
	GetConfig() *Config
}

type EnvLoad struct {
	logger logging.Logger
	Env    *Config
}

func NewEnvLoad(logger logging.Logger) EnvLoader {
	return &EnvLoad{
		logger: logger,
	}
}

func (e *EnvLoad) Load(path string) EnvLoader {
	// Carrega variáveis de ambiente de .env em desenvolvimento
	if err := godotenv.Load(); err != nil {
		e.logger.Info("File .env not found, assuming env is already set")
	}

	var c Config
	if err := envconfig.Process("", &c); err != nil {
		e.logger.Error("Failed to configure environment variables", "error", err.Error())
	}
	e.Env = &c
	return e
}

func (e *EnvLoad) GetConfig() *Config {
	return e.Env
}
