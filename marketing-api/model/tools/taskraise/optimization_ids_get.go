package taskraise

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// OptimizationIDsGetRequest 查询优选起量状态 API Request
type OptimizationIDsGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r OptimizationIDsGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// OptimizationIDsGetResponse 查询优选起量状态 API Response
type OptimizationIDsGetResponse struct {
	model.BaseResponse
	Data struct {
		// Status 优选起量状态，枚举值: DISABLERAISE 不能开启起量、ENABLERAISE 可以开启起量、RAISING 起量中
		Status enum.TaskRaiseStatus `json:"status,omitemtpy"`
	} `json:"data,omitempty"`
}
