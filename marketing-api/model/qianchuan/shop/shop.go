package shop

// Shop shop账号
type Shop struct {
	// ShopID 店铺id
	ShopID uint64 `json:"shop_id,omitempty"`
	// ShopName 店铺名称
	ShopName string `json:"shop_name,omitempty"`
}
