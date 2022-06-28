package dpa

import "github.com/bububa/oceanengine/marketing-api/enum"

// Platform 商品库详情
type Platform struct {
	// PlatformID 商品库ID
	PlatformID uint64 `json:"platform_id,omitempty"`
	// Name 商品库名称
	Name string `json:"name,omitempty"`
	// ProductIndustry 商品库行业类别
	ProductIndustry enum.ProductIndustry `json:"product_industry,omitempty"`
}
