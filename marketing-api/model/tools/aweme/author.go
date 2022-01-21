package aweme

// Author 抖音账户信息
type Author struct {
	// AwemeID 抖音id
	AwemeID string `json:"aweme_id,omitempty"`
	// LabelID 抖音号id
	LabelID uint64 `json:"label_id,omitempty"`
	// Name 抖音作者名
	Name string `json:"author_name,omitempty"`
	// Avatar 抖音头像
	Avatar string `json:"avatar,omitempty"`
	// CategoryName 抖音分类
	CategoryName string `json:"category_name,omitempty"`
	// TotalFansNumStr 粉丝数
	TotalFansNumStr string `json:"total_fans_num_str,omitempty"`
	// CoverNumStr 覆盖人群数
	CoverNumStr string `json:"cover_num_str,omitempty"`
}
