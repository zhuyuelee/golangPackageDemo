package controllers

import (
	"GoSql/EchoDemo/dto"
	"GoSql/EchoDemo/servers"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// GetUser 获取用户
func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	fmt.Println("GetUser parm=", id)
	if err != nil {
		return c.JSON(http.StatusOK, dto.ErrorResult(1, "参数类型错误"))
	}

	userDto, err := servers.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusOK, dto.ErrorResult(1, "获取数据失败"))
	}
	return c.JSON(http.StatusOK, dto.SuccessResult(userDto))
}

// AddUser 新加用户
func AddUser(c echo.Context) error {
	var userDto = new(dto.UserDto)
	err := c.Bind(userDto)
	if err != nil {
		return c.JSON(http.StatusOK, dto.ErrorResult(1, "参数错误"))
	}
	err = servers.AddUser(userDto)

	if err != nil {
		return c.JSON(http.StatusOK, dto.ErrorResult(1, "添加失败"))
	}
	return c.JSON(http.StatusOK, dto.SuccessResult(userDto))
}
