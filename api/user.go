package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net/http"
	"user_web/proto"
)

func GetUserList(ctx *gin.Context) {
	conn, err := grpc.Dial("127.0.0.1:8000", grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("连接服务失败", "msg", err.Error())
	}
	//proto.UserClient()
	userSrvClient := proto.NewUserClient(conn)
	rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{PageSize: 10, PageIndex: 1})
	if err != nil {
		zap.S().Errorw("查询用户列表失败")
		return
	}
	ctx.JSON(http.StatusOK, rsp)
	zap.S().Debug("获取用户列表")
}
