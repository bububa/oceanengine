package enum

// LocationType 位置类型
type LocationType string

const (
	// LocationType_CURRENT 正在该地区的用户
	LocationType_CURRENT LocationType = "CURRENT"
	// LocationType_HOME 居住在该地区的用户
	LocationType_HOME LocationType = "HOME"
	// LocationType_TRAVEL 到该地区旅行的用户
	LocationType_TRAVEL LocationType = "TRAVEL"
	// LocationType_ALL 该地区内的所有用户
	LocationType_ALL LocationType = "ALL"
)
