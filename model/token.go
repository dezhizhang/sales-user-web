package model

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}
