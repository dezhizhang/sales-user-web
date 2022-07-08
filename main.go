package main

import (
	"go.uber.org/zap"
	"user_web/initialize"
)

func main() {
	//初始化日志
	initialize.Logger()

	//初始化路由
	r := initialize.Routers()
	zap.S().Debugf("启动服务，端口:%d", 8082)
	err := r.Run(":8082")

	if err != nil {
		zap.S().Panic("启动失败", err.Error())
	}

}
