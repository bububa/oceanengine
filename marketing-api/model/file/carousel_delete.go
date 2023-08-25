package file

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CarouselDeleteRequest 批量删除图集 API Request
type CarouselDeleteRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CarouselIDs 图集id
	CarouselIDs []uint64 `json:"carousel_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r CarouselDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CarouselDeleteResponse 批量删除图集 API Response
type CarouselDeleteResponse struct {
	model.BaseResponse
	Data *CarouselDeleteResult `json:"data,omitempty"`
}

// CarouselDeleteResult 批量删除图集结果
type CarouselDeleteResult struct {
	// FailedIDs 批量删除失败列表
	FailedIDs []uint64 `json:"failed_ids,omitempty"`
	// SuccessList 批量删除成功列表
	SuccessList []uint64 `json:"success_list,omitempty"`
}
