package creative

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DetailGetRequest 创意详细信息(新)API Request
type DetailGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 计划ID
	AdID uint64 `json:"ad_id,omitempty"`
}

// Encode implement GetRequest interface
func (r DetailGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("ad_id", strconv.FormatUint(r.AdID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// DetailGetResponse 创意详细信息(新)API Response
type DetailGetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *CreativeDetailV2 `json:"data,omitempty"`
}
