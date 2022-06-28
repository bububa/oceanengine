package dpa

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// ProductCreateRequest 创建DPA商品（无商品id） API Request
type ProductCreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PlatformID 商品库ID
	PlatformID uint64 `json:"platform_id,omitempty"`
	// ProductInfo 商品详情，包括商品库基础字段、落地页、品牌、商户、价格以及其他字段信息，以下所有字段均包含在该结构体下，且不同类型商品库填入字段存在差异，具体格式见下方字段详情
	ProductInfo *Product `json:"product_info,omitempty"`
}

// Encode implement PostRequest interface
func (r ProductCreateRequest) Encode() []byte {
	bs, _ := json.Marshal(r)
	return bs
}

// ProductCreateResponse 创建DPA商品（无商品id） API Response
type ProductCreateResponse struct {
	model.BaseResponse
	Data struct {
		// ProductID 商品ID
		ProductID uint64 `json:"product_id,omitempty"`
	} `json:"data,omitempty"`
}
