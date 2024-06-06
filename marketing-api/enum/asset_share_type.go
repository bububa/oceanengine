package enum

// AssetShareType 资产来源
type AssetShareType string

const (
	// AssetShareType_MY_CREATIONS 我创建的
	AssetShareType_MY_CREATIONS AssetShareType = "MY_CREATIONS"
	// AssetShareType_SHARING 共享中
	AssetShareType_SHARING AssetShareType = "SHARING"
	// AssetShareType_SHATE_EXPIRED 共享失效
	AssetShareType_SHATE_EXPIRED AssetShareType = "SHATE_EXPIRED"
)
