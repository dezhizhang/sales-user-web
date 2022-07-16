# user_web

### viper

```go
func main() {
    v := viper.New()
    v.SetConfigFile("config.yaml")
    err := v.ReadInConfig()
    if err != nil {
        panic(err)
    }
    fmt.Println(v.Get("name"))
}
```

### structure

```go
type ServerConfig struct {
Name string `mapstructure:"name"`
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
    fmt.Println(serverConfig.Name)
    //fmt.Println(v.Get("name"))
}
```

### 多结构体

```go
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

```

### 自定义验证器

```go
if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
    v.RegisterValidation("mobile", validator2.ValidateMobile)
    v.RegisterTranslation("mobile", global.Trans, func (ut ut.Translator) error {
    return ut.Add("mobile", "{0} 非法的手机号码", true)
    }, func (ut ut.Translator, fe validator.FieldError) string {
    t, _ := ut.T("mobile", fe.Field())
        return t
    })
}
```
### 服务注册与发现
```go
func Register(address string, port int, name string, tags []string, id string) error {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"

	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	check := &api.AgentServiceCheck{
		HTTP:                           "http://127.0.0.1:8082/health",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	// 生成注册对像

	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Check = check
	registration.Address = address

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	return nil

}

func AllService() {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"

	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}

func main() {
	err := Register("127.0.0.1", 8082, "user_web", []string{"shop", "saas"}, "123456")
	if err != nil {
		panic(err)
	}
	AllService()
}
```