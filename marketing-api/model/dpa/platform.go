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

// Meta 商品库元信息
type Meta struct {
	// Status 元信息状态
	Status int `json:"status,omitempty"`
	// Name 元信息名称
	Name string `json:"name,omitempty"`
	// Title 元信息头
	Title string `json:"title,omitempty"`
	// MediaType 元信息媒体类型
	MediaType int `json:"media_type,omitempty"`
	// FieldGroup 分组
	FieldGroup string `json:"field_group,omitempty"`
	// Type 字段类型
	Type string `json:"type,omitempty"`
}
