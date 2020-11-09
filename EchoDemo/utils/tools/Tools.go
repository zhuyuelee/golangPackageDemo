package tools

import (
	"math/rand"
	"time"
)

const (
	//Numbers 数字
	Numbers = iota
	//Lower 小写 + 数字
	Lower
	//Upper 大写 + 数字
	Upper
	//All 大小写 + 数字
	All
)

//RandStr 随机字符串
func RandStr(lenth, types int) string {

	chars := "0123456789"
	if types == Lower || types == All {
		chars += "abcdefghijklmnopqrstuvwxyz"
	}
	if types == Upper || types == All {
		chars += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	bytes := []byte(chars)
	bytesLen := len(bytes)
	result := make([]byte, lenth)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < lenth; i++ {
		result[i] = bytes[r.Intn(bytesLen)]
	}
	return string(result)
}
