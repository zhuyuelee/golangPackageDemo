package dtos

//UserDto 用户DTO
type UserDto struct {
	BaseDto `mapper:"model"`
	// UserName 用户账号
	UserName string `json:"userName" mapper:"user_name"`
	Password string `json:"password" mapper:"password"`
}
