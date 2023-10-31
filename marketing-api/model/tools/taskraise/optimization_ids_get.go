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
	// PlatformVersion 创建优选起量任务的平台版本，不传值默认为巨量引擎平台原版，允许值：V2 巨量引擎体验版。可用于创建巨量引擎体验版上的账户优选起量任务
	PlatformVersion enum.PlatformVersion `json:"platform_version,omitempty"`
}

// Encode implement GetRequest interface
func (r OptimizationIDsGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.PlatformVersion != "" {
		values.Set("platform_version", string(r.PlatformVersion))
	}
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
