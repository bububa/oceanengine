package model

type Geolocation struct {
	Radius       int64   `json:"radius,omitempty"`        // 半径
	Name         string  `json:"name,omitempty"`          // 地点名称
	Long         float64 `json:"long,omitempty"`          // 经度
	Lat          float64 `json:"lat,omitempty"`           // 纬度
	City         string  `json:"city,omitempty"`          // 城市名
	StreetNumber string  `json:"street_number,omitempty"` // 街道号
	Street       string  `json:"street,omitempty"`        // 街道名
	District     string  `json:"district,omitempty"`      // 区域名
	Province     string  `json:"province,omitempty"`      // 省份名
}
