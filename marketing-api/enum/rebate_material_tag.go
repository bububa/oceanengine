package enum

// RebateMaterialTag 素材标签筛选项
type RebateMaterialTag string

const (
	// RebateMaterialTag_HIGH_QUALITY_MATERIAL（优质素材）
	RebateMaterialTag_HIGH_QUALITY_MATERIAL RebateMaterialTag = "HIGH_QUALITY_MATERIAL"
	// RebateMaterialTag_LOW_QUALITY_MATERIAL（低质素材）
	RebateMaterialTag_LOW_QUALITY_MATERIAL RebateMaterialTag = "LOW_QUALITY_MATERIAL"
	// RebateMaterialTag_FIRST_EFFECTIVE_MATERIAL（首发素材）
	RebateMaterialTag_FIRST_EFFECTIVE_MATERIAL RebateMaterialTag = "FIRST_EFFECTIVE_MATERIAL"
)
