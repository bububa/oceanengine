package enum

// AsyncTaskType 任务类型
type AsyncTaskType string

const (
	// AsyncTaskType_REPORT 普通报表
	AsyncTaskType_REPORT AsyncTaskType = "REPORT"
	// AsyncTaskType_REPORT_DPA DPA 报表
	AsyncTaskType_REPORT_DPA AsyncTaskType = "REPORT_DPA"
	// AsyncTaskType_REPORT_BIDWORD 关键词/搜索词报表
	AsyncTaskType_REPORT_BIDWORD AsyncTaskType = "REPORT_BIDWORD"
)
