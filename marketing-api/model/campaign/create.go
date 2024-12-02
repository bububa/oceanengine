package campaign

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 创建广告组 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignName 广告组名称，长度为1-100个字符，其中1个中文字符算2位
	CampaignName string `json:"campaign_name,omitempty"`
	// Operation 广告组状态, 允许值: "enable","disable"默认值：enable开启状态
	Operation string `json:"operation,omitempty"`
	// BudgetMode 广告组预算类型
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// Budget 广告组预算，取值范围: ≥ 0; 当budget_mode为"BUDGET_MODE_DAY"时,必填,且日预算不少于300	元
	Budget float64 `json:"budget,omitempty"`
	// LandingType 广告组推广目的
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// CampaignType 广告组类型，允许值FEED信息流、SEARCH搜索广告
	CampaignType enum.CampaignType `json:"campaign_type,omitempty"`
	// MarketingPurpose 营销目的，允许值CONVERSION行动转化、INTENTION用户意向、ACKNOWLEDGE品牌认知
	// 创建新版营销链路广告组，该字段必填
	MarketingPurpose enum.MarketingPurpose `json:"marketing_purpose,omitempty"`
	// DeliveryRelatedNum 广告组商品类型
	DeliveryRelatedNum enum.CampaignDPA `json:"delivery_related_num,omitempty"`
	// CampaignBudgetOptimization 支持预算择优分配允许值：
	// ON 开启，OFF 不开启(默认值)
	// 广告主商品类型为CAMPAIGN_DPA_MULTI_DELIVERY时不允许开启预算择优分配
	CampaignBudgetOptimization enum.OnOff `json:"campaign_budget_optimization,omitempty"`
	// SmartBidType 出价方式（投放场景），campaign_budget_optimization为ON时必填，且必须为SMART_BID_NO_BID
	SmartBidType enum.SmartBidType `json:"smart_bid_type,omitempty"`
	// UniqueFk 第三方唯一键，传该值时保证接口重试的幂等性，请注意，带有相同unique_fk的请求服务端会视为同一个广告处理。仅在创建接口传入且无法修改，如果创建时传入了已存在的唯一键值，那么会返回该唯一键值所对应的广告组ID。该值可用于内部系统会生成的唯一ID与头条ID做关联的场景，避免超时重试实际上一次创建请求又成功导致的重复创建问题，通过unique_fk可与内部系统ID实现关联并避免重复创建，可结合实际场景选择使用，广告组中的unique_fk要求不重复，与计划中的unique_fk无相关。
	UniqueFk string `json:"unique_fk,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建广告组 API Response
type CreateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *CreateResponseData `json:"data,omitempty"`
}

// CreateResponseData json返回值
type CreateResponseData struct {
	// CampaignID 广告组id
	CampaignID uint64 `json:"campaign_id,omitempty"`
}
