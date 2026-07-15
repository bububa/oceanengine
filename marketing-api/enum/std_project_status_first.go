package enum

// StdProjectStatusFirst 标准项目一级状态过滤
type StdProjectStatusFirst string

const (
	// StdProjectStatusFirst_PROJECT_STATUS_DELETE 已删除
	StdProjectStatusFirst_PROJECT_STATUS_DELETE StdProjectStatusFirst = "PROJECT_STATUS_DELETE"
	// StdProjectStatusFirst_PROJECT_STATUS_DONE 已完成
	StdProjectStatusFirst_PROJECT_STATUS_DONE StdProjectStatusFirst = "PROJECT_STATUS_DONE"
	// StdProjectStatusFirst_PROJECT_STATUS_DISABLE 未投放
	StdProjectStatusFirst_PROJECT_STATUS_DISABLE StdProjectStatusFirst = "PROJECT_STATUS_DISABLE"
	// StdProjectStatusFirst_PROJECT_STATUS_ENABLE 投放中
	StdProjectStatusFirst_PROJECT_STATUS_ENABLE StdProjectStatusFirst = "PROJECT_STATUS_ENABLE"
	// StdProjectStatusFirst_PROJECT_STATUS_FROZEN 已终止
	StdProjectStatusFirst_PROJECT_STATUS_FROZEN StdProjectStatusFirst = "PROJECT_STATUS_FROZEN"
	// StdProjectStatusFirst_ALL_EXCEPT_DELETE 不限（默认值）
	StdProjectStatusFirst_ALL_EXCEPT_DELETE StdProjectStatusFirst = "ALL_EXCEPT_DELETE"
	// StdProjectStatusFirst_ALL 不限（包含已删除和已终止）
	StdProjectStatusFirst_ALL StdProjectStatusFirst = "ALL"
)
