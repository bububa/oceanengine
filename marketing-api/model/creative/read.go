package creative

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ReadRequest 创意详细信息API Request
type ReadRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 计划ID
	AdID uint64 `json:"ad_id,omitempty"`
}

// Encode implement GetRequest interface
func (r ReadRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("ad_id", strconv.FormatUint(r.AdID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ReadResponse 创意详细信息API Response
type ReadResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *CreativeDetail `json:"data,omitempty"`
}
