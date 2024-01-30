package enum

// ScheduleScene 生效方式
type ScheduleScene string

const (
	// ScheduleScene_NEXT_DAY 次日0点生效
	ScheduleScene_NEXT_DAY ScheduleScene = "NEXT_DAY"
	// ScheduleScene_REALTIME 立即生效
	ScheduleScene_REALTIME ScheduleScene = "REALTIME"
)
