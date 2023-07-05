package taskraise

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// CreateRequest 新建优选起量任务 API Request
type CreateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// BudgetMode 预算设置，允许值: LIMIT 指定预算、NO_LIMIT 不限预算
	BudgetMode enum.TaskRaiseBudgetMode `json:"budget_mode,omitempty"`
	// PlatformVersion 创建优选起量任务的平台版本，不传值默认为巨量引擎平台原版，允许值：V2 巨量引擎体验版。可用于创建巨量引擎体验版上的账户优选起量任务
	PlatformVersion enum.PlatformVersion `json:"platform_version,omitempty"`
	// AllocatedBudget 日预算金额，当budget_mode为LIMIT时必填，范围：1000-9999999
	AllocatedBudget float64 `json:"allocated_budget,omitempty"`
	// EndTime 起量任务结束时间yyyy-mm-dd，传空为不限时长，起量将在指定日期0点结束
	EndTime string `json:"end_time,omitempty"`
	// RaiseMode 起量模式，允许值：常规起量CUSTOM（默认值）、强劲起量STRONG
	RaiseMode enum.TaskRaiseMode `json:"raise_mode,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}

// CreateResponse 新建优选起量任务 API Response
type CreateResponse struct {
	model.BaseResponse
	Data struct {
		// ReportID 任务id
		ReportID uint64 `json:"report_id,omitemtpy"`
	} `json:"data,omitempty"`
}
