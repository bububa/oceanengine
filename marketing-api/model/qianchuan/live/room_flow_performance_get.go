package live

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// RoomFlowPerformanceGetRequest 获取直播间流量表现 API Request
type RoomFlowPerformanceGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// RoomID 直播间id
	RoomID uint64 `json:"room_id,omitempty"`
	// FlowSource 来源
	FlowSource qianchuan.LiveFlowSource `json:"flow_source,omitempty"`
}

// Encode implement GetRequest interface
func (r RoomFlowPerformanceGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("room_id", strconv.FormatUint(r.RoomID, 10))
	if r.FlowSource != "" {
		values.Set("flow_source", string(r.FlowSource))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// RoomFlowPerformanceGetResponse 获取直播间流量表现 API Response
type RoomFlowPerformanceGetResponse struct {
	model.BaseResponse
	Data *RoomFlowPerformance `json:"data,omitempty"`
}

// RoomFlowPerformance 直播间流量表现
type RoomFlowPerformance struct {
	// LiveDouPlus 抖+直播
	LiveDouPlus int64 `json:"live_dou_plus,omitempty"`
	// LiveAwemePromote 抖音号推广直播
	LiveAwemePromote int64 `json:"live_aweme_promote,omitempty"`
	// LiveLuban 鲁班直播
	LiveLuban int64 `json:"live_luban,omitempty"`
	// LiveBrandRedirect 品牌引流直播间
	LiveBrandRedirect int64 `json:"live_brand_redirect,omitempty"`
	// LiveBrandTop 品牌toplive
	LiveBrandTop int64 `json:"live_brand_top,omitempty"`
	// LiveBrandDirect 品牌直投直播间(feedslive)
	LiveBrandDirect int64 `json:"live_brand_direct,omitempty"`
	// LiveOtherBidding 其他竞价直播
	LiveOtherBidding int64 `json:"live_other_bidding,omitempty"`
	// LivePcQianChuan 千川直播PC
	LivePcQianChuan int64 `json:"live_pc_qian_chuan,omitempty"`
	// LiveXiaoDianQianChuan 千川直播小店随心推
	LiveXiaoDianQianChuan int64 `json:"live_xiao_dian_qian_chuan,omitempty"`
	// RoomPromotionLive 直播间推广
	RoomPromotionLive int64 `json:"room_promotion_live,omitempty"`
	// QianChuanBrandFeedsLiveAd 千川品牌feedslive广告
	QianChuanBrandFeedsLiveAd int64 `json:"qian_chuan_brand_feeds_live_ad,omitempty"`
	// QianChuanBrandFeedsLiveSearch 千川品牌feedslive服务
	QianChuanBrandFeedsLiveSearch int64 `json:"qian_chuan_brand_feeds_live_search,omitempty"`
	// Nature 自然流量
	Nature int64 `json:"nature,omitempty"`
	// VideoToLive 短视频引流
	VideoToLive int64 `json:"video_to_live,omitempty"`
	// HomepageHot 推荐feed
	HomepageHot int64 `json:"homepage_hot,omitempty"`
	// LiveMerge 直播广场
	LiveMerge int64 `json:"live_merge,omitempty"`
	// HomepageFresh 同城
	HomepageFresh int64 `json:"homepage_fresh,omitempty"`
	// OhterRecommendLive 其他推荐场景
	OtherRecommendLive int64 `json:"other_recommend_live,omitempty"`
	// HomepageFollow 关注
	HomepageFollow int64 `json:"homepage_follow,omitempty"`
	// GeneralSearch 搜索
	GeneralSearch int64 `json:"general_search,omitempty"`
	// OthersHomepage 个人主页
	OthersHomepage int64 `json:"others_homepage"`
	// DouyinShoppingCenter 抖音商城
	DouyinShoppingCenter int64 `json:"douyin_shopping_center,omitempty"`
	// Activity 活动页
	Activity int64 `json:"activity,omitempty"`
	// TouxiSaas 头条西瓜
	TouxiSaas int64 `json:"touxi_saas,omitempty"`
	// AllSource 全部来源
	AllSource int64 `json:"all_source,omitempty"`
	// AllAdSource 全部广告流量
	AllAdSource int64 `json:"all_ad_source,omitempty"`
	// AllBrand 全部品牌广
	AllBrand int64 `json:"all_brand,omitempty"`
	// AllOtherBidding 全部其他竞价广告
	AllOtherBidding int64 `json:"all_other_bidding,omitempty"`
	// AllBidding 全部竞价广告
	AllBidding int64 `json:"all_bidding,omitempty"`
	// Other 有可能是广告作弊流量
	Other int64 `json:"other,omitempty"`
	// AllEcomFlowSource 广告大屏-流量来源
	AllEcomFlowSource int64 `json:"all_ecom_flow_source,omitempty"`
	// AllEcomGmvSource 广告大屏-成交来源
	AllEcomGmvSource int64 `json:"all_ecom_gmv_source,omitempty"`
	// EcomAdSecond 付费流量
	EcomAdSecond int64 `json:"ecom_ad_second,omitempty"`
	// EcomLiveSecond 直播自然推荐流量
	EcomLiveSecond int64 `json:"ecom_live_second,omitempty"`
	// EcomOther 其他
	EcomOther int64 `json:"ecom_other,omitempty"`
	// EcomAdThirdQcBrand 付费流量-千川品牌
	EcomAdThirdQcBrand int64 `json:"ecom_ad_third_qc_brand,omitempty"`
	// EcomAdThirdQcBrandOtherBidding 付费流量-其他竞价广告
	EcomAdThirdQcBrandOtherBidding int64 `json:"ecom_ad_third_qc_brand_other_bidding,omitempty"`
	// EcomAdThirdQcOtherBrand 付费流量-其他品牌广告
	EcomAdThirdQcOtherBrand int64 `json:"ecom_ad_third_qc_other_brand,omitempty"`
}
