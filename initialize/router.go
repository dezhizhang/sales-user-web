package initialize

import "github.com/gin-gonic/gin"

func Routers() *gin.Engine {
	r := gin.Default()
	ApiGroup := r.Group("/api/v1")

	return r
}
