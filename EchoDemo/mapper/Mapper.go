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

//Map to struct or Slice
func Map(source, target interface{}) (err error) {
	startTime := time.Now()
	defer func() {
		fmt.Println("mapper time=", time.Now().Sub(startTime).Seconds())
	}()
	defer func() {
		if mapErr := recover(); mapErr != nil {
			err = fmt.Errorf("map error %v", mapErr)
		}
	}()

	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr {
		return errors.New("target must pointer")
	}
	if targetValue.IsNil() {
		return errors.New("target must not nil")
	}

	var sources = reflect.ValueOf(source)
	if sources.Kind() == reflect.Ptr {
		sources = sources.Elem()
	}
	switch targetValue.Elem().Kind() {
	case reflect.Slice:
		return toSlice(sources, targetValue)
	case reflect.Struct:
		return toStruct(sources, targetValue)
	}
	err = errors.New("data type only supported struct or Slice")

	return
}

//mapToSlice to Slice
func toSlice(source, target reflect.Value) error {
	//remove pointer
	target = target.Elem()
	targetType := target.Type()
	len := source.Len()
	targetSlice := reflect.MakeSlice(targetType, len, len)
	for i := 0; i < len; i++ {
		value := reflect.New(targetType.Elem())
		toStruct(source.Index(i), value)
		targetSlice.Index(i).Set(value.Elem())
	}
	target.Set(targetSlice)
	return nil
}

//toStruct map to Struct
func toStruct(source, target reflect.Value) (err error) {
	sourceMap, err := toMap(source)
	if err != nil {
		return
	}
	//remove pointer
	target = target.Elem()
	for i := 0; i < target.NumField(); i++ {
		vField := target.Field(i)
		tField := target.Type().Field(i)
		var (
			name string
			ok   bool
		)
		if name, ok = tField.Tag.Lookup("mapper"); !ok {
			name = tField.Name
		}
		if value, ok := sourceMap[name]; ok {
			sType := value.Type()
			vType := tField.Type

			if sType.Kind() == reflect.Ptr {
				sType = sType.Elem()
			}
			if vType.Kind() == reflect.Ptr {
				vType = vType.Elem()
			}
			//type of Struct
			if sType.Kind() == reflect.Struct && sType != vType {
				nValue := reflect.New(vType)
				toStruct(value, nValue)
				vField.Set(nValue.Elem())
				//type of Slice
			} else if sType.Kind() == reflect.Slice && sType != vType {
				nValue := reflect.New(vType)
				toSlice(value, nValue)
				vField.Set(nValue.Elem())
			} else {
				vField.Set(value)
			}
		}
	}
	return
}

//toMap struct to map
func toMap(source reflect.Value) (map[string]reflect.Value, error) {
	m := make(map[string]reflect.Value)
	t := source.Type()

	if source == zeroValue {
		return nil, errors.New("no exists this value")
	}
	for i := 0; i < source.NumField(); i++ {
		field := t.Field(i)
		tag, ok := field.Tag.Lookup(tag)
		if !ok {
			tag = field.Name
		}
		m[tag] = source.Field(i)
	}
	return m, nil
}
