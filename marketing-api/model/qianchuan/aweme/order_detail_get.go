package aweme

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// OrderDetailGetRequest 获取随心推订单详情 API Request
type OrderDetailGetRequest struct {
	// AdvertiserID 千川广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// OrderID 千川小店订单id
	OrderID uint64 `json:"order_id,omitempty"`
}

// Encode implement GetRequest interface
func (r OrderDetailGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("order_id", strconv.FormatUint(r.OrderID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// OrderDetailGetResponse 获取随心推订单详情 API Response
type OrderDetailGetResponse struct {
	model.BaseResponse
	Data *Order `json:"data,omitempty"`
}
