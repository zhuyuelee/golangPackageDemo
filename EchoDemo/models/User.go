package models

import "github.com/jinzhu/gorm"

// User 会员
type User struct {
	gorm.Model `mapper:"model"`
	UserName   string `mapper:"user_name"`
	Password   string `mapper:"password"`
}
