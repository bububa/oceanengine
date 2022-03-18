package tools

// District 行政层级信息
type District struct {
	// Name 行政区域名称
	Name string `json:"name,omitempty"`
	// Level 行政区域层级
	Level string `json:"level,omitempty"`
	// Code 中国大陆行政区域编码
	Code string `json:"code,omitempty"`
	// GeonameID 港澳台、国外行政区域编码
	GeonameID uint64 `json:"geoname_id,omitempty"`
	// SubDistricts 子行政层级信息
	SubDistricts []District `json:"sub_districts,omitempty"`
	// Description 行政区域类型
	// COUNTRY表示国家
	// STATE表示国家
	// REGION表示区域
	Description string `json:"description,omitempty"`
}
