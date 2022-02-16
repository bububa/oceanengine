package sitetemplate

// Template 模板
type Template struct {
	// TemplateID 模板ID
	TemplateID uint64 `json:"template_id,omitempty"`
	// SiteID 站点ID，可通过【橙子建站】平台或【获取橙子建站站点列表】接口获取
	SiteID uint64 `json:"site_id,omitempty"`
	// TemplateName 模板名称
	TemplateName string `json:"template_name,omitempty"`
	// Bricks 组件列表
	Bricks []Brick `json:"bricks,omitempty"`
}
