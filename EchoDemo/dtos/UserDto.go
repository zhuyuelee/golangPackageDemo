package dtos

//UserDto 用户DTO
type UserDto struct {
	BaseDto
	// UserName 用户账号
	UserName string `json:"userName"`
	Password string `json:"password"`
}
