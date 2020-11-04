package servers

import (
	"GoSql/EchoDemo/dao"
	"GoSql/EchoDemo/dtos"
	"GoSql/EchoDemo/mapper"
	"GoSql/EchoDemo/models"
)

// GetUserList 获取用户信息列表
func GetUserList(input *dtos.PageInput) (list []dtos.UserDto, err error) {
	users, err := dao.GetUserList(input)
	if err != nil {
		return nil, err
	}
	list = make([]dtos.UserDto, input.Limit)
	err = mapper.Map(users, &list)
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
