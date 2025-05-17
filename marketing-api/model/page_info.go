package model

// PageInfo 通用翻页数据
type PageInfo struct {
	// Page 当前页码
	Page int `json:"page,omitempty"`
	// PageSize 每页item个数
	PageSize int `json:"page_size,omitempty"`
	// TotalNumber 总item个数
	TotalNumber Int64 `json:"total_number,omitempty"`
	// TotalPage 总页数
	TotalPage int `json:"total_page,omitempty"`
	// HasMore 是否有下一页
	HasMore int `json:"has_more,omitempty"`
	// Count 过滤后返回的视频数量，注意，此处的数量不一定与入参的count一致，因为存在过滤逻辑
	Count int `json:"count,omitempty"`
	// Cursor 下一次分页拉取的游标值
	Cursor int `json:"cursor,omitempty"`
}

// CursorInfo 分页信息
type CursorInfo struct {
	// Cursor 下一次分页拉取的游标值
	Cursor string `json:"cursor,omitempty"`
	// Count
	Count int `json:"count,omitempty"`
	// HasMore 是否有下一页
	HasMore bool `json:"has_more,omitempty"`
}
