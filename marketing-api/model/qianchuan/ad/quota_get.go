package ad

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QuotaGetRequest 获取在投计划配额信息 API Request
type QuotaGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
}

// Encode implement GetRequest interface
func (r QuotaGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id=", strconv.FormatUint(r.AdvertiserID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// QuotaGetResponse 获取在投计划配额信息 API Response
type QuotaGetResponse struct {
	model.BaseResponse
	Data *QuotaGetResult `json:"data,omitempty"`
}

type QuotaGetResult struct {
	// QuotaFeed 通投广告在投计划配额信息
	QuotaFeed *Quota `json:"quota_feed,omitempty"`
	// QuotaSearch 搜索广告在投计划配额信息
	QuotaSearch *Quota `json:"quota_search,omitempty"`
}

// Quota 广告在投计划配额信息
type Quota struct {
	// QuotaGear 当前所在配额档位：1,2,3,4,5,6,7,8,9
	QuotaGear int `json:"quota_gear,omitempty"`
	// MonthCost 当前最高月消耗
	MonthCost float64 `json:"month_cost,omitempty"`
	// QuotaInfo 在投计划数配额信息
	QuotaInfo *QuotaInfo `json:"quota_info,omitempty"`
	// DeliveryInfo 当前在投计划数信息
	DeliveryInfo *DeliveryInfo `json:"delivery_info,omitempty"`
	// StageInfo 当前所在投放阶段信息
	StageInfo *StageInfo `json:"stage_info,omitempty"`
}

// QuotaInfo 在投计划数配额信息
type QuotaInfo struct {
	// TotalNum 总在投计划配额
	TotalNum int `json:"total_num,omitempty"`
}

// DeliveryInfo 当前在投计划数信息
type DeliveryInfo struct {
	// TotalNum 总在投计划数
	TotalNum int `json:"total_num,omitempty"`
	// AdlabNum 托管计划的在投计划数
	AdlabNum int `json:"adlab_num,omitempty"`
	// NoAdlabNum 非托管计划的在投计划数
	NoAdlabNum int `json:"no_adlab_num,omitempty"`
}

// StageInfo 当前所在投放阶段信息
type StageInfo struct {
	// IsPrimary 是否在投放初期
	// 0：否
	// 1：是
	IsPrimary int `json:"is_primary,omitempty"`
	// StartDay 投放初期起始日期
	StartDay string `json:"start_day,omitempty"`
	// EndDay 投放初期截止日期
	EndDay string `json:"end_day,omitempty"`
}
