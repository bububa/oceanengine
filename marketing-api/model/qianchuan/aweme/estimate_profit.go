package aweme

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// EstimateProfitRequest 获取随心推投放效果预估 API Request
type EstimateProfitRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MarketingGoal 营销目标，可选值:
	// LIVE_PROM_GOODS: 直播间带货
	// VIDEO_PROM_GOODS: 短视频带货
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// DeliverySetting 投放设置
	DeliverySetting *DeliverySetting `json:"delivery_setting,omitempty"`
	// Audience 定向设置
	Audience *Audience `json:"audience,omitempty"`
}

// Encode implement GetRequest interface
func (r EstimateProfitRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.MarketingGoal != "" {
		values.Set("marketing_goal", string(r.MarketingGoal))
	}
	if r.DeliverySetting != nil {
		values.Set("delivery_setting", string(util.JSONMarshal(r.DeliverySetting)))
	}
	if r.Audience != nil {
		values.Set("audience", string(util.JSONMarshal(r.Audience)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// EstimateProfitResponse 获取随心推投放效果预估 API Response
type EstimateProfitResponse struct {
	model.BaseResponse
	Data *EstimateProfit `json:"data,omitempty"`
}

type EstimateProfit struct {
	// MinProfit 最小效果
	MinProfit float64 `json:"min_profit,omitempty"`
	// MaxProfit 最大效果
	MaxProfit float64 `json:"max_profit,omitempty"`
}
