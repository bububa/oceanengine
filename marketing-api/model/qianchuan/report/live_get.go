package report

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// LiveGetRequest 获取今日直播数据 API Request
type LiveGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AwemeID 抖音号ID
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// StartTime 直播开始时间，格式 2021-04-05 00:00:00 ，开播时间在筛选日期内的直播间，会被统计
	StartTime string `json:"start_time,omitempty"`
	// EndTime 直播结束时间，格式 2021-04-05 23:59:59，开始时间同一天的23:59:59
	// 即：查询的是一天0点～24点的直播数据
	EndTime string `json:"end_time,omitempty"`
	// Fields 需要查询的消耗指标，具体可参考返回值
	Fields []string `json:"fields,omitempty"`
	// StatsAuthority 需要查询的广告账户维度，直客账户支持查看官方抖音号下同主体全部广告账户的数据表现
	// 同主体账户：QUALIFICATION，仅直客账户支持
	// 当前账户：CURRENT，默认
	StatsAuthority qianchuan.StatsAuthority `json:"stats_authority,omitempty"`
}

// Encode implement GetRequest interface
func (r LiveGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("aweme_id", strconv.FormatUint(r.AwemeID, 10))
	values.Set("start_time", r.StartTime)
	values.Set("end_time", r.EndTime)
	if len(r.Fields) > 0 {
		values.Set("fields", string(util.JSONMarshal(r.Fields)))
	}
	if r.StatsAuthority != "" {
		values.Set("stats_authority", string(r.StatsAuthority))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// LiveGetResponse 获取今日直播数据 API Response
type LiveGetResponse struct {
	model.BaseResponse
	Data *LiveStat `json:"data,omitempty"`
}

// LiveStat 直播数据
type LiveStat struct {
	// StatCost 消耗
	StatCost float64 `json:"stat_cost,omitempty"`
	// CpmPlatform 平均千次展现费用
	CpmPlatform float64 `json:"cpm_platform,omitempty"`
	// ClickCnt 广告点击次数
	ClickCnt int64 `json:"click_cnt,omitempty"`
	// Ctr 广告点击率
	Ctr float64 `json:"ctr,omitempty"`
	// TodayLivePayOrderGpm GPM
	TodayLivePayOrderGpm float64 `json:"today_live_pay_order_gpm,omitempty"`
	// LubanLivePayOrderGpm 广告GPM
	LubanLivePayOrderGpm float64 `json:"luban_live_pay_order_gpm,omitempty"`
	// CpcPlatform 广告点击单价
	CpcPlatform float64 `json:"cpc_platform,omitempty"`
	// ConvertCnt 广告转化数
	ConvertCnt int64 `json:"convert_cnt,omitempty"`
	// ConvertRate 广告转化率
	ConvertRate float64 `json:"convert_rate,omitempty"`
	// CpaPlatform 广告转化成本
	CpaPlatform float64 `json:"cpa_platform,omitempty"`
	// LivePayOrderGmvAlias 累计成交金额
	LivePayOrderGmvAlias float64 `json:"live_pay_order_gmv_alias,omitempty"`
	// LubanLivePayOrderGmv 广告成交金额
	LubanLivePayOrderGmv float64 `json:"luban_live_pay_order_gmv,omitempty"`
	// LivePayOrderGmvRoi 整体ROI
	LivePayOrderGmvRoi float64 `json:"live_pay_order_gmv_roi,omitempty"`
	// AdLivePrepayAndPayOrderGmvRoi 广告ROI
	AdLivePrepayAndPayOrderGmvRoi float64 `json:"ad_live_prepay_and_pay_order_gmv_roi,omitempty"`
	// LiveCreateOrderCountAlias 整体下单数
	LiveCreateOrderCountAlias int64 `json:"live_create_order_count_alias,omitempty"`
	// LiveCreateOrderRate 整体下单率
	LiveCreateOrderRate float64 `json:"live_create_order_rate,omitempty"`
	// LubanLiveOrderCount 广告下单数
	LubanLiveOrderCount int64 `json:"luban_live_order_count,omitempty"`
	// AdLiveCreateOrderRate 广告下单率
	AdLiveCreateOrderRate float64 `json:"ad_live_create_order_rate,omitempty"`
	// LivePayOrderRate 整体成单数
	LivePayOrderRate float64 `json:"live_pay_order_rate,omitempty"`
	// LubanLivePayOrderCount 广告成单数
	LubanLivePayOrderCount int64 `json:"luban_live_pay_order_count,omitempty"`
	// AdLivePayOrderRate 广告成单率
	AdLivePayOrderRate float64 `json:"ad_live_pay_order_rate,omitempty"`
	// LivePayOrderGmvAvg 直播间平均成单金额
	LivePayOrderGmvAvg float64 `json:"live_pay_order_gmv_avg,omitempty"`
	// AdLivePayOrderGmvAvg 广告平均成单金额
	AdLivePayOrderGmvAvg float64 `json:"ad_live_pay_order_gmv_avg,omitempty"`
	// LubanLivePrepayOrderCount 广告预售订单数
	LubanLivePrepayOrderCount int64 `json:"luban_live_prepay_order_count,omitempty"`
	// LubanLivePrepayOrderGmv 广告预售订单金额
	LubanLivePrepayOrderGmv float64 `json:"luban_live_prepay_order_gmv,omitempty"`
	// LivePrepayOrderCountAlias 整体预售订单数
	LivePrepayOrderCountAlias int64 `json:"live_prepay_order_count_alias,omitempty"`
	// LivePrepayOrderGmvAlias 整体预售订单金额
	LivePrepayOrderGmvAlias float64 `json:"live_prepay_order_gmv_alias,omitempty"`
	// LiveOrderPayCouponAmount 成交智能优惠券金额
	LiveOrderPayCouponAmount float64 `json:"live_order_pay_coupon_amount,omitempty"`
	// TotalLiveWatchCnt 直播间观看人次
	TotalLiveWatchCnt int64 `json:"total_live_watch_cnt,omitempty"`
	// TotalLiveFollowCnt 新增粉丝数
	TotalLiveFollowCnt int64 `json:"total_live_follow_cnt,omitempty"`
	// LiveWatchOneMinuteCount 广告引导超1分钟观看人数
	LiveWatchOneMinuteCount int64 `json:"live_watch_one_minute,omitempty"`
	// TotalLiveFansClubJoinCnt 直播间新加团人数
	TotalLiveFansClubJoinCnt int64 `json:"total_live_fans_club_join_cnt,omitempty"`
	// LiveClickCartCountAlias 直播间查看购物车次数
	LiveClickCartCountAlias int64 `json:"Live_click_cart_count_alias,omitempty"`
	// LiveClickProductCountAlias 直播间商品点击次数
	LiveClickProductCountAlias int64 `json:"live_click_product_count_alias,omitempty"`
	// TotalLiveCommentCnt 直播间评论次数
	TotalLiveCommentCnt int64 `json:"total_live_comment,omitempty"`
	// TotalLiveShareCnt 直播间分享次数
	TotalLiveShareCnt int64 `json:"total_live_share_cnt,omitempty"`
	// TotalLiveGiftCnt 直播间打赏次数
	TotalLiveGiftCnt int64 `json:"total_live_gift_cnt,omitempty"`
	// TotalLiveGiftAmount 直播间音浪收入
	TotalLiveGiftAmount float64 `json:"total_live_gift_amount,omitempty"`
}
