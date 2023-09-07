package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	chTranslations "github.com/go-playground/validator/v10/translations/zh"
	"sales-user-web/global"
)

func InitTrans(local string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zhT := zh.New() //chinese
		enT := en.New() //english
		uni := ut.New(enT, zhT, enT)

		global.Trans, ok = uni.GetTranslator(local)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", local)
		}
		// 注册翻译器
		switch local {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			err = chTranslations.RegisterDefaultTranslations(v, global.Trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, global.Trans)
		}
		return
	}
	return
}
