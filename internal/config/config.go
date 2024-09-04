package config

import (
	"github.com/joho/godotenv"
)

var (
	logger = newLogger()
)

func InitConfigs() {
	logger.Info("Starting configs")
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func GetLogger() *Logger {
	return logger
}