package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"sales-user-web/global"
	"sales-user-web/proto"
)

// 初始化grpc链接

func InitSrvConn() {

	var err error
	userSrvAddress := ""

	name := global.ServerConfig.UserSrv.Name
	cfg := api.DefaultConfig()
	consulConfig := global.ServerConfig.ConsulConfig
	cfg.Address = fmt.Sprintf("%s:%d", consulConfig.Host, consulConfig.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err1 := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, name))
	if err1 != nil {
		panic(err1)
	}

	for _, value := range data {
		userSrvAddress = value.Address
		break
	}
	if userSrvAddress == "" {
		zap.S().Errorw("initSrvConn 连接用户服务失败", "msg", err.Error())
	}

	fmt.Println("userSrvAddress", userSrvAddress)

	userConn, err := grpc.Dial(userSrvAddress, grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("initSrvConn 连接用户服务失败", "msg", err.Error())
	}

	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient

}
