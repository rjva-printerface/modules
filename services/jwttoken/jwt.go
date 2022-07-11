package jwttoken

import (
	"os"
	"time"

	"github.com/form3tech-oss/jwt-go"
	"github.com/rjva-printerface/auth-service-go/helpers"
	"github.com/rjva-printerface/auth-service-go/models"
)

type JwtToken struct {
	log *helpers.Log
}

func NewJWT(log *helpers.Log) *JwtToken {
	return &JwtToken{log: log}
}

type Claims struct {
	User *models.UserModel
	jwt.StandardClaims
}

var jwtKey = os.Getenv("JWT_KEY")

func getClaims(usr *models.UserModel) Claims {
	return Claims{
		User: usr,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Now().AddDate(0, 0, 1).Unix()),
		},
	}
}

func (jwtToken *JwtToken) GenerateToken(usr *models.UserModel) string {
	claims := getClaims(usr)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		jwtToken.log.Print(err.Error(), helpers.Red)
		return ""
	}

	return tokenString
}

func (jwtToken *JwtToken) DecodeToken(tokenStr []byte, usr *models.UserModel) *models.UserModel {
	claims := getClaims(usr)
	_, err := jwt.ParseWithClaims(string(tokenStr), &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})

	if err != nil {
		jwtToken.log.Print(err.Error(), helpers.Red)
	}

	return usr
}
