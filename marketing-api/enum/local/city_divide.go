package local

// CityDivide 城市划分类型
type CityDivide string

const (
	// CityDivide_BY_LEVEL 按发展等级划分
	CityDivide_BY_LEVEL CityDivide = "BY_LEVEL"
	// CityDivide_BY_LOCATION 按地理划分（默认值）
	CityDivide_BY_LOCATION CityDivide = "BY_LOCATION"
)
