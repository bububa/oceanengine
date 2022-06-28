package dpa

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// ProductStatusBatchUpdateRequest 批量修改DPA商品状态 API Request
type ProductStatusBatchUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// PlatformID 商品库ID
	PlatformID uint64 `json:"platform_id,omitempty"`
	// ProductIDs 商品ID
	ProductIDs []uint64 `json:"product_ids,omitempty"`
	// OptStatus 操作
	// ENABLE开启投放
	// DISABLE停止投放
	OptStatus string `json:"opt_status,omitempty"`
}

// Encode implement PostRequest interface
func (r ProductStatusBatchUpdateRequest) Encode() []byte {
	bs, _ := json.Marshal(r)
	return bs
}

// ProductStatusBatchUpdateResponse 批量修改DPA商品状态 API Response
type ProductStatusBatchUpdateResponse struct {
	model.BaseResponse
	Data *ProductStatusBatchUpdateResponseData `json:"data,omitempty"`
}

type ProductStatusBatchUpdateResponseData struct {
	// SuccessList 修改状态成功的商品列表
	SuccessList []uint64 `json:"success_list,omitempty"`
	// ErrorList 修改状态失败的商品列表
	ErrorList []UpdateError `json:"error_list,omitempty"`
}

type UpdateError struct {
	// ProductID 修改失败的商品id
	ProductID uint64 `json:"product_id,omitempty"`
	// ErrorMsg 失败原因
	ErrorMsg string `json:"error_msg,omitempty"`
}
