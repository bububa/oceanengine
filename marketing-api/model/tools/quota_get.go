package tools

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QuotaGetRequest 查询在投计划配额 API Request
type QuotaGetRequest struct {
	// AdvertiserID 广告账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignType 广告组类型，FEED：信息流 SEARCH:搜索广告
	CampaignType enum.CampaignType `json:"campaign_type,omitempty"`
	// DeliveryRange
	DeliveryRange enum.AdDeliveryRange `json:"delivery_range,omitempty"`
}

// Encode implement GetRequest interface
func (r QuotaGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("campaign_type", string(r.CampaignType))
	if r.DeliveryRange != "" {
		values.Set("delivery_range", string(r.DeliveryRange))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// QuotaGetResponse 查询在投计划配额 API Response
type QuotaGetResponse struct {
	model.BaseResponse
	Data *QuotaGetResult `json:"data,omitempty"`
}

// QuotaGetResult
type QuotaGetResult struct {
	// UsedQuota 在投计划数
	UsedQuota int64 `json:"used_quota,omitempty"`
	// SumQuota 在投计划配额
	SumQuota int64 `json:"sum_quota,omitempty"`
	// InLearning 是否在学习期
	InLearning bool `json:"in_learning,omitempty"`
	// MaxCost 最高月消耗
	MaxCost float64 `json:"max_cost,omitempty"`
}
