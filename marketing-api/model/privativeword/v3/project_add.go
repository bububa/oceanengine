package v3

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AddRequest 批量新增计划否定词 API Request
type AddRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ProjectList 项目列表
	ProjectList []Word `json:"project_list,omitempty"`
}

// Encode implement PostRequest interface
func (r AddRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AddResponse 批量新增计划否定词 API Response
type AddResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AddResponseData `json:"data,omitempty"`
}

// AddResponseData json返回值
type AddResponseData struct {
	// ProjectErrorList 添加失败的项目id列表
	ProjectErrorList []uint64 `json:"project_error_list,omitempty"`
	// ProjectList 添加成功的项目列表
	ProjectList []Word `json:"project_list,omitempty"`
}
