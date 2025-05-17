package ad

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateRequest 更新计划 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdID 广告计划ID，广告计划id需要属于广告主，且ad_id不能重复，否则会报错！
	AdID uint64 `json:"ad_id,omitempty"`
	// Name 广告计划名称，长度为1-100个字符，其中1个中文字符算2位
	Name string `json:"name,omitempty"`
	// DownloadMode 优先从系统应用商店下载（下载模式），当下载方式为下载链接时，可修改。允许值：APP_STORE_DELIVERY（仅安卓应用下载支持）、 DEFAULT当应用下载时，默认default下载，可选用APP_STORE_DELIVERY（应用商店直投），选择后，将优先跳转目标应用对应手机系统应用商店安装详情页，跳转失败则使用下载链接下载。
	DownloadMode enum.DownloadMode `json:"download_mode,omitempty"`
	// OpenUrl 直达链接（点击唤起APP）直达链接仅支持部分App唤起，点击创意将优先跳转App，再根据投放内容跳转相关链接
	OpenUrl *string `json:"open_url,omitempty"`
	// Ulink 直达备用链接，仅支持穿山甲广告位（不支持搜索广告）
	Ulink *string `json:"ulink,omitempty"`
	// ExternalUrl 落地页链接（支持橙子建站落地页）;对于转化量为目标的计划如OCPM计划不允许更改，非转化为目标的计划如CPC、CPM计划可更改; 获取橙子建站落地页可参考【橙子建站落地页管理】
	ExternalUrl string `json:"external_url,omitempty"`
	// CategoryType 商品目录投放范围,当广告组商品类型选择 DPA 多商品时可修改;允许值：NONE不限，"CATEGORY"选择分类，"PRODUCT"指定商品
	CategoryType string `json:"category_type,omitempty"`
	// DpaCategories 分类列表，category_type更新为"CATEGORY"时可修改，由【DPA商品广告-获取DPA分类】 得到限制个数1~100
	DpaCategories *[]uint64 `json:"dpa_categories,omitempty"`
	// DpaProducts 商品列表，category_type更新为"PRODUCT"时可修改，由【DPA商品广告-获取DPA商品库商品列表】 得到;限制个数1~100
	DpaProducts *[]uint64 `json:"dpa_products,omitempty"`
	// DpaProductTargets 自定义筛选条件（商品投放条件）。用于圈定商品投放范围，结合商品库字段搭配判断条件，圈定商品投放范围。
	DpaProductTarget *[]DpaProductTarget `json:"dpa_product_targets,omitempty"`
	// ParamsType 链接类型(落地页)，当dpa_adtype为"DPA_LINK"时有值，取值: "DPA"商品库所含链接, "CUSTOM"自定义链接
	ParamsType string `json:"params_type,omitempty"`
	// DpaExternalUrlField 落地页链接字段选择，当params_type为"DPA"时有值
	DpaExternalUrlField string `json:"dpa_external_url_field,omitempty"`
	// DpaExternalUrls 落地页链接地址列表，当params_type为"CUSTOM"时有值
	DpaExternalUrls *[]string `json:"dpa_external_urls,omitempty"`
	// AppType 下载类型，当推广类型为应用推广且download_type为DOWNLOAD_URL或者QUICK_APP_URL时或当推广类型为DPA(商品目录推广)且dpa_adtype为DPA_APP有值
	AppType string `json:"app_type,omitempty"`
	// DownloadUrl 应用下载方式，推广目的为APP时有值。返回值：DOWNLOAD_URL下载链接，QUICK_APP_URL快应用+下载链接，EXTERNAL_URL落地页链接
	DownloadUrl string `json:"download_url,omitempty"`
	// Package 应用包名，当推广类型为应用推广且download_type为DOWNLOAD_URL时或当推广类型为DPA(商品目录推广)且dpa_adtype为DPA_APP有值
	Package string `json:"package,omitempty"`
	// DpaOpenUrlType 直达链接类型，取值: "NONE"不启用, "DPA"商品库所含链接, "CUSTOM"自定义链接商品库链接对应商品库内调起字段。
	DpaOpenUrlType string `json:"dpa_open_url_type,omitempty"`
	// DpaOpenUrlField 直达链接字段选择，当dpa_open_url_type为"DPA"时有值
	DpaOpenUrlField string `json:"dpa_open_url_field,omitempty"`
	// DpaOpenUrls 直达链接地址列表，当dpa_open_url_type为"CUSTOM"时有值
	DpaOpenUrls *[]string `json:"dpa_open_urls,omitempty"`
	// ExternalUrlParams 落地页检测参数(DPA推广目的特有,在填写的参数后面添加"=urlencode(开放平台提供的h5链接地址）"，其中urlencode(开放平台提供的h5链接地址）替换为商品库中的h5地址encode的结果)
	ExternalUrlParams string `json:"external_url_params,omitempty"`
	// OpenUrlParams 直达链接检测参数(DPA推广目的特有,在“产品库中提取的scheme地址"后面追加填写的参数)
	OpenUrlParams string `json:"open_url_params,omitempty"`
	// FeedDeliverySearch 搜索快投关键词，HAS_OPEN:启用，DISABLE:不启用
	FeedDeliverySearch string `json:"feed_delivery_search,omitempty"`
	// IntelligentFlowSwitch 智能流量开关，ON:开启，OFF:关闭
	IntelligentFlowSwitch enum.OnOff `json:"intelligent_flow_switch,omitempty"`
	// AudiencePackageID 定向包ID
	AudiencePackageID uint64 `json:"audience_package_id,omitempty"`
	// District 地域;取值: "CITY"省市, "COUNTY"区县, "BUSINESS_DISTRICT"商圈,"NONE"不限，省市传法："city": [12],"district": "CITY",区县的传法："city": [130102],"district": "COUNTY";暂不支持"海外"
	District enum.District `json:"district,omitempty"`
	// RegionVersion 行政区域版本号。通过[【获取行政信息】]https://open.oceanengine.com/doc/index.html?key=ad&type=api&id=1709606596424718)接口获取; district =REGION/OVERSEA时必填
	RegionVersion string `json:"region_version,omitempty"`
	// City 地域定向省市或者区县列表(当传递省份ID时,旗下市县ID可省略不传),当district为"CITY"或"COUNTY"时有值
	City *[]uint64 `json:"city,omitempty"`
	// BusinessIDs 商圈ID数组，district为"BUSINESS_DISTRICT"时有值
	BusinessIDs *[]uint64 `json:"business_ids,omitempty"`
	// Geolocation 从地图添加(地图位置)
	Geolocation *[]model.Geolocation `json:"geolocation,omitempty"`
	// LocationType 位置类型;取值：CURRENT正在该地区的用户，HOME居住在该地区的用户，TRAVEL到该地区旅行的用户，ALL该地区内的所有用户;当city和district有值时返回值
	LocationType enum.LocationType `json:"location_type,omitempty"`
	// Gender 性别
	Gender enum.AudienceGender `json:"gender,omitempty"`
	// Age 年龄
	Age *[]enum.AudienceAge `json:"age,omitempty"`
	// Carreer 职业选项，详见【附录-职业】
	Career *[]string `json:"career,omitempty"`
	// RetargetingTagsInclude 定向人群包列表（自定义人群），内容为人群包id。如果选择"同时定向与排除"，需传入retargeting_tags_include和retargeting_tags_exclude
	RetargetingTagsInclude *[]uint64 `json:"retargeting_tags_include,omitempty"`
	// RetargetingTagsExclude 排除人群包列表（自定义人群），内容为人群包id。如果选择"同时定向与排除"，需传入retargeting_tags_include和retargeting_tags_exclude
	RetargetingTagsExclude *[]uint64 `json:"retargeting_tags_exclude,omitempty"`
	// InterestActionMode 行为兴趣;取值："UNLIMITED"不限,"CUSTOM"自定义,"RECOMMEND"系统推荐。若与自定义人群同时使用，系统推荐("RECOMMEND")不生效;仅推广范围为默认时可填，且不可与老版行为兴趣定向同时填写，否则会报错
	InterestActionMode enum.InterestActionMode `json:"interest_action_mode,omitempty"`
	// ActionScene 行为场景
	ActionScene *[]enum.ActionScene `json:"action_scene,omitempty"`
	// ActionDays 用户发生行为天数，当interest_action_mode传CUSTOM时有效
	ActionDays uint `json:"action_days,omitempty"`
	// ActionCategories 行为类目词，当interest_action_mode传CUSTOM时有效
	ActionCategories *[]uint64 `json:"action_categories,omitempty"`
	// ActionWords 行为关键词，当interest_action_mode传CUSTOM时有效
	ActionWords *[]uint64 `json:"action_words,omitempty"`
	// InterestCategories 兴趣类目词，当interest_action_mode传CUSTOM时有效
	InterestCategories *[]uint64 `json:"interest_categories,omitempty"`
	// InterestWords 兴趣关键词, 传入具体的词id，非兴趣词包id，可以通过词包相关接口或者兴趣关键词word2id接口获取词id，一个计划下最多创建1000个关键词。当interest_action_mode传CUSTOM时有效
	InterestWords *[]uint64 `json:"interest_words,omitempty"`
	// AwemeFanBehaviors 抖音达人互动用户行为类型
	AwemeFanBehaviors *[]enum.Behavior `json:"aweme_fan_behaviors,omitempty"`
	// AwemeFanTimeScope
	AwemeFanTimeScope string `json:"aweme_fan_time_scope,omitempty"`
	// AwemeFanCategories 抖音达人分类ID列表，与aweme_fan_behaviors同时设置才会生效（抖音达人定向）
	AwemeFanCategories *[]uint64 `json:"aweme_fan_categories,omitempty"`
	// AwemeFanAccounts 抖音达人ID列表，与aweme_fan_behaviors同时设置才会生效（抖音达人定向）
	AwemeFanAccounts *[]uint64 `json:"aweme_fan_accounts,omitempty"`
	// FilterAwemeAbnormalActive （抖音号推广特有）过滤高活跃用户; 取值：0表示不过滤，1表示过滤
	FilterAwemeAbnormalActive *int `json:"filter_aweme_abnormal_active,omitempty"`
	// FilterAwemeFansCount （抖音号推广特有）过滤高关注数用户，例如"filter_aweme_fans_count": 1000表示过滤粉丝数在1000以上的用户
	FilterAwemeFansCount *int64 `json:"filter_aweme_fans_count,omitempty"`
	// FilterOwnAwemeFans （抖音号推广特有）过滤自己的粉丝; 取值：0表示不过滤，1表示过滤
	FilterOwnAwemeFans *int `json:"filter_own_aweme_fans,omitempty"`
	// SuperiorPopularityType 媒体定向;
	SuperiorPopularityType enum.SuperiorPopularityType `json:"superior_popularity_type,omitempty"`
	// FlowPackage 定向逻辑
	FlowPackage *[]uint64 `json:"flow_package,omitempty"`
	// ExcludeFlowPackage 排除定向逻辑
	ExcludeFlowPackage *[]uint64 `json:"exclude_flow_package,omitempty"`
	// Platform 平台，当下载方式包含下载链接时，平台类型需与选择的下载链接类型对应，当下载方式不包含下载方式的时候，平台可多选。为保证投放效果,平台类型定向PC与移动端互斥
	Platform *[]enum.AudiencePlatform `json:"platform,omitempty"`
	// AndroidOsv 最低安卓版本，当app_type为"APP_ANDROID"选填,其余情况不填,
	AndroidOsv string `json:"android_osv,omitempty"`
	// IosOsv 最低IOS版本，当app_type为"APP_IOS"选填,其余情况不填
	IosOsv string `json:"ios_osv,omitempty"`
	// DeviceType 设备类型;取值是："MOBILE", "PAD"。缺省表示不限设备类型。穿山甲已经全量，投放范围为默认时需要有白名单权限才可以
	DeviceType *[]string `json:"device_type,omitempty"`
	// Ac 网络类型
	Ac *[]string `json:"ac,omitempty"`
	// Carrier 运营商
	Carrier *[]enum.Carrier `json:"carrier,omitempty"`
	// HideIfExists 过滤已安装，当推广目标为安卓应用下载时可填，0表示不限，1表示过滤，2表示定向。默认为不限;默认值:0; 取值: 0, 1, 2。建议促进app活跃度客户使用定向安装功能。选择定向安装时，向已安装客户展示广告；投放时优先获取直达链接，无直达链接时使用应用包名进行投放；如果无直达链接或应用包名，定向安装选项实际不生效；仅对Android链接生效。
	HideIfExists *int `json:"hide_if_exists,omitempty"`
	// HideIfConverted 过滤已转化用户
	HideIfConverted enum.HideIfConverted `json:"hide_if_converted,omitempty"`
	// ConvertedTimeDuration 过滤时间范围
	ConvertedTimeDuration enum.ConvertedTimeDuration `json:"converted_time_duration,omitempty"`
	// ActivateType 新用户(新用户使用头条的时间)
	ActivateType *[]enum.ActivateType `json:"activate_type,omitempty"`
	// ArticleCategory 文章分类
	ArticleCategory *[]string `json:"article_category,omitempty"`
	// DeviceBrand 手机品牌
	DeviceBrand *[]string `json:"device_brand,omitempty"`
	// LauchPrice 手机价格,传入价格区间，最高传入11000（表示1w以上）;传值示例 "launch_price": [2000, 11000]，表示2000元以上;
	LaunchPrice *[]int `json:"launch_price,omitempty"`
	// AutoExtendEnabled 是否启用智能放量。取值是：0、1。缺省为 0。
	AutoExtendEnabled *int `json:"auto_extend_enabled,omitempty"`
	// AutoExtendTarget 可放开定向。当auto_extend_enabled=1 时选填。详见：【附录-可开放定向】。缺省为全不选。
	AutoExtendTarget *[]string `json:"auto_extend_targets,omitempty"`
	// DpaLbs 地域匹配-LBS;开启时，根据用户的地理位置信息，给用户投放位于其附近的产品 ;取值：0，1（0表示不启用，1表示启用）
	DpaLbs *int `json:"dpa_lbs,omitempty"`
	// DpaCity 地域匹配-商品所在城市;开启时，仅将商品投放给位于该商品设置的可投城市的用户 ;取值：0，1（0表示不启用，1表示启用）
	DpaCity *int `json:"dpa_city,omitempty"`
	// DpaProvince 地域匹配-商品所在省份;开启时，将商品仅投放给位于该商品设置的可投省份的用户 ;取值：0，1（0表示不启用，1表示启用）
	DpaProvince *int `json:"dpa_province,omitempty"`
	// DpaLocationAudience DPA行为重定向，0:不启用，1：启用
	DpaLocationAudience *int `json:"dap_location_audience,omitempty"`
	// IncludeCustomActions 包含人群包((DPA推广目的特有,格式举例[{"days": 7, "code": 1001},]， dpa_local_audience为1时有值; day可选范围:1, 7, 14, 28, 60, 90, 120, 180。
	IncludeCustomActions json.RawMessage `json:"include_custom_actions,omitempty"`
	// ExcludeCustomActions 排除人群包((DPA推广目的特有,格式举例{"days": 7, "code": 1002},]，day可选范围: 1, 7, 14, 28, 60, 90, 120, 180。
	ExcludeCustomActions json.RawMessage `json:"exclude_custom_actions,omitempty"`
	// DpaRecommendType dpa商品重定向推荐类型，dpa_local_audience为1时有值;取值：1（基于重定向推荐更多商品(根据重定向商品和行业特点，推荐更多相关商品投放，包含重定向商品），2仅重定向商品（仅根据重定向人群内定义的重定向行为商品进行投放）
	DpaRecommendType int `json:"dpa_recommend_type,omitempty"`
	// AdjustCpa 是否调整自动出价，意味如果预期成本不在范围内将在此基础上调整，仅OCPM支持，当smart_bid_type=SMART_BID_CONSERVATIVE时选填。当smart_bid_type为"SMART_BID_CONSERVATIVE"且adjust_cpa=0时，cpa_bid由系统自动计算；;当smart_bid_type为"SMART_BID_CONSERVATIVE" 且adjust_cpa=1时，cpa_bid必填
	AdjustCpa *int `json:"adjust_cpa,omitempty"`
	// FlowControlMode 竞价策略(投放方式)
	FlowControlMode enum.FlowControlMode `json:"flow_control_mode,omitempty"`
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
	// Bid 出价
	Bid float64 `json:"bid,omitempty"`
	// CpaBid 目标转化出价/预期成本， 当pricing为"OCPM"、"OCPC"出价方式时有值
	CpaBid float64 `json:"cpa_bid,omitempty"`
	// DeepCpabidid 深度优化出价，deep_bid_type为"DEEP_BID_MIN"时有值。当对应的转化convert_id，设定深度转化目标时才会有效。
	DeepCpabidid float64 `json:"deep_cpabidid,omitempty"`
	// LubanROiGoal 鲁班目标ROI出价策略系数。推广目的为商品推广(GOODS)时可填。当传入该参数时，表示启用鲁班ROI优化，支持范围(0,100]，精度：保留小数点后四位
	LubanRoiGoal float64 `json:"luban_roi_goal,omitempty"`
	// RoiGoal 深度转化ROI系数, 范围(0,5]，精度：保留小数点后四位, deep_bid_type为"ROI_COEFFICIENT"时有值
	RoiGoal float64 `json:"roi_goal,omitempty"`
	// TargetCvr 目标转化率，转化率优化仅支持电商店铺推广，且付费方式为OCPC或OCPM; 允许值：0～100，0表示关闭转化率优化
	TargetCvr float64 `json:"target_cvr,omitempty"`
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
	// AutoUpdateKeyword 否开启自动加词，允许值：ON（开启）、OFF（关闭）
	AutoUpdateKeyword enum.OnOff `json:"auto_update_keyword,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
