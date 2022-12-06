package coupon

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DetailRequest 获取卡券详情 API Request
type DetailRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ActivityID 活动ID
	ActivityID uint64 `json:"activity_id,omitempty"`
}

// Encode implement GetRequest interface
func (r DetailRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("activity_id", strconv.FormatUint(r.ActivityID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// DetailResponse 获取卡券详情 API Response
type DetailResponse struct {
	model.BaseResponse
	// Data 返回数据
	Data *DetailResponseData `json:"data,omitempty"`
}

// DetailResponseData json返回值
type DetailResponseData struct {
	// ActivityID 活动ID
	ActivityID uint64 `json:"activity_id,omitempty"`
	// ActivityType 卡券类型，可选值:
	// DIRECT_NEED_PHONE：直接发券，收集手机号，need_phone须为true
	// DIRECT_NOT_NEED_PHONE ：直接发券，不收集手机号，need_phone须为false 并且 BindFormId会被忽略
	ActivityType ActivityType `json:"activity_type,omitempty"`
	// Coupon 卡券信息
	Coupon *Coupon `json:"coupon,omitempty"`
	// Form 表单信息
	Form *Form `json:"form,omitempty"`
}
