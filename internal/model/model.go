package model

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	UserId int64
	jwt.MapClaims
}
