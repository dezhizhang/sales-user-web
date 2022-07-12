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