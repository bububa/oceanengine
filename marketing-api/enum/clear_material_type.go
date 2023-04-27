package enum

// ClearMaterialType 清理素材类型，
type ClearMaterialType string

const (
	// ClearMaterialType_INEFFICIENT_MATERIAL 低效素材
	ClearMaterialType_INEFFICIENT_MATERIAL ClearMaterialType = "INEFFICIENT_MATERIAL"
	// ClearMaterialType_SIMILAR_MATERIAL 同质化挤压严重素材
	ClearMaterialType_SIMILAR_MATERIAL ClearMaterialType = "SIMILAR_MATERIAL"
	// ClearMaterialType_SIMILAR_QUEUE_MATERIAL 同质化排队素材
	ClearMaterialType_SIMILAR_QUEUE_MATERIAL ClearMaterialType = "SIMILAR_QUEUE_MATERIAL"
)
