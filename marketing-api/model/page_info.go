package model

type PageInfo struct {
	Page        int   `json:"page,omitempty"`
	PageSize    int   `json:"page_size,omitempty"`
	TotalNumber int64 `json:"total_number,omitempty"`
	TotalPage   int   `json:"total_page,omitempty"`
}
