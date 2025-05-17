package poi

// MultiPoi 多门店信息
type MultiPoi struct {
	// MultiPoiID 多门店id
	MultiPoiID uint64 `json:"multi_poi_id,omitempty"`
	// PoiIDs 门店id列表
	PoiIDs []uint64 `json:"poi_ids,omitempty"`
}
