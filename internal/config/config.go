package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type ShopConfig struct {
	Postgres PostgresConfig
}

type PostgresConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

var once sync.Once

var cfg *ShopConfig

func LoadConfig() *ShopConfig {
	once.Do(func() {
		_ = godotenv.Load()

		cfg = &ShopConfig{
			Postgres: PostgresConfig{
				Host:     getEnv("POSTGRES_HOST", ""),
				Password: getEnv("POSTGRES_PASSWORD", ""),
				User:     getEnv("POSTGRES_USER", ""),
				Port:     getEnv("POSTGRES_PORT", ""),
				DBName:   getEnv("POSTGRES_DB", ""),
			},
		}
	})

	return cfg
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}

func (p *PostgresConfig) GenerateDSN() string {
	return "postgres://" + p.User + ":" + p.Password + "@" + p.Host + ":" + p.Port + "/" + p.DBName + "?sslmode=disable"
}
