package liveroom

import (
	"github.com/bububa/oceanengine/marketing-api/model"
)

type Response struct {
	model.BaseResponse
	Data *ResponseData `json:"data,omitempty"`
}

type ResponseData struct {
	List     []ResponseList  `json:"list,omitempty"`
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}

type ResponseList struct {
	AnchorAvatar      string          `json:"anchor_avatar,omitempty"`       // 主播头像url
	AnchorID          uint64          `json:"anchor_id,omitempty"`           // 主播id
	AnchorNick        string          `json:"anchor_nick,omitempty"`         // 主播昵称
	RoomCover         string          `json:"room_cover,omitempty"`          // 直播间封面url
	RoomCreateTime    string          `json:"room_create_time,omitempty"`    // 直播开始时间，格式为%Y-%m-%d %H:%M:%S
	RoomFinishTime    string          `json:"room_finish_time,omitempty"`    // 直播结束时间，格式为%Y-%m-%d %H:%M:%S
	RoomID            uint64          `json:"room_id,omitempty"`             // 直播间id
	RoomQrcode        string          `json:"room_qrcode,omitempty"`         // 直播间二维码url，仅直播状态为直播中的直播间包含该指标
	RoomStatus        string          `json:"room_status,omitempty"`         // 直播间状态，"PREPARING", "LIVING", "PAUSE", "END"其中一种
	RoomTitle         string          `json:"room_title,omitempty"`          // 直播间标题
	FirstFlowCategory string          `json:"first_flow_category,omitempty"` // 一级流量来源，枚举值，0-自然流量、1-Dou+、2-竞价广告、3-品牌广告
	ProductID         uint64          `json:"product_id,omitempty"`          // 商品id
	ProductName       string          `json:"product_name,omitempty"`        // 商品名称
	ProductPicture    string          `json:"product_picture,omitempty"`     // 商品图片url
	ProductCategory   string          `json:"product_category,omitempty"`    // 商品目录
	Age               string          `json:"age,omitempty"`                 // 年龄范围
	City              string          `json:"city,omitempty"`                // 城市
	Gender            string          `json:"gender,omitempty"`              // 性别
	Platform          string          `json:"platform,omitempty"`            // 用户设备平台
	ProvinceName      string          `json:"province_name,omitempty"`       // 省份
	Fields            *ResponseFields `json:"fields,omitempty"`              // 指标
}

type ResponseFields struct {
	ClickCartCount                                int64   `json:"click_cart_count,omitempty"`                                      // 查看购物车数
	ClickProductCount                             int64   `json:"click_product_count,omitempty"`                                   // 商品点击数
	ClickProductUcount                            int64   `json:"json:click_product_ucount,omitempty"`                             // 商品点击人数
	LiveCommentCount                              int64   `json:"live_comment_count,omitempty"`                                    // 评论数
	LiveFansCount                                 int64   `json:"live_fans_count,omitempty"`                                       // 加入粉丝团数
	LiveFollowCount                               int64   `json:"live_follow_count,omitempty"`                                     // 关注数
	LiveGiftCount                                 int64   `json:"live_gift_count,omitempty"`                                       // 打赏次数
	LiveGiftMoney                                 int64   `json:"live_gift_money,omitempty"`                                       // 礼物总金额，单位：音浪
	LiveOrdersCount                               int64   `json:"live_orders_count,omitempty"`                                     // 商品下单数
	OrderUcount                                   int64   `json:"order_ucount,omitempty"`                                          // 商品下单人数
	LiveShareCount                                int64   `json:"live_share_count,omitempty"`                                      // 分享数
	LiveWatchCount                                int64   `json:"live_watch_count,omitempty"`                                      // 观看数
	LiveWatchUcount                               int64   `json:"live_watch_ucount,omitempty"`                                     // 观看人数
	LiveWatchDuration                             float64 `json:"live_watch_duration,omitempty"`                                   // 观看时长
	LiveAvgWatchDuration                          float64 `json:"live_avg_watch_duration,omitempty"`                               // 人均停留时长
	PayOrderCount                                 int64   `json:"pay_order_count,omitempty"`                                       // 商品订单数
	PayOrderUcount                                int64   `json:"pay_order_ucount,omitempty"`                                      // 商品订单人数
	PayOrderGmv                                   float64 `json:"pay_order_gmv,omitempty"`                                         // 商品订单金额
	LivePayAvgOrderGmv                            float64 `json:"live_pay_avg_order_gmv,omitempty"`                                // 场均订单金额
	LiveDuration60sCount                          int64   `json:"live_duration_60s_count,omitempty"`                               // 超1分钟观看数
	LiveDuration60sUcount                         int64   `json:"live_duration_60s_ucount,omitempty"`                              // 超1分钟观看人数
	LiveDuration60sCountToLiveWatchCountRate      float64 `json:"live_duration_60s_count_to_live_watch_count_rate,omitempty"`      // 观看数转化率（超1分钟观看数/观看数）*100%
	CountProductCountToLiveDuration60sCountRate   float64 `json:"click_product_count_to_live_duration_60s_count_rate,omitempty"`   // 超1分钟观看数转化率（商品点击数/超1分钟观看数）*100%
	LiveOrdersCountToClickProductCountRate        float64 `json:"live_orders_count_to_click_product_count_rate,omitempty"`         // 商品点击数转化率（商品下单数/商品点击数）*100%
	PayOrderCountToLiveOrdersCountRate            float64 `json:"pay_order_count_to_live_orders_count_rate,omitempty"`             // 商品下单数转化率（商品订单数/商品下单数）*100%
	LiveDuration60sUcountToLiveWatchUcountRate    float64 `json:"live_duration_60s_ucount_to_live_watch_ucount_rate,omitempty"`    // 观看人数转化率（超1分钟观看人数/观看人数）*100%
	ClickProductUcountToLiveDuration60sUcountRate float64 `json:"click_product_ucount_to_live_duration_60s_ucount_rate,omitempty"` // 超1分钟观看人数转化率（商品点击人数/超1分钟观看人数）*100%
	OrderUcountToClickProductUcountRate           float64 `json:"order_ucount_to_click_product_ucount_rate,omitempty"`             // 商品点击人数转化率（商品下单人数/商品点击人数）*100%
	PayOrderUcountToOrderUcountRate               float64 `json:"pay_order_ucount_to_order_ucount_rate,omitempty"`                 // 商品下单人数转化率（商品订单人数/商品下单人数）*100%
	LiveOnlineUcount                              int64   `json:"live_online_ucount,omitempty"`                                    // 直播间内人数
	PerCustomerTransaction                        float64 `json:"per_customer_transaction,omitempty"`                              // 客单价
}
