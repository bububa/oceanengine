package creativeword

import "github.com/bububa/oceanengine/marketing-api/enum"

// CreativeWord 创意词包
type CreativeWord struct {
	// ID 创意词包ID
	ID uint64 `json:"creative_word_id,omitempty"`
	// Name 创意词包名称
	Name string `json:"name,omitempty"`
	// DefaultWord 默认词
	DefaultWord string `json:"default_word,omitempty"`
	// Words 替换词
	Words []string `json:"words,omitempty"`
	// ContentType 创意词包类型
	ContentType enum.CreativeWordType `json:"content_type,omitempty"`
	// MaxWordLen 替换词最大长度
	MaxWordLen int `json:"max_word_len,omitempty"`
	// MinWordLen 替换词最小长度
	MinWordLen int `json:"min_word_len,omitempty"`
	// Status 创意词包状态
	Status enum.CreativeWordStatus `json:"status,omitempty"`
	// UserRate 创意词包人群覆盖率;取值范围: 0-1
	UserRate float64 `json:"user_rate,omitempty"`
}
