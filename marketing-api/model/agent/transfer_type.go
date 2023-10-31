package agent

// AdvertiserTransferType 转账类型
type AdvertiserTransferType string

const (
	// AdvertiserTransferType_GRANT 赠款
	AdvertiserTransferType_GRANT AdvertiserTransferType = "GRANT"
	// AdvertiserTransferType_PREPAY_UNIVERSAL 通用预付
	AdvertiserTransferType_PREPAY_UNIVERSAL AdvertiserTransferType = "PREPAY_UNIVERSAL"
	// AdvertiserTransferType_PREPAY_BRAND 品牌预付
	AdvertiserTransferType_PREPAY_BRAND AdvertiserTransferType = "PREPAY_BRAND"
	// AdvertiserTransferType_PREPAY_BID 竞价预付
	AdvertiserTransferType_PREPAY_BID AdvertiserTransferType = "PREPAY_BID"
	// AdvertiserTransferType_CREDIT_UNIVERSAL 通用授信
	AdvertiserTransferType_CREDIT_UNIVERSAL AdvertiserTransferType = "CREDIT_UNIVERSAL"
	// AdvertiserTransferType_CREDIT_BRAND 品牌授信
	AdvertiserTransferType_CREDIT_BRAND AdvertiserTransferType = "CREDIT_BRAND"
	// AdvertiserTransferType_CREDIT_BID 竞价授信
	AdvertiserTransferType_CREDIT_BID AdvertiserTransferType = "CREDIT_BID"
)

// FundTransferType 转账类型(方舟)
type FundTransferType string

const (
	// FundTransferType_CASH_DEFAULT 非赠款金额
	FundTransferType_CASH_DEFAULT FundTransferType = "CASH_DEFAULT"
	// FundTransferType_CREDIT_BIDDING 授信-竞价专用金额
	FundTransferType_CREDIT_BIDDING FundTransferType = "CREDIT_BIDDING"
	// FundTransferType_CREDIT_BRAND 授信-品牌专用金额
	FundTransferType_CREDIT_BRAND FundTransferType = "CREDIT_BRAND"
	// FundTransferType_CREDIT_GENERAL 授信-通用金额
	FundTransferType_CREDIT_GENERAL FundTransferType = "CREDIT_GENERAL"
	// FundTransferType_GRANT_GENERAL 赠款金额
	FundTransferType_GRANT_GENERAL FundTransferType = "GRANT_GENERAL"
	// FundTransferType_PREPAY_BIDDING 预付-竞价专用金额
	FundTransferType_PREPAY_BIDDING FundTransferType = "PREPAY_BIDDING"
	// FundTransferType_PREPAY_BRAND 预付-品牌专用金额
	FundTransferType_PREPAY_BRAND FundTransferType = "PREPAY_BRAND"
	// FundTransferType_PREPAY_GENERAL 预付-通用金额
	FundTransferType_PREPAY_GENERAL FundTransferType = "PREPAY_GENERAL"
)
