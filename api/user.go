package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"log"
	"net/http"
	"sales-user-web/forms"
	"sales-user-web/global"
	"sales-user-web/model"
	"sales-user-web/proto"
	"sales-user-web/utils"
	"strconv"
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
	id := ctx.Param("id")
	userSrvClient := proto.NewUserClient(global.Conn)
	rsp, err := userSrvClient.DeleteUser(context.Background(), &proto.IdRequest{
		Id: id,
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
	userSrvClient := proto.NewUserClient(global.Conn)
	rsp, err := userSrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		Id:       utils.SnowflakeId(),
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

func LoginIn(c *gin.Context) {
	loginUserForm := forms.LoginUserForm{}
	if err := c.ShouldBind(&loginUserForm); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"code":    400,
				"msg":     err.Error(),
				"success": false,
				"data":    nil,
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"success": false,
			"msg":     errs.Translate(global.Trans),
			"data":    nil,
		})
		return
	}

	userSrvClient := proto.NewUserClient(global.Conn)
	res, err := userSrvClient.GetUserByExist(context.Background(), &proto.UserLogin{
		Mobile:   loginUserForm.Mobile,
		Password: loginUserForm.Password,
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"msg":     err.Error(),
			"success": false,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"msg":     "获取成功",
		"success": true,
		"data":    res,
	})

}
