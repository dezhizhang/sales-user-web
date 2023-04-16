package api

import (
	"fmt"
	dysmsapi "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"sales-user-web/driver"
)

type Sms struct {
}

func (that *Sms) SendSms(c *gin.Context) {

	client, err := dysmsapi.NewClientWithAccessKey("cn-qingdao", "LTAI5tPSXr6dnDX2a4QmYHSF", "Gpk3pknYy4TyYzK1LKlOuFWmpLlNnx")

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = "15992478448"   //接收短信的手机号码
	request.SignName = "晓智云测试"             //短信签名名称
	request.TemplateCode = "SMS_245930093" //短信模板ID

	_, err = client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	err = driver.RDB.Set("name", "刘德华", 0).Err()
	if err != nil {
		zap.S().Errorf("设置失败%s", err.Error())
	}
	val, err := driver.RDB.Get("name").Result()
	if err != nil {
		zap.S().Errorf("获取redis失败%s", err.Error())
	}
	fmt.Printf(val)
}
