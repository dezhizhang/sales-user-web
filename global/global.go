package global

import (
	"google.golang.org/grpc"
	"user_web/model"
)

var (
	UserWeb *model.ServerConfig
	Conn    *grpc.ClientConn
)
