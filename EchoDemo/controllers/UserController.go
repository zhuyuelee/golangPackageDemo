package controllers

import (
	"GoSql/EchoDemo/dtos"
	"GoSql/EchoDemo/servers"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/labstack/echo"
)

// GetUser 获取用户
func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	fmt.Println("GetUser parm=", id)
	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "参数类型错误"))
	}

	userDto, err := servers.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "获取数据失败"))
	}
	return c.JSON(http.StatusOK, dtos.SuccessDataResult(userDto))
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
	fmt.Println("reflect", reflect.ValueOf(userDto))

	fmt.Println("get parms", userDto)
	err = servers.UpdateUser(&userDto)

	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "添加失败"))
	}
	return c.JSON(http.StatusOK, dtos.SuccessDataResult(userDto))
}

// DeleteUser 删除用户
func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	fmt.Println("DeleteUser parm=", id)
	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "参数类型错误"))
	}

	err = servers.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "删除数据失败"))
	}
	return c.JSON(http.StatusOK, dtos.SuccessResult())
}
