package ad

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

type GetResponse struct {
	model.BaseResponse
	Data *GetResponseData `json:"data,omitempty"`
}

type GetResponseData struct {
	List     []GetResponseList `json:"list,omitempty"`
	PageInfo *model.PageInfo   `json:"page_info,omitempty"`
}

type GetResponseList struct {
	ID                      uint64                   `json:"id,omitempty"`                         // 计划ID
	AdID                    uint64                   `json:"ad_id,omitempty"`                      // 计划ID,返回值同id
	Name                    string                   `json:"name,omitempty"`                       // 计划名称
	AdvertiserID            uint64                   `json:"advertiser_id,omitempty"`              // 广告主ID
	CampaignID              uint64                   `json:"campaign_id,omitempty"`                // 广告组ID
	ModifyTime              string                   `json:"modify_time,omitempty"`                // 计划上次修改时间标识(用于更新计划时提交,服务端判断是否基于最新信息修改)
	AdModifyTime            string                   `json:"ad_modify_time,omitempty"`             // 计划上次修改时间
	AdCreateTime            string                   `json:"ad_create_time,omteimpty"`             // 计划创建时间
	UniqueFk                string                   `json:"unique_fk,omitempty"`                  // 第三方唯一键，传该值时保证接口重试的幂等性，带有相同unique_fk的请求服务端会视为同一个广告处理。仅在创建接口传入且无法修改，如果创建时传入了已存在的唯一键值，那么会返回该唯一键值所对应的广告计划ID。该值可用于内部系统会生成的唯一ID与头条ID做关联的场景，避免超时重试实际上一次创建请求又成功导致的重复创建问题，通过unique_fk可与内部系统ID实现关联并避免重复创建，可结合实际场景选择使用，广告计划中的unique_fk要求不重复，与广告组中的unique_fk无相关。
	Status                  enum.AdStatus            `json:"status,omitempty"`                     // 广告计划投放状态; (进入投放之前,优先披露审核状态,此时优先于启用暂停,启用暂停信息以opt_status为准)
	LearningPhrase          string                   `json:"learning_phrase,omitempty"`            // 学习期状态; 许值：DEFAULT（默认，不在学习期中）、LEARNING（学习中）、LEARNED（学习成功）、LEARN_FAILED（学习失败）;关于学习期，此字段即将废弃关于学习期，此字段即将废弃
	OptStatus               enum.AdOptStatus         `json:"opt_status,omitempty"`                 // 广告计划操作状态, 允许值: "AD_STATUS_ENABLE","AD_STATUS_DISABLE"
	DeliveryRange           enum.AdDeliveryRange     `json:"delivery_range,omitempty"`             // 投放范围
	UnionVideoType          enum.UnionVideoType      `json:"union_video_type,omitempty"`           // 投放形式（穿山甲视频创意类型）;默认值: ORIGINAL_VIDEO原生
	DownloadUrl             string                   `json:"download_url,omitempty"`               // 应用下载方式，推广目的为APP时有值。返回值：DOWNLOAD_URL下载链接，QUICK_APP_URL快应用+下载链接，EXTERNAL_URL落地页链接
	QuickAppUrl             string                   `json:"quick_app_url,omimtepty"`              // 快应用链接，当推广类型为应用推广，且download_type为QUICK_APP_URL时有值
	ExternalUrl             string                   `json:"external_url,omitempty"`               // 落地页链接，投放内容或下载方式为落地页时有值
	AppType                 string                   `json:"app_type,omitempty"`                   // 下载类型，当推广类型为应用推广且download_type为DOWNLOAD_URL或者QUICK_APP_URL时或当推广类型为DPA(商品目录推广)且dpa_adtype为DPA_APP有值
	DownloadMode            string                   `json:"download_mode,omitempty"`              // 优先从系统应用商店下载（下载模式）;允许值：APP_STORE_DELIVERY（仅安卓应用下载支持）、 DEFAULT当应用下载时，默认default下载，可选用APP_STORE_DELIVERY（应用商店直投），当为该值时，将优先跳转目标应用对应手机系统应用商店安装详情页，跳转失败则使用下载链接下载。请确保投放的应用在应用商店内已上架
	ConvertID               int                      `json:"convert_id,omitempty"`                 // 转化目标，其中convert_id数值较小时为预定义转化
	ExternalActions         []string                 `json:"external_actions,omitempty"`           // 转化类型，目前当推广类型为抖音时有值，允许值："AD_CONVERT_TYPE_FOLLOW_ACTION", "AD_CONVERT_TYPE_MESSAGE_ACTION", "AD_CONVERT_TYPE_INTERACTION"
	OpenUrl                 string                   `json:"open_url,omitempty"`                   // 直达链接(点击唤起APP)
	AdvancedCreativeType    string                   `json:"advanced_creative_type,omitempty"`     // 附加创意类型; 允许值: ATTACHED_CREATIVE_GAME_PACKAGE游戏礼包码,ATTACHED_CREATIVE_GAME_FORM游戏表单收集,ATTACHED_CREATIVE_GAME_SUBSCRIBE游戏预约,ATTACHED_CREATIVE_NONE无推广目的为应用推广类型、下载方式选择下载链接且下载链接为安卓应用下载时才可以设置
	GamePackageDesc         string                   `json:"game_package_desc,omitempty"`          // 应用描述,最少1字，最多15字
	GamePackageBatchID      int64                    `json:"game_package_batch_id,omitempty"`      // 游戏礼包码id，目前仅支持直接发券类型
	GamePackageThumbnails   []string                 `json:"game_package_thumbnails,omitempty"`    // 应用图片集，图片URL链接，有且仅有两个该字段10.10日下线
	GamePackageThumbnailIDs []string                 `json:"game_package_thumbnail_ids,omitempty"` // 应用图片集，图片image_id，有且仅有两个，游戏礼包码时有值,可以从 【素材管理-获取图片素材】 接口中获取，建议尺寸 16:9 与game_package_thumbnails同时传入时，以game_package_thumbnail_ids字段为准
	StoreproUnit            string                   `json:"storepro_unit,omitempty"`              // 门店推广-投放内容，当推广目的为STORE(门店推广)时有值。取值: "STORE"门店, "STORE_ACTIVITY"活动;目前暂时不支持线下商品类型
	StoreType               string                   `json:"store_type,omitempty"`                 // 门店类型，（storepro_unit 为 "STORE" 时有值。取值: "STORE_NORMAL"平台通用门店, "STORE_THIRT_PARTY"第三方门店, "STORE_DOUYIN"抖音POI门店
	AdvertiserStoreIDs      []uint64                 `json:"advertiser_store_ids,omitempty"`       // 门店ID列表 （storepro_unit 为 "STORE" 时有值
	StoreproPackID          uint64                   `json:"storepro_pack_id,omitempty"`           // 活动ID （storepro_unit 为 "STORE_ACTIVITY" 时有值
	ProductPlatformID       int64                    `json:"product_platform_id,omitempty"`        // 产品目录ID(ID由查询产品目录接口得到), 当推广目的landing_type为DPA时有值
	ProductID               string                   `json:"product_id,omitempty"`                 // 商品id，当推广目的为 DPA 广告组商品类型为 SDPA 时有值
	CategoryType            string                   `json:"category_type,omitempty"`              // DPA投放范围，取值：NONE不限，"CATEGORY"选择分类，"PRODUCT"指定商品
	DpaCategories           []uint64                 `json:"dpa_categories,omitempty"`             // 分类列表，category_type取值范围为CATEGORY时有值
	DpaProducts             []uint64                 `json:"dpa_products,omitempty"`               // 商品列表，category_type为PRODUCT时有值
	DpaProductTarget        *DpaProductTarget        `json:"dpa_product_target,omitempty"`         // 自定义筛选条件（商品投放条件）。用于圈定商品投放范围，结合商品库字段搭配判断条件，圈定商品投放范围。
	DpaAdtype               []string                 `json:"dpa_adtype,omitempty"`                 // dpa广告类型，取值范围："DPA_LINK"落地页, "DPA_APP"应用下载
	ParamsType              string                   `json:"params_type,omitempty"`                // 链接类型(落地页)，当dpa_adtype为"DPA_LINK"时有值，取值: "DPA"商品库所含链接, "CUSTOM"自定义链接
	DpaExternalUrlField     string                   `json:"dpa_external_url_field,omitempty"`     // 落地页链接字段选择，当params_type为"DPA"时有值
	DpaExternalUrls         []string                 `json:"dpa_external_urls,omitempty"`          // 落地页链接地址列表，当params_type为"CUSTOM"时有值
	Package                 string                   `json:"package,omitempty"`                    // 应用包名，当推广类型为应用推广且download_type为DOWNLOAD_URL时或当推广类型为DPA(商品目录推广)且dpa_adtype为DPA_APP有值
	InventoryType           []enum.StatInventoryType `json:"inventory_type,omitempty"`             // 创意投放位置; 创建选择优选广告位时，此字段回会返回对应的优选广告位
	PromotionType           string                   `json:"promotion_type,omitempty"`             // 投放内容; GOODS：商品推广;LIVE：直播; AWEME_HOME_PAGE：抖音主页;LANDING_PAGE_LINK：落地页
	AwemeAccount            string                   `json:"aweme_account,omitempty"`              // 抖音号
	SubscribeUrl            string                   `json:"subscribe_url,omitempty"`              // 游戏营销场景-预约下载链接
	FormID                  uint64                   `json:"form_id,omitempty"`                    // 游戏营销场景-表单id
	FormIndex               int                      `json:"form_index,omitempty"`                 // 游戏营销场景-表单位置索引
	AppDesc                 string                   `json:"app_desc,omitempty"`                   // 游戏营销场景-应用描述
	AppIntroduction         string                   `json:"app_introduction,omitempty"`           // 游戏营销场景-应用介绍
	AppThumbnails           []string                 `json:"app_thumbnails,omitempty"`             // 游戏营销场景-应用图片集，返回图片集Id
	DpaOpenUrlType          string                   `json:"dpa_open_url_type,omitempty"`          // 直达链接类型，取值: "NONE"不启用, "DPA"商品库所含链接, "CUSTOM"自定义链接商品库链接对应商品库内调起字段。
	DpaOpenUrlField         string                   `json:"dpa_open_url_field,omitempty"`         // 直达链接字段选择，当dpa_open_url_type为"DPA"时有值
	DpaOpenUrls             []string                 `json:"dpa_open_urls,omitempty"`              // 直达链接地址列表，当dpa_open_url_type为"CUSTOM"时有值
	ExternalUrlParams       string                   `json:"external_url_params,omitempty"`        // 落地页检测参数(DPA推广目的特有,在填写的参数后面添加"=urlencode(开放平台提供的h5链接地址）"，其中urlencode(开放平台提供的h5链接地址）替换为商品库中的h5地址encode的结果)
	OpenUrlParams           string                   `json:"open_url_params,omitempty"`            // 直达链接检测参数(DPA推广目的特有,在“产品库中提取的scheme地址"后面追加填写的参数)
	AudiencePackageID       uint64                   `json:"audience_package_id,omitempty"`        // 定向包ID
	Audience                *Audience                `json:"audience,omitempty"`                   // 广告受众, 字典类型, 包含下面字段
	HideIfExists            int                      `json:"hide_if_exists,omitempty"`             // 过滤已安装，当推广目标为安卓应用下载时可填，0表示不限，1表示过滤，2表示定向。默认为不限;默认值:0; 取值: 0, 1, 2。建议促进app活跃度客户使用定向安装功能。选择定向安装时，向已安装客户展示广告；投放时优先获取直达链接，无直达链接时使用应用包名进行投放；如果无直达链接或应用包名，定向安装选项实际不生效；仅对Android链接生效。
	HideIfConverted         string                   `json:"hide_if_converted,omitempty"`          // 过滤已转化用户
	ConvertedTimeDuration   string                   `json:"converted_time_duration,omitempty"`    // 过滤时间范围
	DpaLbs                  int                      `json:"dpa_lbs,omitempty"`                    // 地域匹配-LBS;开启时，根据用户的地理位置信息，给用户投放位于其附近的产品 ;取值：0，1（0表示不启用，1表示启用）
	DpaCity                 int                      `json:"dpa_city,omitempty"`                   // 地域匹配-商品所在城市;开启时，仅将商品投放给位于该商品设置的可投城市的用户 ;取值：0，1（0表示不启用，1表示启用）
	DpaProvince             int                      `json:"dpa_province,omitempty"`               // 地域匹配-商品所在省份;开启时，将商品仅投放给位于该商品设置的可投省份的用户 ;取值：0，1（0表示不启用，1表示启用）
	DpaLocationAudience     int                      `json:"dap_location_audience,omitempty"`      // DPA行为重定向，0:不启用，1：启用
	IncludeCustomActions    json.RawMessage          `json:"include_custom_actions,omitempty"`     // 包含人群包((DPA推广目的特有,格式举例[{"days": 7, "code": 1001},]， dpa_local_audience为1时有值; day可选范围:1, 7, 14, 28, 60, 90, 120, 180。
	ExcludeCustomActions    json.RawMessage          `json:"exclude_custom_actions,omitempty"`     // 排除人群包((DPA推广目的特有,格式举例{"days": 7, "code": 1002},]，day可选范围: 1, 7, 14, 28, 60, 90, 120, 180。
	DpaRecommendType        int                      `json:"dpa_recommend_type,omitempty"`         // dpa商品重定向推荐类型，dpa_local_audience为1时有值;取值：1（基于重定向推荐更多商品(根据重定向商品和行业特点，推荐更多相关商品投放，包含重定向商品），2仅重定向商品（仅根据重定向人群内定义的重定向行为商品进行投放）
	SmartBidType            enum.SmartBidType        `json:"smart_bid_type,omitempty"`             // 投放场景(出价方式)
	AdjustCpa               int                      `json:"adjust_cpa,omitempty"`                 // 是否调整自动出价，意味如果预期成本不在范围内将在此基础上调整，仅OCPM支持，当smart_bid_type=SMART_BID_CONSERVATIVE时选填。当smart_bid_type为"SMART_BID_CONSERVATIVE"且adjust_cpa=0时，cpa_bid由系统自动计算；;当smart_bid_type为"SMART_BID_CONSERVATIVE" 且adjust_cpa=1时，cpa_bid必填
	FlowControlMode         enum.FlowControlMode     `json:"flow_control_mode,omitempty"`          // 竞价策略(投放方式)
	BudgetMode              enum.BudgetMode          `json:"budget_mode,omitempty"`                // 预算类型(创建后不可修改),
	ScheduleType            enum.ScheduleType        `json:"schedule_type,omitempty"`              // 投放时间类型
	StartTime               string                   `json:"start_time,omitempty"`                 // 投放起始时间，当schedule_type为"SCHEDULE_START_END"时取值，形式如：2017-01-01 00:00
	EndTime                 string                   `json:"end_time,omitempty"`                   // 投放结束时间，当schedule_type为"SCHEDULE_START_END"时取值，形式如：2017-01-01 00:00
	ScheduleTime            string                   `json:"schedule_time,omitempty"`              // 投放时段，默认全时段投放，格式是48*7位字符串，且都是0或1。也就是以半个小时为最小粒度，周一至周日每天分为48个区段，0为不投放，1为投放，不传、全传0、全传1均代表全时段投放。
	Pricing                 enum.PricingType         `json:"pricing,omitempty"`                    // 付费方式（计划出价类型）
	Bid                     float64                  `json:"bid,omitempty"`                        // 点击出价/展示出价，当pricing为"CPC"、"CPM"、"CPV"出价方式时有值
	CpaBid                  float64                  `json:"cpa_bid,omitempty"`                    // 目标转化出价/预期成本， 当pricing为"OCPM"、"OCPC"出价方式时有值
	DeepBidType             enum.DeepBidType         `json:"deep_bid_type,omitempty"`              // 深度优化方式
	DeepCpabidid            uint64                   `json:"deep_cpabidid,omitempty"`              // 深度优化出价，deep_bid_type为"DEEP_BID_MIN"时有值。当对应的转化convert_id，设定深度转化目标时才会有效。
	LubanRoiGoal            float64                  `json:"luban_roi_goal,omitempty"`             // 鲁班目标ROI出价策略系数。推广目的为商品推广(GOODS)时可填。当传入该参数时，表示启用鲁班ROI优化，支持范围(0,100]，精度：保留小数点后四位
	RoiGoal                 float64                  `json:"roi_goal,omitempty"`                   // 深度转化ROI系数, 范围(0,5]，精度：保留小数点后四位, deep_bid_type为"ROI_COEFFICIENT"时有值
}

type DpaProductTarget struct {
	Title string `json:"title,omitempty"` // 筛选字段
	Rule  string `json:"rule,omitempty"`  // 定向规则，允许值：'=', '!=', '>', '<', '>=', '<=', 'contain', 'exclude', 'notEmpty
	Type  string `json:"type,omitempty"`  // 字段类型，允许值：'int', 'double', 'long', 'string'
	Value string `json:"value,omitempty"` // 规则值
}

type Audience struct {
	District                  string                  `json:"district,omitempty"`                     // 地域;取值: "CITY"省市, "COUNTY"区县, "BUSINESS_DISTRICT"商圈,"NONE"不限，省市传法："city": [12],"district": "CITY",区县的传法："city": [130102],"district": "COUNTY";暂不支持"海外"
	City                      []uint64                `json:"city,omitempty"`                         // 地域定向省市或者区县列表(当传递省份ID时,旗下市县ID可省略不传),当district为"CITY"或"COUNTY"时有值
	BusinessIDs               []uint64                `json:"business_ids,omitempty"`                 // 商圈ID数组，district为"BUSINESS_DISTRICT"时有值
	Geolocation               []model.Geolocation     `json:"geolocation,omitempty"`                  // 从地图添加(地图位置)
	LocationType              string                  `json:"location_type,omitempty"`                // 位置类型;取值：CURRENT正在该地区的用户，HOME居住在该地区的用户，TRAVEL到该地区旅行的用户，ALL该地区内的所有用户;当city和district有值时返回值
	Gender                    enum.AudienceGender     `json:"gender,omitempty"`                       // 性别
	Age                       []enum.AudienceAge      `json:"age,omitempty"`                          // 年龄
	RetargetingTagsInclude    []uint64                `json:"retargeting_tags_include,omitempty"`     // 定向人群包列表（自定义人群），内容为人群包id。如果选择"同时定向与排除"，需传入retargeting_tags_include和retargeting_tags_exclude
	RetargetingTagsExclude    []uint64                `json:"retargeting_tags_exclude,omitempty"`     // 排除人群包列表（自定义人群），内容为人群包id。如果选择"同时定向与排除"，需传入retargeting_tags_include和retargeting_tags_exclude
	InterestActionMode        string                  `json:"interest_action_mode,omitempty"`         // 行为兴趣;取值："UNLIMITED"不限,"CUSTOM"自定义,"RECOMMEND"系统推荐。若与自定义人群同时使用，系统推荐("RECOMMEND")不生效;仅推广范围为默认时可填，且不可与老版行为兴趣定向同时填写，否则会报错
	ActionScene               []enum.ActionScene      `json:"action_scene,omitempty"`                 // 行为场景
	ActionDays                uint                    `json:"action_days,omitempty"`                  // 用户发生行为天数，当interest_action_mode传CUSTOM时有效
	ActionCategories          []uint64                `json:"action_categories,omitempty"`            // 行为类目词，当interest_action_mode传CUSTOM时有效
	ActionWords               []uint64                `json:"action_words,omitempty"`                 // 行为关键词，当interest_action_mode传CUSTOM时有效
	InterestCategories        []uint64                `json:"Interest_categories,omitempty"`          // 兴趣类目词，当interest_action_mode传CUSTOM时有效
	InterestWords             []uint64                `json:"interest_words,omitempty"`               // 兴趣关键词, 传入具体的词id，非兴趣词包id，可以通过词包相关接口或者兴趣关键词word2id接口获取词id，一个计划下最多创建1000个关键词。当interest_action_mode传CUSTOM时有效
	AdTags                    []uint64                `json:"ad_tags,omitempty"`                      // （老版行为兴趣）兴趣分类,如果传"空数组"表示不限，如果"数组传0"表示系统推荐,如果按兴趣类型传表示自定义
	InterestTags              []uint64                `json:"interest_tags,omitempty"`                // （老版行为兴趣）兴趣关键词, 传入具体的词id，非兴趣词包id，可以通过词包相关接口或者兴趣关键词word2id接口获取词id，一个计划下最多创建1000个关键词。
	AppBehaviorTarget         string                  `json:"app_behavior_target,omitempty"`          // （老版行为兴趣）APP行为; 取值：NONE不限，CATEGORY按分类，APP按APP
	AppCategory               []uint64                `json:"app_category,omitempty"`                 // 老版行为兴趣）APP行为定向——按分类
	AppIDs                    []uint64                `json:"app_ids,omitempty"`                      // （老版行为兴趣）APP行为定向——按APP（请注意如果投放的是"应用下载-IOS"不支持设置APP行为定向，请勿传值。）可通过【工具-查询工具-查询应用信息】获取。当app_behavior_target为APP时有值
	AwemeFanBehavior          []enum.Behavior         `json:"aweme_fan_behavior,omitempty"`           // 抖音达人互动用户行为类型
	AwemeFanCategories        []uint64                `json:"aweme_fan_categories,omitempty"`         // 抖音达人分类ID列表，与aweme_fan_behaviors同时设置才会生效（抖音达人定向）
	AwemeFanAccounts          []uint64                `json:"aweme_fan_accounts,omitempty"`           // 抖音达人ID列表，与aweme_fan_behaviors同时设置才会生效（抖音达人定向）
	AwemeFansNumbers          []int64                 `json:"aweme_fans_numbers,omitempty"`           // （抖音号推广特有）账号粉丝相似人群（添加抖音账号，会将广告投放给对应账号的相似人群粉丝）
	FilterAwemeAbnormalActive int                     `json:"filter_aweme_abnormal_active,omitempty"` // （抖音号推广特有）过滤高活跃用户; 取值：0表示不过滤，1表示过滤
	FilterAwemeFansCount      int64                   `json:"filter_aweme_fans_count,omitempty"`      // （抖音号推广特有）过滤高关注数用户，例如"filter_aweme_fans_count": 1000表示过滤粉丝数在1000以上的用户
	FilterOwnAwemeFans        int                     `json:"filter_own_aweme_fans,omitempty"`        // （抖音号推广特有）过滤自己的粉丝; 取值：0表示不过滤，1表示过滤
	SuperiorPopularityType    string                  `json:"superior_popularity_type,omitempty"`     // 媒体定向;
	FlowPackage               []uint64                `json:"flow_package,omitempty"`                 // 定向逻辑
	ExcludeFlowPackage        []uint64                `json:"exclude_flow_package,omitempty"`         // 排除定向逻辑
	Platform                  []enum.AudiencePlatform `json:"platform,omitempty"`                     // 平台，当下载方式包含下载链接时，平台类型需与选择的下载链接类型对应，当下载方式不包含下载方式的时候，平台可多选。为保证投放效果,平台类型定向PC与移动端互斥
	AndroidOsv                string                  `json:"android_osv,omitempty"`                  // 最低安卓版本，当app_type为"APP_ANDROID"选填,其余情况不填,
	IosOsv                    string                  `json:"ios_osv,omitempty"`                      // 最低IOS版本，当app_type为"APP_IOS"选填,其余情况不填
	DeviceType                []string                `json:"device_type,omitempty"`                  // 设备类型;取值是："MOBILE", "PAD"。缺省表示不限设备类型。穿山甲已经全量，投放范围为默认时需要有白名单权限才可以
	Ac                        []string                `json:"ac,omitempty"`                           // 网络类型
	Carrier                   []enum.Carrier          `json:"carrier,omitempty"`                      // 运营商
	ActivateType              []enum.ActivateType     `json:"activate_type,omitempty"`                // 新用户(新用户使用头条的时间)
	ArticleCategory           []string                `json:"article_category,omitempty"`             // 文章分类
	DeviceBrand               []string                `json:"device_brand,omitempty"`                 // 手机品牌
	LaunchPrice               []int                   `json:"launch_price,omitempty"`                 // 手机价格,传入价格区间，最高传入11000（表示1w以上）;传值示例 "launch_price": [2000, 11000]，表示2000元以上;
	AutoExtendEnabled         int                     `json:"auto_extend_enabled,omitempty"`          // 是否启用智能放量。取值是：0、1。缺省为 0。
	AutoExtendTarget          []string                `json:"auto_extend_targets,omitempty"`          // 可放开定向。当auto_extend_enabled=1 时选填。详见：【附录-可开放定向】。缺省为全不选。
}
