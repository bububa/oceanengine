package promotion

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// Promotion 广告信息
type Promotion struct {
	// PromotionID 广告ID
	PromotionID uint64 `json:"promotion_id,omitempty"`
	// PromotionName 广告名称
	PromotionName string `json:"promotion_name,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ProjectID 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// ModifyTime 计划上次修改时间标识(用于更新计划时提交,服务端判断是否基于最新信息修改)
	ModifyTime string `json:"modify_time,omitempty"`
	// PromotionModifyTime 广告更新时间，格式yyyy-MM-dd HH:mm:ss
	PromotionModifyTime string `json:"promotion_modify_time,omitempty"`
	// AdCreateTime 广告创建时间，格式yyyy-MM-dd HH:mm:ss
	PromotionCreateTime string `json:"promotion_create_time,omitempty"`
	// LearningPhase 学习期状态，枚举值：DEFAULT（默认，不在学习期中）、LEARNING（学习中）、LEARNED（学习成功）、LEARN_FAILED（学习失败)
	LearningPhase enum.LearningPhase `json:"learning_phase,omitempty"`
	// Status 广告状态，枚举值：NOT_DELETED 不限 、ALL 不限（包含已删除）、OK 投放中、DELETED 已删除、PROJECT_OFFLINE_BUDGET 项目超出预算、PROJECT_PREOFFLINE_BUDGET 项目接近预算、TIME_NO_REACH 未到达投放时间、TIME_DONE 已完成、NO_SCHEDULE 不在投放时段、AUDIT 新建审核中、REAUDIT 修改审核中、FROZEN 已终止、AUDIT_DENY 审核不通过、OFFLINE_BUDGET 广告超出预算、OFFLINE_BALANCE 账户余额不足、PREOFFLINE_BUDGET 广告接近预算、DISABLED 已暂停、PROJECT_DISABLED 已被项目暂停、LIVE_ROOM_OFF 关联直播间不可投、PRODUCT_OFFLINE 关联商品不可投，、AWEME_ACCOUNT_DISABLED 关联抖音账号不可投、AWEME_ANCHOR_DISABLED 锚点不可投、DISABLE_BY_QUOTA 已暂停（配额达限）、CREATE 新建、ADVERTISER_OFFLINE_BUDGET 账号超出预算、ADVERTISER_PREOFFLINE_BUDGET 账号接近预算
	Status enum.PromotionStatus `json:"status,omitempty"`
	// OptStatus 操作状态
	OptStatus enum.OptStatus `json:"opt_status,omitempty"`
	// PromotionMaterials 广告素材组合
	PromotionMaterials []PromotionMaterial `json:"promotion_materials,omitempty"`
	// Source 广告来源
	Source string `json:"source,omitempty"`
	// Budget 预算
	Budget float64 `json:"budget,omitempty"`
	// CpaBid 目标转化出价/预期成本
	CpaBid float64 `json:"cpa_bid,omitempty"`
	// DeepCpaBid 深度优化出价
	DeepCpaBid float64 `json:"deep_cpabid,omitempty"`
	// RoiGoal 深度转化ROI系数
	RoiGoal float64 `json:"roi_goal,omitempty"`
	// MaterialScoreInfo 素材评级信息
	MaterialScoreInfo *MaterialScoreInfo `json:"material_score_info,omitempty"`
}

func (p Promotion) Version() model.AdVersion {
	return model.AdVersion_2
}

func (p Promotion) GetID() uint64 {
	return p.PromotionID
}

func (p Promotion) GetName() string {
	return p.PromotionName
}

func (p Promotion) GetCampaignID() uint64 {
	return p.ProjectID
}

func (p Promotion) GetAdvertiserID() uint64 {
	return p.AdvertiserID
}

func (p Promotion) GetOptStatus() enum.AdOptStatus {
	switch p.OptStatus {
	case enum.OptStatus_ENABLE:
		return enum.AdOptStatus_ENABLE
	case enum.OptStatus_DISABLE:
		return enum.AdOptStatus_DISABLE
	}
	return ""
}
func (p Promotion) GetBudget() float64 {
	return p.Budget
}
func (p Promotion) GetCpaBid() float64 {
	return p.CpaBid
}

func (p Promotion) GetDeepCpaBid() float64 {
	return p.DeepCpaBid
}

// PromotionMaterial 广告素材
type PromotionMaterial struct {
	// VideoMaterialList 视频素材信息
	VideoMaterialList []VideoMaterial `json:"video_material_list,omitempty"`
	// ImageMaterialList 创意图片素材
	ImageMaterialList []ImageMaterial `json:"image_material_list,omitempty"`
	// TitleMaterialList 标题素材
	TitleMaterialList []TitleMaterial `json:"title_material_list,omitempty"`
	// ComponentMaterialList 创意组件信息
	ComponentMaterialList []ComponentMaterial `json:"component_material_list,omitempty"`
	// ExternalURLMaterialList 普通落地页链接素材
	ExternalURLMaterialList []string `json:"external_url_material_list,omitempty"`
	// WebURLMaterialList Android应用下载详情页
	WebURLMaterialList []string `json:"web_url_material_list,omitempty"`
	// PlayableURLMaterialList 试玩落地页素材
	PlayableURLMaterialList []string `json:"playable_url_material_list,omitempty"`
	// ProductInfo 产品信息
	ProductInfo *ProductInfo `json:"product_info,omitempty"`
	// CallToActionButtons 行动号召文案
	CallToActionButtons []string `json:"call_to_action_buttons,omitempty"`
}

// VideoMaterial 视频素材信息
type VideoMaterial struct {
	// ImageMode 素材类型
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
	// VideoID 视频ID
	VideoID string `json:"video_id,omitempty"`
	// VideoCoverID 视频封面图片ID
	VideoCoverID string `json:"video_cover_id,omitempty"`
	// MaterialID 素材ID
	MaterialID model.FlexUint64 `json:"material_id,omitempty"`
	// MaterialStatus 素材审核状态，枚举值：
	// MATERIAL_STATUS_OK 投放中、MATERIAL_STATUS_DELETE 已删除、MATERIAL_STATUS_PROJECT_OFFLINE_BUDGET 项目超出预算、MATERIAL_STATUS_PROJECT_PREOFFLINE_BUDGET 项目接近预算、MATERIAL_STATUS_TIME_NO_REACH 未到达投放时间、MATERIAL_STATUS_TIME_DONE 已完成、MATERIAL_STATUS_NO_SCHEDULE 不在投放时段、MATERIAL_STATUS_AUDIT 新建审核中、MATERIAL_STATUS_REAUDIT 修改审核中、MATERIAL_STATUS_OFFLINE_AUDIT 审核不通过、MATERIAL_STATUS_OFFLINE_BUDGET 广告超出预算、MATERIAL_STATUS_ OFFLINE_BALANCE 账户余额不足、MATERIAL_STATUS_ PREOFFLINE_BUGDET 广告接近预算、MATERIAL_STATUS_PROJECT_DISABLE 已被项目暂停、MATERIAL_STATUS_DISABLE 已暂停、MATERIAL_STATUS_PROMOTION_DISABLE 已被广告暂停、MATERIAL_STATUS_MATERIAL_DELETE 已删除
	MaterialStatus enum.MaterialStatus `json:"material_status,omitempty"`
}

// ImageMaterial 创意图片素材
type ImageMaterial struct {
	// ImageMode 素材类型
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
	// Images 图片ID数组
	Images []Image `json:"images,omitempty"`
}

// Image 图片
type Image struct {
	// ImageID 图片ID
	ImageID string `json:"image_id,omitempty"`
	// MaterialID 素材ID
	MaterialID model.FlexUint64 `json:"material_id,omitempty"`
	// MaterialStatus 素材审核状态，枚举值：
	// MATERIAL_STATUS_OK 投放中、MATERIAL_STATUS_DELETE 已删除、MATERIAL_STATUS_PROJECT_OFFLINE_BUDGET 项目超出预算、MATERIAL_STATUS_PROJECT_PREOFFLINE_BUDGET 项目接近预算、MATERIAL_STATUS_TIME_NO_REACH 未到达投放时间、MATERIAL_STATUS_TIME_DONE 已完成、MATERIAL_STATUS_NO_SCHEDULE 不在投放时段、MATERIAL_STATUS_AUDIT 新建审核中、MATERIAL_STATUS_REAUDIT 修改审核中、MATERIAL_STATUS_OFFLINE_AUDIT 审核不通过、MATERIAL_STATUS_OFFLINE_BUDGET 广告超出预算、MATERIAL_STATUS_ OFFLINE_BALANCE 账户余额不足、MATERIAL_STATUS_ PREOFFLINE_BUGDET 广告接近预算、MATERIAL_STATUS_PROJECT_DISABLE 已被项目暂停、MATERIAL_STATUS_DISABLE 已暂停、MATERIAL_STATUS_PROMOTION_DISABLE 已被广告暂停、MATERIAL_STATUS_MATERIAL_DELETE 已删除
	MaterialStatus enum.MaterialStatus `json:"material_status,omitempty"`
}

// TitleMaterial 标题素材
type TitleMaterial struct {
	// Title 创意标题
	Title string `json:"title,omitempty"`
	// WordList 动态词包ID
	WordList []uint64 `json:"word_list,omitempty"`
}

// ComponentMaterial 创意组件信息
type ComponentMaterial struct {
	// ComponentID 组件id
	ComponentID uint64 `json:"component_id,omitempty"`
}

// ProductInfo 产品信息
type ProductInfo struct {
	// Titles 产品名称
	Titles []string `json:"titles,omitempty"`
	// ImageIDs 产品主图
	ImageIDs []string `json:"image_ids,omitempty"`
	// SellingPoints 产品卖点
	SellingPoints []string `json:"selling_points,omitempty"`
}

// MaterialScoreInfo 素材评级信息
type MaterialScoreInfo struct {
	// ScoreNumOfMaterial 素材数量评级分数，枚举值：LEVEL1 较差、LEVEL2 一般、LEVEL3 良好、LEVEL4 极佳、LEVEL5 完美
	ScoreNumOfMaterial string `json:"score_num_of_material,omitempty"`
	// ScoreTypeOfMaterial 素材类型评级分数，枚举值：LEVEL1 较差、LEVEL2 一般、LEVEL3 良好、LEVEL4 极佳、LEVEL5 完美
	ScoreTypeOfMaterial string `json:"score_type_of_material,omitempty"`
	// ScoreValueOfMaterial 素材品质评级分数，枚举值：LEVEL1 较差、LEVEL2 一般、LEVEL3 良好、LEVEL4 极佳、LEVEL5 完美
	ScoreValueOfMaterial string `json:"score_value_of_material,omitempty"`
	// MaterialAdvice 素材评级建议
	MaterialAdvice []string `json:"material_advice,omitempty"`
	// LowQualityMaterialList 低质素材信息
	LowQualityMaterialList []LowQualityMaterial `json:"low_quality_material_list,omitempty"`
}

// LowQualityMaterial 低质素材信息
type LowQualityMaterial struct {
	// LowQualityVideoIDs 低质视频素材
	LowQualityVideoIDs []string `json:"low_quality_video_ids,omitempty"`
	// LowQualityImageIDs 低质图片素材
	LowQualityImageIDs []string `json:"low_quality_image_ids,omitempty"`
}
