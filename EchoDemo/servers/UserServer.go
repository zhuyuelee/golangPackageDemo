package servers

import (
	"GoSql/EchoDemo/dao"
	"GoSql/EchoDemo/dto"
	"GoSql/EchoDemo/models"

	"github.com/devfeel/mapper"
)

func init() {
	mapper.Register(&models.User{})
	mapper.Register(&dto.UserDto{})
}

// GetUser 获取用户信息
func GetUser(id int) (userDto *dto.UserDto, err error) {
	user, err := dao.GetUser(id)
	if err != nil {
		userDto = nil
		return
	}
	userDto = new(dto.UserDto)
	mapper.AutoMapper(&user, userDto)
	return
}

// AddUser 新加用户
func AddUser(userDto *dto.UserDto) (err error) {
	user := &models.User{}
	err = mapper.AutoMapper(userDto, user)
	if err != nil {
		return
	}
	err = dao.AddUser(user)
	if err != nil {
		return
	}
	userDto.ID = user.ID
	userDto.CreatedAt = user.CreatedAt
	userDto.UpdatedAt = user.UpdatedAt
	return
}
