package enum

type DmpDatasourceOperationType int

const (
	DmpDatasourceOperationType_APPEND DmpDatasourceOperationType = 1
	DmpDatasourceOperationType_DELETE DmpDatasourceOperationType = 2
	DmpDatasourceOperationType_RESET  DmpDatasourceOperationType = 3
)
