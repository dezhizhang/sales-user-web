package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"user_web/global"
	"user_web/initialize"
	validator2 "user_web/validator"
)

func main() {
	//初始化日志
	initialize.Logger()

	//初始化配置文件
	initialize.Config()

	//注册初始化翻译器
	initialize.InitTrans("zh")

	//自定义验证
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mobile", validator2.ValidateMobile)
	}

	//初始化grpc链接
	initialize.InitConn()
	//初始化路由
	r := initialize.Routers()
	zap.S().Debugf("启动服务，端口:%d", 8082)
	err := r.Run(fmt.Sprintf(":%s", global.UserWeb.Port))

	if err != nil {
		zap.S().Panic("启动失败", err.Error())
	}

}
