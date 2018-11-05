package model

import "github.com/revel/revel/logger"

var modelLogger = logger.New()
func initLogger(baseLogger logger.MultiLogger) {
	modelLogger = baseLogger.New("module","revel-model")
}