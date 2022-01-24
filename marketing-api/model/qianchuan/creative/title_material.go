package creative

// TitleMaterial 标题素材
type TitleMaterial struct {
	// ID 素材唯一标识
	ID uint64 `json:"id,omitempty"`
	// Title 创意标题
	Title string `json:"title,omitempty"`
	// DynamicWords 动态词包对象列表
	DynamicWords []DynamicWord `json:"dynamic_words,omitempty"`
}
