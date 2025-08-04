package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
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
	viper.SetConfigName("config") // config.yaml
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("..") // for monorepo

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No config.yaml found, relying on env variables only")
	}

	return &Config{
		ServerPort:    os.Getenv("BOOKING_SERVICE_PORT"),
		SQLServerUser: os.Getenv("BOOKING_SQLSERVER_USER"),
		SQLServerPass: os.Getenv("BOOKING_SQLSERVER_PASSWORD"),
		SQLServerHost: os.Getenv("BOOKING_SQLSERVER_HOST"),
		SQLServerPort: os.Getenv("BOOKING_SQLSERVER_PORT"),
		SQLServerDB:   os.Getenv("BOOKING_SQLSERVER_DB"),
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
