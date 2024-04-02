package enum

// IesAccountStatus 抖音号绑定状态
type IesAccountStatus string

const (
	// IesAccountStatus_BINDING 绑定中
	IesAccountStatus_BINDING IesAccountStatus = "BINDING"
	// IesAccountStatus_UNBINDING 解绑
	IesAccountStatus_UNBINDING IesAccountStatus = "UNBINDING"
	// IesAccountStatus_OUT_TIME 过期
	IesAccountStatus_OUT_TIME IesAccountStatus = "OUT_TIME"
	// IesAccountStatus_NO_AUTH 资质不通过
	IesAccountStatus_NO_AUTH IesAccountStatus = "NO_AUTH"
	// IesAccountStatus_PRE_BIND 预绑定
	IesAccountStatus_PRE_BIND IesAccountStatus = "PRE_BIND"
	// IesAccountStatus_COMP_PRE_BIND 企业号预绑定
	IesAccountStatus_COMP_PRE_BIND IesAccountStatus = "COMP_PRE_BIND"
)
