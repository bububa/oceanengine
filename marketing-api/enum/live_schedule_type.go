package enum

// 直播投放日期选择方式
type LiveScheduleType string

const (
	// LiveScheduleType_SCHEDULE_TIME_ALLDAY 全天
	LiveScheduleType_SCHEDULE_TIME_ALLDAY LiveScheduleType = "SCHEDULE_TIME_ALLDAY"
	// LiveScheduleType_SCHEDULE_TIME_WEEKLY_SETTING 指定时间段
	LiveScheduleType_SCHEDULE_TIME_WEEKLY_SETTING LiveScheduleType = "SCHEDULE_TIME_WEEKLY_SETTING"
	// LiveScheduleType_SCHEDULE_TIME_FIXEDRANGE 固定时长
	LiveScheduleType_SCHEDULE_TIME_FIXEDRANGE LiveScheduleType = "SCHEDULE_TIME_FIXEDRANGE"
)
