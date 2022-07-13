package router

import "github.com/gin-gonic/gin"

func IBaseRouter(Router *gin.RouterGroup) {
	baseRouter := Router.Group("/captcha", api)
}
