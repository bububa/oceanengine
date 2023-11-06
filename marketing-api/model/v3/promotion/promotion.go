package promotion

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/file"
	"github.com/bububa/oceanengine/marketing-api/model/v3/project"
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
	// StatusFirst 广告一级状态
	StatusFirst enum.PromotionStatusFirst `json:"status_first,omitempty"`
	// StatusSecond 广告二级状态
	StatusSecond []enum.PromotionStatus `json:"status_second,omitempty"`
	// Status 广告状态，枚举值：NOT_DELETED 不限 、ALL 不限（包含已删除）、OK 投放中、DELETED 已删除、PROJECT_OFFLINE_BUDGET 项目超出预算、PROJECT_PREOFFLINE_BUDGET 项目接近预算、TIME_NO_REACH 未到达投放时间、TIME_DONE 已完成、NO_SCHEDULE 不在投放时段、AUDIT 新建审核中、REAUDIT 修改审核中、FROZEN 已终止、AUDIT_DENY 审核不通过、OFFLINE_BUDGET 广告超出预算、OFFLINE_BALANCE 账户余额不足、PREOFFLINE_BUDGET 广告接近预算、DISABLED 已暂停、PROJECT_DISABLED 已被项目暂停、LIVE_ROOM_OFF 关联直播间不可投、PRODUCT_OFFLINE 关联商品不可投，、AWEME_ACCOUNT_DISABLED 关联抖音账号不可投、AWEME_ANCHOR_DISABLED 锚点不可投、DISABLE_BY_QUOTA 已暂停（配额达限）、CREATE 新建、ADVERTISER_OFFLINE_BUDGET 账号超出预算、ADVERTISER_PREOFFLINE_BUDGET 账号接近预算
	Status enum.PromotionStatus `json:"status,omitempty"`
	// OptStatus 操作状态
	OptStatus enum.OptStatus `json:"opt_status,omitempty"`
	// NativeSetting 原生广告设置
	NativeSetting *NativeSetting `json:"native_setting,omitempty"`
	// PromotionMaterials 广告素材组合
	PromotionMaterials *PromotionMaterial `json:"promotion_materials,omitempty"`
	// Keywords 关键词列表，关键词和智能拓流二者必须开启一个，一个广告最多可添加1000个
	Keywords []project.Keyword `json:"keywords,omitempty"`
	// IsCommentDisable 广告评论，ON为启用，OFF为不启用
	IsCommentDisable model.ReverseOnOffInt `json:"is_comment_disable,omitempty"`
	// AdDownloadStatus 客户端下载视频功能，ON为启用，OFF为不启用
	AdDownloadStatus model.OnOffInt `json:"ad_download_status,omitempty"`
	// MaterialsType 素材类型
	MaterialsType enum.MaterialsType `json:"materials_type,omitempty"`
	// Source 广告来源
	Source string `json:"source,omitempty"`
	// BudgetMode 预算类型
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// Budget 预算
	Budget float64 `json:"budget,omitempty"`
	// Bid 点击出价/展示出价
	Bid float64 `json:"bid,omitempty"`
	// CpaBid 目标转化出价/预期成本
	CpaBid float64 `json:"cpa_bid,omitempty"`
	// DeepCpaBid 深度优化出价
	DeepCpaBid float64 `json:"deep_cpabid,omitempty"`
	// RoiGoal 深度转化ROI系数
	RoiGoal float64 `json:"roi_goal,omitempty"`
	// ScheduleTime 广告的投放时段
	ScheduleTime string `json:"schedule_time,omitempty"`
	// MaterialScoreInfo 素材评级信息
	MaterialScoreInfo *MaterialScoreInfo `json:"material_score_info,omitempty"`
	// CreativeAutoGenerateSwitch 是否开启自动生成素材
	// 默认值：OFF
	// 枚举值：ON开启、OFF不开启
	CreativeAutoGenerateSwitch string `json:"creative_auto_generate_switch,omitempty"`
	// ConfigID 配置ID，开关打开，不传为黑盒明投派生
	ConfigID uint64 `json:"config_id,omitempty"`
	// BrandInfo 品牌信息
	BrandInfo *BrandInfo `json:"brand_info,omitempty"`
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

func (p Promotion) GetExternalURLs() []string {
	if p.PromotionMaterials == nil {
		return nil
	}
	return p.PromotionMaterials.ExternalURLMaterialList
}

func (p Promotion) IsProject() bool {
	return false
}

// PromotionMaterial 广告素材
type PromotionMaterial struct {
	// VideoMaterialList 视频素材信息
	VideoMaterialList []VideoMaterial `json:"video_material_list,omitempty"`
	// ImageMaterialList 创意图片素材
	ImageMaterialList []ImageMaterial `json:"image_material_list,omitempty"`
	// TitleMaterialList 标题素材
	TitleMaterialList []TitleMaterial `json:"title_material_list,omitempty"`
	// TextAbstractList 文本摘要信息，单广告可添加1-10个，长度25-45个字
	TextAbstractList []TextAbstract `json:"text_abstract_list,omitempty"`
	// AnchorMaterialList 原生锚点素材，当 anchor_related_type = SELECT时必填，数量上限为1
	AnchorMaterialList []AnchorMaterial `json:"anchor_material_list,omitempty"`
	// DecorationMaterial 家装卡券素材
	// 仅当landing_type选择LINK且命中白名单可用
	// 需在广告平台签署协议后可用
	DecorationMaterial []DecorationMaterial `json:"decoration_material,omitempty"`
	// AutoExtendTraffic 智能拓流
	// 允许值：ON开启（默认值）； OFF关闭
	AutoExtendTraffic string `json:"auto_extend_traffic,omitempty"`
	// Keywords 关键词列表，关键词和智能拓流二者必须开启一个，一个广告最多可添加1000个
	Keywords []project.Keyword `json:"keywords,omitempty"`
	// ComponentMaterialList 创意组件信息
	ComponentMaterialList []ComponentMaterial `json:"component_material_list,omitempty"`
	// MiniProgramInfo 字节小程序信息，当landing_type = MICRO_GAME且micro_promotion_type = BYTE_APP或BYTE_GAME时有效且必填
	MiniProgramInfo *MiniProgramInfo `json:"mini_program_info,omitempty"`
	// ExternalURLMaterialList 普通落地页链接素材，上限10个
	// 当landing_type = APP ，且 download_type = EXTERNAL_URL时，external_url_material_list 至少传入一个
	// 当landing_type = LINK和SHOP类型，且项目asset_type = ORANGE 时，仅允许传入支持对应优化目标的橙子落地页
	// 当landing_type = LINK和SHOP类型，且项目asset_type = THIRDPARTY 时，仅允许自研落地页
	// 当landing_type = MICRO_GAME和WECHAT_APP类型，仅支持选择含微信小程序/小游戏的橙子建站落地页且如果多选落地页，需要所有建站对应相同的微信小游戏/微信小程序，否则报错
	// 应用直播链路和线索自研落地页直播链路不支持该字段
	ExternalURLMaterialList []string `json:"external_url_material_list,omitempty"`
	// ParamsType 链接类型(落地页)，DPA推广目的下必填允许值: DPA 商品库所含链接、CUSTOM 自定义链接
	ParamsType string `json:"params_type,omitempty"`
	// ExternalURLField 落地页链接字段选择，当params_type为DPA时必填
	ExternalURLField string `json:"external_url_field,omitempty"`
	// ExternalURLParams 落地页检测参数
	ExternalURLParams string `json:"external_url_params,omitempty"`
	// OpenURLType 直达链接类型，允许值: NONE不启用, DPA商品库所含链接, CUSTOM自定义链接商品库链接对应商品库内调起字段，如对接多种调起链接则可选择；自定义链接可自行输入调起链接。
	// DPA推广母的下可以使用
	OpenURLType string `json:"open_url_type,omitempty"`
	// OpenURLField 直达链接字段选择，当dpa_open_url_type为DPA必填
	OpenURLField string `json:"open_url_field,omitempty"`
	// OpenURLParams 直达链接检测参数(DPA推广目的特有,在“产品库中提取的scheme地址"后面追加填写的参数)
	OpenURLParams string `json:"open_url_params,omitempty"`
	// WebURLMaterialList Android应用下载详情页
	WebURLMaterialList []string `json:"web_url_material_list,omitempty"`
	// PlayableURLMaterialList 试玩落地页素材
	PlayableURLMaterialList []string `json:"playable_url_material_list,omitempty"`
	// CarouselMaterialList   图集素材信息，当ad_type=ALL时，支持上限10个；当ad_type=SEARCH时，支持上限30个
	CarouselMaterialList []CarouselMaterial `json:"carousel_material_list,omitempty"`
	// OpenURL 直达链接，用于打开电商app，调起店铺
	// 仅在电商店铺推广目的下有效
	OpenURL string `json:"open_url,omitempty"`
	// OpenURLs 电商优选直达链接组
	OpenURLs []string `json:"open_urls,omitempty"`
	// Ulink 直达备用链接，支持穿山甲和站内广告位
	Ulink string `json:"ulink,omitempty"`
	// DynamicCreateiveSwitch 动态创意开关，允许值：ON开启（默认值），OFF关闭，当ad_type=SEARCH时有效
	DynamicCreateiveSwitch string `json:"dynamic_creative_switch,omitempty"`
	// ProductInfo 产品信息
	ProductInfo *ProductInfo `json:"product_info,omitempty"`
	// CallToActionButtons 行动号召文案
	CallToActionButtons []string `json:"call_to_action_buttons,omitempty"`
	// IntelligentGeneration 智能生成行动号召按钮，开启后即对应的文案自动生成，可选项为OFF（默认）、ON
	IntelligentGeneration string `json:"intelligent_generation,omitempty"`
}

// VideoMaterial 视频素材信息
type VideoMaterial struct {
	// ImageMode 素材类型
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
	// VideoID 视频ID
	VideoID string `json:"video_id,omitempty"`
	// VideoCoverID 视频封面图片ID
	VideoCoverID string `json:"video_cover_id,omitempty"`
	// ItemID 抖音短视频ID
	ItemID model.Uint64 `json:"item_id,omitempty"`
	// MaterialID 素材ID
	MaterialID model.Uint64 `json:"material_id,omitempty"`
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
	MaterialID model.Uint64 `json:"material_id,omitempty"`
	// MaterialStatus 素材审核状态，枚举值：
	// MATERIAL_STATUS_OK 投放中、MATERIAL_STATUS_DELETE 已删除、MATERIAL_STATUS_PROJECT_OFFLINE_BUDGET 项目超出预算、MATERIAL_STATUS_PROJECT_PREOFFLINE_BUDGET 项目接近预算、MATERIAL_STATUS_TIME_NO_REACH 未到达投放时间、MATERIAL_STATUS_TIME_DONE 已完成、MATERIAL_STATUS_NO_SCHEDULE 不在投放时段、MATERIAL_STATUS_AUDIT 新建审核中、MATERIAL_STATUS_REAUDIT 修改审核中、MATERIAL_STATUS_OFFLINE_AUDIT 审核不通过、MATERIAL_STATUS_OFFLINE_BUDGET 广告超出预算、MATERIAL_STATUS_ OFFLINE_BALANCE 账户余额不足、MATERIAL_STATUS_ PREOFFLINE_BUGDET 广告接近预算、MATERIAL_STATUS_PROJECT_DISABLE 已被项目暂停、MATERIAL_STATUS_DISABLE 已暂停、MATERIAL_STATUS_PROMOTION_DISABLE 已被广告暂停、MATERIAL_STATUS_MATERIAL_DELETE 已删除
	MaterialStatus enum.MaterialStatus `json:"material_status,omitempty"`
}

// TitleMaterial 标题素材
type TitleMaterial struct {
	// MaterialID 素材ID
	MaterialID model.Uint64 `json:"material_id,omitempty"`
	// Title 创意标题
	Title string `json:"title,omitempty"`
	// BidwordList 搜索关键词列表
	BidwordList []Bidword `json:"bidword_list,omitempty"`
	// WordList 动态词包ID
	WordList []uint64 `json:"word_list,omitempty"`
}

// TextAbstract 文本摘要信息，单广告可添加1-10个，长度25-45个字
type TextAbstract struct {
	// AbstractText 文本摘要内容,单广告可添加1-10个，长度25-45个字, 两个英文字符占1位。
	// 如果要使用动态词包，格式如下：“XXX{词包名}XXX{词包名}XXX”，请注意当您使用动态词包需在下方 word_list 字段中按顺序传入词包ID，并且在一个文本摘要内容中最多使用两个动态词包。如果要使用搜索关键词，格式如下：“XXX{#关键词#}XXX”，请注意当您使用关键词需在下方 bidword_list 字段中传入关键词，并且在一个文本摘要内容中最多使用一个关键词
	AbstractText string `json:"abstract_text,omitempty"`
	// BidwordList 搜索关键词列表
	BidwordList []Bidword `json:"bidword_list,omitempty"`
	// WordList 动态词包ID，可使用 【查询动态词包接口】 获得，结合标题中的词包格式您需要填写相同个数与顺序的词包ID，如果实际ID顺序与标题中词包名顺序不一致我们将以词包ID顺序为准。
	WordList []uint64 `json:"word_list,omitempty"`
}

// Bidword 搜索关键词列表
type Bidword struct {
	// DefaultWord 关键词
	DefaultWord string `json:"default_word,omitempty"`
}

// ComponentMaterial 创意组件信息
type ComponentMaterial struct {
	// ComponentID 组件id
	ComponentID uint64 `json:"component_id,omitempty"`
	// ExternalURLMaterialList 普通落地页链接素材，上限10个
	// 当landing_type = APP ，且 download_type = EXTERNAL_URL时，external_url_material_list 至少传入一个;
	// 当landing_type = LINK和SHOP类型，且项目asset_type = ORANGE 时，仅允许传入支持对应优化目标的橙子落地页;
	// 当landing_type = LINK和SHOP类型，且项目asset_type = THIRDPARTY 时，仅允许自研落地页;
	// landing_type = MICRO_GAME时，当micropromotion_type = WECHAT_APP或WECHAT_GAME时有效，且如果多选落地页，需要所有建站对应相同的微信小游戏/微信小程序，否则报错；
	// 当micropromotion_type = BYTE_APP或BYTE_GAME时有效，可作为备用链接传入该字段，数量上限为1;
	// 获取对应优化目标的橙子建站落地页可参考【通过优化目标获取橙子落地页站点信息】获取第三方落地页可参考【第三方落地页管理】
	ExternalURLMaterialList []string `json:"external_url_material_list,omitempty"`
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
	LowQualityMaterialList *LowQualityMaterial `json:"low_quality_material_list,omitempty"`
}

// LowQualityMaterial 低质素材信息
type LowQualityMaterial struct {
	// LowQualityVideoIDs 低质视频素材
	LowQualityVideoIDs []string `json:"low_quality_video_ids,omitempty"`
	// LowQualityImageIDs 低质图片素材
	LowQualityImageIDs []string `json:"low_quality_image_ids,omitempty"`
}

// NativeSetting 原生广告设置
type NativeSetting struct {
	// AwemeID 授权抖音号id，可通过【获取抖音号授权关系】接口获得
	// 当传入抖音号表示原生广告开启
	AwemeID string `json:"aweme_id,omitempty"`
	// IsFeedAndFavSee 主页作品列表隐藏广告内容
	// 允选值：OFF（不隐藏），ON（隐藏）
	// 默认值：OFF
	IsFeedAndFavSee string `json:"is_feed_and_fav_see,omitempty"`
	// AnchorRelatedType 原生锚点启用开关，允许值:
	// 不启用 OFF，自动生成 AUTO，手动选择 SELECT
	// 默认值为 OFF
	AnchorRelatedType string `json:"anchor_related_type,omitempty"`
}

// MiniProgramInfo 字节小程序信息，当landing_type = MICRO_GAME且micro_promotion_type = BYTE_APP或BYTE_GAME时有效且必填
type MiniProgramInfo struct {
	// AppID 小程序/小游戏id，当使用 mini_program_info且url不传入时必填
	AppID string `json:"app_id,omitempty"`
	// StartPath 启动路径，小程序类型必传，小游戏类型不传值
	StartPath string `json:"start_path,omitempty"`
	// Params 页面监测参数
	Params string `json:"params,omitempty"`
	// URL 字节小程序调起链接，传输会检查url正确性，不传输由平台自动生成;
	// url传入时，app_id、start_path、params无须传入
	URL string `json:"url,omitempty"`
}

// AnchorMaterial 原生锚点素材
type AnchorMaterial struct {
	// AnchorType 锚点类型，允许值：
	// - 应用下载-游戏：APP_GAME- 应用下载-网服：APP_INTERNET_SERVICE- 应用下载-电商：APP_SHOP- 高级在线预约：ONLINE_SUBSCRIBE当 anchor_related_type = SELECT时必填
	AnchorType enum.AnchorType `json:"anchor_type,omitempty"`
	// AnchorID 原生锚点id可使用【获取账户下的原生锚点】获得
	// 当 anchor_related_type = SELECT时必填
	AnchorID string `json:"anchor_id,omitempty"`
}

// DecorationMaterial 家装卡券素材
type DecorationMaterial struct {
	// MaterialID 素材ID
	MaterialID model.Uint64 `json:"material_id,omitempty"`
	// ImageMode 素材类型，仅支持传入CREATIVE_IMAGE_MODE_DECORATION_COUPON
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
	// ActivityID 活动ID，image_mode为家具卡券素材时填写
	// 通过【获取家装联盟卡券列表】接口获取
	ActivityID string `json:"activity_id,omitempty"`
}

// BrandInfo 品牌信息
type BrandInfo struct {
	// YuntuCategoryID 品牌分类id
	YuntuCategoryID int64 `json:"yuntu_category_id,omitempty"`
	// CdpCategoryID cdp品牌id
	CdpCategoryID int64 `json:"cdp_category_id,omitempty"`
	// EcomCategoryID 电商品牌id
	EcomCategoryID int64 `json:"ecom_category_id,omitempty"`
	// BrandNameID 云图品牌id
	BrandNameID int64 `json:"brand_name_id,omitempty"`
	// CdpBrandID cdp品牌id
	CdpBrandID int64 `json:"cdp_brand_id,omitempty"`
	// CdpBrandName 云图品牌名称
	CdpBrandName string `json:"cdp_brand_name,omitempty"`
	// SubBrandNames 子品牌名称
	SubBrandNames []string `json:"sub_brand_names,omitempty"`
	// SubBrandNameIDs 子品牌id
	SubBrandNameIDs []string `json:"sub_brand_name_ids,omitempty"`
}

// CarouselMaterial   图集素材信息
type CarouselMaterial struct {
	// CarouselID 图集id，可通过【获取图集素材】接口获得
	CarouselID string `json:"carousel_id,omitempty"`
	// ImageID 图片ID列表
	ImageID []string `json:"image_id,omitempty"`
	// AudioID 音频ID
	AudioID string `json:"audio_id,omitempty"`
	// MaterialStatus 素材审核状态
	MaterialStatus string `json:"material_status,omitempty"`
	// CarouselType 图集素材类型
	CarouselType int `json:"carousel_type,omitempty"`
	// ImageSubject 图片主题
	ImageSubject []file.ImageSubject `json:"image_subject,omitempty"`
}
