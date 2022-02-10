package enum

// OrderType 排序方式
type OrderType string

const (
	// DESC 反向排序
	DESC OrderType = "DESC"
	// ASC 正向排序
	ASC OrderType = "ASC"
)
