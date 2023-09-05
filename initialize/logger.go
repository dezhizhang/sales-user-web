package initialize

import "go.uber.org/zap"

func Logger() {

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)

}
