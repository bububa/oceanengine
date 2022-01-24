package qianchuan

// CreativeOptStatus 广告创意操作状态
type CreativeOptStatus string

const (
	// CreativeOptStatus_ENABLE 启用
	CreativeOptStatus_ENABLE CreativeOptStatus = "ENABLE"
	// CreativeOptStatus_DISABLE 暂停
	CreativeOptStatus_DISABLE CreativeOptStatus = "DISABLE"
	// CreativeOptStatus_DELETE 删除
	CreativeOptStatus_DELETE CreativeOptStatus = "DELETE"
)
