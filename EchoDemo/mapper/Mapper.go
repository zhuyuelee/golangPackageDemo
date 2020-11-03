package mapper

import (
	"errors"
	"fmt"
	"reflect"
)

var zeroValue reflect.Value

const tag string = "mapper"

func init() {
	zeroValue = reflect.Value{}

}

//Mapper struct map to struct
func Mapper(source interface{}, to interface{}) error {
	sourceMap, err := toMap(source)
	if err != nil {
		fmt.Println("mapper error=", err)
		return err
	}
	toStruct(sourceMap, to)
	return nil
}

//toMap struct to map
func toMap(obj interface{}) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}
	if v == zeroValue {
		return nil, errors.New("no exists this value")
	}
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag, ok := field.Tag.Lookup(tag)
		if !ok {
			tag = field.Name
		}
		value := v.Field(i)
		m[tag] = value.Interface()
	}
	return m, nil
}

func toStruct(m map[string]interface{}, obj interface{}) {
	var v = reflect.Value{}
	var ok bool
	if v, ok = obj.(reflect.Value); !ok {
		v = reflect.ValueOf(obj)
	}
	t := v.Type()
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = t.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		vField := v.Field(i)
		tField := t.Field(i)
		var name string
		if name, ok = tField.Tag.Lookup("mapper"); !ok {
			name = tField.Name
		}
		if value, ok := m[name]; ok {
			sType := reflect.TypeOf(value)
			vType := tField.Type

			if sType.Kind() == reflect.Ptr {
				sType = sType.Elem()
			}
			if vType.Kind() == reflect.Ptr {
				vType = vType.Elem()
			}
			if sType.Kind() == reflect.Struct && sType != vType {
				nValue := reflect.New(vType)
				fmt.Println("vType=", vType)

				Mapper(value, nValue)
				vField.Set(nValue.Elem())
			} else {
				vField.Set(reflect.ValueOf(value))
			}

		}
	}
}
