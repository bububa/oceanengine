package local

// LocalDeliveryScene 推广目的
type LocalDeliveryScene string

const (
	// LocalDeliveryScene_ALL 不限
	LocalDeliveryScene_ALL LocalDeliveryScene = "ALL"
	// LocalDeliveryScene_CONTENT_HEAT 内容加热
	LocalDeliveryScene_CONTENT_HEAT LocalDeliveryScene = "CONTENT_HEAT"
	// LocalDeliveryScene_POI_RECOMMEND 门店引流
	LocalDeliveryScene_POI_RECOMMEND LocalDeliveryScene = "POI_RECOMMEND"
	// LocalDeliveryScene_PRODUCT_PAY 团购成交
	LocalDeliveryScene_PRODUCT_PAY LocalDeliveryScene = "PRODUCT_PAY"
	// LocalDeliveryScene_EXTERNAL 销售线索收集
	LocalDeliveryScene_EXTERNAL LocalDeliveryScene = "EXTERNAL"
	// LocalDeliveryScene_CLUE 线索
	LocalDeliveryScene_CLUE LocalDeliveryScene = "CLUE"
	// LocalDeliveryScene_POI_CUSTOMER 门店引流
	LocalDeliveryScene_POI_CUSTOMER LocalDeliveryScene = "POI_CUSTOMER"
	// LocalDeliveryScene_PURCHASE 团购成交
	LocalDeliveryScene_PURCHASE LocalDeliveryScene = "PURCHASE"
)
