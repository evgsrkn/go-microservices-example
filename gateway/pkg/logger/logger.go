package logger

import (
	"os"

	"go.uber.org/zap"
)

func init() {
	os.Setenv("TZ", "Europe/Moscow")
}

func New() *zap.Logger {
	logger := zap.Must(zap.NewProduction())
	return logger
}
