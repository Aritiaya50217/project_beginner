package config

import (
	"log"
	"os"
	"time"
)

type Config struct {
	ServerPort    string
	SQLServerUser string
	SQLServerPass string
	SQLServerHost string
	SQLServerPort string
	SQLServerDB   string
	JWTSecret     string
}

func LoadConfig() *Config {
	return &Config{
		ServerPort:    os.Getenv("USER_SERVICE_PORT"),
		SQLServerUser: os.Getenv("USER_SQLSERVER_USER"),
		SQLServerPass: os.Getenv("USER_SQLSERVER_PASSWORD"),
		SQLServerHost: os.Getenv("USER_SQLSERVER_HOST"),
		SQLServerPort: os.Getenv("USER_SQLSERVER_PORT"),
		SQLServerDB:   os.Getenv("USER_SQLSERVER_DB"),
		JWTSecret:     os.Getenv("JWT_SECRET"),
	}
}

func SetTimeZone() {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Fatalf("failed to load location: %v", err)
	}

	time.Local = loc
}
