package mapper

import (
	"fmt"
	"reflect"
)

//Mapper struct map to struct
func Mapper(source interface{}, to *interface{}) {
	suroceMap := toMap(source)
	//toMap := toMap(to)
	fmt.Println(suroceMap)
}

//toMap struct to map
func toMap(obj interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tag, ok := field.Tag.Lookup("mapper"); ok {
			m[tag] = v.Field(i).Interface()
		} else {
			m[field.Name] = v.Field(i).Interface()
		}
	}
	return m
}

// func toStruct(m map[string]interface{}, obj *interface{}) {
// 	v := reflect.ValueOf(obj)
// 	t := reflect.TypeOf(obj)
// 	for i := 0; i < v.NumField(); i++ {
// 		field := v.Field(i)
// 		if tag, ok := t.Field(i).Tag.Lookup("mapper"); ok {
// 			field.Set(m[tag])
// 		} else {
// 			field.Set(m[field.Kind()])
// 		}
// 	}
// 	return m
// }
