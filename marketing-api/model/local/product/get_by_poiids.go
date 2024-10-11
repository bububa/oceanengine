package product

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetByPoiIDsRequest 根据门店ID查询门店下商品ID API Request
type GetByPoiIDsRequest struct {
	// LocalAccountID 本地推广告账户ID
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// PoiIDs 门店ID列表，这里需填入项目所推广的门店
	PoiIDs []uint64 `json:"poi_ids,omitempty"`
}

// Encode implements GetRequest interface
func (r GetByPoiIDsRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("local_account_id", strconv.FormatUint(r.LocalAccountID, 10))
	values.Set("poi_ids", string(util.JSONMarshal(r.PoiIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetByPoiIDsResponse 根据门店ID查询门店下商品ID API Response
type GetByPoiIDsResponse struct {
	model.BaseResponse
	Data struct {
		// ProductIDs 门店下商品ID列表
		ProductIDs []uint64 `json:"product_ids,omitempty"`
	} `json:"data,omitempty"`
}
