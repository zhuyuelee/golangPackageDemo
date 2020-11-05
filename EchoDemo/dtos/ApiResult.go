package dtos

import (
	"encoding/json"
	"fmt"
)

// APIResult 接口返回值
type APIResult struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}

func (a APIResult) Error() (result string) {
	data, err := json.Marshal(a)
	if err == nil {
		result = string(data)
		fmt.Println(result)
	}
	return
}

// APIDataResult 接口返回值
type APIDataResult struct {
	APIResult
	Data interface{} `json:"data"`
}

// SuccessDataResult 返回成功消息
func SuccessDataResult(data interface{}) APIDataResult {
	return APIDataResult{
		APIResult: APIResult{
			ErrorCode: 0,
			Message:   "success",
		},
		Data: data,
	}
}

// SuccessResult 返回成功消息
func SuccessResult() APIResult {
	return APIResult{
		ErrorCode: 0,
		Message:   "success",
	}
}

// ErrorResult 返回错误消息
func ErrorResult(errorCode int, message string) APIResult {
	return APIResult{
		ErrorCode: errorCode,
		Message:   message,
	}
}
