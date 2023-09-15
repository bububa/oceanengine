package report

// Metrics 指标数据
type Metrics struct {
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
	// EcpConvertRate 转化率
	EcpConvertRate float64 `json:"ecp_convert_rate,omitempty"`
	// EcpConvertCnt 转化数
	EcpConvertCnt int64 `json:"ecp_convert_cnt,omitempty"`
	// EcpConvertPlatform 转化成本
	EcpConvertPlatform float64 `json:"ecp_convert_platform,omitempty"`
	// EcpCpaPlatform 转化成本
	EcpCpaPlatform float64 `json:"ecp_cpa_platform,omitempty"`
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
	// ValidPlay 有效播放数
	ValidPlay int64 `json:"valid_play,omitempty"`
	// AveragePlayTimePerPlay 平均单次播放时长
	AveragePlayTimePerPlay float64 `json:"average_play_time_per_play,omitempty"`
	// ValidPlayRate 有效播放率
	ValidPlayRate float64 `json:"valid_play_rate,omitempty"`
	// CpcPlatform 平均点击单价
	CpcPlatform float64 `json:"cpc_platform,omitempty"`
	// DeepConvertCnt 深度转化次数
	DeepConvertCnt int64 `json:"deep_convert_cnt,omitempty"`
	// DeepConvertCost 深度转化成本
	DeepConvertCost float64 `json:"deep_convert_cost,omitempty"`
	// DeepConvertRate 深度转化率
	DeepConvertRate float64 `json:"deep_convert_rate,omitempty"`
	// AttributionConvertCnt 转化数（计费时间）
	AttributionConvertCnt int64 `json:"attribution_convert_cnt,omitempty"`
	// AttributionConvertRate 转化率（计费时间）
	AttributionConvertRate float64 `json:"attribution_convert_rate,omitempty"`
	// AttributionConvertCost 转化成本（计费时间）
	AttributionConvertCost float64 `json:"attribution_convert_cost,omitempty"`
	// AttributionDeepConvertCnt 深度转化次数（计费时间）
	AttributionDeepConvertCnt int64 `json:"attribution_deep_convert_cnt,omitempty"`
	// AttributionDeepConvertCost 深度转化成本（计费时间）
	AttributionDeepCOnvertCost float64 `json:"attribution_deep_convert_cost,omitempty"`
	// AttributionDeepConvertRate 深度转化率（计费时间）
	AttributionDeepConvertRate float64 `json:"attribution_deep_convert_rate,omitempty"`
	// AllOrderCreateRoi7Days 7日总下单ROI
	AllOrderCreateRoi7Days float64 `json:"all_order_create_roi_7days,omitempty"`
	// AllOrderPayRoi7Days 7日总支付ROI
	AllOrderPayRoi7Days float64 `json:"all_order_pay_roi_7days,omitempty"`
	// AllOrderPayCount7Days 7日总成交订单
	AllOrderPayCount7Days int64 `json:"all_order_pay_count_7days,omitempty"`
	// AllOrderPayGmv7Days 7日总成交金额
	AllOrderPayGmv7Days float64 `json:"all_order_pay_gmv_7days,omitempty"`
	// PayOrderCostPerOrder 直接成交客单价
	PayOrderCostPerOrder float64 `json:"pay_order_cost_per_order,omitempty"`
	// CreateOrderCouponAmount 下单智能优惠券金额
	CreateOrderCouponAmount float64 `json:"create_order_coupon_amount,omitempty"`
	// PayOrderCouponAmount 支付智能优惠券金额
	PayOrderCouponAmount float64 `json:"pay_order_coupon_amount,omitempty"`
	// IndirectOrderCreateCount7Days 间接下单订单数
	IndirectOrderCreateCount7Days int64 `json:"indirect_order_create_count_7days,omitempty"`
	// IndirectOrderCrreateGmv7Days 间接下单金额
	IndirectOrderCrreateGmv7Days float64 `json:"indirect_order_crreate_gmv_7days,omitempty"`
	// IndirectOrderPayCount7Days 间接支付订单数
	IndirectOrderPayCount7Days int64 `json:"indirect_order_pay_count_7days,omitempty"`
	// IndirectOrderPayGmv7Days 间接支付金额
	IndirectOrderPayGmv7Days float64 `json:"indirect_order_pay_gmv_7days,omitempty"`
	// IndirectOrderPrePayCount7Days 间接预售订单数
	IndirectOrderPrePayCount7Days int64 `json:"indirect_order_pre_pay_count_7days,omitempty"`
	// IndirectOrderPrePayGmv7Days 间接预售金额
	IndirectOrderPrePayGmv7Days float64 `json:"indirect_order_pre_pay_gmv_7days,omitempty"`
	// QianchuanFistOrderCnt 店铺首单新客人数
	QianchuanFirstOrderCnt int64 `json:"qianchuan_first_order_cnt,omitempty"`
	// QianchuanFistOrderRate 店铺首单新客订单占比
	QianchuanFirstOrderRate float64 `json:"qianchuan_first_order_rate,omitempty"`
	// QianchuanFirstOrderConvertCost 店铺首单新客转化成本
	QianchuanFirstOrderConvertCost float64 `json:"qianchuan_first_order_convert_cost,omitempty"`
	// QianchuanFirstOrderDirectPayGmv 店铺首单新客直接成交金额
	QianchuanFirstOrderDirectPayGmv float64 `json:"qianchuan_first_order_direct_pay_gmv,omitempty"`
	// QianchuanFirstOrderDirectPayOrderRoi 店铺首单新客直接支付ROI
	QianchuanFirstOrderDirectPayOrderRoi float64 `json:"qianchuan_first_order_direct_pay_order_roi,omitempty"`
	// QianchuanFirstOrderLtv30 店铺首单新客30天累计成交金额
	QianchuanFirstOrderLtv30 float64 `json:"qianchuan_first_order_ltv_30,omitempty"`
	// QianchuanFirstOrderRoi30 店铺首单新客30天累计支付ROI
	QianchuanFirstOrderRoi30 float64 `json:"qianchuan_first_order_roi_30,omitempty"`
	// QianchuanBrandFirstOrderRate 品牌首单新客订单占比
	QianchuanBrandFirstOrderRate float64 `json:"qianchuan_brand_first_order_rate,omitempty"`
	// QianchuanBrandFirstOrderConvertCost 品牌首单新客转化成本
	QianchuanBrandFirstOrderConvertCost float64 `json:"qianchuan_brand_first_order_convert_cost,omitempty"`
	// QianchuanBrandFirstOrderLtv30 品牌首单新客30天累计成交金额
	QianchuanBrandFirstOrderLtv30 float64 `json:"qianchuan_brand_first_order_ltv_30,omitempty"`
	// QianchuanBrandFirstOrderDirectPayOrderRoi 品牌首单新客直接支付roi
	QianchuanBrandFirstOrderDirectPayOrderRoi float64 `json:"qianchuan_brand_first_order_direct_pay_order_roi,omitempty"`
	// QianchuanBrandFirstOrderRoi30 品牌首单新客30天累计支付roi
	QianchuanBrandFirstOrderRoi30 float64 `json:"qianchuan_brand_first_order_roi_30,omitempty"`
	// QianchuanBrandFirstOrderCnt 品牌首单新客数
	QianchuanBrandFirstOrderCnt int64 `json:"qianchuan_brand_first_order_cnt,omitempty"`
	// QianchuanBrandFirstOrderDirectPayGmv 品牌首单新客直接成交金额
	QianchuanBrandFirstOrderDirectPayGmv float64 `json:"qianchuan_brand_first_order_direct_pay_gmv,omitempty"`
	// QianchuanAuthorFirstOrderRate 抖音号首单新客订单占比
	QianchuanAuthorFirstOrderRate float64 `json:"qianchuan_author_first_order_rate,omitempty"`
	// QianchuanAuthorFirstOrderRoi30 抖音号首单新客30天累计支付roi
	QianchuanAuthorFirstOrderRoi30 float64 `json:"qianchuan_author_first_order_roi_30,omitempty"`
	// QianchuanAuthorFirstOrderCnt 抖音号首单新客数
	QianchuanAuthorFirstOrderCnt int64 `json:"qianchuan_author_first_order_cnt,omitempty"`
	// QianchuanAuthorFirstOrderConvertCost 抖音号首单新客转化成本
	QianchuanAuthorFirstOrderConvertCost float64 `json:"qianchuan_author_first_order_convert_cost,omitempty"`
	// QianchuanAuthorFirstOrderDirectPayGmv 抖音号首单新客直接成交金额
	QianchuanAuthorFirstOrderDirectPayGmv float64 `json:"qianchuan_author_first_order_direct_pay_gmv,omitempty"`
	// QianchuanAuthorFIrstOrderDirectPayOrderRoi 抖音号首单新客直接支付roi
	QianchuanAuthorFIrstOrderDirectPayOrderRoi float64 `json:"qianchuan_author_first_order_direct_pay_order_roi,omitempty"`
	// QianchuanAuthorFirstOrderLtv30 抖音号首单新客30天累计成交金额
	QianchuanAuthorFirstOrderLtv30 float64 `json:"qianchuan_author_first_order_ltv_30,omitempty"`
	// LiveOrderSettleAmount14d 14天结算金额
	LiveOrderSettleAmount14d float64 `json:"live_order_settle_amount_14d,omitempty"`
	// LiveOrderSettleCount14d 14天结算订单数
	LiveOrderSettleCount14d int64 `json:"live_order_settle_count_14d,omitempty"`
	// LiveOrderSettleCountRate14d 14天结算订单率
	LiveOrderSettleCountRate14d float64 `json:"live_order_settle_count_rate_14d,omitempty"`
	// AdLiveOrderSettleRoi14d 14天结算roi
	AdLiveOrderSettleRoi14d float64 `json:"ad_live_order_settle_roi_14d,omitempty"`
	// AdLiveOrderSettleCost14d 14天结算成本
	AdLiveOrderSettleCost14d float64 `json:"ad_live_order_settle_cost_14d,omitempty"`
	// DislikeCnt 不感兴趣数
	DislikeCnt int64 `json:"dislike_cnt,omitempty"`
	// ReportCnt 举报数
	ReportCnt int64 `json:"report_cnt,omitempty"`
	// UserLoseCnt 流失数
	UserLoseCnt int64 `json:"user_lose_cnt,omitempty"`
}
