package event

// LandingPage 三方落地页数据
type LandingPage struct {
	// AssetID 三方的资产id
	AssetID uint64 `json:"asset_id,omitempty"`
	// AssetName 三方落地页名称
	AssetName string `json:"asset_name,omitempty"`
}
