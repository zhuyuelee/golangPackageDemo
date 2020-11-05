package data

import (
	"GoSql/EchoDemo/models"
	"GoSql/EchoDemo/utils"
	"fmt"

	"github.com/jinzhu/gorm"

	// Register sqlite3 驱动
	_ "github.com/mattn/go-sqlite3"
)

// DbHelper returns gorm.DB对象
func DbHelper() *gorm.DB {
	datapath, err := utils.GetDatabaseConfig()
	if err != nil {
		panic(fmt.Sprintf("database path config erro err:%v", err))
	}
	db, err := gorm.Open("sqlite3", datapath)
	if err != nil {
		panic(fmt.Sprintf("database open error err:%v", err))
	}
	return db
}

//初始化数据库
func init() {
	db := DbHelper()
	defer db.Close()

	//初始化user表
	db.AutoMigrate(models.User{})
	fmt.Println("database init")
}
