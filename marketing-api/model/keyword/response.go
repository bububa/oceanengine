package keyword

import "github.com/bububa/oceanengine/marketing-api/model"

// Response 关键词 API Response
type Response struct {
	model.BaseResponse
	// Data json返回值
	Data *ResponseData `json:"data,omitempty"`
}

// ResponseData json返回值
type ResponseData struct {
	// ErrorList 添加失败的搜索关键词列表
	ErrorList []Keyword `json:"error_list,omitempty"`
	// SuccessList 添加成功的搜索关键词列表
	SuccessList []Keyword `json:"success_list,omitempty"`
}
