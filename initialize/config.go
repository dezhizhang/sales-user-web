package initialize

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"sales-user-web/global"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = v.Unmarshal(&global.ServerConfig)
	if err != nil {
		panic(err)
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置文件发生变化%s", e.Name)
		v.ReadInConfig()
		v.Unmarshal(&global.ServerConfig)
	})
	zap.S().Infof("初始化配置文件")
}
