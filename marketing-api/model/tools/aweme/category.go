package aweme

// Category 抖音类目信息
type Category struct {
	// ID 抖音类目分类
	ID uint64 `json:"id,omitempty"`
	// CoverNumStr 覆盖人群数
	CoverNumStr string `json:"cover_num_str,omitempty"`
	// FansNumStr 粉丝数
	FansNumStr string `json:"fans_num_str,omitempty"`
	// Value 抖音类目名称
	Value string `json:"value,omitempty"`
	// Children 次级分类信息
	Children []Category `json:"children,omitempty"`
}
