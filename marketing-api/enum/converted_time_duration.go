package enum

// ConvertedTimeDuration 过滤已转化投放时间
type ConvertedTimeDuration string

const (
	// ConvertedTimeDuration_SEVEN_DAY 七天（默认值）
	ConvertedTimeDuration_SEVEN_DAY ConvertedTimeDuration = "SEVEN_DAY"
	// ConvertedTimeDuration_ONE_MONTH 1个月
	ConvertedTimeDuration_ONE_MONTH ConvertedTimeDuration = "ONE_MONTH"
	// ConvertedTimeDuration_THREE_MONTH 3个月
	ConvertedTimeDuration_THREE_MONTH ConvertedTimeDuration = "THREE_MONTH"
	// ConvertedTimeDuration_SIX_MONTH 6个月
	ConvertedTimeDuration_SIX_MONTH ConvertedTimeDuration = "SIX_MONTH"
	// ConvertedTimeDuration_TWELVE_MONTH 12个月
	ConvertedTimeDuration_TWELVE_MONTH ConvertedTimeDuration = "TWELVE_MONTH"
	// ConvertedTimeDuration_TODAY 当天
	ConvertedTimeDuration_TODAY ConvertedTimeDuration = "TODAY"
)
