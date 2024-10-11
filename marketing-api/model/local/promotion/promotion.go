package promotion

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/local"
)

// Promotion 广告
type Promotion struct {
	// ProjectID 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// LocalAccountID 广告主ID
	LocalAccountID uint64 `json:"local_account_id,omitempty"`
	// AdType 广告类型，枚举值：
	// GENERAL 通投广告
	// SEARCHING 搜索广告
	AdType local.AdType `json:"ad_type,omitempty"`
	// PromotionID 广告ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// PromotionName 广告名称
	PromotionName string `json:"promotion_name,omitempty"`
	// PromotionCreateTime 广告创建时间，格式yyyy-MM-dd HH:mm:ss
	PromotionCreateTime string `json:"promotion_create_time,omitempty"`
	// PromotionModifyTime 广告修改时间，格式yyyy-MM-dd HH:mm:ss
	PromotionModifyTime string `json:"promotion_modify_time,omitempty"`
	// PromotionStatusFirst 广告一级状态
	PromotionStatusFirst local.PromotionStatusFirst `json:"promotion_status_first,omitempty"`
	// PromotionStatusSecond 广告二级状态
	PromotionStatusSecond []local.PromotionStatusSecond `json:"promotion_status_second,omitempty"`
	// LearningPhase 学习期状态，枚举值：
	// LEARNED 学习期结束
	// LEARNING 学习中
	// LEARN_FAILED 学习失败
	LearningPhase enum.LearningPhase `json:"learning_phase,omitempty"`
	// AwemeID 抖音号
	AwemeID string `json:"aweme_id,omitempty"`
	// AwemeName 抖音号昵称
	AwemeName string `json:"aweme_name,omitempty"`
}

// PromotionDetail 广告详情
type PromotionDetail struct {
	// PromotionID 广告ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// AwemeID 抖音号
	AwemeID string `json:"aweme_id,omitempty"`
	// EnableGraphicDelivery 是否开启团购卡
	EnableGraphicDelivery bool `json:"enable_graphic_delivery,omitempty"`
	// VideoHpVisibility 视频曝光
	VideoHpVisibility enum.VideoHpVisibility `json:"video_hp_visibility,omitempty"`
	// LiveMaterialType 直播素材类型，枚举值：
	// LIVE 直播素材
	// VIDEO 广告素材
	LiveMaterialType enum.MarketingGoal `json:"live_material_type,omitempty"`
	// CustomerMaterialList 自定义素材组合
	CustomerMaterialList []CustomerMaterial `json:"customer_material_list,omitempty"`
}
