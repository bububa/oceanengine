package enum

// IndustryConfigRuleType 规则的资质组成类型
type IndustryConfigRuleType string

const (
	// IndustryConfigRuleType_CHOICE 选择（资质规则内资质类型任选>1个类型填写提交）
	IndustryConfigRuleType_CHOICE IndustryConfigRuleType = "CHOICE"
	// IndustryConfigRuleType_COMPOSE 组合（资质规则内资质类型必须全部填写提交）
	IndustryConfigRuleType_COMPOSE IndustryConfigRuleType = "COMPOSE"
)
