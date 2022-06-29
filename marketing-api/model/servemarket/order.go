package servemarket

// OrderStatus 订单状态
type OrderStatus string

const (
	// UNPAID 待支付
	UNPAID OrderStatus = "UNPAID"
	// CANCELED 取消
	CANCELED OrderStatus = "CANCELED"
	// OVERTIME_CANCELED 超时取消
	OVERTIME_CANCELED OrderStatus = "OVERTIME_CANCELED"
	// SERVING 服务中，应用类订单目前无该状态
	SERVING OrderStatus = "SERVING"
	// REFUND_APPLY 申请退款
	REFUND_APPLY OrderStatus = "REFUND_APPLY"
	// REFUND_REJECT 拒绝退款
	REFUND_REJECT OrderStatus = "REFUND_REJECT"
	// REFUND_APPROVED 退款已审批
	REFUND_APPROVED OrderStatus = "REFUND_APPROVED"
	// REFUND_SUCCESS 退款成功
	REFUND_SUCCESS OrderStatus = "REFUND_SUCCESS"
	// DELIVERED 已发货
	DELIVERED OrderStatus = "DELIVERED"
	// FINISHED 已完成
	FINISHED OrderStatus = "FINISHED"
	// OVERTIME_FINISHED 超时完成
	OVERTIME_FINISHED = "OVERTIME_FINISHED"
	// DELIVERY_REJECT 发货拒绝
	DELIVERY_REJECT = "DELIVERY_REJECT"
)

// SkuType 购买的商品规格类型
type SkuType string

const (
	// FREE 免费
	FREE SkuType = "FREE"
	// TRY 试用
	TRY SkuType = "TRY"
	// PAY 付费
	PAY SkuType = "PAY"
)

// OrderFunction 功能点信息（千川应用暂不支持该类型）
type OrderFunction struct {
	// Key 功能点key（千川应用暂不支持该类型）
	Key string `json:"key,omitempty"`
	// Name 功能点名称（千川应用暂不支持该类型）
	Name string `json:"name,omitempty"`
	// Desc 功能点描述（千川应用暂不支持该类型）
	Desc string `json:"desc,omitempty"`
}

// Order 订单
type Order struct {
	// OrderID 订单id
	OrderID uint64 `json:"order_id,omitempty"`
	// OrderStatus 订单状态
	OrderStatus OrderStatus `json:"order_status,omitempty"`
	// PaidUserID 下单用户uid
	PaidUserID uint64 `json:"paid_user_id,omitempty"`
	// AppID 应用ID
	AppID uint64 `json:"app_id,omitempty"`
	// SkuID 购买的商品规格ID
	SkuID uint64 `json:"sku_id,omitempty"`
	// SkuDescription 购买的商品规格描述
	SkuDescription string `json:"sku_description,omitempty"`
	// SkuType 购买的商品规格类型
	SkuType SkuType `json:"sku_type,omitempty"`
	// IsFunc 是否是付费功能点类型的sku（千川应用暂不支持该类型）
	IsFunc bool `json:"is_func,omitempty"`
	// Function 功能点信息（千川应用暂不支持该类型）
	Function *OrderFunction `json:"function,omitempty"`
	// Fee 应付价格，单位分
	Fee float64 `json:"fee,omitempty"`
	// AppLimitUserCount 购买的商品规格的授权账户数量
	AppLimitUserCount int64 `json:"app_limit_user_count,omitempty"`
	// AppAvailableUserIDs 可使用用户uid list，即下单人填写的应用可使用的帐号列表
	AppAvailableUserIDs []uint64 `json:"app_available_user_ids,omitempty"`
	// AppActiveDays 购买的商品规格的服务周期，单位天
	AppActiveDays int `json:"app_active_days,omitempty"`
	// CreateTime 下单时间（UNIX时间戳，单位ms）
	CreateTime int64 `json:"create_time,omitempty"`
	// PayTime 付款时间（UNIX时间戳，单位ms）
	PayTime int64 `json:"pay_time,omitempty"`
	// BeginTime 购买生效期的开始时间（UNIX时间戳，单位ms）
	BeginTime int64 `json:"begin_time,omitempty"`
	// EndTime 购买生效期的结束时间（UNIX时间戳，单位ms）
	EndTime int64 `json:"end_time,omitempty"`
}
