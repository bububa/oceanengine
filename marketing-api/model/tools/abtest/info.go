package abtest

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// InfoRequest 获取实验详情及报告 API Request
type InfoRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TestID 实验ID
	TestID uint64 `json:"test_id,omitempty"`
}

// Encode implement GetRequest interface
func (r InfoRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("test_id", strconv.FormatUint(r.TestID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// InfoResponse 获取实验详情及报告 API Response
type InfoResponse struct {
	model.BaseResponse
	Data *AbTest `json:"data,omitempty"`
}
