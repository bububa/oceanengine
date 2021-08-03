package campaign

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
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
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
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
