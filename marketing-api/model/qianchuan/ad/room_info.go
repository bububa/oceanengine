package ad

// RoomInfo 直播间
type RoomInfo struct {
	// RoomTitle 直播间名称（若未开播，则返回NULL）
	RoomTitle string `json:"room_title,omitempty"`
	// RoomStatus 直播间状态（若未开播，则返回NULL）
	RoomStatus string `json:"room_status,omitempty"`
	// AnchorID 主播ID
	AnchorID uint64 `json:"anchor_id,omitempty"`
	// AnchorName 主播名称
	AnchorName string `json:"anchor_name,omitempty"`
	// AnchorAvatar 主播头像
	AnchorAvatar string `json:"anchor_avatar,omitempty"`
}
