package enum

// MaterialProperty 素材标签过滤项
type MaterialProperty string

const (
	// MaterialProperty_INEFFICIENT_MATERIAL 低效素材
	MaterialProperty_INEFFICIENT_MATERIAL MaterialProperty = "INEFFICIENT_MATERIAL"
	// MaterialProperty_SIMILAR_MATERIAL 同质化挤压严重素材
	MaterialProperty_SIMILAR_MATERIAL MaterialProperty = "SIMILAR_MATERIAL"
	// MaterialProperty_SIMILAR_QUEUE_MATERIAL 同质化素材风险-排队投放素材
	MaterialProperty_SIMILAR_QUEUE_MATERIAL MaterialProperty = "SIMILAR_QUEUE_MATERIAL"
	// MaterialProperty_SIMILAR_RISK_MATERIAL 同质化素材
	MaterialProperty_SIMILAR_RISK_MATERIAL MaterialProperty = "SIMILAR_RISK_MATERIAL"
	// MaterialProperty_HIGH_QUALITY_MATERIAL 优质素材
	MaterialProperty_HIGH_QUALITY_MATERIAL MaterialProperty = "HIGH_QUALITY_MATERIAL"
	// MaterialProperty_POOR_QUALITY_MATERIAL 低质素材
	MaterialProperty_POOR_QUALITY_MATERIAL MaterialProperty = "POOR_QUALITY_MATERIAL"
	// MaterialProperty_AD_HIGH_QUALITY_MATERIAL AD 优质素材
	MaterialProperty_AD_HIGH_QUALITY_MATERIAL MaterialProperty = "AD_HIGH_QUALITY_MATERIAL"
	// MaterialProperty_ECP_HIGH_QUALITY_MATERIAL 千川优质素材
	MaterialProperty_ECP_HIGH_QUALITY_MATERIAL MaterialProperty = "ECP_HIGH_QUALITY_MATERIAL"
	// MaterialProperty_FIRST_PUBLISH_MATERIAL  首发素材
	MaterialProperty_FIRST_PUBLISH_MATERIAL MaterialProperty = "FIRST_PUBLISH_MATERIAL"
	// MaterialProperty_LOW_EFFICIENCY_MATERIAL 低效素材
	MaterialProperty_LOW_EFFICIENCY_MATERIAL MaterialProperty = "LOW_EFFICIENCY_MATERIAL"
	// MaterialProperty_CARRY_MATERIAL 存在搬运打压风险
	MaterialProperty_CARRY_MATERIAL MaterialProperty = "CARRY_MATERIAL"
)
