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
	// fmt.Println("------------------------")
	//t := reflect.TypeOf(stu)
	// for i := 0; i < t.NumField(); i++ {
	// 	f := t.Field(i)
	// 	fmt.Printf("name:%s tag:%v index:%v path:%s \n", f.Name, f.Tag.Get("mapper"), f.Index, f.PkgPath)
	// }

	fmt.Printf("name:%s age:%d created:%v\n", stu.Name, stu.Age, stu.Created)
	fmt.Println("------------------------")
	v := reflect.ValueOf(&stu).Elem()
	time.Sleep(3 * time.Second)
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		//fmt.Printf("Kind=%v&Name=%v\n", f.Interface())
		switch f.Kind() {
		case reflect.Int:
			f.SetInt(12)
		case reflect.String:
			f.SetString("改名")
		case reflect.Struct:
			if f.Type().Name() == "Time" {
				time, ok := time.Parse("2006-01-02 15:04:05", "2010-12-12 11:02:11")
				if ok != nil {
					fmt.Println("error=", ok)
				}
				f.Set(reflect.ValueOf(time))
			}
		}

	}
	fmt.Printf("name:%s age:%d created:%v\n", stu.Name, stu.Age, stu.Created)
	// e := reflect.ValueOf(&stu).Elem()
	// fmt.Println("------------------------")
	// for i := 0; i < e.NumField(); i++ {
	// 	f := e.Field(i)
	// 	fmt.Println("filed=", f, "\nvalue=", f.Interface())
	// 	t := f.Type()
	// 	fmt.Printf("Kind=%v&key=%v\n", t.Kind(), t.Name())
	// }

}
