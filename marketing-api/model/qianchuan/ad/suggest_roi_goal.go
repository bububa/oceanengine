package ad

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// SuggestRoiGoalRequest 获取支付ROI目标建议 API Request
type SuggestRoiGoalRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 抖音id ，即待推广直播开播的抖音号，可通过【查询可推广抖音号列表】接口获取名下可推广抖音号
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// MarketingGoal 营销目标，可选值:
	// LIVE_PROM_GOODS 推直播间
	// VIDEO_PROM_GOODS 推商品
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// MarketingScene 广告类型， 可选值:
	// FEED 通投
	// SEARCH 搜索
	// SHOPPING_MALL商城广告
	MarketingScene qianchuan.MarketingScene `json:"marketing_scene,omitempty"`
	// ProductID 商品ID
	ProductID uint64 `json:"product_id,omitempty"`
	// ProductNewOpen 是否开启新品加速
	// 开启：true
	// 关闭：false
	ProductNewOpen bool `json:"product_new_open,omitempty"`
	// ExternalAction 短视频带货目的允许值：
	// AD_CONVERT_TYPE_SHOPPING 商品购买
	// AD_CONVERT_TYPE_QC_FOLLOW_ACTION 粉丝提升
	// AD_CONVERT_TYPE_QC_MUST_BUY 点赞评论
	// AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7DAYS7日总成交
	// 注意：搜索广告短视频带货只支持AD_CONVERT_TYPE_SHOPPING
	//
	// 直播带货目的允许值：
	// AD_CONVERT_TYPE_LIVE_ENTER_ACTION 进入直播间
	// AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION 直播间商品点击
	// AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION 直播间下单
	// AD_CONVERT_TYPE_NEW_FOLLOW_ACTION 直播间粉丝提升
	// AD_CONVERT_TYPE_LIVE_COMMENT_ACTION 直播间评论
	// AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY直播间成交
	// AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7DAYS7日总成交
	ExternalAction qianchuan.ExternalAction `json:"external_action,omitempty"`
	// CampaignScene 营销场景，可选值
	// DAILY_SALE 日常销售
	// NEW_CUSTOMER_TRANSFORMATION 新客转化
	CampaignScene qianchuan.CampaignScene `json:"campaign_scene,omitempty"`
}

// Encode implement GetRequest interface
func (r SuggestRoiGoalRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("aweme_id", strconv.FormatUint(r.AwemeID, 10))
	if r.MarketingGoal != "" {
		values.Set("marketing_goal", string(r.MarketingGoal))
	}
	if r.MarketingScene != "" {
		values.Set("marketing_scene", string(r.MarketingScene))
	}
	if r.ProductID > 0 {
		values.Set("product_id", strconv.FormatUint(r.ProductID, 10))
	}
	if r.ProductNewOpen {
		values.Set("product_new_open", "true")
	} else {
		values.Set("product_new_open", "false")
	}
	if r.ExternalAction != "" {
		values.Set("external_action", string(r.ExternalAction))
	}
	if r.CampaignScene != "" {
		values.Set("campaign_scene", string(r.CampaignScene))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// SuggestRoiGoalResponse 获取支付ROI目标建议 API Response
type SuggestRoiGoalResponse struct {
	model.BaseResponse
	Data *SuggestRoiResult `json:"data,omitempty"`
}

type SuggestRoiResult struct {
	// EcpRoiGoal 支付ROI目标建议值的上限，单位：元，精确到两位小数
	EcpRoiGoal float64 `json:"ecp_roi_goal,omitempty"`
	// RoiLowerBound roi下限，在开启新品加速后，需校验上下线
	RoiLowerBound float64 `json:"roi_lower_bound,omitempty"`
	// RoiUpperBound roi上限，在开启新品加速后，需校验上下线
	RoiUpperBound float64 `json:"roi_upper_bound,omitempty"`
}
