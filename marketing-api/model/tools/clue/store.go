package clue

// Store 门店信息
type Store struct {
	// StoreID 门店ID
	StoreID uint64 `json:"store_id,omitempty"`
	// StoreName 门店名称
	StoreName string `json:"store_name,omitempty"`
	// StorePackID 门店活动ID
	StorePackID uint64 `json:"store_pack_id,omitempty"`
	// StoreLocation 门店所在地
	StoreLocation string `json:"store_location,omitempty"`
	// StoreAddress 门店详细地址
	StoreAddress string `json:"store_address,omitempty"`
	// StoreRemark 门店备注
	StoreRemark string `json:"store_remark,omitempty"`
	// StorePackRemark 门店活动备注
	StorePackRemark string `json:"store_pack_remark,omitempty"`
}
