package data

import (
	"GoSql/EchoDemo/models"
	"fmt"

	"github.com/jinzhu/gorm"

	// Register sqlite3 驱动
	_ "github.com/mattn/go-sqlite3"
)

// DbHelper returns gorm.DB对象
func DbHelper() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "./data/data.db")
	if err != nil {
		fmt.Println("database open error", err)
		return nil, err
	}
	fmt.Println("database Open")
	return db, err
}

//初始化数据库
func init() {
	db, err := DbHelper()
	defer db.Close()
	if err != nil {
		return
	}
	//初始化user表
	db.AutoMigrate(models.User{})
	fmt.Println("database init")
}
