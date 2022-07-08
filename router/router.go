package router

import (
	"github.com/gin-gonic/gin"
	"user_web/api"
)

func Router(Router *gin.RouterGroup) {
	userRouter := Router.Group("/user")
	{
		userRouter.GET("/list", api.GetUserList)
	}

}
