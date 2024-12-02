package enum

// HideIfConverted 过滤已转化的维度，可以避免该广告再次投放给已转化过的用户，不传默认不过滤
type HideIfConverted string

const (
	// HideIfConverted_NOT_EXCLUDE 不过滤
	HideIfConverted_NOT_EXCLUDE HideIfConverted = "NOT_EXCLUDE"
	// HideIfConverted_APP 应用
	HideIfConverted_APP HideIfConverted = "APP"
	// HideIfConverted_ORGANIZATION 组织（仅加白广告主可用）
	HideIfConverted_ORGANIZATION HideIfConverted = "ORGANIZATION"
	// HideIfConverted_ADVERTISER 广告主账户
	HideIfConverted_ADVERTISER HideIfConverted = "ADVERTISER"
	// HideIfConverted_CC 组织账户
	HideIfConverted_CC HideIfConverted = "CC"
	// HideIfConverted_CUSTOMER 公司账户
	HideIfConverted_CUSTOMER HideIfConverted = "CUSTOMER"
	// HideIfConverted_PROJECT 项目
	HideIfConverted_PROJECT HideIfConverted = "PROJECT"
	// HideIfConverted_PROMOTION 广告
	HideIfConverted_PROMOTION HideIfConverted = "PROMOTION"
)
