package ad

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

type IAd interface {
	GetID() uint64
	GetName() string
	GetCampaignID() uint64
	GetAdvertiserID() uint64
	GetOptStatus() enum.AdOptStatus
	GetBudget() float64
	GetCpaBid() float64
	GetDeepCpaBid() float64
	GetExternalURLs() []string
	Version() model.AdVersion
}

type WithOpenURL interface {
	GetID() uint64
	GetActionTrackURL() []string
	GetOpenURL() string
	IsProject() bool
}

// Ad 广告信息
type Ad struct {
	// ID 计划ID
	ID uint64 `json:"id,omitempty"`
	// AdID 计划ID,返回值同id
	AdID uint64 `json:"ad_id,omitempty"`
	// Name 计划名称
	Name string `json:"name,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告组ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// ModifyTime 计划上次修改时间标识(用于更新计划时提交,服务端判断是否基于最新信息修改)
	ModifyTime string `json:"modify_time,omitempty"`
	// AdModifyTime 计划上次修改时间
	AdModifyTime string `json:"ad_modify_time,omitempty"`
	// AdCreateTime 计划创建时间
	AdCreateTime string `json:"ad_create_time,omitempty"`
	// UniqueFk 第三方唯一键，传该值时保证接口重试的幂等性，带有相同unique_fk的请求服务端会视为同一个广告处理。仅在创建接口传入且无法修改，如果创建时传入了已存在的唯一键值，那么会返回该唯一键值所对应的广告计划ID。该值可用于内部系统会生成的唯一ID与头条ID做关联的场景，避免超时重试实际上一次创建请求又成功导致的重复创建问题，通过unique_fk可与内部系统ID实现关联并避免重复创建，可结合实际场景选择使用，广告计划中的unique_fk要求不重复，与广告组中的unique_fk无相关。
	UniqueFk string `json:"unique_fk,omitempty"`
	// Status 广告计划投放状态; (进入投放之前,优先披露审核状态,此时优先于启用暂停,启用暂停信息以opt_status为准)
	Status enum.AdStatus `json:"status,omitempty"`
	// LearningPhase 学习期状态; 许值：DEFAULT（默认，不在学习期中）、LEARNING（学习中）、LEARNED（学习成功）、LEARN_FAILED（学习失败）;关于学习期，此字段即将废弃关于学习期，此字段即将废弃
	LearningPhase enum.LearningPhase `json:"learning_phase,omitempty"`
	// OptStatus 广告计划操作状态, 允许值: "AD_STATUS_ENABLE","AD_STATUS_DISABLE"
	OptStatus enum.AdOptStatus `json:"opt_status,omitempty"`
	// DeliveryRange 投放范围
	DeliveryRange enum.AdDeliveryRange `json:"delivery_range,omitempty"`
	// UnionVideoType 投放形式（穿山甲视频创意类型）;默认值: ORIGINAL_VIDEO原生
	UnionVideoType enum.UnionVideoType `json:"union_video_type,omitempty"`
	// DownloadUrl 应用下载方式，推广目的为APP时有值。返回值：DOWNLOAD_URL下载链接，QUICK_APP_URL快应用+下载链接，EXTERNAL_URL落地页链接
	DownloadUrl string `json:"download_url,omitempty"`
	// QuickAppUrl 快应用链接，当推广类型为应用推广，且download_type为QUICK_APP_URL时有值
	QuickAppUrl string `json:"quick_app_url,omitempty"`
	// ExternalUrl 落地页链接，投放内容或下载方式为落地页时有值
	ExternalUrl string `json:"external_url,omitempty"`
	// Ulink 直达备用链接，仅支持穿山甲广告位（不支持搜索广告）
	Ulink string `json:"ulink,omitempty"`
	// TrackUrl 展示（监测链接）
	TrackUrl []string `json:"track_url,omitempty"`
	// ActionTrackUrl 点击（监测链接）
	ActionTrackUrl []string `json:"action_track_url,omitempty"`
	// VideoPlayEffectTrackUrl 视频有效播放（监测链接）
	VideoPlayEffectTrackUrl []string `json:"video_play_effect_track_url,omitempty"`
	// VideoPlayDoneTrackUrl 视频播完（监测链接）
	VideoPlayDoneTrackUrl []string `json:"video_play_done_track_url,omitempty"`
	// VideoPlayTrackUrl 视频播放（监测链接）
	VideoPlayTrackUrl []string `json:"video_play_track_url,omitempty"`
	// TrackUrlSendType 数据发送方式; 允许值: SERVER_SEND(服务器端上传), CLIENT_SEND(客户端上传)
	TrackUrlSendType string `json:"track_url_send_type,omitempty"`
	// DownloadType 应用下载方式，推广目的为APP时有值。返回值：DOWNLOAD_URL下载链接，QUICK_APP_URL快应用+下载链接，EXTERNAL_URL落地页链接
	DownloadType enum.DownloadType `json:"download_type,omitempty"`
	// AppType 下载类型，当推广类型为应用推广且download_type为DOWNLOAD_URL或者QUICK_APP_URL时或当推广类型为DPA(商品目录推广)且dpa_adtype为DPA_APP有值
	AppType string `json:"app_type,omitempty"`
	// DonwloadMode 优先从系统应用商店下载（下载模式）;允许值：APP_STORE_DELIVERY（仅安卓应用下载支持）、 DEFAULT当应用下载时，默认default下载，可选用APP_STORE_DELIVERY（应用商店直投），当为该值时，将优先跳转目标应用对应手机系统应用商店安装详情页，跳转失败则使用下载链接下载。请确保投放的应用在应用商店内已上架
	DownloadMode enum.DownloadMode `json:"download_mode,omitempty"`
	// ConvertID 转化目标，其中convert_id数值较小时为预定义转化
	ConvertID uint64 `json:"convert_id,omitempty"`
	// ExternalAction 转化类型，目前当推广类型为抖音时有值，允许值："AD_CONVERT_TYPE_FOLLOW_ACTION", "AD_CONVERT_TYPE_MESSAGE_ACTION", "AD_CONVERT_TYPE_INTERACTION"
	ExternalAction enum.AdConvertType `json:"external_action,omitempty"`
	// ExternalActions 转化类型，目前当推广类型为抖音时有值，允许值："AD_CONVERT_TYPE_FOLLOW_ACTION", "AD_CONVERT_TYPE_MESSAGE_ACTION", "AD_CONVERT_TYPE_INTERACTION"
	ExternalActions []enum.AdConvertType `json:"external_actions,omitempty"`
	// DeepExternalAction 深度转化目标
	DeepExternalAction enum.DeepExternalAction `json:"deep_external_action,omitempty"`
	// FeedDeliverySearch 搜索快投关键词，HAS_OPEN:启用，DISABLE:不启用
	FeedDeliverySearch string `json:"feed_delivery_search,omitempty"`
	// IntelligentFlowSwitch 智能流量开关，ON:开启，OFF:关闭
	IntelligentFlowSwitch enum.OnOff `json:"intelligent_flow_switch,omitempty"`
	// OpenUrl 直达链接(点击唤起APP)
	OpenUrl string `json:"open_url,omitempty"`
	// AdvancedCreativeType 附加创意类型; 允许值: ATTACHED_CREATIVE_GAME_PACKAGE游戏礼包码,ATTACHED_CREATIVE_GAME_FORM游戏表单收集,ATTACHED_CREATIVE_GAME_SUBSCRIBE游戏预约,ATTACHED_CREATIVE_NONE无推广目的为应用推广类型、下载方式选择下载链接且下载链接为安卓应用下载时才可以设置
	AdvancedCreativeType enum.AdvancedCreativeType `json:"advanced_creative_type,omitempty"`
	// GamePackageDesc 应用描述,最少1字，最多15字
	GamePackageDesc string `json:"game_package_desc,omitempty"`
	// GamePackageBatchID  游戏礼包码id，目前仅支持直接发券类型
	GamePackageBatchID int64 `json:"game_package_batch_id,omitempty"`
	// GamePackageThumbnails  应用图片集，图片URL链接，有且仅有两个该字段10.10日下线
	GamePackageThumbnails []string `json:"game_package_thumbnails,omitempty"`
	// GamePackageThumbnailIDs  应用图片集，图片image_id，有且仅有两个，游戏礼包码时有值,可以从 【素材管理-获取图片素材】 接口中获取，建议尺寸 16:9 与game_package_thumbnails同时传入时，以game_package_thumbnail_ids字段为准
	GamePackageThumbnailIDs []string `json:"game_package_thumbnail_ids,omitempty"`
	// StoreproUnit 门店推广-投放内容，当推广目的为STORE(门店推广)时有值。取值: "STORE"门店, "STORE_ACTIVITY"活动;目前暂时不支持线下商品类型
	StoreproUnit string `json:"storepro_unit,omitempty"`
	// StoreType  门店类型，（storepro_unit 为 "STORE" 时有值。取值: "STORE_NORMAL"平台通用门店, "STORE_THIRT_PARTY"第三方门店, "STORE_DOUYIN"抖音POI门店
	StoreType string `json:"store_type,omitempty"`
	// AdvertiserStoreIDs 门店ID列表 （storepro_unit 为 "STORE" 时有值
	AdvertiserStoreIDs []uint64 `json:"advertiser_store_ids,omitempty"`
	// StoreproPackID 活动ID （storepro_unit 为 "STORE_ACTIVITY" 时有值
	StoreproPackID uint64 `json:"storepro_pack_id,omitempty"`
	// ProductPlatformID  产品目录ID(ID由查询产品目录接口得到), 当推广目的landing_type为DPA时有值
	ProductPlatformID uint64 `json:"product_platform_id,omitempty"`
	// ProductID 商品id，当推广目的为 DPA 广告组商品类型为 SDPA 时有值
	ProductID string `json:"product_id,omitempty"`
	// AssetID 物件id，当广告组商品类型为 SDPA 且商品库为汽车商品库时有值
	AssetID uint64 `json:"asset_id,omitempty"`
	// CategoryType DPA投放范围，取值：NONE不限，"CATEGORY"选择分类，"PRODUCT"指定商品
	CategoryType string `json:"category_type,omitempty"`
	// DpaCategories 分类列表，category_type取值范围为CATEGORY时有值
	DpaCategories []uint64 `json:"dpa_categories,omitempty"`
	// DpaProducts 商品列表，category_type为PRODUCT时有值
	DpaProducts []uint64 `json:"dpa_products,omitempty"`
	// DpaProductTarget 自定义筛选条件（商品投放条件）。用于圈定商品投放范围，结合商品库字段搭配判断条件，圈定商品投放范围。
	DpaProductTarget []DpaProductTarget `json:"dpa_product_target,omitempty"`
	// DpaAdtype dpa广告类型，取值范围："DPA_LINK"落地页, "DPA_APP"应用下载
	DpaAdtype string `json:"dpa_adtype,omitempty"`
	// ParamsType 链接类型(落地页)，当dpa_adtype为"DPA_LINK"时有值，取值: "DPA"商品库所含链接, "CUSTOM"自定义链接
	ParamsType string `json:"params_type,omitempty"`
	// DpaExternalUrlField 落地页链接字段选择，当params_type为"DPA"时有值
	DpaExternalUrlField string `json:"dpa_external_url_field,omitempty"`
	// DpaExternalUrls 落地页链接地址列表，当params_type为"CUSTOM"时有值
	DpaExternalUrls []string `json:"dpa_external_urls,omitempty"`
	// Package 应用包名，当推广类型为应用推广且download_type为DOWNLOAD_URL时或当推广类型为DPA(商品目录推广)且dpa_adtype为DPA_APP有值
	Package string `json:"package,omitempty"`
	// InventoryCatalog 广告位大类。 MANUAL首选媒体 SCENE场景广告位，SMART优选广告位，UNIVERSAL通投智选
	InventoryCatalog enum.InventoryCatalog `json:"inventory_catalog,omitempty"`
	// InventoryType 创意投放位置; 创建选择优选广告位时，此字段回会返回对应的优选广告位
	InventoryType []enum.StatInventoryType `json:"inventory_type,omitempty"`
	// SmartInventory 优选广告位，NORMAL表示不使用优选，SMART表示使用优选，UNIVERSAL表示通投
	SmartInventory string `json:"smart_inventory,omitempty"`
	// SceneInventory 首选场景广告位
	SceneInventory string `json:"scene_inventory,omitempty"`
	// PromotionType 投放内容; GOODS：商品推广;LIVE：直播; AWEME_HOME_PAGE：抖音主页;LANDING_PAGE_LINK：落地页
	PromotionType string `json:"promotion_type,omitempty"`
	// AwemeAccount 抖音号
	AwemeAccount string `json:"aweme_account,omitempty"`
	// SubscribeUrl 游戏营销场景-预约下载链接
	SubscribeUrl string `json:"subscribe_url,omitempty"`
	// FormID 游戏营销场景-表单id
	FormID uint64 `json:"form_id,omitempty"`
	// FormIndex 游戏营销场景-表单位置索引
	FormIndex int `json:"form_index,omitempty"`
	// AppDesc 游戏营销场景-应用描述
	AppDesc string `json:"app_desc,omitempty"`
	// AppIntroduction 游戏营销场景-应用介绍
	AppIntroduction string `json:"app_introduction,omitempty"`
	// AppThumbnails 游戏营销场景-应用图片集，返回图片集Id
	AppThumbnails []string `json:"app_thumbnails,omitempty"`
	// DpaOpenUrlType 直达链接类型，取值: "NONE"不启用, "DPA"商品库所含链接, "CUSTOM"自定义链接商品库链接对应商品库内调起字段。
	DpaOpenUrlType string `json:"dpa_open_url_type,omitempty"`
	// DpaOpenUrlField 直达链接字段选择，当dpa_open_url_type为"DPA"时有值
	DpaOpenUrlField string `json:"dpa_open_url_field,omitempty"`
	// DpaOpenUrls 直达链接地址列表，当dpa_open_url_type为"CUSTOM"时有值
	DpaOpenUrls []string `json:"dpa_open_urls,omitempty"`
	// ExternalUrlParams 落地页检测参数(DPA推广目的特有,在填写的参数后面添加"=urlencode(开放平台提供的h5链接地址）"，其中urlencode(开放平台提供的h5链接地址）替换为商品库中的h5地址encode的结果)
	ExternalUrlParams string `json:"external_url_params,omitempty"`
	// OpenUrlParams 直达链接检测参数(DPA推广目的特有,在“产品库中提取的scheme地址"后面追加填写的参数)
	OpenUrlParams string `json:"open_url_params,omitempty"`
	// Audience 广告受众, 字典类型, 包含下面字段
	Audience *Audience `json:"audience,omitempty"`
	// AudiencePackageID 定向包ID
	AudiencePackageID uint64 `json:"audience_package_id,omitempty"`
	// HideIfExists 过滤已安装，当推广目标为安卓应用下载时可填，0表示不限，1表示过滤，2表示定向。默认为不限;默认值:0; 取值: 0, 1, 2。建议促进app活跃度客户使用定向安装功能。选择定向安装时，向已安装客户展示广告；投放时优先获取直达链接，无直达链接时使用应用包名进行投放；如果无直达链接或应用包名，定向安装选项实际不生效；仅对Android链接生效。
	HideIfExists enum.HideIfExists `json:"hide_if_exists,omitempty"`
	// HideIfConverted 过滤已转化用户
	HideIfConverted enum.HideIfConverted `json:"hide_if_converted,omitempty"`
	// ConvertedTimeDuration 过滤时间范围
	ConvertedTimeDuration enum.ConvertedTimeDuration `json:"converted_time_duration,omitempty"`
	// DpaLbs 地域匹配-LBS;开启时，根据用户的地理位置信息，给用户投放位于其附近的产品 ;取值：0，1（0表示不启用，1表示启用）
	DpaLbs *int `json:"dpa_lbs,omitempty"`
	// DpaCity 地域匹配-商品所在城市;开启时，仅将商品投放给位于该商品设置的可投城市的用户 ;取值：0，1（0表示不启用，1表示启用）
	DpaCity *int `json:"dpa_city,omitempty"`
	// DpaProvince 地域匹配-商品所在省份;开启时，将商品仅投放给位于该商品设置的可投省份的用户 ;取值：0，1（0表示不启用，1表示启用）
	DpaProvince *int `json:"dpa_province,omitempty"`
	// DpaLocationAudience DPA行为重定向，0:不启用，1：启用
	DpaLocationAudience *int `json:"dap_location_audience,omitempty"`
	// AssetIDs 事件管理下资产 id
	AssetIDs []uint64 `json:"asset_ids,omitempty"`
	// ValueOptimizedType 目标优化类型，0表示行为优化，1表示价值优化
	ValueOptimizedType *int `json:"value_optimized_type,omitempty"`
	// ValueOptimizedOpen 价值优选，0表示关闭，1表示开启
	ValueOptimizedOpen *int `json:"value_optimized_open,omitempty"`
	// IncludeCustomActions 包含人群包((DPA推广目的特有,格式举例[{"days": 7, "code": 1001},]， dpa_local_audience为1时有值; day可选范围:1, 7, 14, 28, 60, 90, 120, 180。
	IncludeCustomActions json.RawMessage `json:"include_custom_actions,omitempty"`
	// ExcludeCustomActions 排除人群包((DPA推广目的特有,格式举例{"days": 7, "code": 1002},]，day可选范围: 1, 7, 14, 28, 60, 90, 120, 180。
	ExcludeCustomActions json.RawMessage `json:"exclude_custom_actions,omitempty"`
	// DpaRecommendType dpa商品重定向推荐类型，dpa_local_audience为1时有值;取值：1（基于重定向推荐更多商品(根据重定向商品和行业特点，推荐更多相关商品投放，包含重定向商品），2仅重定向商品（仅根据重定向人群内定义的重定向行为商品进行投放）
	DpaRecommendType int `json:"dpa_recommend_type,omitempty"`
	// SmartBidType 投放场景(出价方式)
	SmartBidType enum.SmartBidType `json:"smart_bid_type,omitempty"`
	// AdjustCpa 是否调整自动出价，意味如果预期成本不在范围内将在此基础上调整，仅OCPM支持，当smart_bid_type=SMART_BID_CONSERVATIVE时选填。当smart_bid_type为"SMART_BID_CONSERVATIVE"且adjust_cpa=0时，cpa_bid由系统自动计算；;当smart_bid_type为"SMART_BID_CONSERVATIVE" 且adjust_cpa=1时，cpa_bid必填
	AdjustCpa int `json:"adjust_cpa,omitempty"`
	// FlowControlMode 竞价策略(投放方式)
	FlowControlMode enum.FlowControlMode `json:"flow_control_mode,omitempty"`
	// BudgetMode 预算类型(创建后不可修改),
	BudgetMode enum.BudgetMode `json:"budget_mode,omitempty"`
	// Budget 预算
	Budget float64 `json:"budget,omitempty"`
	// ScheduleType 投放时间类型
	ScheduleType enum.ScheduleType `json:"schedule_type,omitempty"`
	// StartTime 投放起始时间，当schedule_type为"SCHEDULE_START_END"时取值，形式如：2017-01-01 00:00
	StartTime string `json:"start_time,omitempty"`
	// EndTime 投放结束时间，当schedule_type为"SCHEDULE_START_END"时取值，形式如：2017-01-01 00:00
	EndTime string `json:"end_time,omitempty"`
	// ScheduleTime 投放时段，默认全时段投放，格式是48*7位字符串，且都是0或1。也就是以半个小时为最小粒度，周一至周日每天分为48个区段，0为不投放，1为投放，不传、全传0、全传1均代表全时段投放。
	ScheduleTime string `json:"schedule_time,omitempty"`
	// Pricing 付费方式（计划出价类型）
	Pricing enum.PricingType `json:"pricing,omitempty"`
	// Bid 出价
	Bid float64 `json:"bid,omitempty"`
	// CpaBid 目标转化出价/预期成本， 当pricing为"OCPM"、"OCPC"出价方式时有值
	CpaBid float64 `json:"cpa_bid,omitempty"`
	// DeepBidType 深度优化方式
	DeepBidType enum.DeepBidType `json:"deep_bid_type,omitempty"`
	// DeepCpabidid 深度优化出价，deep_bid_type为"DEEP_BID_MIN"时有值。当对应的转化convert_id，设定深度转化目标时才会有效。
	DeepCpabidid float64 `json:"deep_cpabidid,omitempty"`
	// LubanROiGoal 鲁班目标ROI出价策略系数。推广目的为商品推广(GOODS)时可填。当传入该参数时，表示启用鲁班ROI优化，支持范围(0,100]，精度：保留小数点后四位
	LubanRoiGoal float64 `json:"luban_roi_goal,omitempty"`
	// RoiGoal 深度转化ROI系数, 范围(0,5]，精度：保留小数点后四位, deep_bid_type为"ROI_COEFFICIENT"时有值
	RoiGoal float64 `json:"roi_goal,omitempty"`
	// AutoInheritSwitch 一键继承开关，ON表示开启一键继承，OFF表示关闭一键继承
	AutoInheritSwitch string `json:"auto_inherit_switch,omitempty"`
	// InheritType 一键继承账户类型，auto_inherit_switch为ON时有意义，INHERIT_FROM_ACCOUNT表示从同账户下的优质计划中继承，INHERIT_FROM_CUSTOMER表示从同账户所在组织下的其他账户的优质计划中继承
	InheritType string `json:"inherit_type,omitempty"`
	// InheritedAdvertiserID 一键继承的同组织账户id的list，inherit_type等于INHERIT_FROM_CUSTOMER时有意义
	InheritedAdvertiserID []uint64 `json:"inherited_advertiser_id,omitempty"`
	// DeliveryPhase 计划所处阶段，允许值：FIRST_PHASE第一阶段，SECOND_PHASE第二阶段。当pricing为PRICING_CPC_OCPM时有值
	DeliveryPhase string `json:"delivery_phase,omitempty"`
	// LauhchTargetType 投放类型，LIVE_CONVERT：直播间转化、APP：应用下载、EXTERNAL：线索收集
	LaunchTargetType string `json:"launch_target_type,omitempty"`
	// AutoUpdateKeyword 是否开启自动加词，ON 开启、OFF 关闭
	AutoUpdateKeyword enum.OnOff `json:"auto_update_keyword,omitempty"`
	// LandingPageStayTime 店铺停留时长，单位为毫秒
	LandingPageStayTime int64 `json:"landing_page_stay_time,omitempty"`
	// TargetCvr 目标转化率
	TargetCvr float64 `json:"target_cvr,omitempty"`
}

// DpaProductTarget 商品投放条件
type DpaProductTarget struct {
	// Title 筛选字段
	Title string `json:"title,omitempty"`
	// Rule 定向规则，允许值：'=', '!=', '>', '<', '>=', '<=', 'contain', 'exclude', 'notEmpty
	Rule string `json:"rule,omitempty"`
	// Type 字段类型，允许值：'int', 'double', 'long', 'string'
	Type string `json:"type,omitempty"`
	// Value 规则值
	Value string `json:"value,omitempty"`
}

// Audience 受众
type Audience struct {
	// District 地域;取值: "CITY"省市, "COUNTY"区县, "BUSINESS_DISTRICT"商圈,"NONE"不限，省市传法："city": [12],"district": "CITY",区县的传法："city": [130102],"district": "COUNTY";暂不支持"海外"
	District enum.District `json:"district,omitempty"`
	// RegionVersion 行政区域版本号。通过[【获取行政信息】]https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1709606596424718)接口获取; district =REGION/OVERSEA时必填
	RegionVersion string `json:"region_version,omitempty"`
	// City 地域定向省市或者区县列表(当传递省份ID时,旗下市县ID可省略不传),当district为"CITY"或"COUNTY"时有值
	City []uint64 `json:"city,omitempty"`
	// BusinessIDs 商圈ID数组，district为"BUSINESS_DISTRICT"时有值
	BusinessIDs []uint64 `json:"business_ids,omitempty"`
	// Geolocation 从地图添加(地图位置)
	Geolocation []model.Geolocation `json:"geolocation,omitempty"`
	// LocationType 位置类型;取值：CURRENT正在该地区的用户，HOME居住在该地区的用户，TRAVEL到该地区旅行的用户，ALL该地区内的所有用户;当city和district有值时返回值
	LocationType enum.LocationType `json:"location_type,omitempty"`
	// Gender 性别
	Gender enum.AudienceGender `json:"gender,omitempty"`
	// Age 年龄
	Age []enum.AudienceAge `json:"age,omitempty"`
	// RetargetingTagsInclude 定向人群包列表（自定义人群），内容为人群包id。如果选择"同时定向与排除"，需传入retargeting_tags_include和retargeting_tags_exclude
	RetargetingTagsInclude []uint64 `json:"retargeting_tags_include,omitempty"`
	// RetargetingTagsExclude 排除人群包列表（自定义人群），内容为人群包id。如果选择"同时定向与排除"，需传入retargeting_tags_include和retargeting_tags_exclude
	RetargetingTagsExclude []uint64 `json:"retargeting_tags_exclude,omitempty"`
	// InterestActionMode 行为兴趣;取值："UNLIMITED"不限,"CUSTOM"自定义,"RECOMMEND"系统推荐。若与自定义人群同时使用，系统推荐("RECOMMEND")不生效;仅推广范围为默认时可填，且不可与老版行为兴趣定向同时填写，否则会报错
	InterestActionMode enum.InterestActionMode `json:"interest_action_mode,omitempty"`
	// Action 行为内容
	Action *AudienceAction `json:"action,omitempty"`
	// InterestCategories 兴趣类目词，当interest_action_mode传CUSTOM时有效
	InterestCategories []uint64 `json:"interest_categories,omitempty"`
	// InterestWords 兴趣关键词, 传入具体的词id，非兴趣词包id，可以通过词包相关接口或者兴趣关键词word2id接口获取词id，一个计划下最多创建1000个关键词。当interest_action_mode传CUSTOM时有效
	InterestWords []uint64 `json:"interest_words,omitempty"`
	// AdTags （老版行为兴趣）兴趣分类,如果传"空数组"表示不限，如果"数组传0"表示系统推荐,如果按兴趣类型传表示自定义
	AdTags []uint64 `json:"ad_tag,omitempty"`
	// InterestTags （老版行为兴趣）兴趣关键词, 传入具体的词id，非兴趣词包id，可以通过词包相关接口或者兴趣关键词word2id接口获取词id，一个计划下最多创建1000个关键词。
	InterestTags []uint64 `json:"interest_tags,omitempty"`
	// AppBehaviorTarget （老版行为兴趣）APP行为; 取值：NONE不限，CATEGORY按分类，APP按APP
	AppBehaviorTarget string `json:"app_behavior_target,omitempty"`
	// AppCategory 老版行为兴趣）APP行为定向——按分类
	AppCategory []uint64 `json:"app_category,omitempty"`
	// AppIDs （老版行为兴趣）APP行为定向——按APP（请注意如果投放的是"应用下载-IOS"不支持设置APP行为定向，请勿传值。）可通过【工具-查询工具-查询应用信息】获取。当app_behavior_target为APP时有值
	AppIDs []uint64 `json:"app_ids,omitempty"`
	// AwemeFanBehavior 抖音达人互动用户行为类型
	AwemeFanBehaviors []enum.Behavior `json:"aweme_fan_behaviors,omitempty"`
	// AwemeFanTimeScope
	AwemeFanTimeScope string `json:"aweme_fan_time_scope,omitempty"`
	// AwemeFanCategories 抖音达人分类ID列表，与aweme_fan_behaviors同时设置才会生效（抖音达人定向）
	AwemeFanCategories []uint64 `json:"aweme_fan_categories,omitempty"`
	// AwemeFanAccounts 抖音达人ID列表，与aweme_fan_behaviors同时设置才会生效（抖音达人定向）
	AwemeFanAccounts []uint64 `json:"aweme_fan_accounts,omitempty"`
	// AwemeFansNumbers （抖音号推广特有）账号粉丝相似人群（添加抖音账号，会将广告投放给对应账号的相似人群粉丝）
	AwemeFansNumbers []int64 `json:"aweme_fans_numbers,omitempty"`
	// FilterAwemeAbnormalActive （抖音号推广特有）过滤高活跃用户; 取值：0表示不过滤，1表示过滤
	FilterAwemeAbnormalActive int `json:"filter_aweme_abnormal_active,omitempty"`
	// FilterAwemeFansCount （抖音号推广特有）过滤高关注数用户，例如"filter_aweme_fans_count": 1000表示过滤粉丝数在1000以上的用户
	FilterAwemeFansCount int64 `json:"filter_aweme_fans_count,omitempty"`
	// FilterOwnAwemeFans （抖音号推广特有）过滤自己的粉丝; 取值：0表示不过滤，1表示过滤
	FilterOwnAwemeFans int `json:"filter_own_aweme_fans,omitempty"`
	// SuperiorPopularityType 媒体定向;
	SuperiorPopularityType enum.SuperiorPopularityType `json:"superior_popularity_type,omitempty"`
	// FlowPackage 定向逻辑
	FlowPackage []uint64 `json:"flow_package,omitempty"`
	// ExcludeFlowPackage 排除定向逻辑
	ExcludeFlowPackage []uint64 `json:"exclude_flow_package,omitempty"`
	// Platform 平台，当下载方式包含下载链接时，平台类型需与选择的下载链接类型对应，当下载方式不包含下载方式的时候，平台可多选。为保证投放效果,平台类型定向PC与移动端互斥
	Platform []enum.AudiencePlatform `json:"platform,omitempty"`
	// AndroidOsv 最低安卓版本，当app_type为"APP_ANDROID"选填,其余情况不填,
	AndroidOsv string `json:"android_osv,omitempty"`
	// IosOsv 最低IOS版本，当app_type为"APP_IOS"选填,其余情况不填
	IosOsv string `json:"ios_osv,omitempty"`
	// DeviceType 设备类型;取值是："MOBILE", "PAD"。缺省表示不限设备类型。穿山甲已经全量，投放范围为默认时需要有白名单权限才可以
	DeviceType []string `json:"device_type,omitempty"`
	// Ac 网络类型
	Ac []string `json:"ac,omitempty"`
	// Carrier 运营商
	Carrier []enum.Carrier `json:"carrier,omitempty"`
	// ActivateType 新用户(新用户使用头条的时间)
	ActivateType []enum.ActivateType `json:"activate_type,omitempty"`
	// ArticleCategory 文章分类
	ArticleCategory []string `json:"article_category,omitempty"`
	// DeviceBrand 手机品牌
	DeviceBrand []string `json:"device_brand,omitempty"`
	// LauchPrice 手机价格,传入价格区间，最高传入11000（表示1w以上）;传值示例 "launch_price": [2000, 11000]，表示2000元以上;
	LaunchPrice []int `json:"launch_price,omitempty"`
	// AutoExtendEnabled 是否启用智能放量。取值是：0、1。缺省为 0。
	AutoExtendEnabled int `json:"auto_extend_enabled,omitempty"`
	// AutoExtendTarget 可放开定向。当auto_extend_enabled=1 时选填。详见：【附录-可开放定向】。缺省为全不选。
	AutoExtendTarget []string `json:"auto_extend_targets,omitempty"`
	// DpaRtaSwitch RTA重定向选项，值：ON:开启，OFF：关闭
	DpaRtaSwitch enum.OnOff `json:"dpa_rta_switch,omitempty"`
	// DpaRtaRecommendType RTA推荐逻辑，ONLY:仅RTA推荐商品，MORE：基于RTA推荐更多商品
	DpaRtaRecommendType string `json:"dpa_rta_recommend_type,omitempty"`
}

func (a Ad) Version() model.AdVersion {
	return model.AdVersion_DEFAULT
}

func (a Ad) GetID() uint64 {
	return a.ID
}

func (a Ad) GetName() string {
	return a.Name
}

func (a Ad) GetCampaignID() uint64 {
	return a.CampaignID
}

func (a Ad) GetAdvertiserID() uint64 {
	return a.AdvertiserID
}

func (a Ad) GetOptStatus() enum.AdOptStatus {
	return a.OptStatus
}

func (a Ad) GetBudget() float64 {
	return a.Budget
}

func (a Ad) GetCpaBid() float64 {
	return a.CpaBid
}

func (a Ad) GetDeepCpaBid() float64 {
	return a.DeepCpabidid
}

func (a Ad) GetOpenURL() string {
	return a.OpenUrl
}

func (a Ad) GetExternalURLs() []string {
	return []string{a.ExternalUrl}
}

func (a Ad) GetActionTrackURL() []string {
	return a.ActionTrackUrl
}

func (a Ad) IsProject() bool {
	return false
}

// AudienceAction 行为内容
type AudienceAction struct {
	// ActionScene 行为场景
	ActionScene []enum.ActionScene `json:"action_scene,omitempty"`
	// ActionDays 用户发生行为天数，当interest_action_mode传CUSTOM时有效
	ActionDays uint `json:"action_days,omitempty"`
	// ActionCategories 行为类目词，当interest_action_mode传CUSTOM时有效
	ActionCategories []uint64 `json:"action_categories,omitempty"`
	// ActionWords 行为关键词，当interest_action_mode传CUSTOM时有效
	ActionWords []uint64 `json:"action_words,omitempty"`
}
