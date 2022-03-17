package enum

// AdConvertStatus 转化工具转化状态
type AdConvertStatus string

const (
	// AD_CONVERT_STATUS_ACTIVE 活跃（激活）
	AD_CONVERT_STATUS_ACTIVE AdConvertStatus = "AD_CONVERT_STATUS_ACTIVE"
	// AD_CONVERT_STATUS_INACTIVE 不活跃（未激活）
	AD_CONVERT_STATUS_INACTIVE AdConvertStatus = "AD_CONVERT_STATUS_INACTIVE"
)
