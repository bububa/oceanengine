package dpa

// Template 模板
type Template struct {
	// TemplateID 模板id
	TemplateID uint64 `json:"template_id,omitempty"`
	// TemplateName 模板名称
	TemplateName string `json:"template_name,omitempty"`
	// TemplateMode 模板类型，对应创意素材类型中的大图小图组图
	TemplateMode string `json:"template_mode,omitempty"`
	// Industry 行业
	Industry uint64 `json:"industry,omitempty"`
	// IsPublic 是否是公共模板
	IsPublic bool `json:"is_public,omitempty"`
	// TemplateDataList 模板数据列表
	TemplateDataList []TemplateData `json:"template_data_list,omitempty"`
}

// TemplateData 模板数据
type TemplateData struct {
	// WebURL 背景图
	WebURL string `json:"web_url,omitempty"`
	// TemplateImageURL 预览图
	TemplateImageURL string `json:"template_image_url,omitempty"`
}
