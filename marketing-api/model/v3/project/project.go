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
	// Pricing 出价方式
	Pricing enum.PricingType `json:"pricing,omitempty"`
	// PackageName 应用包名
	PackageName string `json:"package_name,omitempty"`
	// AppName 应用名
	AppName string `json:"app_name,omitempty"`
	// RelatedProduct 关联产品投放相关
	RelatedProduct *RelatedProduct `json:"related_product,omitempty"`
	// AssetType 资产类型
	AssetType enum.AssetType `json:"asset_type,omitempty"`
	// MicroPromotionType 小游戏类型
	MicroPromotionType enum.MicroPromotionType `json:"micropromotion_type,omitempty"`
	// DownloadURL 下载链接
	DownloadURL string `json:"download_url,omitempty"`
	// DownloadType 下载方式，枚举值：DOWNLOAD_URL 直接下载、EXTERNAL_URL 落地页下载
	DownloadType enum.DownloadType `json:"download_type,omitempty"`
	// DownloadMode 优先从系统应用商店下载（下载模式），枚举值：APP_STORE_DELIVERY 优先商店下载、 DEFAULT 默认下载
	DownloadMode enum.DownloadMode `json:"download_mode,omitempty"`
	// LaunchType 调起方式，枚举值： DIRECT_OPEN 直接调起、EXTERNAL_OPEN 落地页调起
	LaunchType enum.LaunchType `json:"launch_type,omitempty"`
	// OpenURL Deeplink直达链接
	OpenURL string `json:"open_url,omitempty"`
	// UlinkURL ulink直达链接
	UlinkURL string `json:"urlink_url,omitempty"`
	// SubscribeURL 预约下载链接
	SubscribeURL string `json:"subscribe_url,omitempty"`
	// OptimizeGoal 优化目标
	OptimizeGoal *OptimizeGoal `json:"optimize_goal,omitempty"`
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
	// ProductSetting 商品库设置，枚举值：SINGLE 启用SDPA、NO_MAP不启用
	ProductSetting enum.ProductSetting `json:"product_setting,omitempty"`
	// ProductPlatformID 商品库ID
	ProductPlatformID uint64 `json:"product_platform_id,omitempty"`
	// ProductID 产品ID
	ProductID string `json:"product_id,omitempty"`
}

// OptimizeGoal 优化目标
type OptimizeGoal struct {
	// AssetIDs 事件管理资产id
	AssetIDs []uint64 `json:"asset_ids,omitempty"`
	// ConvertID 转化跟踪id
	ConvertID uint64 `json:"convert_id,omitempty"`
	// ExternalAction 优化目标
	ExternalAction enum.AdConvertType `json:"external_action,omitempty"`
	// DeepExternalAction 深度转化目标
	DeepExternalAction enum.DeepExternalAction `json:"deep_external_action,omitempty"`
}

// DeliveryRange 广告版位
type DeliveryRange struct {
	// InventoryCatalog 广告位大类，枚举值： MANUAL 首选媒体，UNIVERSAL_SMART 通投智选
	InventoryCatalog enum.InventoryCatalog `json:"inventory_catalog,omitempty"`
	// InventoryType 广告投放位置（首选媒体)
	InventoryType []enum.StatInventoryType `json:"inventory_type,omitempty"`
	// UnionVideoType 投放形式（穿山甲视频创意类型），枚举值：ORIGINAL_VIDEO 原生视频、REWARDED_VIDEO 激励视频、SPLASH_VIDEO 开屏视频
	UnionVideoType *enum.UnionVideoType `json:"union_video_type,omitempty"`
}

// Audience 定向设置
type Audience struct {
	// AudiencePackageID 定向包ID
	// 如果同时传定向包ID和自定义用户定向参数时，仅定向包中的定向生效
	AudiencePackageID uint64 `json:"audience_package_id,omitempty"`
	// District 地域类型，枚举值: CITY 省市、 COUNTY 区县、REGION 行政区域、OVERSEA 海外区域、NONE 不限
	District enum.District `json:"district,omitempty"`
	// Geolocation 从地图添加(地图位置)
	Geolocation []model.Geolocation `json:"geolocation,omitempty"`
	// RegionVersion 行政区域版本号
	RegionVersion string `json:"region_version,omitempty"`
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
	Age enum.AudienceAge `json:"age,omitempty"`
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
	AwemeFanBehaviors *[]string `json:"aweme_fan_behaviors,omitempty"`
	// AwemeFanTimeScope 抖音达人互动行为时间范围，枚举值：FIFTEEN_DAYS 15天、THIRTY_DAYS 30天、SIXTY_DAYS 60天
	AwemeFanTimeScope string `json:"aweme_fan_time_scope,omitempty"`
	// AwemeFanCategories 抖音达人分类ID列表，可通过【工具-抖音达人-查询抖音类目列表】接口获取
	AwemeFanCategories *[]uint64 `json:"aweme_fan_categories,omitempty"`
	// AwemeFanAccounts 抖音达人ID列表，可通过【工具-抖音达人-查询抖音类目下的推荐达人】接口获取。
	AwemeFanAccounts *[]uint64 `json:"aweme_fan_accounts,omitempty"`
	// SuperiorPopularityType 媒体定向，详见【附录-媒体定向】
	SuperiorPopularityType enum.SuperiorPopularityType `json:"superior_popularity_type,omitempty"`
	// FlowPackage 定向逻辑，可通过【工具-穿山甲流量包-获取穿山甲流量包】
	FlowPackage *[]uint64 `json:"flow_package,omitempty"`
	// ExcludeFlowPackage 排除定向逻辑，可通过【工具-穿山甲流量包-获取穿山甲流量包】
	ExcludeFlowPackage *[]uint64 `json:"exclude_flow_package,omitempty"`
	// Platform 投放平台列表，枚举值：ANDROID、IOS
	Platform enum.AudiencePlatform `json:"platform,omitempty"`
	// AndroidOsv 最低安卓版本
	AndroidOsv string `json:"android_osv,omitempty"`
	// IosOsv 最低IOS版本
	IosOsv string `json:"ios_osv,omitempty"`
	// DeviceType 设备类型，枚举值：MOBILE、PAD
	DeviceType string `json:"device_type,omitempty"`
	// Ac 网络类型, 详见【附录-受众网络类型】
	Ac *[]string `json:"ac,omitempty"`
	// Carrier 运营商, 详见【附录-受众运营商类型】
	Carrier *[]enum.Carrier `json:"carrier,omitempty"`
	// HideIfExists 过滤已安装，枚举值：UNLIMITED不限、FILTER 过滤、TARGETING 定向
	HideIfExists string `json:"hide_if_exists,omitempty"`
	// HideIfConverted 过滤已转化用户
	// 枚举值：NO_EXCLUDE 不限制、PROMOTION 广告、PROJECT 推广项目、ADVERTISER 广告账户、APP 应用、CUSTOMER 客户、ORGANIZATION 组织
	HideIfConverted string `json:"hide_if_converted,omitempty"`
	// ConvertedTimeDuration 过滤时间范围，详见 【附录-过滤时间范围】
	ConvertedTimeDuration string `json:"converted_time_duration,omitempty"`
	// DeviceBrand 手机品牌, 详见【附录-手机品牌】
	DeviceBrand *[]string `json:"device_brand,omitempty"`
	// LaunchPrice 手机价格，价格区间，最高11000（表示1w以上）
	LaunchPrice *[]int `json:"launch_price,omitempty"`
	// AutoExtendTargets 可放开定向，枚举值：AGE 年龄、REGION 地域、GENDER 性别、CUSTOM_AUDIENCE 自定人群-定向
	AutoExtendTargets *[]string `json:"auto_extend_targets,omitempty"`
}

// DeliverySetting 投放设置
type DeliverySetting struct {
	// CpaBid 目标转化出价/预期成本(注意：nobid不返回该字段)
	CpaBid float64 `json:"cpa_bid,omitempty"`
	// RoiGoal 深度转化ROI系数(注意：nobid不返回该字段)
	RoiGoal float64 `json:"roi_goal,omitempty"`
	// ScheduleType 投放时间类型，枚举值：SCHEDULE_FROM_NOW 从今天起长期投放、SCHEDULE_START_END 设置开始和结束日期
	ScheduleType enum.ScheduleType `json:"schedule_type,omitempty"`
	// StartTime 投放起始时间，如：2017-01-01 精确到天
	StartTime string `json:"start_time,omitempty"`
	// EndTime 投放结束时间，如：2017-01-01 精确到天
	EndTime string `json:"end_time,omitempty"`
	// ScheduleTime 投放时段
	ScheduleTime string `json:"schedule_time,omitempty"`
	// DeepBidType 深度出价方式
	DeepBidType enum.DeepBidType `json:"deep_bid_type,omitempty"`
	// BidType 竞价策略，枚举值：CUSTOM 稳定成本、NO_BID 最大转化投放
	BidType enum.BidType `json:"bid_type,omitempty"`
	// BudgetMode 项目预算类型， 枚举值：BUDGET_MODE_INFINITE 不限、BUDGET_MODE_DAY 日预算
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// Budget 项目预算
	Budget float64 `json:"budget,omitempty"`
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
	VideoPlayEffectiveTrackURL string `json:"video_play_effective_track_url,omitempty"`
	// VideoPlayDoneTrackURL 视频播完（监测链接）
	VideoPlayDoneTrackURL string `json:"video_play_done_track_url,omitempty"`
	// VideoPlayFirstTrackURL 视频开始播放（监测链接）
	VideoPlayFirstTrackURL string `json:"video_play_first_track_url,omitempty"`
	// SendType 数据发送方式，枚举值：SERVER_SEND 服务器端上传, CLIENT_SEND 客户端上传
	SendType string `json:"send_type,omitempty"`
}
