package advertiser

import "github.com/bububa/oceanengine/marketing-api/enum"

// IndustryConfig 资质规则
type IndustryConfig struct {
	// ConfigID 资质规则id
	ConfigID uint64 `json:"config_id,omitempty"`
	// IndustryStatus 行业状态：1 生效 、2 禁投（不允许提交该行业的资质） 可选值:
	// NONVALID 禁投
	// VALID 生效
	IndustryStatus enum.IndustryStatus `json:"industry_status,omitempty"`
	// IndustryIDs 第一级到第三级行业id
	IndustryIDs []uint64 `json:"industry_ids,omitempty"`
	// IndustryNames 第一级到第三级行业名称
	IndustryNames []string `json:"industry_names,omitempty"`
	// Necessaries 必填资质模块配置
	Necessaries []IndustryConfigNecessary `json:"necessaries,omitempty"`
	// Unnecessaries 选填资质模块配置
	Unnecessaries []IndustryConfigUnnecessary `json:"unnecessaries,omitempty"`
}

// IndustryConfigNecessary 必填资质模块配置
type IndustryConfigNecessary struct {
	// PromotionTypeID 推广类型id
	PromotionTypeID uint64 `json:"promotion_type_id,omitempty"`
	// PromotionTypeName 推广类名称
	PromotionTypeName string `json:"promotion_type_name,omitempty"`
	// Rules 具体的资质规则
	Rules []IndustryConfigRule `json:"rules,omitempty"`
}

// IndustryConfigUnnecessary 选填资质模块配置
type IndustryConfigUnnecessary struct {
	// CombineID 规则组合id
	CombineID uint64 `json:"combine_id,omitempty"`
	// Description 选填资质场景描述，用于引导用户提交
	Description string `json:"description,omitempty"`
	// Rules 具体的资质规则
	Rules []IndustryConfigRule `json:"rules,omitempty"`
}

// IndustryConfigRule 具体的资质规则
type IndustryConfigRule struct {
	// RuleID 原子规则id
	RuleID uint64 `json:"rule_id,omitempty"`
	// Type 规则的资质组成类型：1 组合资质 2 多选一资质 可选值:
	// CHOICE 选择（资质规则内资质类型任选>1个类型填写提交）
	// COMPOSE 组合（资质规则内资质类型必须全部填写提交）
	Type enum.IndustryConfigRuleType `json:"type,omitempty"`
	// Description 原子规则描述，用于引导用户提交
	Description string `json:"description,omitempty"`
	// QualTypes 资质类型
	QualTypes []IndustryConfigQualType `json:"qual_types,omitempty"`
}

// IndustryConfigQualType 资质类型
type IndustryConfigQualType struct {
	// QualType 资质类型id
	QualType uint64 `json:"qual_type,omitempty"`
	// QualTypeName 资质类型名称
	QualTypeName string `json:"qual_type_name,omitempty"`
}
