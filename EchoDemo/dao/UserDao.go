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
	if err != nil {
		return
	}
	if input.Page == 0 {
		input.Page = 1
	}

	if input.Limit == 0 {
		input.Limit = 10
	}
	result := db.Debug().Where("user_name like '%?%'", input.Key).Limit(input.Limit).Offset((input.Page - 1) * input.Limit).Find(list)
	if result.Error != nil {
		err = result.Error
		fmt.Println("GetUserList error=", err)
	}
	fmt.Println("GetUserList list=", list)
	return
}

//GetUser 获取用户信息
func GetUser(id int) (user models.User, err error) {
	db, err := data.DbHelper()
	defer db.Close()
	if err != nil {
		return
	}
	result := db.Debug().First(&user, id)
	if result.Error != nil {
		err = result.Error
		fmt.Println("GetUser error=", err)
	}
	fmt.Println("GetUser user=", user)
	return
}

//AddUser 获取用户信息
func AddUser(user *models.User) (err error) {
	db, err := data.DbHelper()
	defer db.Close()
	if err != nil {
		return
	}
	db.Debug().Create(user)
	return
}

// UpdateUser 修改用户
func UpdateUser(user *models.User) (err error) {
	db, err := data.DbHelper()
	defer db.Close()
	if err != nil {
		return
	}
	db.Debug().Model(&models.User{}).Where("id=?", user.ID).Update(user)
	return
}

//DeleteUser 删除用户
func DeleteUser(id int) (err error) {
	db, err := data.DbHelper()
	defer db.Close()
	if err != nil {
		return
	}
	result := db.Debug().Delete(&models.User{}, id)
	if result.Error != nil {
		err = result.Error
		fmt.Println("DeleteUser error=", err)
	}
	fmt.Println("DeleteUser user=", result.RowsAffected)
	return
}
