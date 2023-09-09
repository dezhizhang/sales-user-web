package initialize

import (
	"github.com/gin-gonic/gin"
	"sales-user-web/middleware"
	"sales-user-web/router"
)

func Routers() *gin.Engine {
	r := gin.Default()

	ApiGroup := r.Group("/api/v1").Use(middleware.Cors())
	router.UserRouter(ApiGroup)
	// router.CommonRouter(ApiGroup)
	return r
}
