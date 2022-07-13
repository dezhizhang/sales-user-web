package router

import (
	"github.com/gin-gonic/gin"
	"user_web/api"
	"user_web/middleware"
)

func Router(Router *gin.RouterGroup) {
	userRouter := Router.Group("/user").Use(middleware.Auth()).Use(middleware.Cors())
	{
		userRouter.POST("/add", api.CreateUser)
		userRouter.GET("/list", api.GetUserList)
		userRouter.POST("/login", api.LoginIn)
		userRouter.DELETE("/delete/:id", api.DeleteUser)
	}

}
