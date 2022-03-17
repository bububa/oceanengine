package keyword

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DeleteRequest 删除关键词 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 待删除关键词的计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// KeywordIDs 待删除的关键词id列表，一次最多批量删除100个关键词
	KeywordIDs []uint64 `json:"keyword_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// DeleteResponse 关键词 API Response
type DeleteResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *DeleteResponseData `json:"data,omitempty"`
}

// DeleteResponseData json返回值
type DeleteResponseData struct {
	// ErrorList 添加失败的搜索关键词列表
	ErrorList []Keyword `json:"error_list,omitempty"`
	// SuccessList 添加成功的搜索关键词列表
	SuccessList []uint64 `json:"success_list,omitempty"`
}
