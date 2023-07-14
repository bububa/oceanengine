package dpa

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// ProductUpdateRequest 创建DPA商品（已有商品id）/修改DPA商品 API Request
type ProductUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PlatformID 商品库ID
	PlatformID uint64 `json:"platform_id,omitempty"`
	// ProductID 商品ID
	ProductID uint64 `json:"product_id,omitempty"`
	// ProductInfo 商品详情，包括商品库基础字段、落地页、品牌、商户、价格以及其他字段信息，以下所有字段均包含在该结构体下，且不同类型商品库填入字段存在差异，具体格式见下方字段详情
	ProductInfo *ProductForUpdate `json:"product_info,omitempty"`
}

// Encode implement PostRequest interface
func (r ProductUpdateRequest) Encode() []byte {
	bs, _ := json.Marshal(r)
	return bs
}

// ProductUpdateResponse 创建DPA商品（已有商品id）/修改DPA商品 API Response
type ProductUpdateResponse struct {
	model.BaseResponse
	Data struct {
		// ProductID 商品ID
		ProductID uint64 `json:"product_id,omitempty"`
	} `json:"data,omitempty"`
}
