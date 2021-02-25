package enum

// 创意操作状态
type CreativeOpStatus string

const (
	CREATIVE_OP_STATUS_ENABLE  CreativeOpStatus = "CREATIVE_STATUS_ENABLE"  // 启用
	CREATIVE_OP_STATUS_DISABLE CreativeOpStatus = "CREATIVE_STATUS_DISABLE" // 暂停
	CREATIVE_OP_STATUS_DELETE  CreativeOpStatus = "CREATIVE_STATUS_DELETE"  // 删除
)
