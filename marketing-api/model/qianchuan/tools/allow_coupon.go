package tools

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AllowCouponRequest 智能优惠券白名单 API Request
type AllowCouponRequest struct {
	// AdvertiserID 千川广告主账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MarketingGoal 营销目标，允许值：
	// VIDEO_PROM_GOODS 推商品
	// LIVE_PROM_GOODS 推直播间
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// CampaignScene 营销场景 ，允许值：
	// DAILY_SALE日常销售
	//  NEW_CUSTOMER_CONVERT新客转化
	CampaignScene qianchuan.CampaignScene `json:"campaign_scene,omitempty"`
	// MarketingScene 广告类型，允许值：
	// FEED 通投广告
	// SEARCH 搜索广告
	// SHOPPING_MALL商城广告
	MarketingScene qianchuan.MarketingScene `json:"marketing_scene,omitempty"`
	// AwemeIDs 抖音id，即商品广告背后关联的抖音号，可通过【查询可推广抖音号列表】接口获取名下可推广抖音号
	AwemeIDs []uint64 `json:"aweme_ids,omitempty"`
	// ProductIDs 商品id列表，即准备推广的商品列表，可通过【查询店铺商品列表】接口获取名下可推广商品
	ProductIDs []uint64 `json:"product_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r AllowCouponRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("marketing_goal", string(r.MarketingGoal))
	values.Set("campaign_scene", string(r.CampaignScene))
	values.Set("marketing_scene", string(r.MarketingScene))
	if len(r.AwemeIDs) > 0 {
		values.Set("aweme_ids", string(util.JSONMarshal(r.AwemeIDs)))
	}
	if len(r.ProductIDs) > 0 {
		values.Set("product_ids", string(util.JSONMarshal(r.ProductIDs)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AllowCouponResponse 智能优惠券白名单 API Response
type AllowCouponResponse struct {
	model.BaseResponse
	Data *AllowCouponResult `json:"data,omitempty"`
}

type AllowCouponResult struct {
	// AdvAllowCoupon 广告主是否支持智能优惠券
	// true：支持
	// fale：不支持
	AdvAllowCoupon bool `json:"adv_allow_coupon,omitempty"`
	// AwemeAllowCoupon 抖音号是否支持智能优惠券信息
	AwemeAllowCoupon []AwemeAllowCoupon `json:"aweme_allow_coupon,omitempty"`
	// ProductAllowCoupon 商品是否支持智能优惠券信息
	ProductAllowCoupon []ProductAllowCoupon `json:"product_allow_coupon,omitempty"`
}

// AwemeAllowCoupon 抖音号是否支持智能优惠券信息
type AwemeAllowCoupon struct {
	// AwemeID 抖音号id，入參包含aweme_id时返回
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// AllowCoupon 抖音号是否支持智能优惠券
	// true：支持
	// fale：不支持
	AllowCoupon bool `json:"allow_coupon,omitempty"`
}

// ProductAllowCoupon 商品是否支持智能优惠券信息
type ProductAllowCoupon struct {
	// ProductID 商品id，入參包含product_id时返回
	ProductID uint64 `json:"product_id,omitempty"`
	// AllowCoupon 抖音号是否支持智能优惠券
	// true：支持
	// fale：不支持
	AllowCoupon bool `json:"allow_coupon,omitempty"`
}
