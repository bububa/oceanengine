package enum

// AdvertiserReportType 自运营报备类型
type AdvertiserReportType string

const (
	// AdvertiserReportType_EMPTY 不报备
	AdvertiserReportType_EMPTY AdvertiserReportType = "EMPTY"
	// AdvertiserReportType_INCREASE_QUANTITY 收量报备
	AdvertiserReportType_INCREASE_QUANTITY AdvertiserReportType = "INCREASE_QUANTITY"
	// AdvertiserReportType_SELF_OPERATION 自运营报备
	AdvertiserReportType_SELF_OPERATION AdvertiserReportType = "SELF_OPERATION"
	// AdvertiserReportType_DECREASE_QUANTITY 走量
	AdvertiserReportType_DECREASE_QUANTITY AdvertiserReportType = "DECREASE_QUANTITY"
)
