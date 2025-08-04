package config

import (
	"log"
	"os"

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
	// Set default paths
	viper.SetConfigName("config") // config.yaml
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("..") // for monorepo

	// Read from .env first
	viper.AutomaticEnv()

	// Read from config.yaml if exists
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("No config.yaml found, relying on env variables only")
	}

	return &Config{
		ServerPort:    getEnvOrViper("USER_SERVICE_PORT"),
		SQLServerUser: getEnvOrViper("SQLSERVER_USER"),
		SQLServerPass: getEnvOrViper("SQLSERVER_PASSWORD"),
		SQLServerHost: getEnvOrViper("SQLSERVER_HOST"),
		SQLServerPort: getEnvOrViper("SQLSERVER_PORT"),
		SQLServerDB:   getEnvOrViper("SQLSERVER_DB"),
		JWTSecret:     getEnvOrViper("JWT_SECRET"),
	}
}

func getEnvOrViper(key string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return viper.GetString(key)
}
