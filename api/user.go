package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/http"
	"sales-user-web/forms"
	"sales-user-web/global"
	"sales-user-web/model"
	"sales-user-web/proto"
	"sales-user-web/utils"
)

// HandleGrpcErrorToHttp 获取用户列表
func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{"msg": e.Message()})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "内部错误"})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{"msg": "参数错误"})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "其它错误"})

			}
		}
		return
	}
}

func GetUserList(ctx *gin.Context) {
	zap.S().Debugf("连接用户服务")
	userConn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("GetUserList 连拉用户服务失败", "msg", err.Error())
	}
	userSrvClient := proto.NewUserClient(userConn)
	list, err1 := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{PageIndex: 1, PageSize: 10})
	if err1 != nil {
		zap.S().Errorw("GetUserList 查询用户列表失败", "msg", err1.Error())
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": list})
	zap.S().Debug("获取用户列表")

	//pageIndex := ctx.DefaultQuery("pageIndex", "1")
	//pageSize := ctx.DefaultQuery("pageSize", "10")
	//pageIndexInt, _ := strconv.Atoi(pageIndex)
	//pageSizeInt, _ := strconv.Atoi(pageSize)
	//userSrvClient := proto.NewUserClient(global.Conn)
	//rsp, err := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{PageSize: uint32(pageSizeInt),
	//	PageIndex: uint32(pageIndexInt)})
	//if err != nil {
	//	zap.S().Errorw("查询用户列表失败")
	//	return
	//}
	//
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code":    200,
	//	"msg":     "获取用户列成功",
	//	"success": true,
	//	"data":    rsp.Data,
	//	"total":   rsp.Total,
	//})
	//zap.S().Debug("获取用户列表")
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
