package local

// ScheduleType 投放日期类型设置
type ScheduleType string

const (
	// ScheduleType_FROM_NOW_ON 从今天起长期投放
	ScheduleType_FROM_NOW_ON ScheduleType = "FROM_NOW_ON"
	// ScheduleType_START_TO_END 设置开始结束时间
	ScheduleType_START_TO_END ScheduleType = "START_TO_END"
	// ScheduleType_FIXED_TIME 固定时长
	ScheduleType_FIXED_TIME ScheduleType = "FIXED_TIME"
)
