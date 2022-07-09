package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	Mysql MysqlConfig `mapstructure:"mysql"`
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
}
