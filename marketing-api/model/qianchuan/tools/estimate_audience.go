package tools

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/ad"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// EstimateAudienceRequest 获取定向受众预估 API Request
type EstimateAudienceRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MarketingGoal 可选值:
	// LIVE_PROM_GOODS: 直播间带货
	// VIDEO_PROM_GOODS: 短视频带货
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// ExternalAction 短视频带货目的允许值：
	// AD_CONVERT_TYPE_SHOPPING 商品购买
	// AD_CONVERT_TYPE_QC_FOLLOW_ACTION 粉丝提升
	// AD_CONVERT_TYPE_QC_MUST_BUY 点赞评论
	// 注意：搜索广告短视频带货只支持AD_CONVERT_TYPE_SHOPPING直播带货目的允许值：
	// AD_CONVERT_TYPE_LIVE_ENTER_ACTION 进入直播间
	// AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION 直播间商品点击
	// AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION 直播间下单
	// AD_CONVERT_TYPE_NEW_FOLLOW_ACTION 直播间粉丝提升
	// AD_CONVERT_TYPE_LIVE_COMMENT_ACTION 直播间评论
	// AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY直播间成交
	ExternalAction qianchuan.ExternalAction `json:"external_action,omitempty"`
	// ProductID 商品ID
	ProductID uint64 `json:"product_id,omitempty"`
	// AwemeID 抖音号ID
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// Audience 定向设置
	Audience *ad.Audience `json:"audience,omitempty"`
}

// Encode implement GetRequest interface
func (r EstimateAudienceRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("marketing_goal", string(r.MarketingGoal))
	values.Set("external_action", string(r.ExternalAction))
	if r.ProductID > 0 {
		values.Set("product_id", strconv.FormatUint(r.ProductID, 10))
	}
	if r.AwemeID > 0 {
		values.Set("aweme_id", strconv.FormatUint(r.AwemeID, 10))
	}
	if r.Audience != nil {
		values.Set("audience", string(util.JSONMarshal(r.Audience)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// EstimateAudienceResponse  获取定向受众预估 API Response
type EstimateAudienceResponse struct {
	model.BaseResponse
	Data *EstimateAudienceResult `json:"data,omitempty"`
}

type EstimateAudienceResult struct {
	// CrowdCoverTotal 用户覆盖数
	CrowdCoverTotal int64 `json:"crowd_cover_total,omitempty"`
	// ShowCntTotal 广告展示数
	ShowCntTotal int64 `json:"show_cnt_total,omitempty"`
}
