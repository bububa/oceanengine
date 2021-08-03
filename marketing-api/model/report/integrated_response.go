package report

import (
	"github.com/bububa/oceanengine/marketing-api/model"
)

// IntegratedResponse 多合一数据报表接口 API Response
type IntegratedResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *IntegratedResponseData `json:"data,omitempty"`
}

// IntegratedResponseData json返回值
type IntegratedResponseData struct {
	// List data list
	List []IntegratedResponseList `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// IntegratedResponseList 详细数据
type IntegratedResponseList struct {
	// Dimensions 维度数据
	Dimensions *Dimensions `json:"dimensions,omitempty"`
	// Metrics 数据指标
	Metrics *Metrics `json:"metrics,omitempty"`
	// SubmitCertificationCount 提交身份认证数-出行行业司机在应用内提交身份认证的次数; 司机端，司机提交身份认证的次数
	SubmitCertificationCount int64 `json:"submit_certification_count,omitempty"`
	// ApprovalCount 通过身份认证数-出行行业司机在应用内成功通过身份认证的次数; 司机端，司机通过身份认证的次数
	ApprovalCount int64 `json:"approval_count,omitempty"`
	// FirstOrderCount 乘客首次下单数-出行行业乘客在应用内首次成功下单的数量; 乘客端，乘客首次发单的次数
	FirstOrderCount int64 `json:"first_order_count,omitempty"`
	// FirstRentalOrderCount 司机首次完单数-出行行业司机在应用内首次成功完成订单的数量; 司机端，司机首次完成订单的次数
	FirstRentalOrderCount int64 `json:"first_rental_order_count,omitempty"`
	// CommuteFirstPayCount 乘客首次支付数-出行行业乘客在应用内首次成功支付订单的数量; 乘客端，乘客首次支付订单的次数
	CommuteFirstPayCount int64 `json:"commute_first_pay_count,omitempty"`
}
