package enterprise

import "github.com/bububa/oceanengine/marketing-api/enum/enterprise"

// OperationLog 推广操作记录
type OperationLog struct {
	// CreateTime 操作时间
	CreateTime string `json:"create_time,omitempty"`
	// BusinessPageOperationType 企业号推广操作类型
	// LIVE 直播开始、VIDEO 发布视频、DOU 开启DOU+推广、AD 开启广告投放
	BusinessPageOperationType enterprise.OperationType `json:"business_page_operation_type,omitempty"`
	// RoomID 直播间ID
	RoomID uint64 `json:"room_id,omitempty"`
	// RoomTitle 直播间名称
	RoomTitle string `json:"room_title,omitempty"`
	// RoomCover 直播间封面
	RoomCover string `json:"room_cover,omitempty"`
	// VideoID 视频ID
	VideoID uint64 `json:"video_id,omitempty"`
	// VideoImage 视频缩略图
	VideoImage string `json:"video_image,omitempty"`
	// Budget DOU+投放金额,元
	Budget float64 `json:"budget,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 广告计划ID列表
	AdID []uint64 `json:"ad_id,omitempty"`
}
