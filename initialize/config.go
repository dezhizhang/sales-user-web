package initialize

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"sales-user-web/global"
	"sales-user-web/model"
)

var UserWeb *model.ServerConfig

func Config() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	global.UserWeb = &model.ServerConfig{}
	err = v.Unmarshal(&global.UserWeb)
	if err != nil {
		panic(err)
	}
	zap.S().Infof("初始化配置文件")
}
