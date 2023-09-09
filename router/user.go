package router

import (
	"github.com/gin-gonic/gin"
	"sales-user-web/api"
)

func UserRouter(Router gin.IRoutes) {
	{
		Router.POST("/manager/add", api.CreateUser)
		Router.POST("/manager/list", api.GetUserList)
		Router.POST("/manager/login", api.UserLoginIn)
		Router.DELETE("/manager/:id", api.DeleteUser)
	}
	//userRouter := Router.Group("/user").Use(middleware.Cors())
	//{
	//	userRouter.POST("/add", api.CreateUser)
	//	userRouter.GET("/list", api.GetUserList)
	//	userRouter.POST("/login", api.UserLoginIn)
	//	userRouter.DELETE("/delete/:id", api.DeleteUser)
	//}

}
