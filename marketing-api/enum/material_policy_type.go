package enum

// MaterialPolicyType 素材政策类型
type MaterialPolicyType string

const (
	// VALID_ORIGINAL_MATERIAL_RATE 有效首投占比
	VALID_ORIGINAL_MATERIAL_RATE MaterialPolicyType = "VALID_ORIGINAL_MATERIAL_RATE"
	// VALID_FIRST_EFFECTIVE_MATERIAL 有效首发素材
	VALID_FIRST_EFFECTIVE_MATERIAL MaterialPolicyType = "VALID_FIRST_EFFECTIVE_MATERIAL"
	// VALID_HIGH_QUALITY_MATERIAL 有效优质素材
	VALID_HIGH_QUALITY_MATERIAL MaterialPolicyType = "VALID_HIGH_QUALITY_MATERIAL"
)