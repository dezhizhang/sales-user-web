package test

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"testing"
)

func TestNacos(t *testing.T) {
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   8848,
		},
	}

	clientConfig := constant.ClientConfig{
		NamespaceId:         "a5445f77-b021-4247-8637-3f27f2ca08fa", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	client, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}

	content, err := client.GetConfig(vo.ConfigParam{
		DataId: "sales-user-web",
		Group:  "dev"})

	if err != nil {
		panic(err)
	}

	fmt.Println("------", content)

}
