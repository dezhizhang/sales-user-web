package driver

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var RDB *redis.Client

func InitDB() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // redis地址
		Password: "",               // redis密码，没有则留空
		DB:       0,                // 默认数据库，默认是0
	})

	_, err := RDB.Ping().Result()
	if err != nil {
		zap.S().Errorf("连接redis失败%s", err.Error())
		return
	}
	zap.S().Info("redis成功")

}

func init() {
	InitDB()
}
