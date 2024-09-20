package enum

// ScheduleType 计划投放时间类型
type ScheduleType string

const (
	// SCHEDULE_7_DAYS 7日稳投 【新增】
	SCHEDULE_7_DAYS ScheduleType = "SCHEDULE_7_DAYS"
	// SCHEDULE_FROM_NOW 从现在开始一直投放
	SCHEDULE_FROM_NOW ScheduleType = "SCHEDULE_FROM_NOW"
	// SCHEDULE_START_END 选择起始时间
	SCHEDULE_START_END ScheduleType = "SCHEDULE_START_END"
	// SCHEDULE_TIME_FIXEDRANGE 固定时长
	SCHEDULE_TIME_FIXEDRANGE ScheduleType = "SCHEDULE_TIME_FIXEDRANGE"
)
