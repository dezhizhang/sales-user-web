package main

import (
	"fmt"
	"go.uber.org/zap"
	"sales-user-web/global"
	"sales-user-web/initialize"
)

func main() {
	//1 初始化日志
	initialize.Logger()

	//2 初始化配置文件
	initialize.InitConfig()

	fmt.Println(global.ServerConfig)

	//3 初始化routers
	router := initialize.Routers()

	// 初始化翻译
	err := initialize.InitTrans("zh")
	if err != nil {
		zap.S().Errorf("初始化翻译器失败%s", err.Error())
	}

	zap.S().Debugf("启动服务器端口运行在:%d", global.ServerConfig.Port)

	err = router.Run(fmt.Sprintf("%s:%d", global.ServerConfig.Host, global.ServerConfig.Port))
	if err != nil {
		//log.Printf("运行失败%s", err.Error())
		zap.S().Panic("启动失败:%s", err.Error())
	}

}
