package dpa

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ClueProductDetailRequest 获取升级版商品详情 API Request
type ClueProductDetailRequest struct {
	// ProductIDs 线索商品ID
	ProductIDs []uint64 `json:"product_ids,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implements GetRequest interface
func (r ClueProductDetailRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("product_ids", string(util.JSONMarshal(r.ProductIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ClueProductDetailResponse 获取升级版商品详情 API Response
type ClueProductDetailResponse struct {
	model.BaseResponse
	Data struct {
		// Products 商品信息
		Products []Product `json:"products,omitempty"`
	} `json:"data,omitempty"`
}
