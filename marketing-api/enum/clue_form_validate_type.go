package enum

// ClueValidateType 青鸟表单-线索验证类型
type ClueValidateType string

const (
	// ClueValidateType_CLUE_PRIORITY 优先线索量 1.优先保证线索提交量；2.仅针对系统判定异常的用户进行短信验证
	ClueValidateType_CLUE_PRIORITY ClueValidateType = "CLUE_PRIORITY"
	// ClueValidateType_VALIDITY_PRIORITY 优先有效性 1.优先保证线索有效性；2.除高质量用户外、其余用户均进行短信验证
	ClueValidateType_VALIDITY_PRIORITY ClueValidateType = "VALIDITY_PRIORITY"
	// ClueValidateType_ALL_VERIFICATION 全量短信验证,所有用户都进行短信验证
	ClueValidateType_ALL_VERIFICATION ClueValidateType = "ALL_VERIFICATION"
	// ClueValidateType_NONE_VERIFICATION 不使用短信验证
	ClueValidateType_NONE_VERIFICATION ClueValidateType = "NONE_VERIFICATION"
	// ClueValidateType_AUTO_VERIFICATION 智能验证，历史逻辑，效果等同于全不出【废弃，仅获取，不支持创建】
	ClueValidateType_AUTO_VERIFICATION ClueValidateType = "AUTO_VERIFICATION"
)
