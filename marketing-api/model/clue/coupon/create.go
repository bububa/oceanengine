package coupon

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 创建卡券 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ActivityType 卡券类型，可选值:
	// DIRECT_NEED_PHONE：直接发券，收集手机号，need_phone须为true
	// DIRECT_NOT_NEED_PHONE ：直接发券，不收集手机号，need_phone须为false 并且 BindFormId会被忽略
	// activity_type和need_phone传入不匹配会报错
	ActivityType ActivityType `json:"activity_type,omitempty"`
	// Coupon 卡券信息
	Coupon *Coupon `json:"coupon,omitempty"`
	// BindFormID 绑定的表单ID。仅need_phone=1时，支持设置此值.
	// 不设置此值时，系统会默认为当前的卡券创建一个用来收集手机号的表单
	// 使用者也可以明确指定要绑定的表单，此表单的验证类型会覆盖need_sms_verify的值
	// 使用者指定的表单必须有且仅有姓名和电话两个字段
	BindFormID uint64 `json:"bind_form_id,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建卡券 API Response
type CreateResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data struct {
		// ActivityID 活动ID
		ActivityID uint64 `json:"activity_id,omitempty"`
	} `json:"data,omitempty"`
}
