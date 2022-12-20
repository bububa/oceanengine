package enum

// AdQualificationType 资质图片类型
type AdQualificationType string

const (
	// AdQualificationType_AUTHORIZATION 肖像、商标或其他授权
	AdQualificationType_AUTHORIZATION AdQualificationType = "AUTHORIZATION"
	// AdQualificationType_CERTIFY 专利证明
	AdQualificationType_CERTIFY AdQualificationType = "CERTIFY"
	// AdQualificationType_TRADEMARK_CERT 商标证
	AdQualificationType_TRADEMARK_CERT AdQualificationType = "TRADEMARK_CERT"
	// AdQualificationType_AD_DATA_CERT 广告内容中的数据证明
	AdQualificationType_AD_DATA_CERT AdQualificationType = "AD_DATA_CERT"
	// AdQualificationType_INSPECT_REPORT 质检报告
	AdQualificationType_INSPECT_REPORT AdQualificationType = "INSPECT_REPORT"
)
