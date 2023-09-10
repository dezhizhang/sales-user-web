package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"sales-user-web/global"
	"sales-user-web/initialize"
	myvalidator "sales-user-web/validator"
)

func main() {
	//1 初始化日志
	initialize.Logger()

	//2 初始化配置文件
	initialize.InitConfig()

	//3 初始化routers
	router := initialize.Routers()

	// 初始化翻译
	err := initialize.InitTrans("zh")
	if err != nil {
		zap.S().Errorf("初始化翻译器失败%s", err.Error())
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号!", true)
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	zap.S().Debugf("启动服务器端口运行在:%d", global.ServerConfig.Port)

	err = router.Run(fmt.Sprintf("%s:%d", global.ServerConfig.Host, global.ServerConfig.Port))
	if err != nil {
		//log.Printf("运行失败%s", err.Error())
		zap.S().Panic("启动失败:%s", err.Error())
	}

}
