package audience

import "github.com/bububa/oceanengine/marketing-api/enum"

// Filtering 过滤条件
type Filtering struct {
	// InterestActionType 行为兴趣类型; 默认值: INTEREST
	InterestActionType enum.InterestActionType `json:"interest_action_type,omitempty"`
	// AudienceLevel 类目词级别; 默认值: FIRST_LEVEL
	AudienceLevel enum.AudienceLevel `json:"audience_level,omitempty"`
	// CampaignIDs 广告组查询列表，长度1-100
	CampaignIDs []uint64 `json:"campaign_ids,omitempty"`
	// AdIDs 广告计划查询列表，长度1-100
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// ActionScene 行为场景，仅在行为兴趣类型为ACTION时传入; 默认值: ["E-COMMERCE", "NEWS", "APP"]
	ActionScene []enum.ActionScene `json:"action_scene,omitempty"`
	// ActionDays 行为天数，仅在行为兴趣类型为ACTION时传入; 默认值: 30
	ActionDays int `json:"action_days,omitempty"`
	// Behaviors 互动类型; 默认值: ["FOLLOWED_USER", "COMMENTED_USER", "LIKED_USER", "SHARED_USER"]
	Behaviors []enum.Behavior `json:"behaviors,omitempty"`
}
