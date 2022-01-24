package campaign

import "github.com/bububa/oceanengine/marketing-api/enum"

// Campaign 广告组信息
type Campaign struct {
	// ID 广告组ID
	ID uint64 `json:"id,omitempty"`
	// Name 广告组名称
	Name string `json:"name,omitempty"`
	// Budget 广告组预算
	Budget float64 `json:"budget,omitempty"`
	// BudgetMode 广告组预算类型
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// MarketingGoal 广告组营销目标，VIDEO_PROM_GOODS：短视频带货、LIVE_PROM_GOODS：直播带货。
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// Status 广告组状态
	Status string `json:"status,omitempty"`
	// CreateDate 广告组创建日期, 格式：yyyy-mm-dd
	CreateDate string `json:"create_date,omitempty"`
}
