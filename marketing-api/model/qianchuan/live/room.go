package live

import "github.com/bububa/oceanengine/marketing-api/enum"

// Room 直播间信息
type Room struct {
	// RoomStatus 直播间状态
	RoomStatus enum.LiveRoomStatus `json:"room_status,omitempty"`
	// RoomTitle 直播间标题
	RoomTitle string `json:"room_title,omitempty"`
	// RoomCover 直播间封面图url列表
	RoomCover string `json:"room_cover,omitempty"`
	// StartTime 直播开始时间
	StartTime string `json:"start_time,omitempty"`
	// EndTime 直播结束时间
	EndTime string `json:"end_time,omitempty"`
	// AwemeName 抖音号名称
	AwemeName string `json:"aweme_name,omitempty"`
	// AwemeAvatar 抖音号头像url列表
	AwemeAvatar string `json:"aweme_avatar,omitempty"`
	// AnchorName 主播名称
	AnchorName string `json:"anchor_name,omitempty"`
	// AnchorAvatar 主播头像
	AnchorAvatar string `json:"anchor_avatar,omitempty"`
	// RoomID 直播间id
	RoomID uint64 `json:"room_id,omitempty"`
	// AwemeID 抖音号id
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// RoomDelivery 在投计划数
	RoomDelivery int64 `json:"room_delivery,omitempty"`
	// AnchorID 主播ID
	AnchorID uint64 `json:"anchor_id,omitempty"`
}
