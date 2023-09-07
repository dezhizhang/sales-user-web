package global

import (
	ut "github.com/go-playground/universal-translator"
	"google.golang.org/grpc"
	"sales-user-web/model"
)

var (
	Conn         *grpc.ClientConn
	Trans        ut.Translator
	ServerConfig *model.ServerConfig = &model.ServerConfig{}
)
