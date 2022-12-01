package enum

// InterestActionMode 行为兴趣
type InterestActionMode string

const (
	// InterestActionMode_UNLIMITED 不限
	InterestActionMode_UNLIMITED InterestActionMode = "UNLIMITED"
	// InterestActionMode_CUSTOM 自定义
	InterestActionMode_CUSTOM InterestActionMode = "CUSTOM"
	// InterestActionMode_RECOMMEND 系统推荐
	InterestActionMode_RECOMMEND InterestActionMode = "RECOMMEND"
)
