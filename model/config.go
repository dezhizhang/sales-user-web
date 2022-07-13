package model

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type JwtConfig struct {
	SigningKey string `mapstructure:"key"`
}

type ServerConfig struct {
	Name    string        `mapstructure:"name"`
	Port    string        `mapstructure:"port"`
	JwtInfo JwtConfig     `mapstructure:"jwt"`
	UserSrv UserSrvConfig `mapstructure:"user_srv"`
}
