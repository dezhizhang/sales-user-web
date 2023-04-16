package router

import (
	"github.com/gin-gonic/gin"
	"sales-user-web/api"
)

func CommonRouter(Router *gin.RouterGroup) {
	common := Router.Group("/common")
	sms := api.Sms{}
	captcha := api.Captcha{}

	{
		common.POST("/send-sms", sms.SendSms)
		common.GET("/captcha", captcha.GetCaptcha)

	}

}
