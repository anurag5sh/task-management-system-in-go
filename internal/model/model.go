package model

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	Uid int64
	jwt.MapClaims
}
