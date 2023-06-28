package aweme

// Coupon 优惠券信息
type Coupon struct {
	// CouponAmount 券的总金额
	CouponAmount float64 `json:"coupon_amount,omitempty"`
	// CouponType 券类型。1表示满减券；4表示满赠券
	CouponType int `json:"coupon_type,omitempty"`
}
