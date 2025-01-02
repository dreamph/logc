package main

import (
	"context"
	"github.com/dreamph/logc"
	"github.com/dreamph/logc/zerolog"
)

func main() {
	logger := zerolog.NewLogger(&logc.Options{
		FilePath: "./app.log",
		Level:    "debug",
		Format:   "json",
		MaxAge:   30,
		MaxSize:  10,
	})
	defer logger.Release()

	logger.Info("Test Info")
	logger.Warn("Test Warn")

	d := map[string]interface{}{
		"requestId": "123",
	}
	log := logger.WithLogger(logc.WithValue(context.Background(), d))
	log.Info("Test")
}
