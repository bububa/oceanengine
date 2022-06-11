package ad

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// CreateRequest 创建计划 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告组ID。注意：广告组ID要求属于广告主ID，且是非删除广告组ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// Name 广告计划名称，长度为1-100个字符，其中1个中文字符算2位
	Name string `json:"name,omitempty"`
	// Operation 计划状态; 默认值： "enable"开启状态; 允许值: "enable"开启,"disable"关闭
	Operation string `json:"operation,omitempty"`
	// DeliveryRange 投放范围。详见【附录-广告投放范围】;默认值: "DEFAULT"默认; 允许值: DEFAULT 默认, UNION 穿山甲、UNIVERSAL 通投智选; 对于投放范围的解释可参考【广告计划】
	DeliveryRange enum.AdDeliveryRange `json:"delivery_range,omitempty"`
	// UnionVideoType 投放形式（穿山甲视频创意类型），当delivery_range为"UNION"时必填。详见【附录-穿山甲视频创意类型】;默认值: ORIGINAL_VIDEO原生; 允许值: ORIGINAL_VIDEO原生, REWARDED_VIDEO激励视频,SPLASH_VIDEO开屏
	UnionVideoType enum.UnionVideoType `json:"union_video_type,omitempty"`
	// DownloadType 下载方式; 默认值：DOWNLOAD_URL下载链接; 可选值：DOWNLOAD_URL下载链接、EXTERNAL_URL落地页链接
	DownloadType string `json:"download_type,omitempty"`
	// DownloadUrl 应用下载方式，推广目的为APP时有值。返回值：DOWNLOAD_URL下载链接，QUICK_APP_URL快应用+下载链接，EXTERNAL_URL落地页链接
	DownloadUrl string `json:"download_url,omitempty"`
	// ExternalUrl 落地页链接（支持橙子建站落地页）;对于转化量为目标的计划如OCPM计划不允许更改，非转化为目标的计划如CPC、CPM计划可更改; 获取橙子建站落地页可参考【橙子建站落地页管理】
	ExternalUrl string `json:"external_url,omitempty"`
	// AppType 下载类型，当推广类型为应用推广且download_type为DOWNLOAD_URL或者QUICK_APP_URL时或当推广类型为DPA(商品目录推广)且dpa_adtype为DPA_APP有值
	AppType string `json:"app_type,omitempty"`
	// Package 应用包名，当推广类型为应用推广且download_type为DOWNLOAD_URL时或当推广类型为DPA(商品目录推广)且dpa_adtype为DPA_APP有值
	Package string `json:"package,omitempty"`
	// DownloadMode 优先从系统应用商店下载（下载模式）;允许值：APP_STORE_DELIVERY（仅安卓应用下载支持）、 DEFAULT当应用下载时，默认default下载，可选用APP_STORE_DELIVERY（应用商店直投），选择后，将优先跳转目标应用对应手机系统应用商店安装详情页，跳转失败则使用下载链接下载。;请确保投放的应用在应用商店内已上架;详情请参照【帮助中心】
	DownloadMode string `json:"download_mode,omitempty"`
	// ConvertID 转化目标， 当出价方式为"OCPM"时必填，当出价方式为CPC和CPM时非必填。可通过【工具-转化目标管理-查询计划可用转化id】查询可用id
	ConvertID uint64 `json:"convert_id,omitempty"`
	// OpenUrl 直达链接（点击唤起APP）直达链接仅支持部分App唤起，点击创意将优先跳转App，再根据投放内容跳转相关链接
	OpenUrl string `json:"open_url,omitempty"`
	// Ulink 直达备用链接，仅支持穿山甲广告位（不支持搜索广告）
	Ulink string `json:"ulink,omitempty"`
	// PromotionType 投放内容，允许值：LIVE：直播; AWEME_HOME_PAGE：抖音主页（默认）;LANDING_PAGE_LINK：落地页
	PromotionType string `json:"promotion_type,omitempty"`
	// AwemeAccount 抖音号，可从【获取绑定抖音号】接口获取，默认取绑定的第一个抖音号
	AwemeAccount string `json:"aweme_account,omitempty"`
	// LandingPageStayTime 店铺停留时长，单位为毫秒，当external_action为AD_CONVERT_TYPE_STAY_TIME时必填
	// 允许值：1000、5000、12000、20000、30000、40000、50000、60000
	LandingPageStayTime int `json:"landing_page_stay_time,omitempty"`
	// InventoryCatalog 广告位大类。允许值 MANUAL首选媒体 SCENE场景广告位，SMART优选广告位，UNIVERSAL通投智选
	InventorCatalog string `json:"inventory_catalog,omitempty"`
	// InventoryType 广告投放位置（首选媒体），详见【附录-首选投放位置】，在没有使用smart_inventory的情况下，当前字段必填。
	InventoryType []enum.StatInventoryType `json:"inventory_type,omitempty"`
	// SmartInventory 优选广告位，允许值NORMAL表示不使用优选，SMART表示使用优选，UNIVERSAL表示通投，使用优选广告位的时候默认忽略inventory_type字段。
	// 默认值: NORMAL
	SmartInventory string `json:"smart_inventory,omitempty"`
	// ExternalAction 预定义转化目标
	ExternalAction enum.AdConvertType `json:"external_action,omitempty"`
	// ExternalActions 转化类型列表，仅抖音号、销售线索收集推广目的支持该字段
	// 可通过【工具-转化目标管理-查询计划可用转化目标】查询可用external_action
	ExternalActions []enum.AdConvertType `json:"external_actions,omitempty"`
	// DeepExternalAction 深度转化目标，可通过【获取优化目标】接口获取
	DeepExternalAction enum.DeepExternalAction `json:"deep_external_action,omitempty"`
	// AssetIDs 资产 id，可通过【获取推广内容】接口获取
	// 当使用事件管理时必填
	// 正常情况下数组限制上限为1个；当landing_type=link时上限为2个
	// 若数组长度为2，需保证有一个小程序资产+落地页资产，顺序不强限制
	AssetIDs []uint64 `json:"asset_ids,omitempty"`
	// TrackURLGroupType 监测链接类型（当推广目的为APP时必填）允许值：TRACK_URL_GROUP_ID 已有链接、 CUSTOM ：自定义
	TrackURLGroupType string `json:"track_url_group_type,omitempty"`
	// TrackURLGroupID 监测链接组ID（当推广目的为APP时，且监测链接类型为TRACK_URL_GROUP_ID 时必填）
	TrackURLGroupID uint64 `json:"track_url_group_id,omitempty"`
	// ValueOptimizedType 目标优化类型，0表示行为优化，1表示价值优化
	// 当前仅支持推广目的为销售线索收集（landing_type=LINK）
	// 当选择价值优化时value_optimized_open价值优化必须为开启，会自动设为1，并将 cpa_bid置为0
	ValueOptimizedType *int `json:"value_optimized_type,omitempty"`
	// ValueOptimizedValue 价值优选，0表示关闭，1表示开启
	// 当前仅支持推广目的为销售线索收集（landing_type=LINK）
	// 目标优化类型为价值优化时必须开启，会自动设为1，并将cpa_bid置为 0
	ValueOptmizedOpen *int `json:"value_optimized_value,omitempty"`
	// ProductPlatformID 商品目录ID(ID由【DPA商品广告-查询商品库】 得到)
	ProductPlatformID uint64 `json:"product_platform_id,omitempty"`
	// ProductID 商品ID，当广告组商品类型选择SDPA时必填(ID由【DPA商品广告-获取DPA商品库商品列表】 得到，创建后不可修改)
	ProductID string `json:"product_id,omitempty"`
	// AssetID 物件ID，可用场景需同时满足：
	// 1. 推广目的为应用推广APP、销售线索收集LINK、电商店铺SHOP中的一种，且广告组商品类型为 SDPA
	// 2. 仅在汽车垂直行业下（商品库行业类型对应AUTO_NEW）可以使用且必填，可通过【商品广告-获取投放条件列表】获取，创建后不可修改。
	AssetID uint64 `json:"asset_id,omitempty"`
	// CategoryType 商品目录投放范围,当广告组商品类型选择 DPA 多商品时可修改;允许值：NONE不限，"CATEGORY"选择分类，"PRODUCT"指定商品
	CategoryType string `json:"category_type,omitempty"`
	// DpaCategories 分类列表，category_type更新为"CATEGORY"时可修改，由【DPA商品广告-获取DPA分类】 得到限制个数1~100
	DpaCategories []uint64 `json:"dpa_categories,omitempty"`
	// DpaProducts 商品列表，category_type更新为"PRODUCT"时可修改，由【DPA商品广告-获取DPA商品库商品列表】 得到;限制个数1~100
	DpaProducts []uint64 `json:"dpa_products,omitempty"`
	// DpaProductTargets 自定义筛选条件（商品投放条件）。用于圈定商品投放范围，结合商品库字段搭配判断条件，圈定商品投放范围。
	DpaProductTarget []DpaProductTarget `json:"dpa_product_targets,omitempty"`
	// DpaAdtype DPA广告类型，允许值: "DPA_LINK"落地页, "DPA_APP"应用下载
	DpaAdtype string `json:"dpa_adtype,omitempty"`
	// ParamsType 链接类型(落地页)，当dpa_adtype为"DPA_LINK"时有值，取值: "DPA"商品库所含链接, "CUSTOM"自定义链接
	ParamsType string `json:"params_type,omitempty"`
	// DpaExternalUrlField 落地页链接字段选择，当params_type为"DPA"时有值
	DpaExternalUrlField string `json:"dpa_external_url_field,omitempty"`
	// DpaExternalUrls 落地页链接地址列表，当params_type为"CUSTOM"时有值
	DpaExternalUrls []string `json:"dpa_external_urls,omitempty"`
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
	// FeedDeliverySearch 搜索快投关键词，HAS_OPEN:启用，DISABLE:不启用
	FeedDeliverySearch string `json:"feed_delivery_search,omitempty"`
	// IntelligentFlowSwitch 智能流量开关，ON:开启，OFF:关闭
	IntelligentFlowSwitch string `json:"intelligent_flow_switch,omitempty"`
	// AudiencePackageID 定向包ID
	AudiencePackageID uint64 `json:"audience_package_id,omitempty"`
	// District 地域;取值: "CITY"省市, "COUNTY"区县, "BUSINESS_DISTRICT"商圈,"NONE"不限，省市传法："city": [12],"district": "CITY",区县的传法："city": [130102],"district": "COUNTY";暂不支持"海外"
	District string `json:"district,omitempty"`
	// RegionVersion 行政区域版本号。通过[【获取行政信息】]https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1709606596424718)接口获取; district =REGION/OVERSEA时必填
	RegionVersion string `json:"region_version,omitempty"`
	// City 地域定向省市或者区县列表(当传递省份ID时,旗下市县ID可省略不传),当district为"CITY"或"COUNTY"时有值
	City []uint64 `json:"city,omitempty"`
	// BusinessIDs 商圈ID数组，district为"BUSINESS_DISTRICT"时有值
	BusinessIDs []uint64 `json:"business_ids,omitempty"`
	// Geolocation 从地图添加(地图位置)
	Geolocation []model.Geolocation `json:"geolocation,omitempty"`
	// LocationType 位置类型;取值：CURRENT正在该地区的用户，HOME居住在该地区的用户，TRAVEL到该地区旅行的用户，ALL该地区内的所有用户;当city和district有值时返回值
	LocationType string `json:"location_type,omitempty"`
	// Gender 性别
	Gender enum.AudienceGender `json:"gender,omitempty"`
	// Age 年龄
	Age []enum.AudienceAge `json:"age,omitempty"`
	// Carreer 职业选项，详见【附录-职业】
	Career []enum.Carrier `json:"career,omitempty"`
	// RetargetingTagsInclude 定向人群包列表（自定义人群），内容为人群包id。如果选择"同时定向与排除"，需传入retargeting_tags_include和retargeting_tags_exclude
	RetargetingTagsInclude []uint64 `json:"retargeting_tags_include,omitempty"`
	// RetargetingTagsExclude 排除人群包列表（自定义人群），内容为人群包id。如果选择"同时定向与排除"，需传入retargeting_tags_include和retargeting_tags_exclude
	RetargetingTagsExclude []uint64 `json:"retargeting_tags_exclude,omitempty"`
	// InterestActionMode 行为兴趣;取值："UNLIMITED"不限,"CUSTOM"自定义,"RECOMMEND"系统推荐。若与自定义人群同时使用，系统推荐("RECOMMEND")不生效;仅推广范围为默认时可填，且不可与老版行为兴趣定向同时填写，否则会报错
	InterestActionMode string `json:"interest_action_mode,omitempty"`
	// ActionScene 行为场景
	ActionScene []enum.ActionScene `json:"action_scene,omitempty"`
	// ActionDays 用户发生行为天数，当interest_action_mode传CUSTOM时有效
	ActionDays uint `json:"action_days,omitempty"`
	// ActionCategories 行为类目词，当interest_action_mode传CUSTOM时有效
	ActionCategories []uint64 `json:"action_categories,omitempty"`
	// ActionWords 行为关键词，当interest_action_mode传CUSTOM时有效
	ActionWords []uint64 `json:"action_words,omitempty"`
	// InterestCategories 兴趣类目词，当interest_action_mode传CUSTOM时有效
	InterestCategories []uint64 `json:"Interest_categories,omitempty"`
	// InterestWords 兴趣关键词, 传入具体的词id，非兴趣词包id，可以通过词包相关接口或者兴趣关键词word2id接口获取词id，一个计划下最多创建1000个关键词。当interest_action_mode传CUSTOM时有效
	InterestWords []uint64 `json:"interest_words,omitempty"`
	// AwemeFanBehaviors 抖音达人互动用户行为类型
	AwemeFanBehaviors []enum.Behavior `json:"aweme_fan_behaviors,omitempty"`
	// AwemeFanTimeScope
	AwemeFanTimeScope string `json:"aweme_fan_time_scope,omitempty"`
	// AwemeFanCategories 抖音达人分类ID列表，与aweme_fan_behaviors同时设置才会生效（抖音达人定向）
	AwemeFanCategories []uint64 `json:"aweme_fan_categories,omitempty"`
	// AwemeFanAccounts 抖音达人ID列表，与aweme_fan_behaviors同时设置才会生效（抖音达人定向）
	AwemeFanAccounts []uint64 `json:"aweme_fan_accounts,omitempty"`
	// FilterAwemeAbnormalActive （抖音号推广特有）过滤高活跃用户; 取值：0表示不过滤，1表示过滤
	FilterAwemeAbnormalActive *int `json:"filter_aweme_abnormal_active,omitempty"`
	// FilterAwemeFansCount （抖音号推广特有）过滤高关注数用户，例如"filter_aweme_fans_count": 1000表示过滤粉丝数在1000以上的用户
	FilterAwemeFansCount *int64 `json:"filter_aweme_fans_count,omitempty"`
	// FilterOwnAwemeFans （抖音号推广特有）过滤自己的粉丝; 取值：0表示不过滤，1表示过滤
	FilterOwnAwemeFans *int `json:"filter_own_aweme_fans,omitempty"`
	// SuperiorPopularityType 媒体定向;
	SuperiorPopularityType string `json:"superior_popularity_type,omitempty"`
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
	// HideIfExists 过滤已安装，当推广目标为安卓应用下载时可填，0表示不限，1表示过滤，2表示定向。默认为不限;默认值:0; 取值: 0, 1, 2。建议促进app活跃度客户使用定向安装功能。选择定向安装时，向已安装客户展示广告；投放时优先获取直达链接，无直达链接时使用应用包名进行投放；如果无直达链接或应用包名，定向安装选项实际不生效；仅对Android链接生效。
	HideIfExists int `json:"hide_if_exists,omitempty"`
	// HideIfConverted 过滤已转化用户
	HideIfConverted string `json:"hide_if_converted,omitempty"`
	// ConvertedTimeDuration 过滤时间范围
	ConvertedTimeDuration string `json:"converted_time_duration,omitempty"`
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
	// DpaLbs 地域匹配-LBS;开启时，根据用户的地理位置信息，给用户投放位于其附近的产品 ;取值：0，1（0表示不启用，1表示启用）
	DpaLbs *int `json:"dpa_lbs,omitempty"`
	// DpaCity 地域匹配-商品所在城市;开启时，仅将商品投放给位于该商品设置的可投城市的用户 ;取值：0，1（0表示不启用，1表示启用）
	DpaCity *int `json:"dpa_city,omitempty"`
	// DpaProvince 地域匹配-商品所在省份;开启时，将商品仅投放给位于该商品设置的可投省份的用户 ;取值：0，1（0表示不启用，1表示启用）
	DpaProvince *int `json:"dpa_province,omitempty"`
	// DpaLocationAudience DPA行为重定向，0:不启用，1：启用
	DpaLocationAudience *int `json:"dap_location_audience,omitempty"`
	// IncludeCustomActions 包含人群包((DPA推广目的特有,格式举例[{"days": 7, "code": 1001},]， dpa_local_audience为1时有值; day可选范围:1, 7, 14, 28, 60, 90, 120, 180。
	IncludeCustomActions interface{} `json:"include_custom_actions,omitempty"`
	// ExcludeCustomActions 排除人群包((DPA推广目的特有,格式举例{"days": 7, "code": 1002},]，day可选范围: 1, 7, 14, 28, 60, 90, 120, 180。
	ExcludeCustomActions interface{} `json:"exclude_custom_actions,omitempty"`
	// DpaRecommendType dpa商品重定向推荐类型，dpa_local_audience为1时有值;取值：1（基于重定向推荐更多商品(根据重定向商品和行业特点，推荐更多相关商品投放，包含重定向商品），2仅重定向商品（仅根据重定向人群内定义的重定向行为商品进行投放）
	DpaRecommendType int `json:"dpa_recommend_type,omitempty"`
	// DpaRtaSwitch RTA重定向选项，值：ON:开启，OFF：关闭
	DpaRtaSwitch string `json:"dpa_rta_switch,omitempty"`
	// DpaRtaRecommendType RTA推荐逻辑，ONLY:仅RTA推荐商品，MORE：基于RTA推荐更多商品
	DpaRtaRecommendType string `json:"dpa_rta_recommend_type,omitempty"`
	// SmartBidType 投放场景(出价方式)，详见【附录-自动出价类型】;允许值: 常规投放"SMART_BID_CUSTOM", 放量投放"SMART_BID_CONSERVATIVE"; 概念解释：常规投放：控制成本，尽量消耗完预算；放量投放：接受成本上浮，尽量消耗更多预算
	SmartBidType enum.SmartBidType `json:"smart_bid_type,omitempty"`
	// AdjustCpa 是否调整自动出价，意味如果预期成本不在范围内将在此基础上调整，仅OCPM支持，当smart_bid_type=SMART_BID_CONSERVATIVE时选填。当smart_bid_type为"SMART_BID_CONSERVATIVE"且adjust_cpa=0时，cpa_bid由系统自动计算；;当smart_bid_type为"SMART_BID_CONSERVATIVE" 且adjust_cpa=1时，cpa_bid必填
	AdjustCpa int `json:"adjust_cpa,omitempty"`
	// FlowControlMode 竞价策略(投放方式)
	FlowControlMode enum.FlowControlMode `json:"flow_control_mode,omitempty"`
	// BudgetMode 预算类型（创建后不可修改）， 详见【附录-预算类型】;允许值: "BUDGET_MODE_DAY"日预算, "BUDGET_MODE_TOTAL"总预算
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
	// Pricing 付费方式（计划出价类型）, 详见【附录-计划出价类型】（目前仅穿山甲类型支持OCPC(具体方式：出价类型传OCPC类型，cpa_bid传值 )）; 决定投放目标的类型，比如CPC表示点击量，OCPM表示转化量
	Pricing enum.PricingType `json:"pricing,omitempty"`
	// Bid 出价
	Bid float64 `json:"bid,omitempty"`
	// CpaBid 目标转化出价/预期成本， 当pricing为"OCPM"、"OCPC"出价方式时有值
	CpaBid float64 `json:"cpa_bid,omitempty"`
	// DeepBidType 深度优化方式，允许值详见【附录-深度优化方式】，对于每次付费的转化，深度优化类型需要设置为BID_PER_ACTION（每次付费出价）具体概念见【深度优化方式】;当转化目标中含有深度转化时，该字段必填。获取方式见【查询深度优化方式】
	DeepBidType enum.DeepBidType `json:"deep_bid_type,omitempty"`
	// DeepCpabidid 深度优化出价，deep_bid_type为"DEEP_BID_MIN"时有值。当对应的转化convert_id，设定深度转化目标时才会有效。
	DeepCpabidid uint64 `json:"deep_cpabidid,omitempty"`
	// LubanROiGoal 鲁班目标ROI出价策略系数。推广目的为商品推广(GOODS)时可填。当传入该参数时，表示启用鲁班ROI优化，支持范围(0,100]，精度：保留小数点后四位
	LubanRoiGoal float64 `json:"luban_roi_goal,omitempty"`
	// RoiGoal 深度转化ROI系数, 范围(0,5]，精度：保留小数点后四位, deep_bid_type为"ROI_COEFFICIENT"时有值
	RoiGoal float64 `json:"roi_goal,omitempty"`
	// TargetCvr 目标转化率，转化率优化仅支持电商店铺推广，且付费方式为OCPC或OCPM
	// 允许值：0～100，0表示关闭转化率优化
	TargetCvr float64 `json:"target_cvr,omitempty"`
	// AutoInheritSwitch 一键继承开关，ON表示开启一键继承，OFF表示关闭一键继承
	AutoInheritSwitch string `json:"auto_inherit_switch,omitempty"`
	// InheritType 一键继承账户类型，auto_inherit_switch为ON时有意义，INHERIT_FROM_ACCOUNT表示从同账户下的优质计划中继承，INHERIT_FROM_CUSTOMER表示从同账户所在组织下的其他账户的优质计划中继承
	InheritType string `json:"inherit_type,omitempty"`
	// InheritedAdvertiserID 一键继承的同组织账户id的list，inherit_type等于INHERIT_FROM_CUSTOMER时有意义
	InheritedAdvertiserID []uint64 `json:"inherited_advertiser_id,omitempty"`
	// TrackURL 展示（监测链接）
	TrackURL *[]string `json:"track_url,omitempty"`
	// ActionTrackURL 点击（监测链接）只允许传入1个（当推广目的为应用下载且创建计划传递了convert_id，系统会自动获取转化中的点击监测链接）
	ActionTrackURL *[]string `json:"action_track_url,omitempty"`
	// VideoPlayEffectiveTrackURL 视频有效播放（监测链接），只允许传入1个，投放范围为穿山甲时暂不支持设置此链接
	VideoPlayEffectiveTrackURL *[]string `json:"video_play_effective_track_url,omitempty"`
	// VideoPlayDoneTrackURL 视频播完（监测链接），只允许传入1个，投放范围为穿山甲时暂不支持设置此链接
	VideoPlayDoneTrackURL *[]string `json:"video_play_done_track_url,omitempty"`
	// VideoPlayTrackURL 视频播放（监测链接），只允许传入1个，投放范围为穿山甲时暂不支持设置此链接
	VideoPlayTrackURL *[]string `json:"video_play_track_url,omitempty"`
	// TrackURLSendType 数据发送方式，不可修改,默认值: SERVER_SEND
	// 允许值: SERVER_SEND(服务器端上传), CLIENT_SEND(客户端上传)
	// 客户端上传是指由客户端直接上报给监测平台的服务器, 只有白名单用户才可使用CLIENT_SEND(客户端上传), 如果需要开通请找对接的销售、运营
	TrackURLSendType string `json:"track_url_send_type,omitempty"`
	// AutoUpdateKeyword 是否开启自动加词，ON 开启、OFF 关闭
	AutoUpdateKeyword string `json:"auto_update_keyword,omitempty"`
}

// Encode implement PostRequest interface
func (r CreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// CreateResponse 创建广告计划 API Response
type CreateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// AdID 广告ID
		AdID uint64 `json:"ad_id,omitempty"`
	} `json:"data,omitempty"`
}
