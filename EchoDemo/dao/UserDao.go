package dao

import (
	"GoSql/EchoDemo/data"
	"GoSql/EchoDemo/dtos"
	"GoSql/EchoDemo/models"
	"fmt"
)

//GetUserList 获取用户信息
func GetUserList(input *dtos.PageInput) (list []models.User, err error) {
	db, err := data.DbHelper()
	defer db.Close()
	if err == nil {
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
	}
	return
}

//GetUser 获取用户信息
func GetUser(id int) (user models.User, err error) {
	db, err := data.DbHelper()
	defer db.Close()
	if err == nil {
		result := db.Debug().First(&user, id)
		if result.Error != nil {
			err = result.Error
		}
	}
	return
}

//AddUser 获取用户信息
func AddUser(user *models.User) (err error) {
	db, err := data.DbHelper()
	defer db.Close()
	if err == nil {
		result := db.Debug().Create(user)
		if result.Error != nil {
			err = result.Error
		}
	}
	return
}

// UpdateUser 修改用户
func UpdateUser(user *models.User) (err error) {
	db, err := data.DbHelper()
	defer db.Close()
	if err == nil {
		result := db.Debug().Model(&models.User{}).Update(user)
		if result.Error != nil {
			err = result.Error
		}
	}
	return
}

//DeleteUser 删除用户
func DeleteUser(id int) (err error) {
	db, err := data.DbHelper()
	defer db.Close()
	if err == nil {
		result := db.Debug().Delete(&models.User{}, id)
		if result.Error != nil {
			err = result.Error
		}
	}
	return
}
