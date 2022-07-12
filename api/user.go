package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"strconv"
	"user_web/global"
	"user_web/model"
	"user_web/proto"
	"user_web/utils"
)

//获取用户列表

func GetUserList(ctx *gin.Context) {

	pageIndex := ctx.DefaultQuery("pageIndex", "1")
	pageSize := ctx.DefaultQuery("pageSize", "10")
	pageIndexInt, _ := strconv.Atoi(pageIndex)
	pageSizeInt, _ := strconv.Atoi(pageSize)
	userSrvClient := proto.NewUserClient(global.Conn)
	rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{PageSize: uint32(pageSizeInt),
		PageIndex: uint32(pageIndexInt)})
	if err != nil {
		zap.S().Errorw("查询用户列表失败")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"msg":     "获取用户列成功",
		"success": true,
		"data":    rsp.Data,
		"total":   rsp.Total,
	})
	zap.S().Debug("获取用户列表")
}

func DeleteUser(ctx *gin.Context) {
	str := ctx.Param("id")
	id, err := strconv.Atoi(str)
	if err != nil {
		zap.S().Errorw("类型转换失败")
	}
	userSrvClient := proto.NewUserClient(global.Conn)
	rsp, err := userSrvClient.DeleteUser(context.Background(), &proto.IdRequest{
		Id: uint64(id),
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "删除用户失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"msg":     "删除用户成功",
		"success": true,
		"data":    rsp,
	})
}

func CreateUser(ctx *gin.Context) {
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		zap.S().Errorw("获取参数失败")
	}
	//user.Id = utils.SnowflakeId()
	userSrvClient := proto.NewUserClient(global.Conn)
	rsp, err := userSrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		Id:       uint64(utils.SnowflakeId()),
		Name:     user.Name,
		Role:     int32(user.Role),
		Gender:   int32(user.Gender),
		Password: user.Password,
		Mobile:   user.Mobile,
		Birthday: uint64(user.Birthday),
	})

	if err != nil {
		log.Fatalf("创建用户失败%s", err.Error())
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"msg":     "创健用户成功",
		"data":    rsp,
		"success": true,
	})

}
