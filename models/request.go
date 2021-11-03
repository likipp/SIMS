package models

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	ID         uint
	Username   string
	NickName   string
	BufferTime int64
	jwt.StandardClaims
}
