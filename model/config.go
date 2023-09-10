package model

type UserSrvConfig struct {
	Name string `mapstructure:"name"`
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}

type ConsulConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	Name         string        `mapstructure:"name"`
	Host         string        `mapstructure:"host"`
	Port         int           `mapstructure:"port"`
	JwtInfo      JWTConfig     `mapstructure:"jwt"`
	ConsulConfig ConsulConfig  `mapstructure:"consul"`
	UserSrv      UserSrvConfig `mapstructure:"user_srv"`
}
