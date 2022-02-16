package coupon

type EmployeeUserType string

const (
	// EmployeeUserType_TOUTIAO 头条用户
	EmployeeUserType_TOUTIAO EmployeeUserType = "TOUTIAO"
	// EmployeeUserType_DOUYIN 抖音用户
	EmployeeUserType_DOUYIN EmployeeUserType = "DOUYIN"
)

// Employee 核销员信息
type Employee struct {
	// UserID 用户ID
	UserID uint64 `json:"user_id,omitempty"`
	// UserType 用户类型
	// TOUTIAO:头条用户
	// DOUYIN:抖音用户
	UserType EmployeeUserType `json:"user_type,omitempty"`
	// StoreID 门店ID，0代表不限制门店
	StoreID uint64 `json:"store_id,omitempty"`
}
