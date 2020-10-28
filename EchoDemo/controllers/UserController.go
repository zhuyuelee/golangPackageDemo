package controllers

import (
	"GoSql/EchoDemo/dao"
	"GoSql/EchoDemo/dto"
	"GoSql/EchoDemo/models"
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
	user, err := dao.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusOK, dto.ErrorResult(1, "获取数据失败"))
	}
	userDto := dto.UserDto{
		BaseDto: dto.BaseDto{
			UpdatedAt: user.UpdatedAt,
			CreatedAt: user.CreatedAt,
			ID:        user.ID,
		},
		UserName: user.UserName,
		Password: user.Password,
	}

	return c.JSON(http.StatusOK, dto.SuccessResult(userDto))
}

// AddUser 新加用户
func AddUser(c echo.Context) error {
	var user = new(dto.UserDto)
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusOK, dto.ErrorResult(1, "参数错误"))
	}
	var model = new(models.User)
	model.UserName = user.UserName
	model.Password = user.Password
	err = dao.AddUser(model)
	if err != nil {
		return c.JSON(http.StatusOK, dto.ErrorResult(1, "添加失败"))
	}
	return c.JSON(http.StatusOK, dto.SuccessResult(model))
}
