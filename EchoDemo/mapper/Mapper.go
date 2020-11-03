package mapper

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

var zeroValue reflect.Value

const tag string = "mapper"

func init() {
	zeroValue = reflect.Value{}
}

//Mapper struct map to struct
func Mapper(source interface{}, to interface{}) error {
	startTime := time.Now()
	defer func() {
		fmt.Println("mapper time=", time.Now().Sub(startTime).Seconds())
	}()

	return mapperToStruct(source, reflect.ValueOf(to))
}

//mapperToStruct struct map to struct
func mapperToStruct(source interface{}, to reflect.Value) error {
	sourceMap, err := toMap(source)
	if err != nil {
		fmt.Println("mapper error=", err)
	} else {
		toStruct(sourceMap, to)
	}
	return err
}

//toMap struct to map
func toMap(source interface{}) (map[string]reflect.Value, error) {
	m := make(map[string]reflect.Value)
	value := reflect.ValueOf(source)
	t := value.Type()
	if t.Kind() == reflect.Ptr {
		value = value.Elem()
		t = t.Elem()
	}
	if value == zeroValue {
		return nil, errors.New("no exists this value")
	}
	for i := 0; i < value.NumField(); i++ {
		field := t.Field(i)
		tag, ok := field.Tag.Lookup(tag)
		if !ok {
			tag = field.Name
		}
		m[tag] = value.Field(i)
	}
	return m, nil
}

func toStruct(m map[string]reflect.Value, to reflect.Value) {

	if to.Kind() == reflect.Ptr {
		to = to.Elem()
	}
	for i := 0; i < to.NumField(); i++ {
		vField := to.Field(i)
		tField := to.Type().Field(i)
		var (
			name string
			ok   bool
		)
		if name, ok = tField.Tag.Lookup("mapper"); !ok {
			name = tField.Name
		}
		if value, ok := m[name]; ok {
			sType := value.Type()
			vType := tField.Type

			if sType.Kind() == reflect.Ptr {
				sType = sType.Elem()
			}
			if vType.Kind() == reflect.Ptr {
				vType = vType.Elem()
			}
			if sType.Kind() == reflect.Struct && sType != vType {
				nValue := reflect.New(vType)
				mapperToStruct(value.Interface(), nValue)
				vField.Set(nValue.Elem())
			} else {
				vField.Set(value)
			}

		}
	}
}
