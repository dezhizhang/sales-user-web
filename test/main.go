package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//type MysqlConfig struct {
//	Host string `mapstructure:"host"`
//	Port int    `mapstructure:"port"`
//}
//
//type ServerConfig struct {
//	Mysql MysqlConfig `mapstructure:"mysql"`
//	Name  string      `mapstructure:"name"`
//}
//
//func main() {
//	v := viper.New()
//	v.SetConfigFile("config.yaml")
//	err := v.ReadInConfig()
//	if err != nil {
//		panic(err)
//	}
//	serverConfig := ServerConfig{}
//	err = v.Unmarshal(&serverConfig)
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(serverConfig)
//}

type MySqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	Mysql MySqlConfig `mapstructure:"mysql"`
	Name  string      `mapstructure:"name"`
}

func main() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	serverConfig := ServerConfig{}
	err = v.Unmarshal(&serverConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println(serverConfig)

	// 动态临控变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		v.ReadInConfig()
		v.Unmarshal(&serverConfig)
	})
	fmt.Println(serverConfig)
}
