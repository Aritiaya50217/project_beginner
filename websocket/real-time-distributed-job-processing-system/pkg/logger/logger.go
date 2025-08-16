package logger

import "go.uber.org/zap"

var Log *zap.Logger

func Init() {
	logs, _ := zap.NewProduction()
	Log = logs
}
