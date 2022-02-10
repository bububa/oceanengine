package enum

// ScheduleType 计划投放时间类型
type ScheduleType string

const (
	// SCHEDULE_FROM_NOW 从现在开始一直投放
	SCHEDULE_FROM_NOW ScheduleType = "SCHEDULE_FROM_NOW"
	// SCHEDULE_START_END 选择起始时间
	SCHEDULE_START_END ScheduleType = "SCHEDULE_START_END"
)
