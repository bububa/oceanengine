package coupon

// CouponStatus 活动状态
type CouponStatus string

const (
	// CouponStatus_UNAUDITED 新建未审核
	CouponStatus_UNAUDITED CouponStatus = "UNAUDITED"
	// CouponStatus_NORMAL 正常
	CouponStatus_NORMAL CouponStatus = "NORMAL"
	// CouponStatus_PAUSE 暂停
	CouponStatus_PAUSE CouponStatus = "PAUSE"
	// CouponStatus_AUDIT_DOING 审核中
	CouponStatus_AUDIT_DOING CouponStatus = "AUDIT_DOING"
	// CouponStatus_AUDIT_FAIL 审核失败
	CouponStatus_AUDIT_FAIL CouponStatus = "AUDIT_FAIL"
	// CouponStatus_OFFLINE 已下线
	CouponStatus_OFFLINE CouponStatus = "OFFLINE"
	// CouponStatus_DELETED 已删除
	CouponStatus_DELETED CouponStatus = "DELETED"
)

// Coupon 卡券
type Coupon struct {
	// Title 活动标题，不超过15字
	Title string `json:"title,omitempty"`
	// DeliverStart 活动开始时间，格式：%Y-%m-%d，
	// 小于结束时间
	// 若小于当天则立即生效
	DeliverStart string `json:"deliver_start,omitempty"`
	// DeliverEnd 活动结束时间，格式：%Y-%m-%d
	DeliverEnd string `json:"deliver_end,omitempty"`
	// NeedPhone 是否收集手机号
	// 1：收集
	// 0：不收集(默认)
	NeedPhone int `json:"need_phone,omitempty"`
	// NeedSmsVerify 是否对手机号进行验证
	// 1：验证
	// 0：不验证
	// need_phone=0时，该字段无意义
	NeedSmsVerify int `json:"need_sms_verify,omitempty"`
	// CreateTime 创建时间，格式：%Y-%m-%d %H:%M:%S
	CreateTime string `json:"create_time,omitempty"`
	// UpdateTime 更新时间，格式：%Y-%m-%d %H:%M:%S
	UpdateTime string `json:"update_time,omitempty"`
	// Status 活动状态:
	// UNAUDITED：新建未审核
	// NORMAL：正常
	// PAUSE：暂停
	// AUDIT_DOING：审核中
	// AUDIT_FAIL：审核失败
	// OFFLINE： 已下线
	// DELETED：已删除
	Status CouponStatus `json:"status,omitempty"`
	// GlobalLimit 总库存量限制
	GlobalLimit *LimitSetting `json:"global_limit,omitempty"`
	// UserLimit 单个用户的库存量限制
	UserLimit *LimitSetting `json:"user_limit,omitempty"`
	// ResourceList 优惠券列表
	ResourceList []Resource `json:"resource_list,omitempty"`
}
