package enum

// RebateCalcPolicyType 政策类型
type RebateCalcPolicyType string

const (
	// RebateCalcPolicyType_NORMAL_POLICY 综代政策
	RebateCalcPolicyType_NORMAL_POLICY RebateCalcPolicyType = "NORMAL_POLICY"
	// RebateCalcPolicyType_EXCLUSIVE_POLICY 优代政策
	RebateCalcPolicyType_EXCLUSIVE_POLICY RebateCalcPolicyType = "EXCLUSIVE_POLICY"
	// RebateCalcPolicyType_CAR_POLICY 汽车政策
	RebateCalcPolicyType_CAR_POLICY RebateCalcPolicyType = "CAR_POLICY"
)
