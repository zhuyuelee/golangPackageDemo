package servers

import (
	"GoSql/EchoDemo/dao"
	"GoSql/EchoDemo/dtos"
	"GoSql/EchoDemo/mapper"
	"GoSql/EchoDemo/models"
	"fmt"
)

// GetUser 获取用户信息
func GetUser(id int) (userDto *dtos.UserDto, err error) {
	user, err := dao.GetUser(id)
	if err != nil {
		return nil, err
	}
	userDto = new(dtos.UserDto)
	mapper.Mapper(user, userDto)
	return
}

// AddUser 新加用户
func AddUser(userDto *dtos.UserDto) (err error) {
	user := &models.User{}

	mapper.Mapper(userDto, user)

	err = dao.AddUser(user)
	if err != nil {
		return
	}
	userDto.ID = user.ID
	userDto.CreatedAt = user.CreatedAt
	userDto.UpdatedAt = user.UpdatedAt
	return
}

// UpdateUser 修改用户
func UpdateUser(userDto *dtos.UserDto) (err error) {
	user := &models.User{}
	fmt.Println("UserDto", userDto)
	mapper.Mapper(userDto, user)
	fmt.Println("UpdateUser", user)
	if err != nil {
		return
	}
	err = dao.UpdateUser(user)
	if err != nil {
		return
	}
	return
}

// DeleteUser 获取用户信息
func DeleteUser(id int) (err error) {
	err = dao.DeleteUser(id)
	return
}
