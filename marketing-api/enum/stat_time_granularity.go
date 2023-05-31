package enum

// StatTimeGranularity 时间粒度
type StatTimeGranularity string

const (
	// STAT_TIME_GRANULARITY_DAILY 按天维度
	STAT_TIME_GRANULARITY_DAILY StatTimeGranularity = "STAT_TIME_GRANULARITY_DAILY"
	// STAT_TIME_GRANULARITY_HOURLY 按小时维度
	STAT_TIME_GRANULARITY_HOURLY StatTimeGranularity = "STAT_TIME_GRANULARITY_HOURLY"
	// STAT_TIME_GRANULARITY_WEEK 按星期粒度
	STAT_TIME_GRANULARITY_WEEK StatTimeGranularity = "STAT_TIME_GRANULARITY_WEEK"
	// STAT_TIME_GRANULARITY_MONTH 按月粒度
	STAT_TIME_GRANULARITY_MONTH StatTimeGranularity = "STAT_TIME_GRANULARITY_MONTH"
)

// TimeGranularity 时间粒度
type TimeGranularity string

const (
	// TIME_GRANULARITY_DAILY 按天维度
	TIME_GRANULARITY_DAILY TimeGranularity = "TIME_GRANULARITY_DAILY"
	// TIME_GRANULARITY_HOURLY 按小时维度
	TIME_GRANULARITY_HOURLY TimeGranularity = "TIME_GRANULARITY_HOURLY"
	// TIME_GRANULARITY_WEEK 按星期粒度
	TIME_GRANULARITY_WEEK TimeGranularity = "TIME_GRANULARITY_WEEK"
	// TIME_GRANULARITY_MONTH 按月粒度
	TIME_GRANULARITY_MONTH TimeGranularity = "TIME_GRANULARITY_MONTH"
)
