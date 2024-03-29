package test

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"sales-user-web/model"
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
		NamespaceId:         "a5445f77-b021-4247-8637-3f27f2ca08fa",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	client, _ := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	content, _ := client.GetConfig(vo.ConfigParam{
		DataId: "sales-user-web",
		Group:  "dev",
	})

	serverSrvConfigs := model.ServerConfig{}
	json.Unmarshal([]byte(content), &serverSrvConfigs)

	//client.ListenConfig(vo.ConfigParam{
	//	DataId: "sales-user-web",
	//	Group:  "dev",
	//	OnChange: func(namespace, group, dataId, data string) {
	//		fmt.Println("配置文件发生变化")
	//	},
	//})
	//time.Sleep(5000 * time.Second)

	fmt.Println(serverSrvConfigs)

}
