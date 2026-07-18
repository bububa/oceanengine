package stdproject

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
)

// VideoMaterial 视频素材信息
type VideoMaterial struct {
	// VideoID 视频id
	VideoID string `json:"video_id,omitempty"`
	// VideoCoverURI 视频封面图片ID
	VideoCoverURI string `json:"video_cover_uri,omitempty"`
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// MaterialOptStatus 视频素材操作状态，枚举值：DISABLE 暂停、ENABLE 启用
	MaterialOptStatus enum.MaterialOptStatus `json:"material_opt_status,omitempty"`
	// MaterialStatusFirst 素材一级状态，可选值：
	// STATUS_DELETE 已删除、STATUS_DISABLE 未投放、STATUS_FROZEN 已终止、STATUS_OK 投放中、STATUS_TIME_DONE 已完成
	// 注意：官方应答字段名为 material_status_frist（拼写如此），此处 json tag 保留原样
	MaterialStatusFirst enum.MaterialStatusFirst `json:"material_status_frist,omitempty"`
	// MaterialStatusSecond 素材二级状态
	MaterialStatusSecond []string `json:"material_status_second,omitempty"`
	// ImageMode 素材类型，可选值：
	// CREATIVE_IMAGE_MODE_CALIBRATION_FREE_VIDEO 穿山甲免校验
	// CREATIVE_IMAGE_MODE_UNION_SPLASH_VIDEO 互动创意
	// CREATIVE_IMAGE_MODE_VIDEO 视频
	// CREATIVE_IMAGE_MODE_VIDEO_VERTICAL 抖音&火山竖版视频素材
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
}

// ImageMaterial 图片素材信息
type ImageMaterial struct {
	// ImageMode 素材类型，可选值：
	// CREATIVE_IMAGE_MODE_LARGE、CREATIVE_IMAGE_MODE_LARGE_VERTICAL 大图竖版图、
	// CREATIVE_IMAGE_MODE_SMALL、SEARCH_AD_SMALL_IMAGE 搜索小图、TOUTIAO_SEARCH_AD_IMAGE 搜索大图
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
	// ImageIDs 图片id
	ImageIDs []string `json:"image_ids,omitempty"`
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// MaterialStatusFirst 素材一级状态（官方应答字段名为 material_status_frist）
	MaterialStatusFirst enum.MaterialStatusFirst `json:"material_status_frist,omitempty"`
	// MaterialStatusSecond 素材二级状态
	MaterialStatusSecond []string `json:"material_status_second,omitempty"`
}

// CarouselMaterial 图文素材信息
type CarouselMaterial struct {
	// CarouselID 图文id
	CarouselID uint64 `json:"carousel_id,omitempty"`
	// ItemID 抖音主页图文id
	ItemID uint64 `json:"item_id,omitempty"`
	// VideoHpVisibility 主页可见性，可选值：ALWAYS_VISIBLE 主页始终可见、HIDE_VIDEO_ON_HP 主页隐藏（默认值）
	VideoHpVisibility enum.VideoHpVisibility `json:"video_hp_visibility,omitempty"`
	// MaterialStatusFirst 素材一级状态（官方应答字段名为 material_status_frist）
	MaterialStatusFirst enum.MaterialStatusFirst `json:"material_status_frist,omitempty"`
	// MaterialStatusSecond 素材二级状态
	MaterialStatusSecond []string `json:"material_status_second,omitempty"`
}

// TrialPlayMaterial 试玩素材信息
type TrialPlayMaterial struct {
	// AppPlayURI 试玩素材uri
	AppPlayURI string `json:"app_play_uri,omitempty"`
	// GuideVideoID 引导视频ID
	GuideVideoID string `json:"guide_video_id,omitempty"`
	// ImageMode 素材类型，可选值：
	// CREATIVE_IMAGE_MODE_TRIAL_PLAY_LARGE 试玩横屏、CREATIVE_IMAGE_MODE_TRIAL_PLAY_VERTICAL 试玩竖屏
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
	// AppPlayVersion 试玩素材版本号
	AppPlayVersion string `json:"app_play_version,omitempty"`
	// FileName 素材名称
	FileName string `json:"file_name,omitempty"`
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// MaterialStatusFirst 素材一级状态（官方应答字段名为 material_status_frist）
	MaterialStatusFirst enum.MaterialStatusFirst `json:"material_status_frist,omitempty"`
	// MaterialStatusSecond 素材二级状态
	MaterialStatusSecond []string `json:"material_status_second,omitempty"`
}

// InstantPlayMaterial 直玩素材信息
type InstantPlayMaterial struct {
	// AppPlayURI 直玩素材uri
	AppPlayURI string `json:"app_play_uri,omitempty"`
	// ImageMode 直玩素材类型，可选值：CREATIVE_IMAGE_MODE_INSTANT_PLAY 直玩
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
	// CoverURI 封面图
	CoverURI string `json:"cover_uri,omitempty"`
	// FileName 封面名称
	FileName string `json:"file_name,omitempty"`
	// MaterialID 素材id
	MaterialID uint64 `json:"material_id,omitempty"`
	// MaterialStatusFirst 素材一级状态（官方应答字段名为 material_status_frist）
	MaterialStatusFirst enum.MaterialStatusFirst `json:"material_status_frist,omitempty"`
	// MaterialStatusSecond 素材二级状态
	MaterialStatusSecond []string `json:"material_status_second,omitempty"`
}

// TitleMaterial 标题素材信息
type TitleMaterial struct {
	// Title 创意标题
	Title string `json:"title,omitempty"`
	// MaterialID 创意id
	MaterialID uint64 `json:"material_id,omitempty"`
}

// ComponentMaterial 创意组件信息
type ComponentMaterial struct {
	// ComponentID 组件id
	ComponentID uint64 `json:"component_id,omitempty"`
}

// AnchorMaterial 原生锚点素材
type AnchorMaterial struct {
	// AnchorID 原生锚点id
	AnchorID string `json:"anchor_id,omitempty"`
	// AnchorType 锚点类型，可选值：
	// APP_GAME、APP_INTERNET_SERVICE、APP_SHOP、INSURANCE、ONLINE_SUBSCRIBE、PRIVATE_CHAT、SHOPPING_CART
	AnchorType enum.AnchorType `json:"anchor_type,omitempty"`
	// AnchorRelatedType 锚点关联类型，可选值：AUTO 自动生成、OFF 不启用、SELECT 手动选择
	AnchorRelatedType enum.AnchorRelatedType `json:"anchor_related_type,omitempty"`
	// OpenURL 直达链接
	OpenURL string `json:"open_url,omitempty"`
	// UlinkURLType 直达链接备用链接类型，可选值：UNIVERSAL_LINK、APPLINK
	UlinkURLType string `json:"ulink_url_type,omitempty"`
	// UlinkURL 直达链接备用链接
	UlinkURL string `json:"ulink_url,omitempty"`
}

// PlayletSeriesURLInfo 短剧专辑链接信息
type PlayletSeriesURLInfo struct {
	// PlayletSeriesURLType 短剧专辑类型，可选值：IAA、IAP
	PlayletSeriesURLType enum.PlayletSeriesURLType `json:"playlet_series_url_type,omitempty"`
	// PlayletSeriesURL 短剧专辑链接
	PlayletSeriesURL string `json:"playlet_series_url,omitempty"`
}

// ProductInfo 产品信息
type ProductInfo struct {
	// Titles 产品名称
	Titles []string `json:"titles,omitempty"`
	// ImageIDs 产品主图
	ImageIDs []string `json:"image_ids,omitempty"`
	// SellingPoints 产品卖点
	SellingPoints []string `json:"selling_points,omitempty"`
	// CallToActionButtons 行动号召
	CallToActionButtons []string `json:"call_to_action_buttons,omitempty"`
}

// MiniProgramInfo 字节小程序&小游戏信息
type MiniProgramInfo struct {
	// URL 字节小程序&小游戏调起链接
	URL string `json:"url,omitempty"`
	// AppID 小程序/小游戏原始id
	AppID string `json:"app_id,omitempty"`
	// StartPath 启动路径
	StartPath string `json:"start_path,omitempty"`
	// Params 页面检测参数
	Params string `json:"params,omitempty"`
}
