package api

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"sales-user-web/global"
	"sales-user-web/middleware"
	"sales-user-web/model"
	"sales-user-web/proto"
	"sales-user-web/utils"
	"time"
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

// GetUserList 获取用户列表
func GetUserList(ctx *gin.Context) {
	zap.S().Debugf("连接用户服务")
	host := global.ServerConfig.UserSrv.Host
	port := global.ServerConfig.UserSrv.Port
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("GetUserList 连拉用户服务失败", "msg", err.Error())
	}
	userSrvClient := proto.NewUserClient(userConn)
	rsp, err1 := userSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		PageIndex: uint32(utils.TransformStringToInt(ctx.DefaultQuery("pageIndex", "1"))),
		PageSize:  uint32(utils.TransformStringToInt(ctx.DefaultQuery("pageSize", "10"))),
	})
	if err1 != nil {
		zap.S().Errorw("GetUserList 查询用户列表失败", "msg", err1.Error())
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		user := model.UserResponse{
			Id:       value.Id,
			Name:     value.Name,
			Birthday: int64(value.Birthday),
			Gender:   int(value.Gender),
			Mobile:   value.Mobile,
		}
		result = append(result, user)
	}
	utils.ResponseSuccessJson(ctx, "获取用户成功", int(rsp.Total), result)

}

// UserLoginIn 用户登录
func UserLoginIn(ctx *gin.Context) {
	var loginUserForm model.LoginUserForm
	if err := ctx.ShouldBindJSON(&loginUserForm); err != nil {
		utils.ValidatorError(ctx, err)
		return
	}
	host := global.ServerConfig.UserSrv.Host
	port := global.ServerConfig.UserSrv.Port
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("GetUserList 连拉用户服务失败", "msg", err.Error())
	}
	userSrvClient := proto.NewUserClient(userConn)
	rsp, err1 := userSrvClient.GetUserByExist(context.Background(), &proto.UserLogin{
		Mobile:   loginUserForm.Mobile,
		Password: loginUserForm.Password,
	})
	if err1 != nil {
		utils.ResponseErrorJson(ctx, http.StatusBadRequest, "登录失败")
		return
	}
	j := middleware.NewJWT()
	claims := model.CustomClaims{
		Id:   rsp.Id,
		Name: rsp.Name,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*24,
			Issuer:    "sales-user-web",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		utils.ResponseErrorJson(ctx, http.StatusInternalServerError, "内部错误")
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "登录成功", "token": token, "data": rsp})
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
	//var user model.User
	//err := ctx.BindJSON(&user)
	//if err != nil {
	//	zap.S().Errorw("获取参数失败")
	//}
	//userSrvClient := proto.NewUserClient(global.Conn)
	//rsp, err := userSrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
	//	Id:       utils.SnowflakeId(),
	//	Name:     user.Name,
	//	Role:     int32(user.Role),
	//	Gender:   int32(user.Gender),
	//	Password:u]\y7yuiopiop[+	qwr]\user.Password,
	//	Mobile:   user.Mobile,
	//	Birthday: uint64(user.Birthday),
	//})
	//
	//if err != nil {
	//	log.Fatalf("创建用户失败%s", err.Error())
	//}
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code":    200,
	//	"msg":     "创健用户成功",
	//	"data":    rsp,
	//	"success": true,
	//})

}
