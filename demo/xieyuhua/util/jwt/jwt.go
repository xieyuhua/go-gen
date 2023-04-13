/*
 * @Descripttion: 
 * @version: 
 * @Author: seaslog
 * @Date: 2022-02-08 10:43:56
 * @LastEditors: 谢余华
 * @LastEditTime: 2022-02-08 14:11:18
 */
package jwt

import (
	"errors"
	"time"
	"github.com/spf13/viper"
	jwt "github.com/dgrijalva/jwt-go"
)
// 一些常量
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "newtrekWang"
)

var jwtSecret = []byte(viper.GetString("jwt.JwtSecret"))

type Claims struct {
	UserID    int    `json:"user_id"`
	UserEmail string `json:"user_email"`
	jwt.StandardClaims
}


func GenerateToken(id int, email string) (string, error) {
	JwtTokenExpire := viper.GetInt64("jwt.JwtTokenExpire")
	JwtIssuer := viper.GetString("jwt.JwtIssuer")
	claims := Claims{
		id,
		email,
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(0 * time.Second).Unix()) + JwtTokenExpire,
			Issuer:    JwtIssuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ParseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
