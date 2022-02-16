package coupon

// Gift 游戏礼包
type Gift struct {
	// Count 礼品数量
	Count int `json:"count,omitempty"`
	// IconURL 礼品icon地址
	IconURL string `json:"icon_url,omitempty"`
	// Name 礼品名称，不超过12字
	Name string `json:"name,omitempty"`
}
