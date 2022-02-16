package clue

// KeyAction 活动记录
type KeyAction struct {
	// AdvID 广告主ID
	AdvID uint64 `json:"adv_id,omitempty"`
	// Name 活动名
	Name string `json:"name,omitempty"`
	// AdPlanName 计划名
	AdPlanName string `json:"ad_plan_name,omitempty"`
	// AdPlanID 计划plan id
	AdPlanID uint64 `json:"ad_plan_id,omitempty"`
	// ClueID 线索id
	ClueID uint64 `json:"clue_id,omitempty"`
	// ReqID 当前线索对应广告的请求id
	ReqID string `json:"req_id,omitempty"`
	// SiteID 站点id
	SiteID uint64 `json:"site_id,omitempty"`
	// Cid cid
	Cid uint64 `json:"cid,omitempty"`
	// CreateTime 创建时间
	CreateTime string `json:"create_time,omitempty"`
}
