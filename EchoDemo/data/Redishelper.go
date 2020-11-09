package data

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

// redisHelper returns edis.Conn对象
func redisHelper() (conn redis.Conn) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}
	return
}

//Set 设置Set
func Set(key string, val interface{}, duration time.Duration) (err error) {
	conn := redisHelper()
	defer conn.Close()
	_, err = conn.Do("Set", key, val, 100)
	return
}

//GetInt 获取Int
func GetInt(key string) (val int, err error) {
	conn := redisHelper()
	defer conn.Close()

	val, err = redis.Int(conn.Do("Get", key))
	return
}

//GetBool 获取Bool
func GetBool(key string) (val bool, err error) {
	conn := redisHelper()
	defer conn.Close()

	val, err = redis.Bool(conn.Do("Get", key))
	return
}

//Get  获取String
func Get(key string) (val string, err error) {
	conn := redisHelper()
	defer conn.Close()
	val, err = redis.String(conn.Do("Get", key))
	return
}

//GetStruct  获取Struct
// func GetStruct(key string, val interface{}) (err error) {

// json.Unmarshal()
// 	return
// }
