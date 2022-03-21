package creative

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CustomCreativeCreateRequest 创建自定义创意（营销链路） API Request
type CustomCreativeCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 广告计划ID，计划ID要属于广告主ID，且非删除计划，否则会报错
	AdID uint64 `json:"ad_id,omitempty"`
	// CreativeList 自定义素材信息, 最多支持10个创意。投放位置和创意类型决定素材规格。
	CreativeList []CreativeInfo `json:"creative_list,omitempty"`
	// AdData 广告计划数据
	AdData *AdData `json:"ad_data,omitempty"`
}

// Encode implement PostRequest interface
func (r CustomCreativeCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CustomCreativeCreateResponse 创建自定义创意 API Response
type CustomCreativeCreateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *CustomCreativeCreateResponseData `json:"data,omitempty"`
}

// CustomCreativeCreateResponseData json返回值
type CustomCreativeCreateResponseData struct {
	// CreativeIDs 创意ID列表，若部分失败，则对应项为null
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
	// Errors 每个创意对应的错误信息，若部分成功，则对应项为null
	Errors []Error `json:"errors,omitempty"`
}
