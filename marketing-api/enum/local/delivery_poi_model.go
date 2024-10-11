package local

// DeliveryPoiMode 是否投放全部门店
type DeliveryPoiMode string

const (
	// DeliveryPoiMode_ALL 投放全部门店
	DeliveryPoiMode_ALL DeliveryPoiMode = "ALL"
	// DeliveryPoiMode_PART 投放指定门店
	DeliveryPoiMode_PART DeliveryPoiMode = "PART"
)
