package dtos

import "github.com/dgrijalva/jwt-go"

// LoginInput 登录
type LoginInput struct {
	// UserName 用户名
	UserName string `json:"userName" validate:"required,min=6,max=10"`
	// Password 密码
	Password string `json:"password" validate:"required,min=6,max=10"`
}

//TokenDto JWTToken
type TokenDto struct {
	ID       uint
	UserName string `json:"userName"`
	Token    string `json:"token"`
}

//ClaimsDto Claims
type ClaimsDto struct {
	ID       uint
	UserName string
	jwt.StandardClaims
}
