package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	url := "https://imooc.com"

	// Sugar mode
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
		"mode", "sugar.Infow()",
	)
	sugar.Infof("Failed to fetch URL: %s", url,
		"mode", "sugar.Infof()",
	)

	// Logger mode
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
		zap.String("mode", "logger.Info()"),
	)
}
