package ad

// ProductInfo 商品
type ProductInfo struct {
	// ID 商品id
	ID uint64 `json:"id,omitempty"`
	// Name 商品名称
	Name string `json:"name,omitempty"`
	// DiscountPrice 售价，已废弃
	DiscountPrice float64 `json:"discount_price,omitempty"`
	// Img 商品主图
	Img string `json:"img,omitempty"`
	// MarketPrice 原价，单位为元
	MarketPrice float64 `json:"market_price,omitempty"`
	// DiscountLowerPrice 折扣价区间最小值，单位为元
	DiscountLowerPrice float64 `json:"discount_lower_price,omitempty"`
	// DiscountHigherPrice 折扣价区间最大值，单位为元
	DiscountHigherPrice float64 `json:"discount_higher_price,omitempty"`
}
