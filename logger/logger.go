package main

import (
	"go.uber.org/zap"
	"time"
)

func NewLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{
		"./myproject.log",
		"stderr",
	}
	return config.Build()
}

func main() {
	logger, err := NewLogger()
	if err != nil {
		panic(err)
	}
	su := logger.Sugar()
	defer su.Sync()
	su.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "https",
		"attempt", 3,
		"backoff", time.Second,
	)
	su.Infof("Failed to fetch URL: %s", "htps")

}
