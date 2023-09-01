package test

import (
	"go.uber.org/zap"
	"log"
	"testing"
	"time"
)

func TestZap(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	url := "https://imooc.com"

	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL", "url", url, "attempt", 3)
	sugar.Infof("Failed to fetch URL:%s", url)
}

func TestZapFile(t *testing.T) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{
		"./myproject.log",
	}
	logger, err := config.Build()
	if err != nil {
		log.Printf("写入日志失败%s", err.Error())
	}
	su := logger.Sugar()
	defer logger.Sync()

	url := "https://imooc.com"
	su.Info("failed to fetch URL",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
