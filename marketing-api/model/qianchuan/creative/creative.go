package creative

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
)

// Creative 创意
type Creative struct {
	// AdID 计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// CreativeID 创意ID
	CreativeID uint64 `json:"creative_id,omitempty"`
	// CreativeMaterialMode 创意呈现方式，CUSTOM_CREATIVE: 自定义创意，PROGRAMMATIC_CREATIVE: 程序化创意
	CreativeMaterialMode enum.CreativeMaterialMode `json:"creative_material_mode,omitempty"`
	// Status 创意状态
	Status qianchuan.CreativeStatus `json:"status,omitempty"`
	// OptStatus 创意操作状态
	OptStatus qianchuan.CreativeOptStatus `json:"opt_status,omitempty"`
	// CreativeCreateTime 创意创建时间
	CreativeCreateTime string `json:"creative_create_time,omitempty"`
	// CreativeModifyTime 创意修改时间
	CreativeModifyTime string `json:"creative_modify_time,omitempty"`
	// ImageMode 素材类型，支持视频和图片
	ImageMode enum.MaterialMode `json:"image_mode,omitempty"`
	// VideoMaterial 视频素材
	VideoMaterial *VideoMaterial `json:"video_material,omitempty"`
	// ImageMaterial 图片素材
	ImageMaterial *ImageMaterial `json:"image_material,omitempty"`
	// TitleMaterial 标题素材
	TitleMaterial *TitleMaterial `json:"title_material,omitempty"`
	// CarouselMaterial 图文信息，对应image_mode=CAROUSEL
	CarouselMaterial *CarouselMaterial `json:"carousel_material,omitempty"`
	// LabAdType 托管计划类型，NOT_LAB_AD：非托管计划，LAB_AD：托管计划
	LabAdType enum.AdLabType `json:"ad_lab_type,omitempty"`
	// VideoMaterialList 视频素材列表
	VideoMaterialList []VideoMaterial `json:"video_material_list,omitempty"`
	// ImageMaterialList 图片素材列表
	ImageMaterialList []ImageMaterial `json:"image_material_list,omitempty"`
	// TitleMaterialList 标题素材列表
	TitleMaterialList []TitleMaterial `json:"title_material_list,omitempty"`
	// PromotionCardMaterial 推广卡片信息
	PromotionCardMaterial *PromotionCardMaterial `json:"promotion_card_material,omitempty"`
}
