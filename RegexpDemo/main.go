package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/zhuyuelee/mapper"
)

//User User
type User struct {
	gorm.Model
	UserName string `gorm:"column=user_name;type=20"`
	Password string `gorm:"type=20"`
	Age      int
	FullName string
}

func main() {
	columns := getColumns(reflect.TypeOf(User{}))

	fmt.Println("colmous=", columns)
}

//getColumns 根据gorm规则，获取实体的字段
func getColumns(types reflect.Type) []string {
	if types.Kind() == reflect.Ptr {
		types = types.Elem()
	}
	columns := make([]string, 0)
	for i := 0; i < types.NumField(); i++ {
		t := types.Field(i)
		if t.Type.Kind() == reflect.Struct && t.Type.Name() != "Time" {
			columns = append(columns, getColumns(t.Type)...)
		} else {
			tag, ok := t.Tag.Lookup("gorm")
			if ok && strings.Index(tag, "column=") > -1 {
				tag = getTag(tag)
			} else {
				tag = getFieldName(t.Name)
			}
			//排除删除和创建字段
			if !strings.Contains("deleted_at|created_at", tag) {
				columns = append(columns, tag)
			}
		}
	}
	return columns
}

//getTag 获取tag中的字段名
func getTag(tag string) string {
	re := regexp.MustCompile(`column=([\w]+);?`)
	tags := re.FindAllStringSubmatch(tag, 1)
	if len(tags) > 0 {
		tag = tags[0][1]
	}
	return tag
}

//getTag 获取tag中的字段名
func getFieldName(name string) string {
	chars := make([]rune, 0)
	upper := false
	for index, c := range name {
		if c >= 'A' && c <= 'Z' {
			if index > 0 && !upper {
				chars = append(chars, '_')
			}
			chars = append(chars, c+32)
			upper = true
		} else {
			chars = append(chars, c)
			upper = false
		}
	}

	source := make([]User, 0)

	to := make([]User, 0)

	mapper.Map(source, &to)

	return string(chars)
}
