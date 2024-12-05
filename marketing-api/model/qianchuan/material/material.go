package material

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/report"
)

// MaterialType 素材类型
type MaterialType string

const (
	// MaterialType_IMAGE 图片，图文
	MaterialType_IMAGE MaterialType = "IMAGE"
	// MaterialType_TITLE 标题
	MaterialType_TITLE MaterialType = "TITLE"
	// MaterialType_LIVE_ROOM 直播间画面
	MaterialType_LIVE_ROOM MaterialType = "LIVE_ROOM"
	// MaterialType_VIDEO 视频
	MaterialType_VIDEO MaterialType = "VIDEO"
)

// AdMaterialInfo 素材信息
type AdMaterialInfo struct {
	// MaterialInfo 素材信息
	MaterialInfo *MaterialInfo `json:"material_info,omitempty"`
	// Metrics 指标信息
	Metrics *report.Metrics `json:"metrics,omitempty"`
	// MaterialDeliveryType 素材投放状态
	MaterialDeliveryType string `json:"material_delivery_type,omitempty"`
	// AuditStatus 审核状态 可选值:
	// PASS 审核通过
	// REJECT 审核拒绝
	// IN_PROGRESS 审核中
	AuditStatus string `json:"audit_status,omitempty"`
	// CreativeIDs 关联的创意id
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
	// IsDel 是否删除
	IsDel bool `json:"is_del,omitempty"`
	// IsAutoGenerate 是否派生
	IsAutoGenerate bool `json:"is_auto_generate,omitempty"`
}

// MaterialInfo 素材信息
type MaterialInfo struct {
	// VideoMaterial 视频素材
	VideoMaterial *VideoMaterial `json:"video_material,omitempty"`
	// ImageMaterial 图片素材
	ImageMaterial *ImageMaterial `json:"image_material,omitempty"`
	// TitleMaterial 标题素材
	TitleMaterial *TitleMaterial `json:"title_material,omitempty"`
	// RoomMaterial 直播间画面用户信息
	RoomMaterial *RoomMaterial `json:"room_material,omitempty"`
	// MaterialType 素材类型
	MaterialType MaterialType `json:"material_type,omitempty"`
}

// VideoMaterial 视频素材
type VideoMaterial struct {
	// CoverImage 视频封面图片
	CoverImage *MaterialImage `json:"cover_image,omitempty"`
	// VideoID 视频 id
	VideoID string `json:"video_id,omitempty"`
	// Title 视频标题
	Title string `json:"title,omitempty"`
	// Source 视频来源
	Source enum.MaterialSource `json:"source,omitempty"`
	// ImageMode 素材样式
	ImageMode enum.MaterialMode `json:"image_mode,omitempty"`
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// VideoDuration 视频时长
	VideoDuration float64 `json:"video_duration,omitempty"`
}

// ImageMaterial 图片素材
type ImageMaterial struct {
	// Title 标题
	Title string `json:"title,omitempty"`
	// MusicURL 图文音乐播放链接
	MusicURL string `json:"music_url,omitempty"`
	// Description 图文描述
	Description string `json:"description,omitempty"`
	// ImageMode 素材样式
	ImageMode enum.MaterialMode `json:"image_mode,omitempty"`
	// Source 来源
	Source enum.MaterialSource `json:"source,omitempty"`
	// Images 图片
	Images []MaterialImage `json:"images,omitempty"`
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
}

// TitleMaterial 标题素材
type TitleMaterial struct {
	// Title 标题
	Title string `json:"title,omitempty"`
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
}

// RoomMaterial 直播间画面用户信息
type RoomMaterial struct {
	// Name 直播间名称
	Name string `json:"name,omitempty"`
	// AwemeAvatar 头像
	AwemeAvatar string `json:"aweme_avatar,omitempty"`
	// ID 直播间id
	ID uint64 `json:"id,omitempty"`
}

// MaterialImage 图片信息
type MaterialImage struct {
	// WebURL 图片url
	WebURL string `json:"web_url,omitempty"`
	// ImageURL 图片url
	ImageURL string `json:"image_url,omitempty"`
	// ID 图片id
	ID string `json:"id,omitempty"`
	// Height 图片高度
	Height int `json:"height,omitempty"`
	// Width 图片宽度
	Width int `json:"width,omitempty"`
}
