package unipromotion

import "github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"

// Material 素材
type Material struct {
	// MaterialInfo 素材信息
	MaterialInfo *MaterialInfo `json:"material_info,omitempty"`
	// MaterialStatus 素材状态，可选值:
	// DELIVERY_OK 投放中
	// DELETED 已删除
	MaterialStatus string `json:"material_status,omitempty"`
	// DeliveryNotReason 素材不可投原因
	DeliveryNotReason []string `json:"delivery_not_reason,omitempty"`
	// AuditStatus 审核状态，可选值:
	// PASS 审核通过
	// REJECT 审核拒绝
	// IN_PROGRESS 审核中
	AuditStatus string `json:"audit_status,omitempty"`
	// IsDeleted 是否已删除
	IsDeleted bool `json:"is_deleted,omitempty"`
	// MaterialSelectType  素材类型，可选值:
	// CUSTOM 自选投放素材
	// AUTO 智能优选素材
	MaterialSelectType string `json:"material_select_type,omitempty"`
	// ProductInfo 商品信息
	ProductInfo []ProductInfo `json:"product_info,omitempty"`
	// StatsInfo 指标
	StatsInfo *report.UniPromotionStats `json:"stats_info,omitempty"`
}

// MaterialInfo 素材信息
type MaterialInfo struct {
	// MaterialType 素材类型，可选值:
	// IMAGE 图片
	// LIVE_ROOM 直播间画面
	// TITLE 标题
	// VIDEO 视频
	// 图文CAROUSEL
	MaterialType string `json:"material_type,omitempty"`
	// VideoMaterial 视频素材
	VideoMaterial *VideoMaterial `json:"video_material,omitempty"`
	// TitleMaterial 标题素材
	TitleMaterial *TitleMaterial `json:"title_material,omitempty"`
	// ImageMaterial 图片素材
	ImageMaterial *ImageMaterial `json:"image_material,omitempty"`
	// RoomMaterial 直播间信息
	RoomMaterial *RoomMaterial `json:"room_material,omitempty"`
	// CarouselMaterial 图文素材
	CarouselMaterial *CarouselMaterial `json:"carousel_material,omitempty"`
}

// RoomMaterial 直播间信息
type RoomMaterial struct {
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// AwemeID 抖音号id
	AwemeID uint64 `json:"aweme_id,omitempty"`
	// AwemeName 抖音号名称
	AwemeName string `json:"aweme_name,omitempty"`
	// AwemeAvatar 抖音号头像url列表
	AwemeAvatar string `json:"aweme_avatar,omitempty"`
}
