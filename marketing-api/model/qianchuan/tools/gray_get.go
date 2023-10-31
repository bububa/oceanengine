package tools

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/tools"
)

// GrayGetResponse 查询白名单能力 API Response
type GrayGetResponse struct {
	model.BaseResponse
	Data *GrayGetResult `json:"data,omitempty"`
}

type GrayGetResult struct {
	// SuccessList 命中白名单请求
	SuccessList []tools.GrayItem `json:"success_list,omitempty"`
	// ErrorList 如果用户传入的key不存在/错误，会在该list中返回
	ErrorList []string `json:"error_list,omitempty"`
}
