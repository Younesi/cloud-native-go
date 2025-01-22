package config

import (
	"log/slog"
	"os"
	"time"

	"github.com/joeshaw/envdecode"
)

type Conf struct {
	Server ConfServer
	DB     ConfDB
}

type ConfServer struct {
	Port         int           `env:"SERVER_PORT,required"`
	TimeoutRead  time.Duration `env:"SERVER_TIMEOUT_READ,required"`
	TimeoutWrite time.Duration `env:"SERVER_TIMEOUT_WRITE,required"`
	TimeoutIdle  time.Duration `env:"SERVER_TIMEOUT_IDLE,required"`
	Debug        bool          `env:"SERVER_DEBUG"`
}

type ConfDB struct {
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	Username string `env:"DB_USER,required"`
	Password string `env:"DB_PASS,required"`
	DBName   string `env:"DB_NAME,required"`
	Debug    bool   `env:"DB_DEBUG"`
}

func New() *Conf {
	var c Conf
	if err := envdecode.StrictDecode(&c); err != nil {
		slog.Error("Failed to decode config", slog.Any("error", err))
	}

	// Set default value for SERVER_DEBUG if not provided
	if _, exists := os.LookupEnv("SERVER_DEBUG"); !exists {
		c.Server.Debug = false // Default value
	}
	return &c
}

func NewDB() *ConfDB {
	var c ConfDB
	if err := envdecode.StrictDecode(&c); err != nil {
		slog.Error("Failed to decode config", slog.Any("error", err))
	}

	// Set default value for DB_DEBUG if not provided
	if _, exists := os.LookupEnv("DB_DEBUG"); !exists {
		c.Debug = false
	}

	return &c
}
