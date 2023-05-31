package analyse

// Product 商品
type Product struct {
	// ProductID 商品id
	ProductID uint64 `json:"product_id,omitempty"`
	// SaleTime 上架时间，时间戳
	SaleTime int64 `json:"sale_time,omitempty"`
	// ProductRate 好评率
	ProductRate float64 `json:"product_rate,omitempty"`
	// ProductName 商品名称
	ProductName string `json:"product_name,omitempty"`
	// ProductImg 商品图片url列表
	ProductImg []string `json:"product_img,omitempty"`
	// DiscountPrice 售卖价格
	DiscountPrice float64 `json:"discount_price,omitempty"`
	// MarketPrice 原价
	MarketPrice float64 `json:"market_price,omitempty"`
	// Recommendation 推荐语
	Recommendation string `json:"recommendation,omitempty"`
	// CategoryName 一级类目
	CategoryName string `json:"category_name,omitempty"`
	Metrics
}
