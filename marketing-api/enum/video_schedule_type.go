package enum

// VideoScheduleType 短视频投放日期选择方式
type VideoScheduleType string

const (
	//VideoScheduleType_SCHEDULE_FROM_NOW 从现在开始一直投放
	VideoScheduleType_SCHEDULE_FROM_NOW VideoScheduleType = "SCHEDULE_FROM_NOW"
	//VideoScheduleType_SCHEDULE_START_END 选择起始时间
	VideoScheduleType_SCHEDULE_START_END VideoScheduleType = "SCHEDULE_START_END"
)
