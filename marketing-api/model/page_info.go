package model

// PageInfo 通用翻页数据
type PageInfo struct {
	// Page 当前页码
	Page int `json:"page,omitempty"`
	// PageSize 每页item个数
	PageSize int `json:"page_size,omitempty"`
	// TotalNumber 总item个数
	TotalNumber int64 `json:"total_number,omitempty"`
	// TotalPage 总页数
	TotalPage int `json:"total_page,omitempty"`
}
