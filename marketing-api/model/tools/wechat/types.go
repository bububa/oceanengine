package wechat

// ListTimeRange 按创建时间查询的时间范围
type ListTimeRange struct {
	// StartTime   创建起始时间，格式：%Y-%m-%d
	StartTime string `json:"start_time,omitempty"`
	// EndTime 创建结束时间，格式：%Y-%m-%d
	EndTime string `json:"end_time,omitempty"`
}
