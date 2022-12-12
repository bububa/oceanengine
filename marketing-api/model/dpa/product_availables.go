package dpa

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ProductAvailablesRequest 获取商品库信息 API Request
type ProductAvailablesRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r ProductAvailablesRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
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
