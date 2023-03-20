package nativeanchor

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
)

// NativeAnchor 广告账户下原生锚点
type NativeAnchor struct {
	// AnchorID 锚点id
	AnchorID string `json:"anchor_id,omitempty"`
	// AnchorName 锚点名称
	AnchorName string `json:"anchor_name,omitempty"`
	// ToolTitle 锚点工具名称（内部管理展示）
	ToolTitle string `json:"tool_title,omitempty"`
	// AnchorType 锚点类型
	AnchorType enum.AnchorType `json:"anchor_type,omitempty"`
	// Status 锚点审核状态
	Status string `json:"status,omitempty"`
	// Source 锚点来源
	Source string `json:"source,omitempty"`
	// CreateTime 锚点创建日期，格式：yyyy-MM-dd
	CreateTime string `json:"create_time,omitempty"`
	// AndroidPackageName 安卓应用包名
	AndroidPackageName string `json:"android_package_name,omitempty"`
	// IosPackageName ios应用包名
	IosPackageName string `json:"ios_package_name,omitempty"`
	// AppEcommerceAnchor 电商下载锚点
	AppEcommerceAnchor *AppEcommerceAnchor `json:"app_ecommerce_anchor,omitempty"`
	// NetServiceAnchor 网服下载锚点
	NetSerivceAnchor *NetServiceAnchor `json:"net_service_anchor,omitempty"`
	// GameAnchor 游戏锚点
	GameAnchor *GameAnchor `json:"game_anchor,omitempty"`
}

// AppEcommerceAnchor 电商下载锚点
type AppEcommerceAnchor struct {
	// IosAnchorTitle ios 锚点入口标题字段，长度：1～12
	IosAnchorTitle string `json:"ios_anchor_title,omitempty"`
	// AndoridAnchorTitel 安卓锚点入口标题字段，长度：1～12
	AndroidAnchorTitle string `json:"android_anchor_title,omitempty"`
	// ProductInfo 商品大图，推荐比例 1：1
	ProductImage *Image `json:"product_info,omitempty"`
	// ProductName 商品名称，长度1～10
	ProductName string `json:"product_name,omitempty"`
	// ProductTags 商品标签列表，每个标签长度：1～4，标签个数：1～3
	ProductTags []string `json:"product_tags,omitempty"`
	// ProductPrice 商品价格（整数，且单位：元）
	ProductPrice float64 `json:"product_price,omitempty"`
	// ActivityInfo 活动信息文案，长度：1～20
	ActivityInfo string `json:"activity_info,omitempty"`
	// ServiceInfo 服务信息文案，长度：1～20
	ServiceInfo string `json:"service_info,omitempty"`
	// DetailInfo 详情信息文案，长度：1～20
	DetailInfo string `json:"detail_info,omitempty"`
	// AppName 应用名称，长度：1～6
	AppName string `json:"app_name,omitempty"`
	// AppIcon 应用图标，推荐比例1：1
	AppIcon *Image `json:"app_icon,omitempty"`
	// DownloadGuideText APP下载引导文案，长度：1～12
	DownloadGuideText string `json:"download_guide_text,omitempty"`
	// ExternalType 跳转链接类型枚举，1：橙子建站，橙子建站落地页设置字段必填；2：第三方落地页，第三方落地页设置字段必填
	ExternalType int `json:"external_type,omitempty"`
	// OrangeSiteInfo 橙子建站落地页设置
	OrangeSiteInfo *SiteInfo `json:"orange_site_info,omitempty"`
	// ThirdSiteInfo 第三方落地页设置
	ThirdSiteInfo *SiteInfo `json:"third_site_info,omitempty"`
	// AppOpenURL 应用吊起链接
	AppOpenURL string `json:"app_open_url,omitempty"`
	// AndroidPkgName 安卓包名
	AndroidPkgName string `json:"android_pkg_name,omitempty"`
	// OfficialActiBannerImage 官方互动banner图，推荐尺寸1032*360px
	OfficialActiBannerImage *Image `json:"official_acti_banner_image,omitempty"`
	// OfficialActiText 官方活动描述详情，长度：1～15
	OfficialActiText string `json:"official_acti_text,omitempty"`
}

// NetServiceAnchor 网服下载锚点
type NetServiceAnchor struct {
	// NetServiceType 推广内容，当前锚点必填，允许值：
	// 'GENERAL' ：常规 （默认值）
	// 'MICRO_APP' ：微信小程序
	NetServiceType string `json:"net_service_type,omitempty"`
	// PlatformType 配置平台，net_service_type为微信小程序场景下不用传入（1:不限,2:安卓,3:iOS）不限：安卓下载链接和iOS下载链接必填；安卓：安卓下载链接必填，iOS下载链接不填写；iOS：iOS下载链接必填
	PlatformType int `json:"platform_type,omitempty"`
	// AndroidDownloadURL 安卓下载链接，net_service_type为微信小程序场景下不用传入
	AndroidDownloadURL string `json:"android_download_url,omitempty"`
	// AndroidAnchorTitle 安卓锚点入口标题字段
	AndroidAnchorTitle string `json:"android_anchor_title,omitempty"`
	// IosDownloadURL iOS下载链接，net_service_type为微信小程序场景下不用传入
	IosDownloadURL string `json:"ios_download_url,omitempty"`
	// IosAnchorTitle iOS 锚点入口标题字段
	IosAnchorTitle string `json:"ios_anchor_title,omitempty"`
	// AppOpenURL app调起链接
	AppOpenURL string `json:"app_open_url,omitempty"`
	// HeadImageList 锚点头部图片list，推荐尺寸为2：1的横图
	HeadImageList []Image `json:"head_image_list,omitempty"`
	// InstanceID 微信小程序ID，当前锚点类型且net_service_type为'MICRO_APP' 必填
	InstanceID string `json:"instance_id,omitempty"`
	// PathParam 微信小游戏/小程序路径参数
	PathParam string `json:"path_param,omitempty"`
	// IconImages 小程序图片list，当前锚点类型且net_service_type为'MICRO_APP' 必填，比例为1：1，最多一个
	IconImages []Image `json:"icon_images,omitempty"`
	// AppTags APP标签列表，每个标签长度：1～4；标签个数：1～3
	AppTags []string `json:"app_tags,omitempty"`
	// GuideText 引导文案，长度：1～15
	GuideText string `json:"guide_text,omitempty"`
	// AnchorImageMode APP图片比例，100：尺寸为2：1的横图，200：尺寸为3：5的竖图
	AnchorImageMode int `json:"anchor_image_mode,omitempty"`
	// AppImages APP图片，图片个数 3～8
	AppImages []Image `json:"app_images,omitempty"`
	// AppDescription APP详情，1～100
	AppDescription string `json:"app_description,omitempty"`
}

// GameAnchor 游戏锚点
type GameAnchor struct {
	// GameType 推广内容，当前锚点必填，允许值：
	// 'GENERAL' ：常规 （默认值）
	// 'MICRO_GAME' ：微信小游戏
	GameType string `json:"game_type,omitempty"`
	// PlatformType 配置平台，game_type为微信小游戏场景下不用传入（1:不限,2:安卓,3:iOS）不限：安卓下载链接和iOS下载链接必填；安卓：安卓下载链接必填，iOS下载链接不填写；iOS：iOS下载链接必填
	PlatformType int `json:"platform_type,omitempty"`
	// AndroidDownloadURL 安卓下载链接，game_type为微信小游戏场景下不用传入
	AndroidDownloadURL string `json:"android_download_url,omitempty"`
	// AndroidAnchorTitle 安卓锚点入口标题字段，长度1～12
	AndroidAnchorTitle string `json:"android_anchor_title,omitempty"`
	// IosDownloadURL iOS下载链接，game_type为微信小游戏场景下不用传入
	IosDownloadURL string `json:"ios_download_url,omitempty"`
	// IosAnchorTitle iOS锚点入口标题字段，长度1～12
	IosAnchorTitle string `json:"ios_anchor_title,omitempty"`
	// HeadImageList 锚点头部图片list，推荐尺寸为2：1的横图
	HeadImageList []Image `json:"head_image_list,omitempty"`
	// InstanceID 微信小程序ID，当前锚点类型且net_service_type为'MICRO_APP' 必填
	InstanceID string `json:"instance_id,omitempty"`
	// PathParam 微信小游戏/小程序路径参数
	PathParam string `json:"path_param,omitempty"`
	// IconImages 小程序图片list，当前锚点类型且net_service_type为'MICRO_APP' 必填，比例为1：1，最多一个
	IconImages []Image `json:"icon_images,omitempty"`
	// AppTags APP标签列表，每个标签长度：1～4；标签个数：1～3
	AppTags []string `json:"app_tags,omitempty"`
	// GuideText 引导文案，长度：1～15
	GuideText string `json:"guide_text,omitempty"`
	// AnchorImageMode APP图片比例，100：尺寸为2：1的横图，200：尺寸为3：5的竖图
	AnchorImageMode int `json:"anchor_image_mode,omitempty"`
	// AppImages APP图片，图片个数 3～8
	AppImages []Image `json:"app_images,omitempty"`
	// GameDescription 游戏简介，长度 1～45
	GameDescription string `json:"game_description,omitempty"`
	// GameCharatoristic 游戏特色，长度 1～45
	GameCharatoristic string `json:"game_charatoristic,omitempty"`
	// OtherDescription 其他说明，长度 1～200
	OtherDescription string `json:"other_description,omitempty"`
}

// Image 商品大图，推荐比例 1：1
type Image struct {
	// Uri 图片image_id
	Uri string `json:"uri,omitempty"`
	// Width 图片宽度
	Width float64 `json:"width,omitempty"`
	// Height 图片高度
	Height float64 `json:"height,omitempty"`
}

// SiteInfo 站点设置
type SiteInfo struct {
	// IosExternalURL iOS端落地页链接，内部需要包含应用下载链接
	IosExternalURL string `json:"ios_external_url,omitempty"`
	// AndroidExternalURL 安卓端落地页链接，内部需要包含应用下载链接
	AndroidExternalURL string `json:"android_external_url,omitempty"`
	// AndroidDownloadURL 安卓下载链接
	AndroidDownloadURL string `json:"android_download_url,omitempty"`
	// IosDonwloadURL ios下载链接
	IosDownloadURL string `json:"ios_download_url,omitempty"`
	// ExternalURL 第三方落地页
	ExternalURL string `json:"external_url,omitempty"`
}
