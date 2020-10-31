package mapper

import (
	"fmt"
	"reflect"
)

//Mapper struct map to struct
func Mapper(source interface{}, to interface{}) {
	toStruct(toMap(source), to)
}

//toMap struct to map
func toMap(obj interface{}) map[string]reflect.Value {
	m := make(map[string]reflect.Value)
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	v := reflect.ValueOf(obj)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag, ok := field.Tag.Lookup("mapper")
		if !ok {
			tag = field.Name
		}
		// if field.Type.Kind() == reflect.Struct {
		// 	value := toMap(v.Field(i).Interface())
		// 	m[tag] = reflect.ValueOf(value)

		// } else {
		// 	m[tag] = v.Field(i)
		// }
		m[tag] = v.Field(i)
		fmt.Printf("value:%v\n", v.Field(i).Interface())

	}
	return m
}

func toStruct(m map[string]reflect.Value, obj interface{}) {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		valueField := v.Field(i)
		field := t.Field(i)
		fmt.Printf("Field:%d\n", i)
		var name string
		var ok bool
		if name, ok = field.Tag.Lookup("mapper"); !ok {
			name = field.Name
		}
		fmt.Printf("mapvalue %v \n", m[name])
		// valueField.Set(m[name])

		if value, ok := m[name]; ok {
			// if field.Type.Kind() == reflect.Struct {
			// 	value := toMap(v.Field(i).Interface())
			// 	m[tag] = reflect.ValueOf(value)

			// } else {

			// 	valueField.Set(value)
			// 	fmt.Printf("name:%s value:%v \n", name, value)
			// }

			valueField.Set(value)
			fmt.Printf("name:%s value:%v \n", name, value)
		}
	}
	return
}
