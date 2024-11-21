package enum

// StarAdUniteTaskItemStatus 视频状态 可选值:
type StarAdUniteTaskItemStatus string

const (
	// StarAdUniteTaskItemStatus_AUDIT_PASS 审核通过
	StarAdUniteTaskItemStatus_AUDIT_PASS StarAdUniteTaskItemStatus = "AUDIT_PASS"
	// StarAdUniteTaskItemStatus_CREATED 视频已发布
	StarAdUniteTaskItemStatus_CREATED StarAdUniteTaskItemStatus = "CREATED"
	// StarAdUniteTaskItemStatus_PRIVATE_AREA 仅个人可见
	StarAdUniteTaskItemStatus_PRIVATE_AREA StarAdUniteTaskItemStatus = "PRIVATE_AREA"
	// StarAdUniteTaskItemStatus_USER_DELETED 用户删除
	StarAdUniteTaskItemStatus_USER_DELETED StarAdUniteTaskItemStatus = "USER_DELETED"
)
