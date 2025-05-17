package product

// Product 商品信息
type Product struct {
	// ProductID 商品id
	ProductID uint64 `json:"product_id,omitempty"`
	// ProductName 商品名称
	ProductName string `json:"product_name,omitempty"`
	// ProductImageURL 商品图片
	ProductImageURL string `json:"product_image_url,omitempty"`
	// Price 商品价格，单位元
	Price float64 `json:"price,omitempty"`
	// ProductPics 商品头图，有效期2h
	ProductPics []string `json:"product_pics,omitempty"`
	// ApplicablePoiNum 商品适用门店数
	ApplicablePoiNum int `json:"applicable_poi_num,omitempty"`
}
