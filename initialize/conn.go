package initialize

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"user_web/global"
)

// 初始化grpc链接

func InitConn() {
	var err error
	global.Conn, err = grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("连接服务失败", "msg", err.Error())
	}
}
