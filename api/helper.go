package api

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetSecretKey() []byte {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Default().Println("Error loading .env file")
	}
	key := os.Getenv("JWT_SECRET_KEY")
	if key == "" {
		panic("JWT_SECRET_KEY environment variable is not set")
	}
	return []byte(key)
}
