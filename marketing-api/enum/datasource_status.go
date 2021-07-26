package enum

type DataSourceStatus int

const (
	DataSourceStatus_CREATED   DataSourceStatus = 0
	DataSourceStatus_PARSING   DataSourceStatus = 1
	DataSourceStatus_COMPLETED DataSourceStatus = 2
	DataSourceStatus_FAILED    DataSourceStatus = 3
)
