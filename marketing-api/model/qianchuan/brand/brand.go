package brand

// Brand 授权品牌
type Brand struct {
	// BrandID 品牌ID
	BrandID uint64 `json:"brand_id,omitempty"`
	// BrandName 品牌名称
	BrandName string `json:"brand_name,omitempty"`
	// BrandLogo 品牌logo
	BrandLogo string `json:"brand_logo,omitempty"`
}
