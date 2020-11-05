package utils

import (
	"GoSql/EchoDemo/dtos"
	"GoSql/EchoDemo/models"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//tokenSecret 密钥
const tokenSecret string = "GoSql_Demo"

//GetToken 获取Token
func GetToken(c echo.Context) (token *dtos.ClaimsDto, err error) {
	user := c.Get("user")
	if user == nil {
		err = errors.New("未登录")
		return
	}
	token = user.(*jwt.Token).Claims.(*dtos.ClaimsDto)
	return
}

//CreateToken 生成Token
func CreateToken(user *models.User) (token *dtos.TokenDto, err error) {
	// Set custom claims
	claims := &dtos.ClaimsDto{
		ID:       user.ID,
		UserName: user.UserName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	// Create token with claims
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token = new(dtos.TokenDto)
	// Generate encoded token and send it as response.
	token.Token, err = tokenClaims.SignedString([]byte(tokenSecret))
	if err == nil {
		token.UserName = claims.UserName
		token.ID = claims.ID
	}
	return
}

// JWTConfig JWT auth middleware with config.
func JWTConfig() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &dtos.ClaimsDto{},
		SigningKey: []byte(tokenSecret),
		// ErrorHandler: func(e error) error {
		// 	return dtos.APIResult{
		// 		ErrorCode: http.StatusUnauthorized,
		// 		Message:   e.Error(),
		// 	}
		// },
	})
}
