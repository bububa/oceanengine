package form

// CountType 计数方式
type CountType string

const (
	// COUNT_TYPE_INCREMENT 递增计数
	COUNT_TYPE_INCREMENT CountType = "COUNT_TYPE_INCREMENT"
	// COUNT_TYPE_DECREMENT 递减计数
	COUNT_TYPE_DECREMENT CountType = "COUNT_TYPE_DECREMENT"
)

// SignType 表单报名配置展示形式
type SignType string

const (
	// SIGN_TYPE_SCROLL_WALL 滚动墙展示
	SIGN_TYPE_SCROLL_WALL SignType = "SIGN_TYPE_SCROLL_WALL"
	// SIGN_TYPE_SCROLL_BAR 滚动条展示
	SIGN_TYPE_SCROLL_BAR SignType = "SIGN_TYPE_SCROLL_BAR"
)

// ExtendInfo 扩展属性
type ExtendInfo struct {
	// CountConfig 计数配置，包括递增/递减、计数开始值、计数文案模版
	CountConfig *CountConfig `json:"count_config,omitempty"`
	// SignUpConfig 报名信息配置
	SignUpConfig *SignUpConfig `json:"sign_up_config,omitempty"`
	// SuccessTip 提交成功后的文案提示，文案长度不超过128个字符
	SuccessTip string `json:"success_tip,omitempty"`
}

// CountConfig 计数配置，包括递增/递减、计数开始值、计数文案模版
type CountConfig struct {
	// CountType 计数方式; 可选值：COUNT_TYPE_INCREMENT（递增计数）、COUNT_TYPE_DECREMENT（递减计数）
	CountType CountType `json:"count_type,omitempty"`
	// Prefix 计数文案前缀，例如：“截止目前已有”，不超过100个字符
	Prefix string `json:"prefix,omitempty"`
	// StartNum 起始数值
	StartNum int `json:"start_num,omitempty"`
	// Suffix 计数文案后缀，例如：“名用户参加”，不超过100个字符
	Suffix string `json:"suffix,omitempty"`
}

// SignUpConfig 报名信息配置
type SignUpConfig struct {
	// SignType 表单报名配置展示形式; SIGN_TYPE_SCROLL_WALL（滚动墙展示）、SIGN_TYPE_SCROLL_BAR（滚动条展示）
	SignType SignType `json:"sign_type,omitempty"`
}
