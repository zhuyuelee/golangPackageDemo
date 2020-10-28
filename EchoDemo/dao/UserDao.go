package dao

import (
	"GoSql/EchoDemo/data"
	"GoSql/EchoDemo/models"
	"fmt"
)

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
	db.Create(user)
	return
}
