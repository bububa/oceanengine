package aweme

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// OrderQuotaGetRequest 查询随心推使用中订单配额信息 API Request
type OrderQuotaGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r OrderQuotaGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// OrderQuotaGetResponse 查询随心推使用中订单配额信息 API Response
type OrderQuotaGetResponse struct {
	model.BaseResponse
	Data *OrderQuota `json:"data,omitempty"`
}
