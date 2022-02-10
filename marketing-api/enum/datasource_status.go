package enum

// DataSourceStatus 人群数据更新状态
type DataSourceStatus int

const (
	// DataSourceStatus_CREATED 新建
	DataSourceStatus_CREATED DataSourceStatus = 0
	// DataSourceStatus_PARSING 处理中
	DataSourceStatus_PARSING DataSourceStatus = 1
	// DataSourceStatus_COMPLETED 生效
	DataSourceStatus_COMPLETED DataSourceStatus = 2
	// DataSourceStatus_FAILED 失败
	DataSourceStatus_FAILED DataSourceStatus = 3
)
