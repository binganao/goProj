package common

import (
	"mall/global"

	"github.com/golang-jwt/jwt"
)

var SigningKey = []byte(global.Config.Jwt.SigningKey)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string)(string, error){
	claims:=Claims{username, jwt.StandardClaims{
		ExpiresAt: ,
	}}
}