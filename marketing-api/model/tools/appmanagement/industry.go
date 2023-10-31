package appmanagement

// Industry 应用分类及题材标签信息
type Industry struct {
	// ID 分类id
	ID string `json:"id,omitempty"`
	// Level 级别
	Level int `json:"level,omitempty"`
	// Name 分类名称
	Name string `json:"name,omitempty"`
	// Children 子分类
	Children []Industry `json:"children,omitempty"`
	// ThemeTags 题材标签列表
	ThemeTags []ThemeTag `json:"theme_tags,omitempty"`
}

// ThemeTag 题材标签
type ThemeTag struct {
	// ID 题材id
	ID string `json:"id,omitempty"`
	// Name 题材名称
	Name string `json:"name,omitempty"`
}
