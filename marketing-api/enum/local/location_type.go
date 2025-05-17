package local

// LocationType 区域内人群定向设置
type LocationType string

const (
	// LocationType_ALL 该地区的所有用户
	LocationType_ALL LocationType = "ALL"
	// LocationType_CURRENT 正在该地区的用户（默认值）
	LocationType_CURRENT LocationType = "CURRENT"
	// LocationType_HOME 居住在该地区的用户
	LocationType_HOME LocationType = "HOME"
	// LocationType_TRAVEL 到该地区旅行的用户
	LocationType_TRAVEL LocationType = "TRAVEL"
)
