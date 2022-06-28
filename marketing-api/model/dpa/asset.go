package dpa

// Asset 物件
type Asset struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AssetID 物件id
	AssetID uint64 `json:"asset_id,omitempty"`
	// PlatformID 商品库id
	PlatformID uint64 `json:"platform_id,omitempty"`
	// AssetType 物件类型，AUTO为汽车
	AssetType string `json:"asset_type,omitempty"`
	// ProductID 商品id
	ProductID uint64 `json:"product_id,omitempty"`
	// Status 物件状态，DISABLE代表暂停， ENABLE代表启用
	Status string `json:"status,omitempty"`
	// Source 来源，MANUAL为用户操作，AUTO为系统生成
	Source string `json:"source,omitempty"`
	// AssetCreateTime 物件创建时间，格式: yyyy-MM-DD
	AssetCreateTime string `json:"asset_create_time,omitempty"`
	// AssetModifyTime 物件最近一次修改时间，格式: yyyy-MM-DD
	AssetModifyTime string `json:"asset_modify_time,omitempty"`
}
