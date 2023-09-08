package model

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type JWTConfig struct {
}

type ServerConfig struct {
	Name    string        `mapstructure:"name"`
	Host    string        `mapstructure:"host"`
	Port    int           `mapstructure:"port"`
	UserSrv UserSrvConfig `mapstructure:"user_srv"`
}
