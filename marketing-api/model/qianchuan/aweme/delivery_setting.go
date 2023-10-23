package aweme

import "github.com/bububa/oceanengine/marketing-api/enum/qianchuan"

// DeliverySetting 投放设置
type DeliverySetting struct {
	// ExternalAction 优化目标
	// 当营销目标为直播带货时，允许值为：
	// AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION: 直播间商品点击
	// AD_CONVERT_TYPE_LIVE_COMMENT_ACTION: 直播间评论
	// AD_CONVERT_TYPE_LIVE_ENTER_ACTION: 进入直播间
	// AD_CONVERT_TYPE_LIVE_ROI: 支付ROI
	// AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION: 直播间下单
	// AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY: 直播间成交
	// AD_CONVERT_TYPE_NEW_FOLLOW_ACTION: 直播间粉丝提升
	// 当营销目标为短视频带货时，允许值：
	// AD_CONVERT_TYPE_QC_FOLLOW_ACTION: 粉丝提升
	// AD_CONVERT_TYPE_QC_MUST_BUY: 点赞评论
	// AD_CONVERT_TYPE_SHOPPING: 商品购买
	// AD_CONVERT_TYPE_LIVE_ROI: 商品支付ROI
	ExternalAction qianchuan.ExternalAction `json:"external_action,omitempty"`
	// 期望曝光时长，允许值：
	// 当营销目标为VIDEO_PROM_GOODS时，支持2、6、12、24、72、120
	// 当营销目标为LIVE_PROM_GOODS时，支持0.5、1、1.5、2、2.5、3、3.5、4、4.5、5、5.5、6、12、24
	DeliveryTime float64 `json:"delivery_time,omitempty"`
	// BidMode 出价模式，可选值:
	// PRICING_ACTION: 按优化目标出价
	// PRICING_PLAY: 按播放量出价
	// 注意：
	// 1. 当营销目标为VIDEO_PROM_GOODS时，两者均支持；当视频为图文时，仅支持PRICING_ACTION
	// 2. 当营销目标为LIVE_PROM_GOODS时，仅支持PRICING_ACTION
	BidMode qianchuan.BidMode `json:"bid_mode,omitempty"`
	// BidType 出价类型，可选值:
	// AUTO_BID: 自动出价
	// MANUAL_BID: 手动出价
	// 注意：
	// 出价模式为“按优化目标出价”时，必填
	// 当营销目标为LIVE_PROM_GOODS&优化目标不是AD_CONVERT_TYPE_LIVE_ROI时，仅支持AUTO_BID
	// 当marketing_goal=VIDEO_PROM_GOODS&external_action=AD_CONVERT_TYPE_LIVE_ROI时，仅支持MANUAL_BID
	// 当marketing_goal=LIVE_PROM_GOODS&external_action=AD_CONVERT_TYPE_LIVE_ROI时，支持自动出价和手动出价
	BidType qianchuan.BidType `json:"bid_type,omitempty"`
	// CpaBid 手动出价金额，0.1～10000，不高于投放金额
	// 注意：仅出价方式为“手动出价”时，必填
	CpaBid float64 `json:"cpa_bid,omitempty"`
	// RoiGoal 支付ROI目标，最多支持两位小数，0.01～100
	// 注意：仅出价方式为“手动出价”且external_action=AD_CONVERT_TYPE_LIVE_ROI时必填
	RoiGoal float64 `json:"roi_goal,omitempty"`
	// LiveroomHeatMode 直播间加热方式，可选值:
	// BY_ROOM: 直接加热直播间
	// BY_VIDEO: 选择视频加热直播间
	// 注意：仅营销目标为LIVE_PROM_GOODS时，必填
	LiveroomHeatMode qianchuan.LiveroomHeatMode `json:"liveroom_heat_mode,omitempty"`
	// Amount 投放金额
	// 注意：1、 数值需为100-5000000之间的10的倍数
	Amount int64 `json:"amount,omitempty"`
	// PaymentAmount 实际支付金额
	PaymentAmount float64 `json:"payment_amount,omitempty"`
	// PayRoi 支付ROI目标，最多支持两位小数，0.01～100
	// 注意：仅出价方式为“手动出价”且external_action=AD_CONVERT_TYPE_LIVE_ROI时必填
	PayRoi float64 `json:"pay_roi,omitempty"`
}
