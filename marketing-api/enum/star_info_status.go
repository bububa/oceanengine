package enum

// StarInfoStatus 账户状态
type StarInfoStatus string

const (
	// StarInfoStatus_DELETED 已删除
	StarInfoStatus_DELETED StarInfoStatus = "DELETED"
	// StarInfoStatus_ENABLE 有效
	StarInfoStatus_ENABLE StarInfoStatus = "ENABLE"
	// StarInfoStatus_FROZEN 禁用
	StarInfoStatus_FROZEN StarInfoStatus = "FROZEN"
	// StarInfoStatus_NEW_PROTOCOL 待同意新协议
	StarInfoStatus_NEW_PROTOCOL StarInfoStatus = "NEW_PROTOCOL"
	// StarInfoStatus_PUNISH 惩罚
	StarInfoStatus_PUNISH StarInfoStatus = "PUNISH"
	// StarInfoStatus_QUALIFICATION_VERIFICATION 资质验证
	StarInfoStatus_QUALIFICATION_VERIFICATION StarInfoStatus = "QUALIFICATION_VERIFICATION"
	// StarInfoStatus_UN_PROTOCOL 未同意协议
	StarInfoStatus_UN_PROTOCOL StarInfoStatus = "UN_PROTOCOL"
)
