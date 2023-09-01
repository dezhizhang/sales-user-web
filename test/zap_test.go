package test

import (
	"go.uber.org/zap"
	"testing"
)

func TestZap(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	url := "https://imooc.com"

	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL", "url", url, "attempt", 3)
	sugar.Infof("Failed to fetch URL:%s", url)
}
