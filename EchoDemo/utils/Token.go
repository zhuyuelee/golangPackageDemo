package utils

import "GoSql/EchoDemo/dtos"

//JWTTokenSecret 密钥
const JWTTokenSecret string = "GoSql_Demo"

//GetToken 获取Token
func GetToken(userDto *dtos.UserDto) (token string, err error) {
	return "", nil
}
