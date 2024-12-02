package enum

// AudiencePackageFiltetEvent 自定义过滤事件
type AudiencePackageFiltetEvent string

const (
	// AudiencePackageFiltetEvent_AD_CONVERT_EXTERNAL_ACTION 优化目标
	AudiencePackageFiltetEvent_AD_CONVERT_EXTERNAL_ACTION AudiencePackageFiltetEvent = "AD_CONVERT_EXTERNAL_ACTION"
	// AudiencePackageFiltetEvent_AD_CONVERT_TYPE_ACTIVE 激活
	AudiencePackageFiltetEvent_AD_CONVERT_TYPE_ACTIVE AudiencePackageFiltetEvent = "AD_CONVERT_TYPE_ACTIVE"
	// AudiencePackageFiltetEvent_AD_CONVERT_TYPE_ACTIVE_REGISTER 注册
	AudiencePackageFiltetEvent_AD_CONVERT_TYPE_ACTIVE_REGISTER AudiencePackageFiltetEvent = "AD_CONVERT_TYPE_ACTIVE_REGISTER"
	// AudiencePackageFiltetEvent_AD_CONVERT_TYPE_PAY 付费
	AudiencePackageFiltetEvent_AD_CONVERT_TYPE_PAY AudiencePackageFiltetEvent = "AD_CONVERT_TYPE_PAY"
)
