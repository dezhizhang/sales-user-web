package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"sales-user-web/global"
	"sales-user-web/model"
	"sales-user-web/utils"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("token过期了")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInValid     = errors.New("Could handle this token")
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		fmt.Println(c.Request.Header.Get("path"))
		if token == "" {
			utils.ResponseErrorJson(c, http.StatusUnauthorized, "未登录")

			c.Abort()
			return
		}
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusUnauthorized, gin.H{
					"msg":     "授权已过期",
					"code":    400,
					"data":    nil,
					"success": false,
				})
				c.Abort()
				return
			} else if err == TokenNotValidYet {
				c.JSON(http.StatusUnauthorized, gin.H{
					"msg":     "未登录",
					"code":    400,
					"data":    nil,
					"success": false,
				})
			}
		}

		c.Set("name", claims.Name)
		c.Set("userId", claims.Id)
		c.Next()

	}
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.ServerConfig.JwtInfo.SigningKey),
	}
}

// 生成token

func (that *JWT) CreateToken(claims model.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(that.SigningKey)
}

// 解析token

func (that *JWT) ParseToken(tokenString string) (*model.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return that.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInValid
			}
		}
	}

	if token != nil {
		if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInValid
	}

	return nil, TokenInValid
}
