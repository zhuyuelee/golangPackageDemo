package controllers

import (
	"GoSql/EchoDemo/dtos"
	"GoSql/EchoDemo/servers"
	"GoSql/EchoDemo/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//Login 登录
func Login(c echo.Context) error {
	input := new(dtos.LoginInput)
	err := c.Bind(input)
	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "参数错误"))
	}
	result, err := servers.Login(input)
	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "获取数据失败"))
	}
	return c.JSON(http.StatusOK, dtos.SuccessDataResult(result))
}

// GetUserList 获取用户列表
func GetUserList(c echo.Context) error {
	var input = new(dtos.PageInput)
	err := c.Bind(input)

	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "参数错误"))
	}

	list, err := servers.GetUserList(input)
	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "获取数据失败"))
	}
	return c.JSON(http.StatusOK, dtos.SuccessDataResult(list))
}

// GetUser 获取用户
func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "参数类型错误"))
	}

	userDto, err := servers.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "获取数据失败"))
	}
	return c.JSON(http.StatusOK, dtos.SuccessDataResult(userDto))
}

//Welcome 欢迎
func Welcome(c echo.Context) error {
	claims, err := utils.GetToken(c)
	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "获取数据失败"))
	}
	return c.JSON(http.StatusOK, dtos.SuccessDataResult(claims))
}

// AddUser 新加用户
func AddUser(c echo.Context) error {
	var userDto = new(dtos.UserDto)
	err := c.Bind(userDto)
	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "参数错误"))
	}
	err = servers.AddUser(userDto)

	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "添加失败"))
	}
	return c.JSON(http.StatusOK, dtos.SuccessDataResult(userDto))
}

// UpdateUser 修改用户
func UpdateUser(c echo.Context) error {
	var userDto = dtos.UserDto{}
	err := c.Bind(&userDto)
	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "参数错误"))
	}
	err = servers.UpdateUser(&userDto)

	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "添加失败"))
	}
	return c.JSON(http.StatusOK, dtos.SuccessDataResult(userDto))
}

// DeleteUser 删除用户
func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "参数类型错误"))
	}

	err = servers.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "删除数据失败"))
	}
	return c.JSON(http.StatusOK, dtos.SuccessResult())
}
