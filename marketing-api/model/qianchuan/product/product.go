package product

// Product 商品
type Product struct {
	// ID 商品id
	ID uint64 `json:"id,omitempty"`
	// Img 主图
	Img string `json:"img,omitempty"`
	// Inventory 库存
	Inventory int64 `json:"inventory,omitempty"`
	// Name 商品名称
	Name string `json:"name,omitempty"`
	// CategoryName 分类
	CategoryName string `json:"category_name,omitempty"`
	// ProductRate 好评率
	ProductRate float64 `json:"product_rate,omitempty"`
	// SaleTime 上架时间
	SaleTime string `json:"sale_time,omitempty"`
	// DiscountPrice 售价，单位：元，已废弃
	DiscountPrice float64 `json:"discount_price,omitempty"`
	// MarketPrice 原价，单位为元
	MarketPrice float64 `json:"market_price,omitempty"`
	// DiscountLowerPrice 折扣价区间最小值，单位为元
	DiscountLowerPrice float64 `json:"discount_lower_price,omitempty"`
	// DiscountHigherPrice 折扣价区间最大值，单位为元
	DiscountHigherPrice float64 `json:"discount_higher_price,omitempty"`
	// Tags 是否可投猜喜可选值:
	// 1: 可投猜喜
	// 0: 不可投猜喜
	Tags int `json:"tags,omitempty"`
	// SupportProductNewOpen 当前商品是否支持开启新品加速
	// 支持：true
	// 不支持：false
	SupportProductNewOpen bool `json:"support_product_new_open,omitempty"`
}
