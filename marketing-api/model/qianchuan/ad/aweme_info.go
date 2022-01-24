package ad

// AwemeInfo 计划中关联的抖音号信息
type AwemeInfo struct {
	// AwemeID 抖音ID
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// AwemeShowID 抖音号，即客户在手机端感知到的抖音号，向客户批量抖音号时请使用该字段
	AwemeShowID string `json:"aweme_show_id,omitempty"`
	// AwemeName 抖音号昵称
	AwemeName string `json:"aweme_name,omitempty"`
	// AwemeAvatar 抖音号头像
	AwemeAvatar string `json:"aweme_avatar,omitempty"`
}
