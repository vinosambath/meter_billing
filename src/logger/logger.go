package logger

import (
	"MeterBilling/src/configuration"
	"sync"

	"go.uber.org/zap"
)

var loggerInstance *zap.Logger
var loggerInstanceSingleton sync.Once

func InitLogger() {
	loggerInstanceSingleton.Do(func() {
		config := zap.NewProductionConfig()
		config.OutputPaths = []string{configuration.GetConfig().Logger.Path}
		logger, err := config.Build()

		if err != nil {
			panic(err)
		}
		loggerInstance = logger
	})
}

func GetLogger() *zap.Logger {
	return loggerInstance
}
