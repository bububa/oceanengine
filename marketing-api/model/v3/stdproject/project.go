package stdproject

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// StdProject 标准项目
type StdProject struct {
	// ProjectID 项目id
	ProjectID uint64 `json:"project_id,omitempty"`
	// AdvertiserID 投放账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Name 项目名称
	Name string `json:"name,omitempty"`
	// AdType 营销类型，可选值：ALL 信息流、SEARCH 搜索
	AdType enum.CampaignType `json:"ad_type,omitempty"`
	// LandingType 营销目的，可选值：APP 应用、LINK 销售线索、MICRO_GAME 小程序、SHOP 电商
	LandingType enum.LandingType `json:"landing_type,omitempty"`
	// MarketingGoal 营销场景，可选值：LIVE 直播、VIDEO_AND_IMAGE 短视频图文
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// DeliveryMode 投放模式，可选值：PROCEDURAL 自动
	DeliveryMode enum.DeliveryMode `json:"delivery_mode,omitempty"`
	// AppPromotionType 子目标，可选值：DOWNLOAD 应用下载、LAUNCH 应用调起、RESERVE 预约下载
	AppPromotionType enum.AppPromotionType `json:"app_promotion_type,omitempty"`
	// StatusFirst 项目状态（一级），可选值：
	// PROJECT_STATUS_DELETE 已删除
	// PROJECT_STATUS_DISABLE 未投放
	// PROJECT_STATUS_DONE 已完成
	// PROJECT_STATUS_ENABLE 启用中
	// PROJECT_STATUS_FROZEN 已终止
	StatusFirst enum.StdProjectStatusFirst `json:"status_first,omitempty"`
	// StatusSecond 项目二级状态
	StatusSecond []string `json:"status_second,omitempty"`
	// OptStatus 项目操作状态，可选值：DISABLE、ENABLE
	OptStatus enum.OptStatus `json:"opt_status,omitempty"`
	// ProjectCreateTime 项目创建时间
	ProjectCreateTime string `json:"project_create_time,omitempty"`
	// ProjectModifyTime 项目修改时间
	ProjectModifyTime string `json:"project_modify_time,omitempty"`
	// ProductID 产品ID
	ProductID string `json:"product_id,omitempty"`
	// ProductPlatformID 商品库ID
	ProductPlatformID uint64 `json:"product_platform_id,omitempty"`
	// UniqueProductID 升级版商品库产品ID
	UniqueProductID uint64 `json:"unique_product_id,omitempty"`
	// DeliveryMedium 投放载体，可选值：
	// APP 应用、AWEME 抖音号、BYTE_APP 字节小程序、BYTE_GAME 字节小游戏、ENTERPRISE 企业号落地页、
	// ORANGE 橙子落地页、THIRDPARTY 自研落地页、WECHAT_APP 微信小程序、WECHAT_GAME 微信小游戏
	DeliveryMedium enum.DeliveryMedium `json:"delivery_medium,omitempty"`
	// DownloadURL 下载链接
	DownloadURL string `json:"download_url,omitempty"`
	// SubscribeURL 预约链接
	SubscribeURL string `json:"subscribe_url,omitempty"`
	// AppName 应用名称
	AppName string `json:"app_name,omitempty"`
	// PackageName 应用包名
	PackageName string `json:"package_name,omitempty"`
	// ExternalAction 优化目标
	ExternalAction string `json:"external_action,omitempty"`
	// DeepExternalAction 深度优化目标
	DeepExternalAction string `json:"deep_external_action,omitempty"`
	// DeepBidType 深度优化方式
	DeepBidType string `json:"deep_bid_type,omitempty"`
	// NativeType 投放身份
	NativeType enum.NativeType `json:"native_type,omitempty"`
	// InventoryCatalog 投放版位，可选值：UNIVERSAL_SMART 通投智选
	InventoryCatalog enum.InventoryCatalog `json:"inventory_catalog,omitempty"`
	// AwemeID 抖音号ID
	AwemeID string `json:"aweme_id,omitempty"`
	// ScheduleType 投放时间类型，允许值：
	// SCHEDULE_FROM_NOW 从今天起长期投放
	// SCHEDULE_START_END 设置开始和结束日期
	// SCHEDULE_7_DAYS 7日稳投
	ScheduleType enum.ScheduleType `json:"schedule_type,omitempty"`
	// StartTime 投放起始时间
	StartTime string `json:"start_time,omitempty"`
	// EndTime 投放结束时间
	EndTime string `json:"end_time,omitempty"`
	// ScheduleTime 投放时段，默认全时段投放，格式是48*7位字符串，且都是0或1
	ScheduleTime string `json:"schedule_time,omitempty"`
	// SearchContinueDelivery 续投，允许值：OFF 关闭、ON 开启
	SearchContinueDelivery enum.OnOff `json:"search_continue_delivery,omitempty"`
	// BidType 竞价策略，可选值：CUSTOM 稳定成本、NO_BID 最大转化投放
	BidType enum.BidType `json:"bid_type,omitempty"`
	// BudgetMode 预算类型，可选值：BUDGET_MODE_DAY 日预算、BUDGET_MODE_INFINITE 不限、BUDGET_MODE_TOTAL 总预算
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// Budget 预算
	Budget float64 `json:"budget,omitempty"`
	// Pricing 付费方式，可选值：PRICING_CPC、PRICING_CPM、PRICING_OCPC、PRICING_OCPM
	Pricing enum.PricingType `json:"pricing,omitempty"`
	// Bid 出价
	Bid float64 `json:"bid,omitempty"`
	// CpaBid 目标转化出价
	CpaBid float64 `json:"cpa_bid,omitempty"`
	// DeepCpaBid 深度转化出价
	DeepCpaBid float64 `json:"deep_cpabid,omitempty"`
	// RoiGoal ROI系数
	RoiGoal float64 `json:"roi_goal,omitempty"`
	// FirstRoiGoal 首日roi系数
	FirstRoiGoal float64 `json:"first_roi_goal,omitempty"`
	// DownloadType 下载类型，可选值：DOWNLOAD_URL 下载链接、EXTERNAL_URL 包含下载链接的落地页
	DownloadType enum.DownloadType `json:"download_type,omitempty"`
	// DownloadMode 下载模式（优先从应用商店下载），允许值：APP_STORE_DELIVERY 优先商店下载、DEFAULT 默认下载
	DownloadMode enum.DownloadMode `json:"download_mode,omitempty"`
	// LaunchType 调起方式，允许值：DIRECT_OPEN 直接调起、EXTERNAL_OPEN 落地页调起
	LaunchType enum.LaunchType `json:"launch_type,omitempty"`
	// IsCommentDisable 评论管理，可选值：OFF 不允许评论、ON 允许评论
	IsCommentDisable enum.OnOff `json:"is_comment_disable,omitempty"`
	// BlueFlowKeywordName 蓝海关键词
	BlueFlowKeywordName []string `json:"blue_flow_keyword_name,omitempty"`
	// StarTaskIDList 星广联投任务ID
	StarTaskIDList []uint64 `json:"star_task_id_list,omitempty"`
	// StarAutoDeliverySwitch 星广联投全自动化开关，OFF 关闭、ON 开启
	StarAutoDeliverySwitch enum.OnOff `json:"star_auto_delivery_switch,omitempty"`
	// AigcDynamicCreativeSwitch 是否开启AIGC动态创意，OFF 关闭、ON 开启
	AigcDynamicCreativeSwitch enum.OnOff `json:"aigc_dynamic_creative_switch,omitempty"`
	// TrackURLSetting 监测链接
	TrackURLSetting *StdTrackURLSetting `json:"track_url_setting,omitempty"`
	// Audience 定向类型
	Audience *StdAudience `json:"audience,omitempty"`
	// BrandInfo 品牌信息
	BrandInfo *StdBrandInfo `json:"brand_info,omitempty"`
	// GameAddictionID 关键行为ID
	GameAddictionID string `json:"game_addiction_id,omitempty"`
	// MultiDeliveryMedium 多投放载体，可选值：ORANGE_AND_AWEME 建站落地页和抖音主页
	MultiDeliveryMedium enum.MultiDeliveryMedium `json:"multi_delivery_medium,omitempty"`
	// InstanceID 字节小游戏/小程序、微信小游戏/小程序 资产id
	InstanceID uint64 `json:"instance_id,omitempty"`
	// ShopPlatform 电商平台
	ShopPlatform enum.ShopPlatform `json:"shop_platform,omitempty"`
	// DeliveryType 投放类型，可选值：NORMAL 常规投放、UBX_INTELLIGENT 智能托管
	DeliveryType enum.DeliveryType `json:"delivery_type,omitempty"`
	// LiveDuration 直播每日投放时长
	LiveDuration float64 `json:"live_duration,omitempty"`
	// SevenRoiGoal 7日Roi
	SevenRoiGoal float64 `json:"seven_roi_goal,omitempty"`
	// LayerRoiSwitch ROI分层出价开关，ON、OFF
	LayerRoiSwitch enum.OnOff `json:"layer_roi_switch,omitempty"`
	// LayerRoiGoal ROI分层出价
	LayerRoiGoal []LayerRoiGoal `json:"layer_roi_goal,omitempty"`
	// LandingPageStayTime 店铺停留时长，单位为毫秒
	LandingPageStayTime int `json:"landing_page_stay_time,omitempty"`
}

// LayerRoiGoal ROI分层出价
type LayerRoiGoal struct {
	// Lower 下限区间
	Lower int `json:"lower,omitempty"`
	// Upper 上限区间
	Upper int `json:"upper,omitempty"`
	// RoiGoal roi系数
	RoiGoal float64 `json:"roi_goal,omitempty"`
}

// StdTrackURLSetting 监测链接
type StdTrackURLSetting struct {
	// TrackURLType 监测链接类型，可选值：CUSTOM 自定义链接、GROUP_ID 监测链接组
	TrackURLType string `json:"track_url_type,omitempty"`
	// TrackURLGroupID 监测链接组id，通过【事件资产下创建监测链接组】获得
	TrackURLGroupID uint64 `json:"track_url_group_id,omitempty"`
	// TrackURL 展示（监测链接）
	TrackURL []string `json:"track_url,omitempty"`
	// ActionTrackURL 点击（监测链接）
	ActionTrackURL []string `json:"action_track_url,omitempty"`
	// ActiveTrackURL 激活检测链接
	ActiveTrackURL []string `json:"active_track_url,omitempty"`
	// VideoPlayEffectiveTrackURL 视频有效播放（监测链接）
	VideoPlayEffectiveTrackURL []string `json:"video_play_effective_track_url,omitempty"`
	// VideoPlayDoneTrackURL 视频播完（监测链接）
	VideoPlayDoneTrackURL []string `json:"video_play_done_track_url,omitempty"`
	// VideoPlayFirstTrackURL 视频开始播放（监测链接）
	VideoPlayFirstTrackURL []string `json:"video_play_first_track_url,omitempty"`
	// SendType 数据发送方式，允许值：SERVER_SEND 服务器端上传、CLIENT_SEND 客户端上传
	SendType enum.SendType `json:"send_type,omitempty"`
}

// StdAudience 定向类型
type StdAudience struct {
	// AudienceType 定向设置方式，允许值：UNLIMITED 不限、CUSTOM 自定义
	AudienceType string `json:"audience_type,omitempty"`
	// District 地域类型，允许值：BUSINESS_DISTRICT 商圈、REGION 行政区域、OVERSEA 海外区域、NONE 不限
	District enum.District `json:"district,omitempty"`
	// Geolocation 从地图添加(地图位置)，district为"BUSINESS_DISTRICT"时填写，最多允许添加1000个
	Geolocation []model.Geolocation `json:"geolocation,omitempty"`
	// RegionVersion 行政区域版本号
	RegionVersion string `json:"region_version,omitempty"`
	// City 地域定向省市或者区县列表
	City []uint64 `json:"city,omitempty"`
	// LocationType 位置类型，允许值：CURRENT、HOME、TRAVEL、ALL
	LocationType enum.LocationType `json:"location_type,omitempty"`
	// Gender 性别，允许值：GENDER_FEMALE、GENDER_MALE、NONE
	Gender enum.AudienceGender `json:"gender,omitempty"`
	// Age 年龄区间，允许值：AGE_BETWEEN_18_23 等
	Age []enum.AudienceAge `json:"age,omitempty"`
	// RetargetingTagsInclude 定向人群包列表（自定义人群）
	RetargetingTagsInclude []uint64 `json:"retargeting_tags_include,omitempty"`
	// RetargetingTagsExclude 排除人群包列表（自定义人群）
	RetargetingTagsExclude []uint64 `json:"retargeting_tags_exclude,omitempty"`
	// HideIfExists 过滤已安装，允许值：UNLIMITED、FILTER、TARGETING
	HideIfExists string `json:"hide_if_exists,omitempty"`
	// HideIfConverted 过滤已转化用户，允许值：NO_EXCLUDE、PROJECT、ADVERTISER、APP、CUSTOMER、ORGANIZATION、GLOBAL_APP
	HideIfConverted string `json:"hide_if_converted,omitempty"`
	// ConvertedTimeDuration 过滤时间范围
	ConvertedTimeDuration enum.ConvertedTimeDuration `json:"converted_time_duration,omitempty"`
	// Platform 投放平台列表，允许值：ANDROID、IOS、HARMONY
	Platform []enum.AudiencePlatform `json:"platform,omitempty"`
	// AndroidOsv 最低安卓版本
	AndroidOsv string `json:"android_osv,omitempty"`
	// IosOsv 最低IOS版本
	IosOsv string `json:"ios_osv,omitempty"`
	// HarmonyOsv 鸿蒙版本
	HarmonyOsv string `json:"harmony_osv,omitempty"`
	// FilterOwnAwemeFans 过滤自己的粉丝，可选值：OFF 不过滤、ON 过滤
	FilterOwnAwemeFans enum.OnOff `json:"filter_own_aweme_fans,omitempty"`
	// Ac 网络，可选值：2G、3G、4G、5G、WIFI
	Ac []string `json:"ac,omitempty"`
	// LaunchPrice 手机价格
	LaunchPrice []int `json:"launch_price,omitempty"`
	// SmartExtend 智能拓展开关，可选值：OFF 不启用、ON 启用
	SmartExtend enum.OnOff `json:"smart_extend,omitempty"`
	// InterestActionMode 行为兴趣，可选值：CUSTOM、RECOMMEND、UNLIMITED
	InterestActionMode enum.InterestActionMode `json:"interest_action_mode,omitempty"`
	// InterestCategories 兴趣类目词
	InterestCategories []uint64 `json:"interest_categories,omitempty"`
	// InterestWords 兴趣关键词
	InterestWords []uint64 `json:"interest_words,omitempty"`
	// ActionDays 用户发生行为天数
	ActionDays int `json:"action_days,omitempty"`
	// ActionCategories 行为类目词
	ActionCategories []uint64 `json:"action_categories,omitempty"`
	// ActionWords 行为关键词
	ActionWords []uint64 `json:"action_words,omitempty"`
}

// StdBrandInfo 品牌信息
type StdBrandInfo struct {
	// YuntuCategoryID 所属类别（品牌分类ID）
	YuntuCategoryID uint64 `json:"yuntu_category_id,omitempty"`
	// CdpBrandID 品牌id
	CdpBrandID uint64 `json:"cdp_brand_id,omitempty"`
	// BrandNameID 品牌名称
	BrandNameID uint64 `json:"brand_name_id,omitempty"`
	// CdpBrandName 产品系列ID
	CdpBrandName string `json:"cdp_brand_name,omitempty"`
	// SubBrandNames 产品系列名称
	SubBrandNames []string `json:"sub_brand_names,omitempty"`
}
