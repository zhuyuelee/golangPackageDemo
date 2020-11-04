package servers

import (
	"GoSql/EchoDemo/dao"
	"GoSql/EchoDemo/dtos"
	"GoSql/EchoDemo/mapper"
	"GoSql/EchoDemo/models"
	"GoSql/EchoDemo/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//Login 登录
func Login(input *dtos.LoginInput) (token *dtos.TokenDto, err error) {
	user, err := dao.Login(input)
	if err == nil {
		// Set custom claims
		claims := struct {
			UserName string
			jwt.StandardClaims
		}{
			user.UserName,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		// Create token with claims
		tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
		token = new(dtos.TokenDto)
		// Generate encoded token and send it as response.
		token.Token, err = tokenClaims.SignedString([]byte(utils.JWTTokenSecret))
	}
	return
}

// GetUserList 获取用户信息列表
func GetUserList(input *dtos.PageInput) (list []dtos.UserDto, err error) {
	users, err := dao.GetUserList(input)
	if err != nil {
		list = make([]dtos.UserDto, input.Limit)
		err = mapper.Map(users, &list)
	}

	return
}

// GetUser 获取用户信息
func GetUser(id int) (userDto *dtos.UserDto, err error) {
	user, err := dao.GetUser(id)
	if err == nil {
		userDto = new(dtos.UserDto)
		err = mapper.Map(user, userDto)
	}
	return
}

// AddUser 新加用户
func AddUser(userDto *dtos.UserDto) (err error) {
	user := &models.User{}
	err = mapper.Map(userDto, user)
	if err != nil {
		err = dao.AddUser(user)
		if err == nil {
			err = mapper.Map(user, userDto)
		}
	}
	return
}

// UpdateUser 修改用户
func UpdateUser(userDto *dtos.UserDto) (err error) {
	user := &models.User{}
	mapper.Map(userDto, user)
	if err == nil {
		err = dao.UpdateUser(user)
		if err == nil {
			err = mapper.Map(user, userDto)
		}
	}

	return
}

// DeleteUser 获取用户信息
func DeleteUser(id int) (err error) {
	err = dao.DeleteUser(id)
	return
}
