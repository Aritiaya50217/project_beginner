package config

import (
	"log"
	"os"
	"strings"
	"time"
)

type Config struct {
	ServerPort     string
	SQLServerUser  string
	SQLServerPass  string
	SQLServerHost  string
	SQLServerPort  string
	SQLServerDB    string
	JWTSecret      string
	UserserviceUrl string
	KafkaBrokers   []string
	FinnhubAPIKey  string
}

func LoadConfig() *Config {
	cfg := &Config{
		ServerPort:     os.Getenv("STOCK_SERVICE_PORT"),
		SQLServerUser:  os.Getenv("STOCK_SQLSERVER_USER"),
		SQLServerPass:  os.Getenv("STOCK_SQLSERVER_PASSWORD"),
		SQLServerHost:  os.Getenv("STOCK_SQLSERVER_HOST"),
		SQLServerPort:  os.Getenv("STOCK_SQLSERVER_PORT"),
		SQLServerDB:    os.Getenv("STOCK_SQLSERVER_DB"),
		JWTSecret:      os.Getenv("JWT_SECRET"),
		UserserviceUrl: os.Getenv("USER_SERVICE_URL"),
		KafkaBrokers:   strings.Split(os.Getenv("KAFKA_BROKERS"), ","),
		FinnhubAPIKey:  os.Getenv("FINNHUB_API_KEY"),
	}

	if len(cfg.KafkaBrokers) == 0 {
		log.Fatal("KAFKA_BROKERS is not set")
	}

	return cfg

}

func SetTimeZone() {
	loc, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		log.Fatalf("failed to load location: %v", err)
	}

	time.Local = loc
}
