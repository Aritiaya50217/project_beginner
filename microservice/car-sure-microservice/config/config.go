package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DB_DSN    string
	JWTSecret string
	Port      string
}

func LoadConfig() *Config {
	// อ่านจาก environment variables โดยตรง
	viper.AutomaticEnv()

	cfg := &Config{
		DB_DSN:    viper.GetString("DB_DSN"),
		JWTSecret: viper.GetString("JWT_SECRET"),
		Port:      viper.GetString("PORT"),
	}

	return cfg
}
