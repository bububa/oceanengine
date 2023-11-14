package enum

// DeliveryPkgStatus 推广产品整组的审核状态，标识该推广产品是否审核通过
type DeliveryPkgStatus string

const (
	// DeliveryPkgStatus_STATUS_CONFIRM 审核通过
	DeliveryPkgStatus_STATUS_CONFIRM DeliveryPkgStatus = "STATUS_CONFIRM"
	// DeliveryPkgStatus_STATUS_CONFIRM_FAIL 审核不通过
	DeliveryPkgStatus_STATUS_CONFIRM_FAIL DeliveryPkgStatus = "STATUS_CONFIRM_FAIL"
	// DeliveryPkgStatus_STATUS_NOT_SUBMIT 未提交
	DeliveryPkgStatus_STATUS_NOT_SUBMIT DeliveryPkgStatus = "STATUS_NOT_SUBMIT"
	// DeliveryPkgStatus_STATUS_PENDING_CONFIRM 审核中
	DeliveryPkgStatus_STATUS_PENDING_CONFIRM DeliveryPkgStatus = "STATUS_PENDING_CONFIRM"
	// DeliveryPkgStatus_STATUS_WAIT_CONFIRM 待审核
	DeliveryPkgStatus_STATUS_WAIT_CONFIRM DeliveryPkgStatus = "STATUS_WAIT_CONFIRM"
)
