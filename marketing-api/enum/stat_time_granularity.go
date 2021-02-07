package enum

// 时间粒度
type StatTimeGranularity string

const (
	STAT_TIME_GRANULARITY_DAILY  StatTimeGranularity = "STAT_TIME_GRANULARITY_DAILY"  // 按天维度
	STAT_TIME_GRANULARITY_HOURLY StatTimeGranularity = "STAT_TIME_GRANULARITY_HOURLY" // 按小时维度
	STAT_TIME_GRANULARITY_WEEK   StatTimeGranularity = "STAT_TIME_GRANULARITY_WEEK"   // 按星期粒度
	STAT_TIME_GRANULARITY_MONTH  StatTimeGranularity = "STAT_TIME_GRANULARITY_MONTH"  // 按月粒度
)
