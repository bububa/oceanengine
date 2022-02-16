package coupon

// LimitSetting 限制设置
type LimitSetting struct {
	// DayLimit 每日投放的总库存，不填或0表示不限制
	DayLimit *int `json:"day_limit,omitempty"`
	// TotalLimit 总投放库存限制，不填或0表示不限制
	TotalLimit *int `json:"total_limit,omitempty"`
}
