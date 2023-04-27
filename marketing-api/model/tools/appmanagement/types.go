package appmanagement

// TimeRange 时间范围
type TimeRange struct {
	// StartTIme 创建起始时间，默认为空，格式：%Y-%m-%d
	StartTime string `json:"start_time,omitempty"`
	// EndTime 创建结束时间，默认为当天，格式：%Y-%m-%d
	EndTime string `json:"end_time,omitempty"`
}
