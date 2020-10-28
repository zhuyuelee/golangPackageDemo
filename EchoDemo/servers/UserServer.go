package servers

import (
	"GoSql/EchoDemo/dao"
	"GoSql/EchoDemo/dto"
)

// GetUser 获取用户信息
// OutPut 会员DTO
func GetUser(id int) (dto dto.UserDto) {
	user, err := dao.GetUser(id)

	if err != nil {
		return
	}
	dto.CreatedAt = user.CreatedAt.Time
	dto.ID = user.ID
	dto.UserName = user.UserName
	dto.Password = user.Password
	return
}
