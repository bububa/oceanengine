package campaign

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// GetResponse 获取广告组 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *GetResponseData `json:"data,omitempty"`
}

// GetResponseData json返回值
type GetResponseData struct {
	// List campaign list
	List []Campaign `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

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
	// LandingType 广告组推广目的
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// Status 广告组状态
	Status enum.CampaignStatus `json:"status,omitempty"`
	// CampaignCreateTime 广告组创建时间, 格式：yyyy-mm-dd hh:MM:ss
	CampaignCreateTime string `json:"campaign_create_time,omitempty"`
	// CampaignModifyTime 广告组修改时间, 格式：yyyy-mm-dd hh:MM:ss
	CampaignModifyTime string `json:"campaign_modify_time,omitempty"`
	// MarketingPurpose 营销目的，允许值"CONVERSION"行动转化、"INTENTION"用户意向、"ACKNOWLEDGE"品牌认知
	MarketingPurpose enum.MarketingPurpose `json:"marketing_purpose,omitempty"`
	// DeliveryRelatedNum 广告组商品类型
	DeliveryRelatedNum enum.CampaignDPA `json:"delivery_related_num,omitempty"`
	// DeliveryMode 投放类型，允许值：MANUAL（手动）、PROCEDURAL（自动，投放管家）
	DeliveryMode string `json:"delivery_mode,omitempty"`
	// CampaignBudgetOptimization 支持预算择优分配允许值：
	// ON 开启，OFF 不开启(默认值)
	// 广告主商品类型为CAMPAIGN_DPA_MULTI_DELIVERY时不允许开启预算择优分配
	CampaignBudgetOptimization enum.OnOff `json:"campaign_budget_optimization,omitempty"`
	// SmartBidType 出价方式（投放场景），campaign_budget_optimization为ON时必填，且必须为SMART_BID_NO_BID
	SmartBidType enum.SmartBidType `json:"smart_bid_type,omitempty"`
}
