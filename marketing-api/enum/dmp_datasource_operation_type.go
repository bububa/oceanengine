package enum

// DmpDatasourceOperationType 人群包更新操作类型
type DmpDatasourceOperationType int

const (
	// DmpDatasourceOperationType_CREATE 新建
	DmpDatasourceOperationType_CREATE DmpDatasourceOperationType = 0
	// DmpDatasourceOperationType_APPEND 添加
	DmpDatasourceOperationType_APPEND DmpDatasourceOperationType = 1
	// DmpDatasourceOperationType_DELETE 删除
	DmpDatasourceOperationType_DELETE DmpDatasourceOperationType = 2
	// DmpDatasourceOperationType_RESET 重置
	DmpDatasourceOperationType_RESET DmpDatasourceOperationType = 3
)
