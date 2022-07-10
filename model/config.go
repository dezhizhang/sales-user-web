package model

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	Name    string        `mapstructure:"name"`
	Port    string        `mapstructure:"port"`
	UserSrv UserSrvConfig `mapstructure:"user_srv"`
}
