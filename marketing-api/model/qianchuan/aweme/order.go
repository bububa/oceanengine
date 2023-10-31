package aweme

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/live"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/product"
)

// Order 订单详情
type Order struct {
	// OrderID 订单id
	OrderID uint64 `json:"order_id,omitempty"`
	// AdID 广告计划id，获取订单数据时使用
	AdID uint64 `json:"ad_id,omitempty"`
	// MarketingGoal 营销目标
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// Status 订单状态 可选值:
	// AUDIT: 审核中
	// BOOK: 未开播
	// CREATING: 已支付，计划创建中
	// DELIVERY_OK: 投放中
	// FAILED: 计划创建失败
	// FINISHED: 投放完成
	// FROZEN: 投放终止
	// OFFLINE_AUDIT: 审核不通过
	// OVER: 投放结束
	// TIMEOUT: 未支付超过15分钟，订单关闭
	// UNPAID: 未支付
	// UNPAIDCANCEL: 未支付取消订单
	Status qianchuan.OrderStatus `json:"status,omitempty"`
	// OrderCreateTime 订单创建时间
	OrderCreateTime string `json:"order_create_time,omitempty"`
	// AwemeInfo 素材作者信息
	AwemeInfo *Aweme `json:"aweme_info,omitempty"`
	// ProductInfo 商品信息
	Product *product.Product `json:"product,omitempty"`
	// RoomInfo 直播间信息
	RoomInfo *live.Room `json:"room_info,omitempty"`
	// VideoInfo 视频信息
	VideoInfo *Video `json:"video_info,omitempty"`
	// DeliverySetting 投放设置
	DeliverySetting *DeliverySetting `json:"delivery_setting,omitempty"`
	// Audience 人群定向
	Audience *Audience `json:"audience,omitempty"`
	// AuditRecord 建议详细内容
	AuditRecord *AuditRecord `json:"audit_record,omitempty"`
	// CouponInfo 优惠券信息
	CouponInfo *Coupon `json:"coupon_info,omitempty"`
}
