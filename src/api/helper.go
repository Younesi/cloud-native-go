package api

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

func GetSecretKey() []byte {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("Error loading .env file", slog.Any("error", err))
	}
	key := os.Getenv("JWT_SECRET_KEY")
	if key == "" {
		panic("JWT_SECRET_KEY environment variable is not set")
	}
	return []byte(key)
}
