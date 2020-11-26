package controllers

import (
	"GoSql/EchoDemo/dtos"
	"GoSql/EchoDemo/servers"
	"GoSql/EchoDemo/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/zhuyuelee/mapper"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

//Login 登录
func Login(c echo.Context) error {
	input := new(dtos.LoginInput)
	err := c.Bind(input)
	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "参数错误"))
	}

	if err = validate.Struct(input); err != nil {

		// if _, ok := err.(*validator.InvalidValidationError); ok {
		// 	fmt.Println(err)
		// 	return err
		// }
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, err.Error()))
	}

	result, err := servers.Login(input)
	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, err.Error()))
	}
	return c.JSON(http.StatusOK, dtos.SuccessDataResult(result))
}

//Register 注册
func Register(c echo.Context) error {
	input := new(dtos.LoginInput)
	err := c.Bind(input)
	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "参数错误"))
	}

	if err = validate.Struct(input); err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return c.JSON(http.StatusOK, dtos.ErrorResult(1, "注册数据验证失败"))
		}

		return c.JSON(http.StatusOK, dtos.ErrorResult(1, fmt.Sprintf("注册数据验证失败 %v", err)))
	}

	userDto := &dtos.UserDto{}
	mapper.Map(input, userDto)

	err = servers.AddUser(userDto)
	if err != nil {
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, "会员注册失败"))
	}
	return c.JSON(http.StatusOK, dtos.SuccessDataResult(userDto))
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
		return c.JSON(http.StatusOK, dtos.ErrorResult(1, err.Error()))
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
