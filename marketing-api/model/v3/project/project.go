package project

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// Project 项目
type Project struct {
	// ProjectID 项目ID
	ProjectID uint64 `json:"project_id,omitempty"`
	// AdvertiserID 广告账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// DeliveryMode 投放模式，允许值：MANUAL手动投放、PROCEDURAL自动投放
	DeliveryMode enum.DeliveryMode `json:"delivery_mode,omitempty"`
	// DeliveryType 投放类型，当前仅搜索广告ad_type = SEARCH下可设置，允许值
	// NORMAL 常规投放（默认）
	// DURATION 周期稳投
	DeliveryType enum.DeliveryType `json:"delivery_type,omitempty"`
	// LandingType 推广目的，枚举值：APP 应用推广、LINK 销售线索推广、MICRO_GAME小游戏
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// AppPromotionType 子目标，枚举值：DOWNLOAD 应用下载、LAUNCH 应用调用、RESERVE 预约下载
	AppPromotionType enum.AppPromotionType `json:"app_promotion_type,omitempty"`
	// MarketingGoal 营销场景，枚举值：VIDEO_AND_IMAGE 短视频/图片
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// AdType 广告类型，枚举值：ALL
	AdType enum.CampaignType `json:"ad_type,omitempty"`
	// OptStatus 目标操作，枚举值：ENABLE 启用项目、DISABLE暂停项目
	OptStatus enum.OptStatus `json:"opt_status,omitempty"`
	// Name 项目名称
	Name string `json:"name,omitempty"`
	// ProjectCreateTime 项目创建时间，格式yyyy-MM-dd HH:mm:ss
	ProjectCreateTime string `json:"project_create_time,omitempty"`
	// ProjectModifyTime 项目更新时间，格式yyyy-MM-dd HH:mm:ss
	ProjectModifyTime string `json:"project_modify_time,omitempty"`
	// Status 项目状态
	Status enum.ProjectStatus `json:"status,omitempty"`
	// StatusFirst 项目一级状态过滤，允许值：
	// PROJECT_STATUS_DELETE已删除
	// PROJECT_STATUS_DONE已完成
	// PROJECT_STATUS_DISABLE未投放
	// PROJECT_STATUS_ENABLE启用中
	// 不传默认返回不限（不含已删除）
	StatusFirst enum.ProjectStatus `json:"status_first,omitempty"`
	// StatusSecond 项目二级状态过滤，允许值：PROJECT_STATUS_STOP 已暂停
	// PROJECT_STATUS_BUDGET_EXCEED 项目超出预算
	// PROJECT_STATUS_NOT_START 未达投放时间
	// PROJECT_STATUS_NO_SCHEDULE 不在投放时段
	// 当status_first = PROJECT_STATUS_DISABLE时传入有效
	StatusSecond []enum.ProjectStatus `json:"status_second,omitempty"`
	// Pricing 出价方式
	Pricing enum.PricingType `json:"pricing,omitempty"`
	// PackageName 应用包名
	PackageName string `json:"package_name,omitempty"`
	// AppName 应用名
	AppName string `json:"app_name,omitempty"`
	// FeedDeliverySearch 搜索快投关键词功能，HAS_OPEN:启用，DISABLED:未启用
	FeedDeliverySearch enum.FeedDeliverySearch `json:"feed_delivery_search,omitempty"`
	// SearchBidRatio 出价系数
	SearchBidRatio float64 `json:"search_bid_ratio,omitempty"`
	// AudienceExtend 定向拓展
	AudienceExtend enum.OnOff `json:"audience_extend,omitempty"`
	// Keywords 搜索关键词列表
	Keywords []Keyword `json:"keywords,omitempty"`
	// AutoTraficExtend 智能拓流 ，允许值：ON开启； OFF关闭
	// 仅支持ad_type = SEARCH时设置
	// 自定义关键词和智能拓流二者必须开启一个：若keywords为空，智能拓流 auto_extend_traffic需为ON
	// 对于搜索极速智投项目，若设置blue_flow_keyword_name蓝海关键词，智能拓流默认值为ON，且不得设置为OFF
	AutoExtendTraffic enum.OnOff `json:"auto_extend_traffic,omitempty"`
	// BlueFlowPackage 搜索蓝海流量投放相关参数
	BlueFlowPackage *BlueFlowPackage `json:"blue_flow_package,omitempty"`
	// RelatedProduct 关联产品投放相关
	RelatedProduct *RelatedProduct `json:"related_product,omitempty"`
	// DpaCategories 商品投放范围，分类列表，由【DPA商品广告-获取DPA分类】 得到
	// 个数限制 [0, 1000]
	// 不传和传空数组即为不限商品投放范围
	// DPA推广目的下有效
	DpaCategories []uint64 `json:"dpa_categories,omitempty"`
	// DpaProductTarget 自定义筛选条件（商品投放条件）。用于圈定商品投放范围，结合商品库字段搭配判断条件，圈定商品投放范围。获取商品库元信息-商品广告-商业开放平台
	// 数组长度限制：最大5条
	// DPA推广目的下有效
	DpaProductTarget []DpaProductTarget `json:"dpa_product_target,omitempty"`
	// DeliveryProduct 推广产品，枚举值
	// NONE：无产品
	// APP ：应用
	// PRODUCT：商品
	// WECHAT_GAME：微信小游戏
	// WECHAT_APP：微信小程序
	// BYTE_GAME：字节小游戏
	// BYTE_APP：字节小程序
	// QUICK_APP：快应用
	// AWEME：抖音号
	DeliveryProduct enum.DeliveryProduct `json:"delivery_product,omitempty"`
	// DeliveryMedium 单投放载体，枚举值
	// WECHAT_GAME：微信小游戏
	// WECHAT_APP：微信小程序
	// BYTE_GAME：字节小游戏
	// BYTE_APP：字节小程序
	// PRODUCT：商品
	// ORANGE： 橙子落地页
	// THIRDPARTY ：自研落地页
	// ENTERPRISE ：企业号落地页
	// AWEME： 抖音号
	// QUICK_APP：快应用
	// APP：应用
	// LANDING_PAGE_LINK：落地页
	DeliveryMedium enum.DeliveryMedium `json:"delivery_medium,omitempty"`
	// MultiDeliveryMedium 多投放载体，仅当landing_type = LINK 销售线索推广目的下会返回
	// 枚举值：
	// ORANGE_AND_AWEME优选投放橙子落地页和抖音主页
	MultiDeliveryMedium enum.MultiDeliveryMedium `json:"multi_delivery_medium,omitempty"`
	// DownloadURL 下载链接，landing_type=APP 子目标为 DOWNLOAD 或者LAUNCH 时有效且必填
	// - 下载、调用场景传入说明：
	// IOS：需要为iTunes官方地址
	// Android：需要为「应用管理中心」提供的下载链接，获取下载链接可参考【应用管理】
	DownloadURL string `json:"download_url,omitempty"`
	// DownloadType 载方式，landing_type=APP 子目标为 DOWNLOAD 时有效，子目标为LAUNCH时无需传入
	// 允许值：DOWNLOAD_URL 直接下载（默认值）、EXTERNAL_URL 落地页下载
	DownloadType enum.DownloadType `json:"download_type,omitempty"`
	// DownloadMode 优先从系统应用商店下载（下载模式），landing_type=APP 子目标为 DOWNLOAD 时有效，子目标为LAUNCH时无需传入
	// 允许值：APP_STORE_DELIVERY 优先商店下载、 DEFAULT 默认下载（默认值）
	// 仅安卓应用下载支持，请确保投放的应用在应用商店内已上架选择后，将优先跳转目标应用对应手机系统应用商店安装详情页，跳转失败则使用下载链接下载
	DownloadMode enum.DownloadMode `json:"download_mode,omitempty"`
	// QuickAppId 快应用资产id ，从【查询快应用信息】接口获取，仅支持已通过审核的快应用资产
	QuickAppId uint64 `json:"quick_app_id,omitempty"`
	// MicroAppInstanceID 字节小程序/小游戏资产id
	MicroAppInstanceID uint64 `json:"micro_app_instance_id,omitempty"`
	// LaunchType 调起方式， landing_type = APP 且子目标为LAUNCH有效
	// 允许值： DIRECT_OPEN 直接调起（默认值）、EXTERNAL_OPEN 落地页调起
	LaunchType enum.LaunchType `json:"launch_type,omitempty"`
	// PromotionType 投放内容，允许值：AWEME_HOME_PAGE：抖音主页（默认）LANDING_PAGE_LINK：落地页
	// 当landing_type = NATIVE_ACTION && marketing_goal=短视频/图文时有效
	PromotionType enum.PromotionType `json:"promotion_type,omitempty"`
	// OpenURL Deeplink直达链接，landing_type = APP 且子目标为 LAUNCH 时有效且必填
	// 直达链接仅支持部分App唤起（点击唤起APP），点击创意将优先跳转App，再根据投放内容跳转相关链接
	OpenURL string `json:"open_url,omitempty"`
	// UlinkURL ulink直达链接，landing_type = APP 且子目标为LAUNCH 时有效
	// 仅支持穿山甲广告位
	// 当ad_type=SEARCH搜索广告时不支持ulink
	UlinkURL string `json:"ulink_url,omitempty"`
	// SubscribeURL 预约下载链接，landing_type = APP 且子目标为 RESERVE 时有效且必填
	SubscribeURL string `json:"subscribe_url,omitempty"`
	// AssetType 资产类型，landing_type = LINK 或SHOP时有效且必填
	// 允许值： ORANGE 橙子落地页、THIRDPARTY 自研落地页
	// ENTERPRISE 企业号落地页
	// 企业号落地页仅landing_type = LINK&&marketing_goal= LIVE 时支持，仅当adv_id下绑定的抖音号为企业号时支持传入
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	// MicroPromotionType 小程序类型，landing_type = MICRO_GAME 时有效且必填
	// 允许值： WECHAT_GAME 微信小游戏、WECHAT_APP微信小程序、BYTE_GAME字节小游戏、BYTE_APP字节小程序
	MicroPromotionType enum.MicroPromotionType `json:"micro_promotion_type,omitempty"`
	// DpaAdType DPA广告类型，
	// 允许值: DPA_LINK 落地页
	// 当landing_type为dpa时有效且必填
	DpaAdType enum.DpaAdType `json:"dpa_adtype,omitempty"`
	// OptimizeGoal 优化目标
	OptimizeGoal *OptimizeGoal `json:"optimize_goal,omitempty"`
	// LandingPageStayTime 店铺停留时长，单位为毫秒
	LandingPageStayTime enum.LandingPageStayTime `json:"landing_page_stay_time,omitempty"`
	// DeliveryRange 广告版位
	DeliveryRange *DeliveryRange `json:"delivery_range,omitempty"`
	// Audience 定向设置
	Audience *Audience `json:"audience,omitempty"`
	// DeliverySetting 投放设置
	DeliverySetting *DeliverySetting `json:"delivery_setting,omitempty"`
	// TrackURLSetting 监测链接设置
	TrackURLSetting *TrackURLSetting `json:"track_url_setting,omitempty"`
}

func (p Project) GetID() uint64 {
	return p.ProjectID
}

func (p Project) GetOpenURL() string {
	return p.OpenURL
}

func (p Project) GetActionTrackURL() []string {
	if p.TrackURLSetting == nil {
		return nil
	}
	if p.TrackURLSetting.ActionTrackURL == nil {
		return nil
	}
	return *p.TrackURLSetting.ActionTrackURL
}

func (p Project) IsProject() bool {
	return true
}

// RelatedProduct 关联产品投放相关
type RelatedProduct struct {
	// ProductSetting 商品库设置
	// 允许值：SINGLE 启用SDPA、NO_MAP不启用（默认值）
	// 当delivery_mode选择PROCEDURAL且landing_type选择LINK时，该字段必须传入SINGLE
	ProductSetting enum.ProductSetting `json:"product_setting,omitempty"`
	// ProductPlatformID 商品库ID，当启用商品库时必填，可通过【商品广告-查询商品库】查询，创建后不可修改
	// 当delivery_mode选择PROCEDURAL且landing_type选择LINK时，传入报错
	ProductPlatformID uint64 `json:"product_platform_id,omitempty"`
	// 产品ID，当启用商品库时必填，可通过【商品广告-获取商品列表】 查询，创建后不可修改
	// 当delivery_mode选择PROCEDURAL且landing_type选择LINK时，传入报错
	ProductID model.JSONUint64 `json:"product_id,omitempty"`
	// UniqueProductID 线索版产品ID，可通过【商品广告-获取线索商品列表】查询获取id（该接口下的product_id就是unique_product_id），创建后不可修改
	// 如果投放线索版商品，只需要传入unique_product_id
	UniqueProductID uint64 `json:"unique_product_id,omitempty"`
	// AssetID 物件ID，可通过【商品广告-获取投放条件列表】获取，创建后不可修改。
	AssetID uint64 `json:"asset_id,omitempty"`
	// Products 产品ID列表，上限为10
	// 当delivery_mode选择PROCEDURAL且landing_type选择LINK时有效且必填
	Products []RelatedProduct `json:"products,omitempty"`
}

// DpaProductTarget 自定义筛选条件（商品投放条件）。用于圈定商品投放范围，结合商品库字段搭配判断条件，圈定商品投放范围。获取商品库元信息-商品广告-商业开放平台
type DpaProductTarget struct {
	// Title 筛选字段
	Title string `json:"title,omitempty"`
	// Rule 定向规则，允许值：type 为int、 double、 long其中一种时支持=、 !=、 >、 <；
	// type = string时支持=、 !=、 notEmpty
	Rule string `json:"rule,omitempty"`
	// Type 字段类型，允许值：int、 double、 long、 string
	Type string `json:"type,omitempty"`
	// Value 规则值
	Value string `json:"value,omitempty"`
}

// OptimizeGoal 优化目标
type OptimizeGoal struct {
	// AssetIDs 事件管理资产 id，list长度上限为 1（landing_type=MICRO_GAME，小程序类型为 BYTE_GAME、BYTE_APP时上限为2）
	// landing_type=APP 子目标为 DOWNLOAD 或者LAUNCH 时有效且必填;
	// landing_type=LINK或SHOP或DPA 落地页类型为 THIRDPARTY 时有效且必填;
	// landing_type=LINK 、SHOP或DPA 落地页类型为 ORANGE 时无需传入，仅需传入external_action;
	// landing_type=MICRO_GAME 小程序类型为 WECHAT_GAME、WECHAT_APP 时无需传入，仅需传入external_action；
	// landing_type=MICRO_GAME 小程序类型为 BYTE_APP、BYTE_GAME时有效且必填，且需要传入两个资产id，数组第一个为主资产id，第二个为备用资产id（兜底落地页，可不传）；
	// landing_type=QUICK_APP 时有效且必填
	AssetIDs []uint64 `json:"asset_ids,omitempty"`
	// ConvertID 转化跟踪id
	ConvertID uint64 `json:"convert_id,omitempty"`
	// ExternalAction 优化目标，可通过【资产-事件管理-获取可用优化目标(体验版)】查询获取
	// landing_type=APP 子目标为 DOWNLOAD 或者LAUNCH 时有效且必填;
	// landing_type=LINK 、SHOP、MICRO_GAME 或DPA 时有效且必填;
	// landing_type= QUICK_APP时有效且必填；
	// 当delivery_mode = PROCEDURAL &&landing_type = APP/MICRO_GAME时，external_action=付费时，deep_external_action不传/传空，deep_bid_type仅支持每次付费/首次付费，当external_action=付费，deep_external_action=付费ROI，deep_bid_type仅支持ROI系数
	ExternalAction enum.AdConvertType `json:"external_action,omitempty"`
	// DeepExternalAction 深度转化目标，可通过【资产-事件管理-获取可用优化目标(体验版)】查询获取，当delivery_mode = PROCEDURAL &&landing_type = APP/MICRO_GAME时可传空/传0/「付费ROI」
	DeepExternalAction enum.DeepExternalAction `json:"deep_external_action,omitempty"`
	// ValueOptimizedType 目标优化类型，
	// 允许值：OFF表示不启用，ACTION表示行为优化，VALUE表示价值优化
	// 默认值：OFF
	// 当前仅支持推广目的为销售线索收集（landing_type=LINK）时传入ACTION，VALUE
	ValueOptimizedType string `json:"value_optimized_type,omitempty"`
	// PaidSwitch 字节提供的归因方式，允许值：
	// 1启用；2 不启用（默认值）
	// 仅当landing_type=SHOP，external_action=AD_CONVERT_TYPE_APP_ORDER app内下单（电商）时有效
	PaidSwitch int `json:"paid_switch,omitempty"`
}

// DeliveryRange 广告版位
type DeliveryRange struct {
	// InventoryCatalog 广告位大类，允许值： MANUAL 首选媒体，UNIVERSAL_SMART 通投智选
	// 当delivery_mode = PROCEDURAL &&landing_type = APP/MICRO_GAME/LINK时仅支持通投智选；
	// 当 marketing_goal= LIVE时，仅支持 MANUAL首选媒体
	// 当 ad_type=SEARCH时，仅支持 MANUAL首选媒体
	// 当landing_type=NATIVE_ACTION时仅支持MANUAL首选媒体
	InventoryCatalog enum.InventoryCatalog `json:"inventory_catalog,omitempty"`
	// InventoryType 广告投放位置（首选媒体），inventory_catalog = MANUAL 有效且必填，允许值：
	// INVENTORY_FEED 今日头条
	// INVENTORY_VIDEO_FEED 西瓜视频
	// INVENTORY_AWEME_FEED 抖音短视频
	// INVENTORY_TOMATO_NOVEL 番茄小说
	// INVENTORY_UNION_SLOT 穿山甲
	// UNION_BOUTIQUE_GAME ohayoo精品游戏
	// INVENTORY_SEARCH 搜索广告
	// 当 marketing_goal= LIVE时，仅支持INVENTORY_AWEME_FEED 抖音短视频INVENTORY_UNION_SLOT穿山甲;
	// 当 ad_type=SEARCH时，仅支持 INVENTORY_SEARCH 搜索广告;
	// 当landing_type =NATIVE_ACTION&&marketing_goal=短视频/图文时，仅支持INVENTORY_AWEME_FEED 抖音短视频；
	// 当landing_type =NATIVE_ACTION&&marketing_goal=LIVE时，支持INVENTORY_AWEME_FEED 抖音短视频和INVENTORY_UNION_SLOT 穿山甲
	InventoryType []enum.StatInventoryType `json:"inventory_type,omitempty"`
	// UnionVideoType 投放形式（穿山甲视频创意类型）
	// inventory_catalog = MANUAL && inventory_type 仅有 INVENTORY_UNION_SLOT 时 有效且必填；
	// 当 marketing_goal=LIVE时，不支持该字段；
	// 允许值：ORIGINAL_VIDEO 原生视频、REWARDED_VIDEO 激励视频、SPLASH_VIDEO 开屏视频
	UnionVideoType *enum.UnionVideoType `json:"union_video_type,omitempty"`
}

// Audience 定向设置
type Audience struct {
	// AudiencePackageID 定向包ID，定向包ID由【工具-定向包管理-获取定向包】获取
	// 如果同时传定向包ID和自定义用户定向参数时，仅定向包中的定向生效
	// 当delivery_mode = PROCEDURAL &&landing_type = APP/MICRO_GAME时不支持定向包
	// 当landing_type=NATIVE_ACTION时仅支持抖音号定向包
	// 当 ad_type=SEARCH时仅支持搜索定向包
	AudiencePackageID uint64 `json:"audience_package_id,omitempty"`
	// District 地域类型
	// 允许值: CITY 省市、 COUNTY 区县、BUSINESS_DISTRICT商圈、REGION 行政区域、OVERSEA 海外区域、NONE 不限
	// 使用省市示例：{"district": "CITY","city": [12]}
	// 使用区县示例：{"district": "COUNTY","city": [130102]}
	// 使用行政区域示例：{"district":"REGION": "city":[31], "region_versio":"1.0.0"}
	// 使用海外区域示例：{"district":"OVERSEA": "city":[3041566], "region_versio":"1.0.0"}
	// 当ad_type=SEARCH时不支持OVERSEA海外
	District enum.District `json:"district,omitempty"`
	// Geolocation 从地图添加(地图位置)，district为"BUSINESS_DISTRICT"时填写，最多允许添加1000个
	Geolocation []model.Geolocation `json:"geolocation,omitempty"`
	// RegionVersion 行政区域版本号，通过【获取行政信息】接口获取
	// district =REGION/OVERSEA时必填
	RegionVersion string `json:"region_version,omitempty"`
	// RegionRecommend 地域智能放量定向，ON为开启、OFF为关闭
	RegionRecommend enum.OnOff `json:"region_recommend,omitempty"`
	// City 地域定向省市或者区县列表，当district=CITY/COUNTY/REGION/OVERSEA时
	// district =CITY/COUNTY时，详见【附件-city.json】
	// district =REGION/OVERSEA时，通过【获取行政信息】接口获取
	City *[]uint64 `json:"city,omitempty"`
	// LocationType 位置类型
	// 枚举值：CURRENT正在该地区的用户，HOME居住在该地区的用户，TRAVEL到该地区旅行的用户，ALL该地区内的所有用户
	LocationType enum.LocationType `json:"location_type,omitempty"`
	// Gender 性别， 详见【附录-受众性别】
	Gender enum.AudienceGender `json:"gender,omitempty"`
	// Age 年龄， 详见【附录-受众年龄区间】
	Age *[]enum.AudienceAge `json:"age,omitempty"`
	// RetargetingTageInclude 定向人群包列表（自定义人群）
	RetargetingTagsInclude *[]uint64 `json:"retargeting_tags_include,omitempty"`
	// RetargetingTagsExclude 排除人群包列表（自定义人群）
	RetargetingTagsExclude *[]uint64 `json:"retargeting_tags_exclude,omitempty"`
	// InterestActionMode 行为兴趣，枚举值：UNLIMITED 不限、CUSTOM 自定义、 RECOMMEND系统推荐
	InterestActionMode enum.InterestActionMode `json:"interest_action_mode,omitempty"`
	// ActionScene 行为场景，枚举值：E-COMMERCE 电商互动行为、 NEWS 资讯互动行为、 APP APP推广互动行为行为场景
	ActionScene *[]enum.ActionScene `json:"action_scene,omitempty"`
	// ActionDays 用户发生行为天数，枚举值：7、 15、 30、 60、 90、 180、 365用户发生行为天数
	ActionDays int `json:"action_days,omitempty"`
	// ActionCategories 行为类目词，行为类目可以通过【工具-行为兴趣词管理-行为类目查询】获取行为类目词
	ActionCategories *[]uint64 `json:"action_categories,omitempty"`
	// ActionWords 行为关键词，行为类目可以通过【工具-行为兴趣词管理-行为关键词查询】获取行为关键词
	ActionWords *[]uint64 `json:"action_words,omitempty"`
	// InterestCategories 兴趣类目词
	InterestCategories *[]uint64 `json:"interest_categories,omitempty"`
	// InterestWords 兴趣关键词
	InterestWords *[]uint64 `json:"interest_words,omitempty"`
	// AwemeFanBehaviors 抖音达人互动用户行为类型, 详见【附录-抖音用户行为类型】
	AwemeFanBehaviors *[]enum.Behavior `json:"aweme_fan_behaviors,omitempty"`
	// AwemeFanTimeScope 抖音达人互动行为时间范围，枚举值：FIFTEEN_DAYS 15天、THIRTY_DAYS 30天、SIXTY_DAYS 60天
	AwemeFanTimeScope string `json:"aweme_fan_time_scope,omitempty"`
	// AwemeFanCategories 抖音达人分类ID列表，可通过【工具-抖音达人-查询抖音类目列表】接口获取
	AwemeFanCategories *[]uint64 `json:"aweme_fan_categories,omitempty"`
	// AwemeFanAccounts 抖音达人ID列表，可通过【工具-抖音达人-查询抖音类目下的推荐达人】接口获取。
	AwemeFanAccounts *model.Uint64s `json:"aweme_fan_accounts,omitempty"`
	// SuperiorPopularityType 媒体定向，详见【附录-媒体定向】
	SuperiorPopularityType enum.SuperiorPopularityType `json:"superior_popularity_type,omitempty"`
	// FlowPackage 定向逻辑，可通过【工具-穿山甲流量包-获取穿山甲流量包】
	FlowPackage *[]uint64 `json:"flow_package,omitempty"`
	// ExcludeFlowPackage 排除定向逻辑，可通过【工具-穿山甲流量包-获取穿山甲流量包】
	ExcludeFlowPackage *[]uint64 `json:"exclude_flow_package,omitempty"`
	// Platform 投放平台列表，枚举值：ANDROID, IOS, HARMONY
	Platform *[]enum.AudiencePlatform `json:"platform,omitempty"`
	// AndroidOsv 最低安卓版本
	AndroidOsv string `json:"android_osv,omitempty"`
	// IosOsv 最低IOS版本
	IosOsv string `json:"ios_osv,omitempty"`
	// HarmonyOsv 鸿蒙版本，允许值：5.0
	// 仅支持定向- platform为鸿蒙时进行鸿蒙版本设置，当platform为ANDROID、IOS时不支持设置
	// 当platform为鸿蒙时有效且必填
	HarmonyOsv string `json:"harmony_osv,omitempty"`
	// DeviceType 设备类型，枚举值：MOBILE、PAD
	DeviceType *[]string `json:"device_type,omitempty"`
	// Ac 网络类型, 详见【附录-受众网络类型】
	Ac *[]string `json:"ac,omitempty"`
	// Carrier 运营商, 详见【附录-受众运营商类型】
	Carrier *[]enum.Carrier `json:"carrier,omitempty"`
	// HideIfExists 过滤已安装，枚举值：UNLIMITED不限、FILTER 过滤、TARGETING 定向
	HideIfExists enum.HideIfConverted `json:"hide_if_exists,omitempty"`
	// HideIfConverted 过滤已转化用户
	// 枚举值：NO_EXCLUDE 不限制、PROMOTION 广告、PROJECT 推广项目、ADVERTISER 广告账户、APP 应用、CUSTOMER 客户、ORGANIZATION 组织
	HideIfConverted enum.HideIfConverted `json:"hide_if_converted,omitempty"`
	// ConvertedTimeDuration 过滤时间范围，详见 【附录-过滤时间范围】
	ConvertedTimeDuration enum.ConvertedTimeDuration `json:"converted_time_duration,omitempty"`
	// FilterAwemeAbnormalActive 过滤高活跃用户，即过滤关注、点赞、评论行为高活跃的用户允许值：
	// ON 过滤
	// OFF不过滤（默认值）
	// 当marketing_goal= LIVE 且inventory_type非仅穿山甲时，支持该字段
	FilterAwemeAbnormalActive enum.OnOff `json:"filter_aweme_abnormal_active,omitempty"`
	// FilterAwemeFansCount 过滤高关注数用户，例如"filter_aweme_fans_count": 1000表示过滤粉丝数在1000以上的用户
	// 允许值：1000、500、200
	// 当marketing_goal= Live 且inventory_type非仅穿山甲时，支持该字段
	FilterAwemeFansCount int `json:"filter_aweme_fans_count,omitempty"`
	// 过滤自己的粉丝，允许值：
	// ON 过滤
	// OFF不过滤（默认值）
	// 当marketing_goal= Live 且inventory_type非仅穿山甲时，支持该字段
	FilterOwnAwemeFans enum.OnOff `json:"filter_own_aweme_fans,omitempty"`
	// DeviceBrand 手机品牌, 详见【附录-手机品牌】
	DeviceBrand *[]string `json:"device_brand,omitempty"`
	// LaunchPrice 手机价格，价格区间，最高11000（表示1w以上）
	LaunchPrice *[]int `json:"launch_price,omitempty"`
	// AutoExtendTargets 可放开定向，枚举值：AGE 年龄、REGION 地域、GENDER 性别、CUSTOM_AUDIENCE 自定人群-定向
	AutoExtendTargets *[]string `json:"auto_extend_targets,omitempty"`
	// DpaCity 地域匹配-商品所在城市开启时，仅将商品投放给位于该商品设置的可投城市的用户默认值：OFF允许值：OFF，ON（OFF表示不启用，ON表示启用）
	// DPA推广目的下有效
	DpaCity enum.OnOff `json:"dpa_city,omitempty"`
	// DpaRtaSwitch RTA重定向开关，
	// 默认值：OFF允许值：OFF，ON（OFF表示不启用，ON表示启用）
	// 启用后，需通过【设置账户下RTA策略生效范围-工具-商业开放平台】绑定rta策略
	// DPA推广目的下有效
	DpaRtaSwitch enum.OnOff `json:"dpa_rta_switch,omitempty"`
	// RtaID RTA策略ID，通过【获取可用的RTA策略】接口获取
	// 开启RTA重定向开关时必填
	// DPA推广目的下有效
	RtaID uint64 `json:"rta_id,omitempty"`
	// DpaRtaRecommedType RTA推荐逻辑，ONLY仅RTA推荐商品，MORE基于RTA推荐更多商品，开启RTA重定向开关时必填
	// DPA推广目的下有效
	DpaRtaRecommendType enum.DpaRtaRecommendType `json:"dpa_rta_recommend_type,omitempty"`
}

// DeliverySetting 投放设置
type DeliverySetting struct {
	// ScheduleType 投放时间类型，枚举值：SCHEDULE_FROM_NOW 从今天起长期投放、SCHEDULE_START_END 设置开始和结束日期
	ScheduleType enum.ScheduleType `json:"schedule_type,omitempty"`
	// StartTime 投放起始时间，如：2017-01-01 精确到天
	StartTime string `json:"start_time,omitempty"`
	// EndTime 投放结束时间，如：2017-01-01 精确到天
	EndTime string `json:"end_time,omitempty"`
	// ScheduleTime 投放时段
	ScheduleTime string `json:"schedule_time,omitempty"`
	// FilterNightSwitch 过滤夜间投放，枚举值：ON过滤开启 OFF过滤不开启
	// 默认值：OFF过滤不开启
	// schedule_time传入，且filter_night_switch传入ON时，以filter_night_switch为准
	// schedule_time传入，且filter_night_switch传入OFF时，以schedule_time为准
	// 仅当landing_type = link && delivery_mode = procedural 时传入有效
	FilterNightSwitch string `json:"filter_night_switch,omitempty"`
	// ProjectCustom 项目成本稳投，当ad_type=SEARCH&&bid_type=CUSTOM 稳定成本时有效，当dea有值时，不支持项目成本稳投
	// 允许值：
	// ON 开启（默认值），
	// OFF 不开启
	ProjectCustom enum.OnOff `json:"project_custom,omitempty"`
	// Bid 点击出价/展示出价，当delivery_mode = MANUAL&&项目成本稳投开启&&pricing=CPC时填写有效；取值范围：0.2-999
	Bid float64 `json:"bid,omitempty"`
	// DeepBidType 深度优化方式，当转化目标中含有深度转化时，该字段必填
	// 允许值：
	// DEEP_BID_MIN 自定义手动出价、
	// ROI_COEFFICIENT ROI系数出价、
	// BID_PER_ACTION 每次付费出价、
	// SOCIAL_ROIROI三出价
	// 对于每次付费的转化，深度优化类型需要设置为BID_PER_ACTION（每次付费出价）
	// FORM_BID优选表单出价（landing_type=link&&external_action=表单提交/多转化&&deep_external_action为空时，支持优选表单出价/不启用）
	// PHONE_CONNECT_BID电话接通出价
	// (landing_type=link&&external_action=表单提交，deep_external_action=电话接通时，deep_bid_type仅支持PHONE_CONNECT_BID）
	// DEEP_BID_DEFAULT  不启用
	DeepBidType enum.DeepBidType `json:"deep_bid_type,omitempty"`
	// BidType 竞价策略，允许值：CUSTOM 稳定成本、NO_BID 最大转化投放、UPPER_CONTROL控制成本上限、CONSERVATIVE放量投放；
	// 当delivery_mode = PROCEDURAL &&landing_type = APP/MICRO_GAME时不支持UPPER_CONTROL；
	// 当delivery_mode = PROCEDURAL &&landing_type = LINK时不支持UPPER_CONTROL和NO_BID；
	// 选择深度优化目标后，不支持UPPER_CONTROL；
	// 当external_action=AD_CONVERT_TYPE_CLICK_NUM 点击量 或 AD_CONVERT_TYPE_SHOW_OFF_NUM 展示量时，不支持NO_BID；
	// 当 ad_type=SEARCH时仅支持稳定成本和最大转化
	BidType enum.BidType `json:"bid_type,omitempty"`
	// BidSpeed 投放速度，允许值：BALANCE 匀速,FAST 加速
	BidSpeed enum.BidSpeed `json:"bid_speed,omitempty"`
	// BudgetMode 项目预算类型， 枚举值：BUDGET_MODE_INFINITE 不限、BUDGET_MODE_DAY 日预算, BUDGET_MODE_TOTAL 总预算
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// Budget 项目预算
	Budget float64 `json:"budget,omitempty"`
	// Pricing 计费方式，允许值：PRICING_CPM 按展示付费,PRICING_CPC 按点击付费,PRICING_OCPM 目标转化出价-按展示付费（默认值）,PRICING_OCPC 目标转化出价-按点击付费
	// 当ea=AD_CONVERT_TYPE_SHOW_OFF_NUM 展示量时：仅支持CPM；
	// 当ea= AD_CONVERT_TYPE_CLICK_NUM点击量时：仅支持CPC；
	// 当ea !=AD_CONVERT_TYPE_SHOW_OFF_NUM 展示量 或AD_CONVERT_TYPE_CLICK_NUM 点击量时，在固定支持oCPM的情况下，以下情况额外支持oCPC：1)首选媒体仅选择穿山甲 2)广告类型为搜索直投；
	// 当ad_type=SEARCH时，不支持PRICING_CPM
	Pricing enum.PricingType `json:"pricing,omitempty"`
	// CpaBid 目标转化出价/预期成本(注意：nobid不返回该字段)
	CpaBid float64 `json:"cpa_bid,omitempty"`
	// DeepCpaBid 深度优化出价，deep_bid_type = DEEP_BID_MIN 时必填。取值范围：0.1~10000元 注意：当 bid_type = NO_BID时，不填写该字段，否则会报错
	DeepCpaBid float64 `json:"deep_cpabid,omitempty"`
	// RoiGoal 深度转化ROI系数(注意：nobid不返回该字段)
	RoiGoal float64 `json:"roi_goal,omitempty"`
	// BudgetOptimizeSwitch 支持预算择优分配，枚举值： ON 开启，OFF 不开启
	BudgetOptimizeSwitch enum.OnOff `json:"budget_optimize_switch,omitempty"`
	// SearchContinueDelivery 续投，仅当delivery_type = DURATION搜索广告周期投放时必填，允许值：
	// OFF:关闭，关闭表示周期稳投7天后投放将自动终止
	// ON:开启，开启表示投放结束后将继续维持7天固定周期的投放，跑量更加稳定，可以延续跑量
	// 仅支持周期稳投链路，其他链路下传入该参数不生效
	SearchContinueDelivery enum.OnOff `json:"search_continue_delivery,omitempty"`
	// ShopMultiRoiGoals 多ROI系数
	// 条件必填，object[]，多ROI系数设置，表示引流电商多平台投放ROI系数及平台信息，广告主可按照电商平台分别确定ROI系数，分平台调控出价。list长度最长为4
	// 多平台优选投放白名单内客户，在以下组合场景时shop_multi_roi_goals有效且必填
	// 推广目的 = 电商（landing_type = SHOP）
	// 投放方式 = 自动投放(delivery_mode = MANUAL)
	// 优化目标 = APP 内下单(external_action = AD_CONVERT_TYPE_APP_ORDER)
	// 深度优化方式 = ROI系数(deep_bid_type = ROI_DIRECT_MAIL)
	ShopMultiRoiGoals []ShopMultiRoiGoal `json:"shop_multi_roi_goals,omitempty"`
}

// TrackURLSetting 监测链接设置
type TrackURLSetting struct {
	// TrackURLType 监测链接类型，区分使用监测链接组或者自定义链接
	// 枚举值：CUSTOM 自定义链接、GROUP_ID 监测链接组
	TrackURLType string `json:"track_url_type,omitempty"`
	// TrackURLGroupID 监测链接组id
	TrackURLGroupID uint64 `json:"track_url_group_id,omitempty"`
	// TrackURL 展示（监测链接）
	TrackURL *[]string `json:"track_url,omitempty"`
	// ActionTrackURL 点击（监测链接）
	ActionTrackURL *[]string `json:"action_track_url,omitempty"`
	// VideoPlayEffectiveTrackURL 视频有效播放（监测链接）
	VideoPlayEffectiveTrackURL *[]string `json:"video_play_effective_track_url,omitempty"`
	// VideoPlayDoneTrackURL 视频播完（监测链接）
	VideoPlayDoneTrackURL *[]string `json:"video_play_done_track_url,omitempty"`
	// VideoPlayFirstTrackURL 视频开始播放（监测链接）
	VideoPlayFirstTrackURL *[]string `json:"video_play_first_track_url,omitempty"`
	// SendType 数据发送方式，枚举值：SERVER_SEND 服务器端上传, CLIENT_SEND 客户端上传
	SendType enum.SendType `json:"send_type,omitempty"`
}

// Keyword 搜索关键词
type Keyword struct {
	// Word 关键词
	Word string `json:"word,omitempty"`
	// BidType 出价类型。 允许值:FEED_TO_SEARCH 搜索快投
	BidType string `json:"bid_type,omitempty"`
	// MatchType 匹配类型，允许值: PHRASE短语匹配，EXTENSIVE广泛匹配，PRECISION精准匹配
	MatchType enum.KeywordMatchType `json:"match_type,omitempty"`
	// Bid 出价。取值范围：0.2至999.0。
	// 当pricing为PRICING_OCPC/PRICING_OCPM时，不支持出价
	Bid float64 `json:"bid,omitempty"`
}

// ShopMultiRoiGoals 多ROI系数
type ShopMultiRoiGoal struct {
	// RoiGoal ROI系数，范围(0.01,100]，精度：最多保留小数点后四位
	RoiGoal float64 `json:"roi_goal,omitempty"`
	// ShopPlatform ROI系数所属平台，允许值：PDD 拼多多、TB 淘宝、JD 京东、OTHER 其他
	ShopPlatform enum.ShopPlatform `json:"shop_platform,omitempty"`
}
