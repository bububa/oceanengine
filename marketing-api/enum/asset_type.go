package enum

// AssetType 资产类型
type AssetType string

const (
	// AssetType_THIRD_EXTERNAL 三方落地页
	AssetType_THIRD_EXTERNAL AssetType = "THIRD_EXTERNAL"
	// AssetType_TETRIS_EXTERNAL 建站
	AssetType_TETRIS_EXTERNAL AssetType = "TETRIS_EXTERNAL"
	// AssetType_APP 应用
	AssetType_APP AssetType = "APP"
	// AssetType_QUICK_APP 快应用
	AssetType_QUICK_APP AssetType = "QUICK_APP"
	// AssetType_MINI_PROGRAME 字节小程序
	AssetType_MINI_PROGRAM AssetType = "MINI_PROGRAM"
)
