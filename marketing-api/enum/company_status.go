package enum

// CompanyStatus 公司主体状态
type CompanyStatus string

const (
	// CompanyStatus_EXPIRED 已过期
	CompanyStatus_EXPIRED CompanyStatus = "EXPIRED"
	// CompanyStatus_FAILED 审核拒绝
	CompanyStatus_FAILED CompanyStatus = "FAILED"
	// CompanyStatus_NOT_STARTED 待提交
	CompanyStatus_NOT_STARTED CompanyStatus = "NOT_STARTED"
	// CompanyStatus_PROCESSING 审核中
	CompanyStatus_PROCESSING CompanyStatus = "PROCESSING"
	// CompanyStatus_SUCCESS 审核通过
	CompanyStatus_SUCCESS CompanyStatus = "SUCCESS"
	// CompanyStatus_WAITING 待审核
	CompanyStatus_WAITING CompanyStatus = "WAITING"
)
