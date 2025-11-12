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
	// AdvertiserID 广告主账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AnchorType 锚点类型
	AnchorType enum.AnchorType `json:"anchor_type,omitempty"`
	// Status 锚点审核状态
	Status string `json:"status,omitempty"`
	// Source 锚点来源
	Source string `json:"source,omitempty"`
	// CreateTime 锚点创建日期，格式：yyyy-MM-dd
	CreateTime string `json:"create_time,omitempty"`
	// ModifyTime 锚点更新时间
	ModifyTime string `json:"modify_time,omitempty"`
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
	// PrivateChatAnchor 咨询锚点，当anchor_type=PRIVATE_CHAT时返回的锚点
	PrivateChatAnchor *PrivateChatAnchor `json:"private_chat_anchor,omitempty"`
	// ShoppingCartAnchor 购物车锚点
	ShoppingCartAnchor *ShoppingCartAnchor `json:"shopping_cart_anchor,omitempty"`
	// InsuranceEnterpriseAnchor 外跳锚点，当anchor_type=INSURANCE时返回的详情
	InsuranceEnterpriseAnchor *InsuranceEnterpriseAnchor `json:"insurance_enterprise_anchor,omitempty"`
	// AnchorShareType 锚点分享类型
	AnchorShareType enum.AnchorShareType `json:"anchor_share_type,omitempty"`
}

// AppEcommerceAnchor 电商下载锚点
type AppEcommerceAnchor struct {
	// IosAnchorTitle ios 锚点入口标题字段，长度：1～12
	IosAnchorTitle string `json:"ios_anchor_title,omitempty"`
	// AndoridAnchorTitel 安卓锚点入口标题字段，长度：1～12
	AndroidAnchorTitle string `json:"android_anchor_title,omitempty"`
	// HarmonyAnchorTitle 鸿蒙下载链接
	HarmonyAnchorTitle string `json:"harmony_anchor_title,omitempty"`
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
	// 推广内容，网服锚点必填，允许值：
	// GENERAL：常规应用下载 （默认值）
	// 如选择“跳转微信”推广内容，则直接传入对应的跳转场景枚举：
	// 跳转场景=添加微信号，传WECHAT_PACKAGE
	// 跳转场景=进入微信小程序，传MICRO_APP
	// 跳转场景=企业微信客服，传WECOM_PACKAGE
	// 跳转场景= 跳转微信链接, 传WECHAT_EXTERNAL_URL
	// QUICK_APP：快应用
	NetServiceType enum.NetServiceType `json:"net_service_type,omitempty"`
	// PlatformType APP下载配置，net_service_type为GENERAL常规应用下载时需要填入
	// 1不限；2安卓；3IOS；4鸿蒙；5鸿蒙+安卓；6鸿蒙+IOS；7安卓+IOS
	// 不限：安卓下载链接和iOS下载链接必填（若账号在HarmonyOS白名单中，则设置不限需要填写安卓、IOS、鸿蒙下载链接）
	// 安卓：安卓下载链接必填
	// iOS：iOS下载链接必填
	// 鸿蒙：鸿蒙下载链接必填
	PlatformType int `json:"platform_type,omitempty"`
	// AndroidDownloadURL 安卓下载链接，net_service_type为微信小程序场景下不用传入
	AndroidDownloadURL string `json:"android_download_url,omitempty"`
	// AndroidAnchorTitle 安卓锚点入口标题字段
	AndroidAnchorTitle string `json:"android_anchor_title,omitempty"`
	// IosDownloadURL iOS下载链接，net_service_type为微信小程序场景下不用传入
	IosDownloadURL string `json:"ios_download_url,omitempty"`
	// IosAnchorTitle iOS 锚点入口标题字段
	IosAnchorTitle string `json:"ios_anchor_title,omitempty"`
	// HarmonyDownloadURL 鸿蒙下载链接
	HarmonyDownloadURL string `json:"harmony_download_url,omitempty"`
	// harmonyAnchorTitle 鸿蒙锚点标题
	HarmonyAnchorTitle string `json:"harmony_anchor_title,omitempty"`
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
	// QuickAppID 快应用ID，网服锚点类型且net_service_type为QUICK_APP必填
	QuickAppID uint64 `json:"quick_app_id,omitempty"`
	// NovelChapter 小说章节，非必填，填写小说章节预览，字数1~9999
	NovelChapter string `json:"novel_chapter,omitempty"`
	// WechatPackageID net_service_type为WECHAT_PACKAGE时，该参数代表微信号码包ID
	// net_service_type为WECOM_PACKAGE时，该参数代表企业加粉方案ID
	WechatPackageID uint64 `json:"wechat_package_id,omitempty"`
	// WechatExternalURL 微信跳转链接，当net_service_type = WECHAT_EXTERNAL_URL时必填
	WechatExternalURL string `json:"wechat_external_url,omitempty"`
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
	// HarmonyDonwloadURL 鸿蒙下载链接
	HarmonyDownloadURL string `json:"harmony_download_url,omitempty"`
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
	// GameBonus 游戏福利，可选择 TRUE 启用、FALSE 不启用
	GameBonus bool `json:"game_bonus,omitempty"`
	// GamePackageList 游戏礼包列表，当game_bonus为TRUE 时必填，数量限制0-3
	// （排在首位的礼包，将会在搜索结果页面展示快捷入口）
	GamePackageList []GamePackage `json:"game_package_list,omitempty"`
	// AppOpenURL App调起链接
	AppOpenURL string `json:"app_open_url,omitempty"`
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
	// HarmonyExternalURL 鸿蒙下载链接
	HarmonyExternalURL string `json:"harmony_external_url,omitempty"`
	// ExternalURL 第三方落地页
	ExternalURL string `json:"external_url,omitempty"`
}

// ShoppingCartAnchor 购物车锚点，当anchor_type=SHOPPING_CART：购物车锚点时必填
type ShoppingCartAnchor struct {
	// Title 购物车小标题，不超过10个字
	Title string `json:"title,omitempty"`
	// ProductImages 商品图片，比例1:1，至少1张
	ProductImages []Image `json:"product_images,omitempty"`
	// ProductPrice 商品价格，最多两位小数
	ProductPrice float64 `json:"product_price,omitempty"`
	// ProductTitle 商品标题，不超过35个字
	ProductTitle string `json:"product_title,omitempty"`
	// ProductSource 商品来源，只支持填写淘宝/天猫/京东/拼多多/唯品会/得物
	ProductSource string `json:"product_source,omitempty"`
	// ExternalURL 落地页链接
	ExternalURL string `json:"external_url,omitempty"`
	// OpenURL 直达链接
	OpenURL string `json:"open_url,omitempty"`
	// LinkType 跳转类型，枚举值一跳：ONE_JUMP、二跳：TWO_JUMP
	LinkType string `json:"link_type,omitempty"`
}

// GamePackage 游戏礼包
type GamePackage struct {
	// GamePackageName 游戏礼包名称，字符限制0～14
	GamePackageName string `json:"game_package_name,omitempty"`
	// AndroidPackageCode 安卓礼包码，字符限制0～20
	AndroidPackageCode string `json:"android_package_code,omitempty"`
	// IosPackageCode ios礼包码，字符限制0～20
	IosPackageCode string `json:"ios_package_code,omitempty"`
	// Gift 礼包内的礼品配置，数量限制0～8
	Gift []GamePackageGift `json:"gift,omitempty"`
	// GiftStartDate 礼包使用开始期限，格式 yyyy.MM.dd
	GiftStartDate string `json:"gift_start_date,omitempty"`
	// GiftEndDate 礼包使用结束期限，格式 yyyy.MM.dd
	GiftEndDate string `json:"gift_end_date,omitempty"`
	// GameGiftRegulation 礼包使用规则，字符限制0～30
	// 礼包码的具体兑换路径，如人物等级10级-左侧菜单-福利-输入礼包码兑换
	GameGiftRegulation string `json:"game_gift_regulation,omitempty"`
}

// GamePackageGift 礼包内的礼品配置
type GamePackageGift struct {
	// GiftName 礼品名称，字符限制0～8
	GiftName string `json:"gift_name,omitempty"`
	// GiftImageURL 礼品图片
	GiftImageURL string `json:"gift_image_url,omitempty"`
	// GiftAmount 礼品数量，0～6
	GiftAmount int `json:"gift_amount,omitempty"`
	// GiftUnit 礼品单位 可选值:
	// INDIVIDUAL
	// MYRIAD
	GiftUnit string `json:"gift_unit,omitempty"`
}

// PrivateChatAnchor 咨询锚点
type PrivateChatAnchor struct {
	// ChatGuide 咨询引导文案，如私信获取一对一解答，不超过9个字
	ChatGuide string `json:"chat_guide,omitempty"`
	// Button 按钮文案，可选私信商家、立即咨询、咨询顾问、咨询设计师、问问老师 可选值:
	// PRIVATE_MESSAGE私信商家
	// CONSULT_NOW立即咨询
	// CONSULT_ADVISOR咨询顾问
	// CONSULT_DESIGNER咨询设计师
	// ASK_TEACHER问问老师
	Button enum.PrivateChatAnchorButton `json:"button,omitempty"`
}

// InsuranceEnterpriseAnchor 外跳锚点，当anchor_type=INSURANCE时返回的详情
type InsuranceEnterpriseAnchor struct {
	// ProductImage 服务主图
	ProductImage *Image `json:"product_image,omitempty"`
	// ProductTitle 产品名称，1-15字
	ProductTitle string `json:"product_title,omitempty"`
	// DetailURL 点击按钮时的跳转链接，此处填写的跳转链接会应用到“转化按钮”与“详情介绍”，以http开头
	DetailURL string `json:"detail_url,omitempty"`
	// ProductTags 产品特点，最多返回3个
	ProductTags []string `json:"product_tags,omitempty"`
	// ProductServiceDescription 服务描述，1-6字
	ProductServiceDescription string `json:"product_service_description,omitempty"`
	// ConversionBtn 转化按钮，仅支持1个枚举 可选值:
	// 查看详情 VIEW_DETIALS
	// 立即购买 BUY_NOW
	// 立即完善 IMPROVE_NOW
	// 免费领取 FREE_RECEIVE
	ConversionBtn string `json:"conversion_btn,omitempty"`
	// ProductDescriptions 产品描述
	ProductDescriptions []string `json:"product_descriptions,omitempty"`
	// BannerImages 详情介绍banner,大小不超过10M的图片，尺寸1032*360px
	BannerImages []Image `json:"banner_images,omitempty"`
	// BannerDescription 详情介绍，1-18个字
	BannerDescription string `json:"banner_description,omitempty"`
	// ProductName 产品详情下的产品名称
	ProductName string `json:"product_name,omitempty"`
	// SingleProductInfo 单项服务名称，5-10组
	SingleProductInfo []SingleProductInfo `json:"single_product_info,omitempty"`
}

// SingleProductInfo 单项服务名称
type SingleProductInfo struct {
	// SingleProductName 单项服务名称，0/15
	SingleProductName string `json:"single_product_name,omitempty"`
	// SingleProductDetail 服务描述，0/10
	SingleProductDetail string `json:"single_product_detail,omitempty"`
}
