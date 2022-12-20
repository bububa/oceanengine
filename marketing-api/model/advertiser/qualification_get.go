package advertiser

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QualificationGetRequest 获取广告主资质（新版） API Request
type QualificationGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r QualificationGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// QualificationGetResponse 获取广告主资质（新版） API Response
type QualificationGetResponse struct {
	model.BaseResponse
	// Data 返回数据
	Data *Qualification `json:"data,omitempty"`
}
