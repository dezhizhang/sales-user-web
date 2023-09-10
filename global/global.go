package global

import (
	ut "github.com/go-playground/universal-translator"
	"google.golang.org/grpc"
	"sales-user-web/model"
	"sales-user-web/proto"
)

var (
	Conn          *grpc.ClientConn
	Trans         ut.Translator
	UserSrvClient proto.UserClient
	ServerConfig  *model.ServerConfig = &model.ServerConfig{}
)
