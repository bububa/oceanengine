package coupon

// CodeStatus 券码状态
type CodeStatus string

const (
	// CodeStatus_VALID 有效中
	CodeStatus_VALID CodeStatus = "VALID"
	// CodeStatus_EXPIRED 已过期
	CodeStatus_EXPIRED CodeStatus = "EXPIRED"
	// CodeStatus_USED 已使用
	CodeStatus_USED CodeStatus = "USED"
	// CodeStatus_ABANDONED 已废弃
	CodeStatus_ABANDONED CodeStatus = "ABANDONED"
)

// Code 券码记录
type Code struct {
	// CodeID 券码ID
	CodeID uint64 `json:"code_id,omitempty"`
	// Code 券码
	Code string `json:"code,omitempty"`
	// ActivityID 活动ID
	ActivityID uint64 `json:"activity_id,omitempty"`
	// CouponID 卡券活动ID
	CouponID uint64 `json:"coupon_id,omitempty"`
	// ResourceID 优惠券ID
	ResourceID uint64 `json:"resource_id,omitempty"`
	// ResourceTitle 优惠券标题
	ResourceTitle string `json:"resource_title,omitempty"`
	// ValidStart 有效期开始时间 格式：%Y-%m-%d
	ValidStart string `json:"valid_start,omitempty"`
	// ValidEnd 有效期结束时间 格式：%Y-%m-%d
	ValidEnd string `json:"valid_end,omitempty"`
	// Status 券码状态
	// VALID:有效中
	// EXPIRED:已过期
	// USED:已使用
	// ABANDONED:已废弃
	Status CodeStatus `json:"status,omitempty"`
	// UpdateTime 更新时间，格式：%Y-%m-%d %H:%M:%S
	// USED状态下，等于核销时间，否则等于创建时间
	UpdateTime string `json:"update_time,omitempty"`
	// CreateTime 创建时间，格式：%Y-%m-%d %H:%M:%S
	CreateTime string `json:"create_time,omitempty"`
}
