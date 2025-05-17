package enum

// CapitalType 可转资金类型
type CapitalType string

const (
	// CapitalType_CREDIT_BIDDING 授信竞价
	CapitalType_CREDIT_BIDDING CapitalType = "CREDIT_BIDDING"
	// CapitalType_CREDIT_BRAND 授信品牌
	CapitalType_CREDIT_BRAND CapitalType = "CREDIT_BRAND"
	// CaptialType_CREDIT_GENERAL 授信通用
	CaptialType_CREDIT_GENERAL CapitalType = "CREDIT_GENERAL"
	// CapitalType_GRANT_COMMON 信息流赠款
	CapitalType_GRANT_COMMON CapitalType = "GRANT_COMMON"
	// CapitalType_GRANT_DEFAULT 通用赠款
	CapitalType_GRANT_DEFAULT CapitalType = "GRANT_DEFAULT"
	// CapitalType_GRANT_SEARCH 搜索赠款
	CapitalType_GRANT_SEARCH CapitalType = "GRANT_SEARCH"
	// CapitalType_GRANT_UNION 穿山甲赠款
	CapitalType_GRANT_UNION CapitalType = "GRANT_UNION"
	// CaptialType_PREPAY_BIDDING 预付竞价
	CaptialType_PREPAY_BIDDING CapitalType = "PREPAY_BIDDING"
	// CaptialType_PREPAY_BRAND 预付品牌
	CaptialType_PREPAY_BRAND CapitalType = "PREPAY_BRAND"
	// CaptialType_PREPAY_GENERAL 预付通用
	CaptialType_PREPAY_GENERAL CapitalType = "PREPAY_GENERAL"
)
