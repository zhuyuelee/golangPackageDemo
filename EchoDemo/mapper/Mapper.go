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

//Mapper to struct or Slice
func Mapper(source interface{}, target interface{}) error {
	startTime := time.Now()
	defer func() {
		fmt.Println("mapper time=", time.Now().Sub(startTime).Seconds())
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
	switch sources.Kind() {
	case reflect.Slice:
		return mapperToSlice(sources, targetValue)
	case reflect.Struct:
		return mapperToStruct(sources, targetValue)
	}
	return errors.New("Source data only struct or Slice supported")
}

//mapperToSlice to Slice
func mapperToSlice(source reflect.Value, target reflect.Value) error {
	// direct := reflect.Indirect(target)

	// for i := 0; i < sources.Len(); i++ {
	// 	s := sources.Index(i)
	// 	d := reflect.New(target.Index(0).Type())
	// 	mapperToStruct(s, d)
	// 	append(target, d)
	// }

	// sourceMap, err := toMap(sources)
	// if err != nil {
	// 	fmt.Println("mapper error=", err)
	// } else {
	// 	toStruct(sourceMap, target)
	// }
	// return err

	return nil
}

//mapperToStruct struct map to struct
func mapperToStruct(source reflect.Value, target reflect.Value) error {
	sourceMap, err := toMap(source)
	if err != nil {
		fmt.Println("mapper error=", err)
	} else {
		toStruct(sourceMap, target)
	}
	return err
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
				mapperToStruct(value, nValue)
				vField.Set(nValue.Elem())
			} else {
				vField.Set(value)
			}

		}
	}
}
