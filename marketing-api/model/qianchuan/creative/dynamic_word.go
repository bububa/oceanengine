package creative

// DynamicWord 动态词包对象
type DynamicWord struct {
	// WordID 动态词包ID
	WordID uint64 `json:"word_id,omitempty"`
	// DictName 创意词包名称
	DictName string `json:"dict_name,omitempty"`
	// DefaultWord 创意词包默认词
	DefaultWord string `json:"default_word,omitempty"`
}
