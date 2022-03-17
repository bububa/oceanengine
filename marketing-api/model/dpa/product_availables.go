package dpa

import (
	"net/url"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// ProductAvailablesRequest 获取商品库信息 API Request
type ProductAvailablesRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r ProductAvailablesRequest) Encode() string {
	values := &url.Values{}
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	return values.Encode()
}

// ProductAvailablesResponse 获取商品库信息 API Response
type ProductAvailablesResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// List
		List []Platform `json:"list,omitempty"`
	} `json:"data,omitempty"`
}
