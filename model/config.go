package model

type UserSrvConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type ServerConfig struct {
	Name    string        `mapstructure:"name"`
	Host    string        `mapstructure:"host"`
	Port    int           `mapstructure:"port"`
	JwtInfo JWTConfig     `mapstructure:"jwt"`
	UserSrv UserSrvConfig `mapstructure:"user_srv"`
}
