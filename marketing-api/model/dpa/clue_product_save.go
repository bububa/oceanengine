package dpa

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ClueProductSaveRequest 创建/编辑升级版商品 API Request
type ClueProductSaveRequest struct {
	// Product 商品信息
	Product *Product `json:"product,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implements PostRequest interface
func (r ClueProductSaveRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ClueProductSaveResponse 创建/编辑升级版商品 API Response
type ClueProductSaveResponse struct {
	model.BaseResponse
	Data *ClueProductSaveResult `json:"data,omitempty"`
}

type ClueProductSaveResult struct {
	// PlatformProductID 商品ID
	PlatformProductID uint64 `json:"platform_product_id,omitempty"`
	// OuterID 商品外部ID
	OuterID string `json:"outer_id,omitempty"`
	// StoreID 店铺ID
	StoreID string `json:"store_id,omitempty"`
}
