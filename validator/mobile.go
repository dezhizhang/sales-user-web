package validator

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"regexp"
)

func ValidateMobile(f1 validator.FieldLevel) bool {
	mobile := f1.Field().String()
	ok, err := regexp.MatchString(`^1(3[0-9]|4[01456879]|5[0-35-9]|6[2567]|7[0-8]|8[0-9]|9[0-35-9])\d{8}$`, mobile)
	if err != nil {
		zap.S().Errorw("验证手机号正则失败")
	}
	return !!ok
}
