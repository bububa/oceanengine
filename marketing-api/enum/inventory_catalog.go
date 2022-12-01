package enum

// InventoryCatalog 广告位大类
type InventoryCatalog string

const (
	// InventoryCatalog_MANUAL 首选媒体
	InventoryCatalog_MANUAL InventoryCatalog = "MANUAL"
	// InventoryCatalog_SCENE 场景广告位
	InventoryCatalog_SCENE InventoryCatalog = "SCENE"
	// InventoryCatalog_SMART 优选广告位
	InventoryCatalog_SMART InventoryCatalog = "SMART"
	// InventoryCatalog_UNIVERSAL 通投智选
	InventoryCatalog_UNIVERSAL InventoryCatalog = "UNIVERSAL"
	// InventoryCatalog_UNIVERSAL_SMART 通投智选
	InventoryCatalog_UNIVERSAL_SMART InventoryCatalog = "UNIVERSAL_SMART"
)
