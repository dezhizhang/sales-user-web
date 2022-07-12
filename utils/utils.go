package utils

import (
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

func SnowflakeId() string {
	node, err := snowflake.NewNode(1)
	if err != nil {
		zap.S().Errorw("生成雪花算法失败")
	}
	return node.Generate().String()
}
