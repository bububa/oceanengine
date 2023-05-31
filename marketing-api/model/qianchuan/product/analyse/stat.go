package analyse

// Metrics 指标数据
type Metrics struct {
	// StatCost 消耗
	StatCost float64 `json:"stat_cost,omitempty"`
	// DirectOrderPayGmv 直接成交金额
	DirectOrderPayGmv float64 `json:"direct_order_pay_gmv,omitempty"`
	// DirectOrderPrepayAndPayRoi 广告支付ROI
	DirectOrderPrepayAndPayRoi float64 `json:"direct_order_prepay_and_pay_roi,omitempty"`
	// Ctr 点击率
	Ctr float64 `json:"ctr,omitempty"`
	// DirectOrderPayRate 成交转化率
	DirectOrderPayRate float64 `json:"direct_order_rate,omitempty"`
	// DirectOrderPayCost 订单转化成本
	DirectOrderPayCost float64 `json:"direct_order_pay_cost,omitempty"`
	// DirectOrderPayCount 直接成交订单数
	DirectOrderPayCount int64 `json:"direct_order_pay_count,omitempty"`
	// DirectOrderPayCostPerOrder 广告成交客单价
	DirectOrderPayCostPerOrder float64 `json:"direct_order_pay_cost_per_order,omitempty"`
	// PayOrderCouponAmount 成交智能优惠券金额
	PayOrderCouponAmount float64 `json:"pay_order_coupon_amount,omitempty"`
	// TotalPlay 播放量
	TotalPlay int64 `json:"total_play,omitempty"`
	// DyShare 分享数
	DyShare int64 `json:"dy_share,omitempty"`
	// DyLike 点赞数
	DyLike int64 `json:"dy_like,omitempty"`
	// DyComment 评论数
	DyComment int64 `json:"dy_comment,omitempty"`
	// DyFollow 新增粉丝数
	DyFollow int64 `json:"dy_follow,omitempty"`
}
