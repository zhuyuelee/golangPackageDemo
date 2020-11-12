package main

import (
	"fmt"
	"reflect"
)

func main() {

	user := User{
		Addr: Addr{
			Addrss: "郑州",
		},
		UserName: "bac",
	}
	toUser := &User1{}
	m, err := toMap(reflect.ValueOf(user))
	if err != nil {
		fmt.Println("err=", err)
		return
	}
	toStruct(m, reflect.ValueOf(toUser), "")

	fmt.Printf("user=%+v \n", user)
	fmt.Printf("toUser=%+v \n", toUser)
}

//Addr Addr
type Addr struct {
	Addrss string
}

//Addr1 Addr
type Addr1 struct {
	Addrss string
}

//User User
type User struct {
	Addr     Addr `mapper:"addr"`
	UserName string
}

//User1 User
type User1 struct {
	Addr1    Addr `mapper:"addr"`
	UserName string
}

//var tagName = "mapper"
//var zeroValue = reflect.Value{}

//toMap struct to map
// func toMap1(source reflect.Value) (map[string]reflect.Value, error) {

// 	m := make(map[string]reflect.Value)
// 	t := source.Type()

// 	if source == zeroValue {
// 		return nil, errors.New("no exists this value")
// 	}
// 	for i := 0; i < source.NumField(); i++ {
// 		field := t.Field(i)
// 		tag, ok := field.Tag.Lookup(tagName)
// 		if !ok {
// 			tag = field.Name
// 		}
// 		//Anonymous Field
// 		if field.Anonymous {
// 			childMap, err := toMap(source.Field(i))
// 			if err == nil && len(childMap) > 0 {
// 				for key, val := range childMap {
// 					if ok {
// 						m[fmt.Sprintf("%s.%s", tag, key)] = val
// 					} else {
// 						m[key] = val
// 					}
// 				}
// 			}
// 		} else {
// 			m[tag] = source.Field(i)
// 		}

// 	}
// 	fmt.Printf("childMap=%+v \n", m)
// 	return m, nil
// }

// //toStruct map to Struct
// func toStruct1(source map[string]reflect.Value, target reflect.Value, parentTag string) (err error) {
// 	//remove pointer
// 	target = target.Elem()
// 	for i := 0; i < target.NumField(); i++ {
// 		vField := target.Field(i)
// 		tField := target.Type().Field(i)
// 		var (
// 			tag string
// 			ok  bool
// 		)
// 		if tag, ok = tField.Tag.Lookup(tagName); !ok {
// 			tag = tField.Name
// 		}

// 		if tField.Anonymous {
// 			nValue := reflect.New(tField.Type)
// 			toStruct(source, nValue, tag)
// 			vField.Set(nValue.Elem())
// 			return
// 		}
// 		if parentTag != "" {
// 			tag = fmt.Sprintf("%s.%s", parentTag, tag)
// 		}
// 		if value, ok := source[tag]; ok {
// 			sType := value.Type()
// 			vType := tField.Type
// 			if sType.Kind() == reflect.Ptr {
// 				sType = sType.Elem()
// 			}
// 			if vType.Kind() == reflect.Ptr {
// 				vType = vType.Elem()
// 			}
// 			if sType.Kind() == reflect.Struct && sType != vType {
// 				nValue := reflect.New(vType)
// 				childMap, err := toMap(value)
// 				if err == nil {
// 					toStruct(childMap, nValue, "")
// 					vField.Set(nValue.Elem())
// 				}
// 				//type of Slice
// 				// } else if sType.Kind() == reflect.Slice && sType != vType {
// 				// 	nValue := reflect.New(vType)
// 				// 	toSlice(value, nValue)
// 				// 	vField.Set(nValue.Elem())
// 			} else {
// 				vField.Set(value)
// 			}
// 		}
// 	}
// 	return
// }
