package dpa

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
