package dao

import (
	"GoSql/EchoDemo/data"
	"GoSql/EchoDemo/dtos"
	"GoSql/EchoDemo/models"
	"GoSql/EchoDemo/utils"
	"crypto/md5"
	"errors"
	"fmt"
)

func getPassword(pwd string) (password string) {
	salt, _ := utils.GetSalt()
	password = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s_%s", pwd, salt))))
	return
}

//Login 登录
func Login(input *dtos.LoginInput) (user models.User, err error) {
	db := data.DbHelper()
	defer db.Close()
	pwd := getPassword(input.Password)
	fmt.Printf("pwd:=%s \n", pwd)
	result := db.Debug().Where("user_name=? and password=?", input.UserName, pwd).Find(&user)
	if result.RowsAffected == 0 {
		err = errors.New("用户不存在")
	} else if result.Error != nil {
		err = result.Error
	}
	fmt.Printf("result:=%+v \n", result)
	return
}

//GetUserList 获取用户信息
func GetUserList(input *dtos.PageInput) (list []models.User, err error) {
	db := data.DbHelper()
	defer db.Close()
	if input.Page == 0 {
		input.Page = 1
	}
	if input.Limit == 0 {
		input.Limit = 10
	}
	list = make([]models.User, input.Limit)
	result := db.Debug()
	if input.Key != "" {
		result = result.Where("user_name like ?", fmt.Sprintf("%%%s%%", input.Key))
	}
	result = result.Limit(input.Limit).Offset((input.Page - 1) * input.Limit).Find(&list)

	if result.Error != nil {
		err = result.Error
	}
	return
}

//GetUser 获取用户信息
func GetUser(id int) (user models.User, err error) {
	db := data.DbHelper()
	defer db.Close()
	result := db.Debug().First(&user, id)
	if result.Error != nil {
		err = result.Error
	}
	return
}

//AddUser 获取用户信息
func AddUser(user *models.User) (err error) {
	db := data.DbHelper()
	defer db.Close()
	user.Password = getPassword(user.Password)
	result := db.Debug().Create(user)
	if result.Error != nil {
		err = result.Error
	}
	return
}

// UpdateUser 修改用户
func UpdateUser(user *models.User) (err error) {
	db := data.DbHelper()
	defer db.Close()
	if user.Password != "" {
		user.Password = getPassword(user.Password)
	}
	result := db.Debug().Model(&models.User{}).Update(user)
	if result.Error != nil {
		err = result.Error
	} else if result.RowsAffected == 0 {
		err = errors.New("数据不存在")
	}
	fmt.Printf("result:=%+v \n", result)
	return
}

//DeleteUser 删除用户
func DeleteUser(id int) (err error) {
	db := data.DbHelper()
	defer db.Close()
	result := db.Debug().Delete(&models.User{}, id)
	if result.Error != nil {
		err = result.Error
	}
	return
}
