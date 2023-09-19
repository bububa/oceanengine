package enterprise

// TimeDimension 是否需要分天数据
type TimeDimension string

const (
	// TimeDimension_SUM 合计数据
	TimeDimension_SUM TimeDimension = "SUM"
	// TimeDimension_DAILY 分天数据
	TimeDimension_DAILY TimeDimension = "DAILY"
)
