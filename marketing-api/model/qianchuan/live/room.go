package live

import (
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
)

// Room 直播间信息
type Room struct {
	// RoomID 直播间id
	RoomID uint64 `json:"room_id,omitempty"`
	// RoomStatus 直播间状态
	RoomStatus qianchuan.RoomStatus `json:"room_status,omitempty"`
	// RoomTitle 直播间标题
	RoomTitle string `json:"room_title,omitempty"`
	// RoomCover 直播间封面图url列表
	RoomCover string `json:"room_cover,omitempty"`
	// StartTime 直播开始时间
	StartTime string `json:"start_time,omitempty"`
	// EndTime 直播结束时间
	EndTime string `json:"end_time,omitempty"`
	// AwemeID 抖音号id
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// AwemeName 抖音号名称
	AwemeName string `json:"aweme_name,omitempty"`
	// AwemeAvatar 抖音号头像url列表
	AwemeAvatar string `json:"aweme_avatar,omitempty"`
	// RoomDelivery 在投计划数
	RoomDelivery int64 `json:"room_delivery,omitempty"`
}
