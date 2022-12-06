package eventmanager

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AvailableEventsGetRequest 获取可创建事件列表 API Request
type AvailableEventsGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AssetID 资产ID
	AssetID uint64 `json:"asset_id,omitempty"`
}

// Encode implement GetRequest interface
func (r AvailableEventsGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Add("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Add("asset_id", strconv.FormatUint(r.AssetID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AvailableEventsGetResponse 获取可创建事件列表 API Response
type AvailableEventsGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// List 可创建事件列表
		List []EventConfig `json:"event_configs,omitempty"`
	} `json:"data,omitempty"`
}
