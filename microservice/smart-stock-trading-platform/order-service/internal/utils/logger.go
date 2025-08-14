package utils

import "go.uber.org/zap"

var Logger *zap.Logger

func InitLogger() {
	logs, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	Logger = logs
}
