package initialize

import (
	"github.com/gin-gonic/gin"
	"user_web/router"
)

func Routers() *gin.Engine {
	r := gin.Default()
	ApiGroup := r.Group("/api/v1")
	router.UserRouter(ApiGroup)
	router.CommonRouter(ApiGroup)
	return r
}
