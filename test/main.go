package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

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
	return err
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
