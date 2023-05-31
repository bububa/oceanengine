package report

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// LongTransferOrderGetRequest 长周期转化价值-订单明细 API Request
type LongTransferOrderGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartDate 开始时间，格式 2021-04-05
	StartDate time.Time `json:"start_date,omitempty"`
	// EndDate 结束时间，格式 2021-04-05，时间跨度不能超过90天
	EndDate time.Time `json:"end_date,omitempty"`
	// OptimizationTimeScope 可选值: 1、15、3、30、7
	OptimizationTimeScope int `json:"optimization_time_scope,omitempty"`
	// Filtering 过滤条件
	Filtering *LongTransferOrderGetFilter `json:"filtering,omitempty"`
	// OrderType 可选值: ASC默认、DESC
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// Page 页码，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，允许值：10, 20, 50, 100, 200，默认值：10
	PageSize int `json:"page_size,omitempty"`
}

// LongTransferOrderGetFilter 过滤条件
type LongTransferOrderGetFilter struct {
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// MarketingGoal 可选值:
	// ALL: 所有
	// LIVE_PROM_GOODS: 直播间带货
	// VIDEO_PROM_GOODS: 短视频带货
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
}

func (r LongTransferOrderGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("start_date", r.StartDate.Format("2006-01-02"))
	values.Set("end_date", r.EndDate.Format("2006-01-02"))
	values.Set("optimization_time_scope", strconv.Itoa(r.OptimizationTimeScope))
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
	}
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// LongTransferOrderGetResponse 长周期转化价值-订单明细 API Response
type LongTransferOrderGetResponse struct {
	model.BaseResponse
	Data *LongTransferOrderGetData `json:"data,omitempty"`
}

type LongTransferOrderGetData struct {
	// List 订单数据列表
	List []Order `json:"list,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// Order 订单数据
type Order struct {
	// OrderID 订单编号
	OrderID uint64 `json:"order_id,omitempty"`
	// AdID 计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// AdName 计划名称
	AdName string `json:"ad_name,omitempty"`
	// OrderPlaceTime 用户下单时间
	OrderPlaceTime string `json:"order_place_time,omitempty"`
	// IsPay 支付属性，PAID: 已支付、UNPAID: 未支付
	IsPay qianchuan.PayStatus `json:"is_pay,omitempty"`
	// PayType 订单类型，DIRECT: 直接转换订单、INDIRECT: 间接转化订单
	PayType qianchuan.PayType `json:"pay_type,omitempty"`
	// OrderFlowSource 订单来源
	// ACTIVITY: 直播-活动页面
	// DOUYIN_SHOPPING_CENTER: 直播-抖音商城
	// GENERAL_SEARCH: 直播-搜索
	// GUESS_YOU_LIKE: 短视频-猜你喜欢
	// HOMEPAGE_FOLLOW: 直播-关注来源
	// LIVE: 直播
	// LIVE_OTHER: 直播-其他
	// OTHERS_HOMEPAGE: 直播-个人主页
	// OTHER_PROFILE: 短视频-达人个人主页
	// PRODUCT_CARD: 商品卡
	// PRODUCT_CARD_GENERAL_SEARCH: 商户卡-搜索
	// PRODUCT_CARD_OTHER: 商户卡-购物车/收藏
	// QIANCHUAN_PROMOTE: 千川推广
	// RECOMMEND_LIVE: 直播-自然推荐
	// RECOMMEND_PRODUCT: 商户卡-商城/商品推荐
	// RECOMMEND_VIDEO: 短视频-短视频推荐
	// SHOP_WINDOW: 商户卡-橱窗/店铺
	// UNKNOWN: 未知
	// VIDEO: 短视频
	// VIDEO_ACTIVITY: 短视频-活动页
	// VIDEO_GENERAL_SEARCH: 短视频-搜索
	// VIDEO_HOMEPAGE_FOLLOW: 短视频-关注来源
	// VIDEO_OTHER: 短视频-其他
	// VIDEO_TO_LIVE: 直播-短视频引流
	OrderFlowSource qianchuan.OrderFlowSource `json:"order_flow_source,omitempty"`
	// OrderAmount 订单金额
	OrderAmount float64 `json:"order_amount,omitempty"`
}
