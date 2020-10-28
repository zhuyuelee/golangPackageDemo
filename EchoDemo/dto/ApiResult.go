package dto

// APIResult 接口返回值
type APIResult struct {
	ErrorCode int         `json:"errorCode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

// APIDataResult 接口返回值
type APIDataResult struct {
	APIResult
	Data interface{} `json:"data"`
}

// SuccessResult 返回成功消息
func SuccessResult(data interface{}) APIDataResult {
	return APIDataResult{
		APIResult: APIResult{
			ErrorCode: 0,
			Message:   "success",
		},
		Data: data,
	}
}

// ErrorResult 返回错误消息
func ErrorResult(errorCode int, message string) APIResult {
	return APIResult{
		ErrorCode: errorCode,
		Message:   message,
	}
}
