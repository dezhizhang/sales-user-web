package global

import (
	ut "github.com/go-playground/universal-translator"
	"google.golang.org/grpc"
	"user_web/model"
)

var (
	UserWeb *model.ServerConfig
	Conn    *grpc.ClientConn
	Trans   ut.Translator
)
