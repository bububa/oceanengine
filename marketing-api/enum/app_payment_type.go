package enum

// AppPaymentType 付费类型
type AppPaymentType string

const (
	// AppPaymentType_FREE 免费下载，且不包含付费内容
	AppPaymentType_FREE AppPaymentType = "FREE"
	// AppPaymentType_PRODUCT 免费下载，但应用内包含付费内容，如道具等虚拟物品
	AppPaymentType_PRODUCT AppPaymentType = "PRODUCT"
	// AppPaymentType_TRIAL_AND_LIMIT_FUNCTIONS 付费使用完整功能，且应用内包含付费内容，如道具等虚拟物品
	AppPaymentType_TRIAL_AND_LIMIT_FUNCTIONS AppPaymentType = "TRIAL_AND_LIMIT_FUNCTIONS"
	// AppPaymentType_TRIAL_AND_PURCHASE 付费使用完整功能，但应用内不包含任何付费内容
	AppPaymentType_TRIAL_AND_PURCHASE AppPaymentType = "TRIAL_AND_PURCHASE"
)
