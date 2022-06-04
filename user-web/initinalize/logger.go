package initinalize

import "go.uber.org/zap"

func InitLogger() {
	logger, _ := zap.NewProduction()
	zap.ReplaceGlobals(logger)
}
