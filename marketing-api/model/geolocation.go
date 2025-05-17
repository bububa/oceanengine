package model

// Geolocation 地理位置信息
type Geolocation struct {
	// City 城市名
	City string `json:"city,omitempty"`
	// StreetNumber 街道号
	StreetNumber string `json:"street_number,omitempty"`
	// Street 街道名
	Street string `json:"street,omitempty"`
	// District 区域名
	District string `json:"district,omitempty"`
	// Province 省份名
	Province string `json:"province,omitempty"`
	// Name 地点名称
	Name string `json:"name,omitempty"`
	// Radius 半径
	Radius int64 `json:"radius,omitempty"`
	// Long 经度
	Long float64 `json:"long,omitempty"`
	// Lat 纬度
	Lat float64 `json:"lat,omitempty"`
}
