package liveroom

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// Response 直播分析数据报表 API Response
type Response struct {
	model.BaseResponse
	// Data json返回值
	Data *ResponseData `json:"data,omitempty"`
}

// ResponseData json返回值
type ResponseData struct {
	// List 返回数据
	List []ResponseList `json:"list,omitempty"`
	// PageInfo 翻页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

// ResponseList 详细数据
type ResponseList struct {
	// AnchorAvatar 主播头像url
	AnchorAvatar string `json:"anchor_avatar,omitempty"`
	// AnchorID 主播id
	AnchorID uint64 `json:"anchor_id,omitempty"`
	// AnchorNick 主播昵称
	AnchorNick string `json:"anchor_nick,omitempty"`
	// RoomCover 直播间封面url
	RoomCover string `json:"room_cover,omitempty"`
	// RoomCreateTime 直播开始时间，格式为%Y-%m-%d %H:%M:%S
	RoomCreateTime string `json:"room_create_time,omitempty"`
	// RoomFinishTime 直播结束时间，格式为%Y-%m-%d %H:%M:%S
	RoomFinishTime string `json:"room_finish_time,omitempty"`
	// RoomID 直播间id
	RoomID uint64 `json:"room_id,omitempty"`
	// RoomQrcode 直播间二维码url，仅直播状态为直播中的直播间包含该指标
	RoomQrcode string `json:"room_qrcode,omitempty"`
	// RoomStatus 直播间状态，"PREPARING", "LIVING", "PAUSE", "END"其中一种
	RoomStatus enum.LiveRoomStatus `json:"room_status,omitempty"`
	// RoomTitle 直播间标题
	RoomTitle string `json:"room_title,omitempty"`
	// FirstFlowCategory 一级流量来源，枚举值，0-自然流量、1-Dou+、2-竞价广告、3-品牌广告
	FirstFlowCategory string `json:"first_flow_category,omitempty"`
	// ProductID 商品id
	ProductID uint64 `json:"product_id,omitempty"`
	// ProductName 商品名称
	ProductName string `json:"product_name,omitempty"`
	// ProductPicture 商品图片url
	ProductPicture string `json:"product_picture,omitempty"`
	// ProductCategory 商品目录
	ProductCategory string `json:"product_category,omitempty"`
	// Age 年龄范围
	Age string `json:"age,omitempty"`
	// City 城市
	City string `json:"city,omitempty"`
	// Gender 性别
	Gender string `json:"gender,omitempty"`
	// Platform 用户设备平台
	Platform string `json:"platform,omitempty"`
	// ProvinceName 省份
	ProvinceName string `json:"province_name,omitempty"`
	// Fields 指标
	Fields *ResponseFields `json:"fields,omitempty"`
}

// ResponseFields 指标
type ResponseFields struct {
	// ClickCartCount 查看购物车数
	ClickCartCount int64 `json:"click_cart_count,omitempty"`
	// ClickProductCount 商品点击数
	ClickProductCount int64 `json:"click_product_count,omitempty"`
	// ClickProductUcount 商品点击人数
	ClickProductUcount int64 `json:"json:click_product_ucount,omitempty"`
	// LiveCommentCount 评论数
	LiveCommentCount int64 `json:"live_comment_count,omitempty"`
	// LiveFansCount 加入粉丝团数
	LiveFansCount int64 `json:"live_fans_count,omitempty"`
	// LiveFollowCount 关注数
	LiveFollowCount int64 `json:"live_follow_count,omitempty"`
	// LiveGiftCount 打赏次数
	LiveGiftCount int64 `json:"live_gift_count,omitempty"`
	// LiveGiftMoney 礼物总金额，单位：音浪
	LiveGiftMoney int64 `json:"live_gift_money,omitempty"`
	// LiveOrdersCount 商品下单数
	LiveOrdersCount int64 `json:"live_orders_count,omitempty"`
	// OrderUcount 商品下单人数
	OrderUcount int64 `json:"order_ucount,omitempty"`
	// LiveShareCount 分享数
	LiveShareCount int64 `json:"live_share_count,omitempty"`
	// LiveWatchCount 观看数
	LiveWatchCount int64 `json:"live_watch_count,omitempty"`
	// LiveWatchUcount 观看人数
	LiveWatchUcount int64 `json:"live_watch_ucount,omitempty"`
	// LiveWatchDuration 观看时长
	LiveWatchDuration float64 `json:"live_watch_duration,omitempty"`
	// LiveAvgWatchDuration 人均停留时长
	LiveAvgWatchDuration float64 `json:"live_avg_watch_duration,omitempty"`
	// PayOrderCount 商品订单数
	PayOrderCount int64 `json:"pay_order_count,omitempty"`
	// PayOrderUcount 商品订单人数
	PayOrderUcount int64 `json:"pay_order_ucount,omitempty"`
	// PayOrderGmv 商品订单金额
	PayOrderGmv float64 `json:"pay_order_gmv,omitempty"`
	// LivePayAvgOrderGmv 场均订单金额
	LivePayAvgOrderGmv float64 `json:"live_pay_avg_order_gmv,omitempty"`
	// LiveDuration60sCount 超1分钟观看数
	LiveDuration60sCount int64 `json:"live_duration_60s_count,omitempty"`
	// LiveDuration60sUcount 超1分钟观看人数
	LiveDuration60sUcount int64 `json:"live_duration_60s_ucount,omitempty"`
	// LiveDuration60sCountToLiveWatchCountRate 观看数转化率（超1分钟观看数/观看数）*100%
	LiveDuration60sCountToLiveWatchCountRate float64 `json:"live_duration_60s_count_to_live_watch_count_rate,omitempty"`
	// CountProductCountToLiveDuration60sCountRate 超1分钟观看数转化率（商品点击数/超1分钟观看数）*100%
	CountProductCountToLiveDuration60sCountRate float64 `json:"click_product_count_to_live_duration_60s_count_rate,omitempty"`
	// LiveOrdersCountToClickProductCountRate 商品点击数转化率（商品下单数/商品点击数）*100%
	LiveOrdersCountToClickProductCountRate float64 `json:"live_orders_count_to_click_product_count_rate,omitempty"`
	// PayOrderCountToLiveOrdersCountRate 商品下单数转化率（商品订单数/商品下单数）*100%
	PayOrderCountToLiveOrdersCountRate float64 `json:"pay_order_count_to_live_orders_count_rate,omitempty"`
	// LiveDuration60sUcountToLiveWatchUcountRate 观看人数转化率（超1分钟观看人数/观看人数）*100%
	LiveDuration60sUcountToLiveWatchUcountRate float64 `json:"live_duration_60s_ucount_to_live_watch_ucount_rate,omitempty"`
	// ClickProductUcountToLiveDuration60sUcountRate 超1分钟观看人数转化率（商品点击人数/超1分钟观看人数）*100%
	ClickProductUcountToLiveDuration60sUcountRate float64 `json:"click_product_ucount_to_live_duration_60s_ucount_rate,omitempty"`
	// OrderUcountToClickProductUcountRate 商品点击人数转化率（商品下单人数/商品点击人数）*100%
	OrderUcountToClickProductUcountRate float64 `json:"order_ucount_to_click_product_ucount_rate,omitempty"`
	// PayOrderUcountToOrderUcountRate 商品下单人数转化率（商品订单人数/商品下单人数）*100%
	PayOrderUcountToOrderUcountRate float64 `json:"pay_order_ucount_to_order_ucount_rate,omitempty"`
	// LiveOnlineUcount 直播间内人数
	LiveOnlineUcount int64 `json:"live_online_ucount,omitempty"`
	// PerCustomerTransaction 客单价
	PerCustomerTransaction             float64 `json:"per_customer_transaction,omitempty"`
	LiveAppActivePayCount              int64   `json:"live_app_active_pay_count,omitempty"`
	LiveAppDownloadFinishCount         int64   `json:"live_app_download_finish_count,omitempty"`
	LiveGroupbuyOrderCount             int64   `json:"live_groupbuy_order_count,omitempty"`
	LiveComponentCtr                   float64 `json:"live_component_ctr,omitempty"`
	LiveCount                          int64   `json:"live_count,omitempty"`
	LiveAppDownloadStartCount          int64   `json:"live_app_download_start_count,omitempty"`
	LiveAppGamePayCount                int64   `json:"live_app_game_pay_count,omitempty"`
	LiveCardIconComponentShowCount     int64   `json:"live_card_icon_component_show_count,omitempty"`
	StatLiveGroupbuyOrderGmv           float64 `json:"stat_live_groupbuy_order_gmv,omitempty"`
	OrderConvertRate                   float64 `json:"order_convert_rate,omitempty"`
	LiveCardIconComponentClickCount    int64   `json:"live_card_icon_component_click_count,omitempty"`
	LiveFormSubmitCount                int64   `json:"live_form_submit_count,omitempty"`
	LiveAppActiveRegisterCount         int64   `json:"live_app_active_register_count,omitempty"`
	LiveAppActiveCount                 int64   `json:"live_app_active_count,omitempty"`
	LiveAppAttributionNextDayOpenRate  float64 `json:"live_app_attribution_next_day_open_rate,omitempty"`
	LiveAppAttributionNextDayOpenCount int64   `json:"live_app_attribution_next_day_open_count,omitempty"`
	LiveAppInstallFinishCount          int64   `json:"live_app_install_finish_count,omitempty"`
	LiveAvgWatchCount                  int64   `json:"live_avg_watch_count,omitempty"`
}
