package utils

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func SnowflakeId() string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		zap.S().Errorw("生成雪花算法失败")
	}
	return node.Generate().String()
}

// ResponseSuccessJson 成功时返回公共方法
func ResponseSuccessJson(ctx *gin.Context, msg string, total int, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": msg, "total": total, "data": data})
}

// ResponseErrorJson 失败时返回
func ResponseErrorJson(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{"code": code, "msg": msg})
}

// TransformStringToInt 字符串转换成int
func TransformStringToInt(value string) int {
	number, err := strconv.Atoi(value)
	if err != nil {
		zap.S().Errorw("TransformStringToInt 类型转换失败%s", err.Error())

	}
	return number
}

func ValidatorError(ctx *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if ok {
		ResponseErrorJson(ctx, http.StatusOK, err.Error())
		return
	}
	ResponseErrorJson(ctx, http.StatusBadGateway, errs.Error())
}
