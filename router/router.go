package router

import (
	"github.com/gin-gonic/gin"
	"user_web/api"
)

func Router(Router *gin.RouterGroup) {
	userRouter := Router.Group("/user")
	{
		userRouter.POST("/add", api.CreateUser)
		userRouter.GET("/list", api.GetUserList)
		userRouter.POST("/login", api.LoginIn)
		userRouter.DELETE("/delete/:id", api.DeleteUser)
	}

}
