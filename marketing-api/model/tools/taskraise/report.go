package taskraise

import "github.com/bububa/oceanengine/marketing-api/enum"

// Report 任务详情
type Report struct {
	// ReportID 任务id
	ReportID uint64 `json:"report_id,omitempty"`
	// Name 任务名字
	Name string `json:"name,omitempty"`
	// RaiseMode 起量模式，允许值：常规起量CUSTOM、强劲起量STRONG
	RaiseMode enum.TaskRaiseMode `json:"raise_mode,omitempty"`
	// StartTime 任务开始时间
	StartTime string `json:"start_time,omitempty"`
	// Status 任务状态，枚举值: RAISING起量中、STOP 已结束
	Status enum.TaskRaiseStatus `json:"status,omitempty"`
	// BudgetMode 预算设置，枚举值: LIMIT 指定预算、NO_LIMIT 不限预算
	BudgetMode enum.TaskRaiseBudgetMode `json:"budget_model,omitemtpy"`
	// AllocatedBudget 日预算金额
	AllocatedBudget float64 `json:"allocated_budget,omitempty"`
	// IncreaseCost 消耗
	IncreaseCost float64 `json:"increase_cost,omitempty"`
	// Duration 生效天数
	Duration int `json:"duration,omitempty"`
}
