package utils

import (
	"bytes"
	"encoding/binary"
	"math/rand"
)

//IntToBytes Int To Bytes
func IntToBytes(n uint32) []byte {
	// data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, n)
	bytes := bytebuf.Bytes()
	return bytes
}

//BytesToInt Bytes To uint32
func BytesToInt(bys []byte) uint32 {
	bytebuff := bytes.NewBuffer(bys)
	var data uint32
	binary.Read(bytebuff, binary.BigEndian, &data)
	return data
}

//RandStrType 随机数据类型
type RandStrType byte

const (
	//Numbers 数字
	Numbers RandStrType = iota
	//Lower 小写 + 数字
	Lower
	//Upper 大写 + 数字
	Upper
	//All 大小写 + 数字
	All
)

//RandStr 随机字符串
func RandStr(lenth int, types RandStrType) string {

	chars := "0123456789"
	if types == Lower || types == All {
		chars += "abcdefghijklmnopqrstuvwxyz"
	}
	if types == Upper || types == All {
		chars += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	if types == All {
		chars += "~!@#$%^&*()_+,./;"
	}
	bytes := []byte(chars)
	bytesLen := len(bytes)
	result := make([]byte, lenth)
	for i := 0; i < lenth; i++ {
		result[i] = bytes[rand.Intn(bytesLen)]
	}
	return string(result)
}
