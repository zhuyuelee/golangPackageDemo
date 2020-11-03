package dtos

//PageInput 分页参数
type PageInput struct {
	// Limit 每页多少条
	Limit uint
	// Page 第几页
	Page uint
	// Key 关键字
	Key string
}
