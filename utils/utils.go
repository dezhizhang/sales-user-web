package utils

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func SnowflakeId() string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		zap.S().Errorw("生成雪花算法失败")
	}
	return node.Generate().String()
}

func ResponseSuccessJson(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": msg, "data": data})
}
