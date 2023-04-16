package router

import (
	"github.com/gin-gonic/gin"
	"sales-user-web/api"
	"sales-user-web/middleware"
)

func UserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("/user").Use(middleware.Auth()).Use(middleware.Cors())
	{
		userRouter.POST("/add", api.CreateUser)
		userRouter.GET("/list", api.GetUserList)
		userRouter.POST("/login", api.LoginIn)
		userRouter.DELETE("/delete/:id", api.DeleteUser)
	}

}
