package campaign

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateRequest 修改广告组 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告组ID，广告组ID需要属于广告主ID，否则会报错！
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CampaignName 广告组名称，长度为1-100个字符，其中1个中文字符算2位
	CampaignName string `json:"campaign_name,omitempty"`
	// BudgetMode 广告组预算类型
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// Budget 广告组预算，取值范围: ≥ 0; 当budget_mode为"BUDGET_MODE_DAY"时,必填,且日预算不少于300	元
	Budget float64 `json:"budget,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
