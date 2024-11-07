package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	AppEnv      string
}

func LoadConfig() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	envFile := fmt.Sprintf("configs/.env.%s", env)
	if err := godotenv.Load(filepath.Clean(envFile)); err != nil {
		log.Printf("Warning: %sファイルが見つかりません。環境変数を直接使用します。\n", envFile)
	}

	return &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		AppEnv:      os.Getenv("APP_ENV"),
	}
}
