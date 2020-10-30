package main

import (
	"fmt"
	"reflect"
	"time"
)

type (
	//Student 学生
	Student struct {
		Name    string    `mapper:"name"`
		Age     int       `mapper:"age"`
		Created time.Time `mapper:"created"`
	}
	//Teacher 老师
	Teacher struct {
		Name    string
		Age     int
		Created time.Time
	}
)

func main() {
	stu := Student{
		Name:    "张三",
		Age:     19,
		Created: time.Now(),
	}
	t := reflect.TypeOf(stu)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("name:%s tag:%v index:%v path:%s \n", f.Name, f.Tag, f.Index, f.PkgPath)
	}
	fmt.Println("------------------------")
	v := reflect.ValueOf(stu)
	for i := 0; i < v.NumField(); i++ {
		fmt.Println("filed=", v.Field(i))
	}
}
