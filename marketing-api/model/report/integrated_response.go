package report

import (
	"github.com/bububa/oceanengine/marketing-api/model"
)

type IntegratedResponse struct {
	model.BaseResponse
	Data *IntegratedResponseData `json:"data,omitempty"`
}

type IntegratedResponseData struct {
	List     []IntegratedResponseList `json:"list,omitempty"`
	PageInfo *model.PageInfo          `json:"page_info,omitempty"`
}

type IntegratedResponseList struct {
	Dimensions               *Dimensions `json:"dimensions,omitempty"`
	Metrics                  *Metrics    `json:"metrics,omitempty"`
	SubmitCertificationCount int64       `json:"submit_certification_count,omitempty"` // 提交身份认证数-出行行业司机在应用内提交身份认证的次数; 司机端，司机提交身份认证的次数
	ApprovalCount            int64       `json:"approval_count,omitempty"`             // 通过身份认证数-出行行业司机在应用内成功通过身份认证的次数; 司机端，司机通过身份认证的次数
	FirstOrderCount          int64       `json:"first_order_count,omitempty"`          // 乘客首次下单数-出行行业乘客在应用内首次成功下单的数量; 乘客端，乘客首次发单的次数
	FirstRentalOrderCount    int64       `json:"first_rental_order_count,omitempty"`   // 司机首次完单数-出行行业司机在应用内首次成功完成订单的数量; 司机端，司机首次完成订单的次数
	CommuteFirstPayCount     int64       `json:"commute_first_pay_count,omitempty"`    // 乘客首次支付数-出行行业乘客在应用内首次成功支付订单的数量; 乘客端，乘客首次支付订单的次数
}
