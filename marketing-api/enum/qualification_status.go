package enum

// QualificationStatus 广告主资质审核状态
type QualificationStatus string

const (
	// QualificationStatus_STATUS_NOT_SUBMIT 未提交
	QualificationStatus_STATUS_NOT_SUBMIT QualificationStatus = "STATUS_NOT_SUBMIT"
	// QualificationStatus_STATUS_WAIT_CONFIRM 待审核
	QualificationStatus_STATUS_WAIT_CONFIRM QualificationStatus = "STATUS_WAIT_CONFIRM"
	// QualificationStatus_STATUS_PENDING_CONFIRM 审核中
	QualificationStatus_STATUS_PENDING_CONFIRM QualificationStatus = "STATUS_PENDING_CONFIRM"
	// QualificationStatus_STATUS_CONFIRM 审核通过
	QualificationStatus_STATUS_CONFIRM QualificationStatus = "STATUS_CONFIRM"
	// QualificationStatus_STATUS_CONFIRM_FAIL 审核不通过
	QualificationStatus_STATUS_CONFIRM_FAIL QualificationStatus = "STATUS_CONFIRM_FAIL"
)
