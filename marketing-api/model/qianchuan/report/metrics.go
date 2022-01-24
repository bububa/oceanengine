package report

// Metrics 维度数据
type Metrics struct {
	// StatDatetime 数据起始时间，分组条件包含 STAT_GROUP_BY_FIELD_STAT_TIME 时返回，格式为：yyyy-MM-dd HH:mm:ss
	StatDatetime string `json:"stat_datetime,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告组id
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// CreativeID 创意id
	CreativeID uint64 `json:"creative_id,omitempty"`
	// StatCost 消耗
	StatCost float64 `json:"stat_cost,omitempty"`
	// ShowCnt 展示数
	ShowCnt int64 `json:"show_cnt,omitempty"`
	// Ctr 点击率
	Ctr float64 `json:"ctr,omitempty"`
	// CpmPlatform 平均千次展示成本
	CpmPlatform float64 `json:"cpm_platform,omitempty"`
	// ClickCnt 点击数
	ClickCnt int64 `json:"click_cnt,omitempty"`
	// PayOrderCount 成交订单数
	PayOrderCount int64 `json:"pay_order_count,omitempty"`
	// CreateOrderAmount 下单成交金额
	CreateOrderAmount float64 `json:"create_order_amount,omitempty"`
	// CreateOrderCount 下单订单数
	CreateOrderCount int64 `json:"create_order_count,omitempty"`
	// PayOrderAmount 成交订单金额
	PayOrderAmount float64 `json:"pay_order_amount,omitempty"`
	// CreateOrderROI 下单roi
	CreateOrderROI float64 `json:"create_order_roi,omitempty"`
	// PrepayAndPayOrderROI 支付roi
	PrepayAndPayOrderROI float64 `json:"prepay_and_pay_order_roi,omitempty"`
	// PrepayOrderCount 广告预售订单数
	PrepayOrderCount int64 `json:"prepay_order_count,omitempty"`
	// PrepayOrderAmount 广告预售订单金额
	PrepayOrderAmount float64 `json:"prepay_order_amount,omitempty"`
	// ConvertCnt 转化数
	ConvertCnt int64 `json:"convert_cnt,omitempty"`
	// ConversionCost 转化成本
	ConversionCost float64 `json:"conversion_cost,omitempty"`
	// ConversionRate 转化率
	ConversionRate float64 `json:"conversion_rate,omitempty"`
	// DyFollow 新增关注数
	DyFollow int64 `json:"dy_follow,omitempty"`
	// DyLike 点赞数
	DyLike int64 `json:"dy_like,omitempty"`
	// DyComment 评论提交数
	DyComment int64 `json:"dy_comment,omitempty"`
	// DyShare 分享数
	DyShare int64 `json:"dy_share,omitempty"`
	// LivePayOrderCostPerOrder 成交客单价。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LivePayOrderCostPerOrder int64 `json:"live_pay_order_cost_per_order,omitempty"`
	// LubanLiveEnterCnt 直播间观看人次。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveEnterCnt int64 `json:"luban_live_enter_cnt,omitempty"`
	// LiveWatchOneMinuteCount 直播间超过1分钟观看人次。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LiveWatchOneMinuteCount int64 `json:"live_watch_one_minute_count,omitempty"`
	// LiveFansClubJoinCnt 直播间新加团人次。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LiveFansClubJoinCnt int64 `json:"live_fans_club_join_cnt,omitempty"`
	// LubanLiveSlidecartClickCnt 直播间查看购物车次数。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveSlidecartClickCnt int64 `json:"luban_live_slidecart_click_cnt,omitempty"`
	// LubanLiveClickProductCnt 直播间商品点击次数。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveClickProductCnt int64 `json:"luban_live_click_product_cnt,omitempty"`
	// LubanLiveCommentCnt 直播间评论次数。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveCommentCnt int64 `json:"luban_live_comment_cnt,omitempty"`
	// LubanLiveShareCnt 直播间分享次数。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveShareCnt int64 `json:"luban_live_share_cnt,omitempty"`
	// LubanLiveGiftCnt 直播间打赏次数。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveGiftCnt int64 `json:"luban_live_gift_cnt,omitempty"`
	// LubanLiveGiftAmount 直播间音浪收入。短视频带货：VIDEO_PROM_GOODS 不支持该指标
	LubanLiveGiftAmount float64 `json:"luban_live_gift_amount,omitempty"`
	// TotalPlay 播放数
	TotalPlay int64 `json:"total_play,omitempty"`
	// PlayDuration3s 3s播放数
	PlayDuration3s int64 `json:"play_duration_3s,omitempty"`
	// Play25FeedBreak 25%进度播放数
	Play25FeedBreak int64 `json:"play_25_feed_break,omitempty"`
	// Play50FeedBreak 50%进度播放数
	Play50FeedBreak int64 `json:"play_50_feed_break,omitempty"`
	// Play75FeedBreak 75%进度播放数
	Play75FeedBreak int64 `json:"play_75_feed_break,omitempty"`
	// PlayOver 播放完成数
	PlayOver int64 `json:"play_over,omitempty"`
	// PlayOverRate 完播率
	PlayOverRate float64 `json:"play_over_rate,omitempty"`
}
