package audience

import "github.com/bububa/oceanengine/marketing-api/enum"

type Filtering struct {
	InterestActionType enum.InterestActionType `json:"interest_action_type,omitempty"` // 行为兴趣类型; 默认值: INTEREST
	AudienceLevel      enum.AudienceLevel      `json:"audience_level,omitempty"`       // 类目词级别; 默认值: FIRST_LEVEL
	CampaignIDs        []uint64                `json:"campaign_ids,omitempty"`         // 广告组查询列表，长度1-100
	AdIDs              []uint64                `json:"ad_ids,omitempty"`               // 广告计划查询列表，长度1-100
	ActionScene        []enum.ActionScene      `json:"action_scene,omitempty"`         // 行为场景，仅在行为兴趣类型为ACTION时传入; 默认值: ["E-COMMERCE", "NEWS", "APP"]
	ActionDays         int                     `json:"action_days,omitempty"`          // 行为天数，仅在行为兴趣类型为ACTION时传入; 默认值: 30
	Behaviors          []enum.Behavior         `json:"behaviors,omitempty"`            // 互动类型; 默认值: ["FOLLOWED_USER", "COMMENTED_USER", "LIKED_USER", "SHARED_USER"]
}
