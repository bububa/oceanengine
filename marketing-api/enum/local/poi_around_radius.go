package local

// PoiAroundRadius 半径，门店附近仅支持固定半径的设置
type PoiAroundRadius string

const (
	// PoiAroundRadius_KM_6 门店附近6km
	PoiAroundRadius_KM_6 PoiAroundRadius = "KM_6"
	// PoiAroundRadius_KM_8 门店附近8km
	PoiAroundRadius_KM_8 PoiAroundRadius = "KM_8"
	// PoiAroundRadius_KM_10 门店附近10km（默认值）
	PoiAroundRadius_KM_10 PoiAroundRadius = "KM_10"
	// PoiAroundRadius_KM_12 门店附近12km
	PoiAroundRadius_KM_12 PoiAroundRadius = "KM_12"
	// PoiAroundRadius_KM_15 门店附近15km
	PoiAroundRadius_KM_15 PoiAroundRadius = "KM_15"
	// PoiAroundRadius_KM_20 门店附近20km
	PoiAroundRadius_KM_20 PoiAroundRadius = "KM_20"
)
