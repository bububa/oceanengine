package tools

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util/query"
)

// EstimateAudienceRequest 查询受众预估结果 API Request
type EstimateAudienceRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// RetargetingType 定向人群包类型,详见【附录-定向人群包类型】,即将下线
	// 允许值: "RETARGETING_INCLUDE", "RETARGETING_EXCLUDE","RETARGETING_NONE"
	RetargetingType enum.RetargetingType `json:"retargeting_type,omitempty"`
	// RetargetingTags 当选择使用人群包定向时填写，内容为人群包id，即将下线
	RetargetingTags []uint64 `json:"retargeting_tags,omitempty"`
	// RetargetingTagsInclude 定向人群包列表（自定义人群），内容为人群包id。如果选择"同时定向与排除"，需传入retargeting_tags_include和retargeting_tags_exclude
	RetargetingTagsInclude []uint64 `json:"retargeting_tags_include,omitempty"`
	// RetargetingTagsExclude 排除人群包列表（自定义人群），内容为人群包id。如果选择"同时定向与排除"，需传入retargeting_tags_include和retargeting_tags_exclude
	RetargetingTagsExclude []uint64 `json:"retargeting_tags_exclude,omitempty"`
	// Gender 性别
	Gender enum.AudienceGender `json:"gender,omitempty"`
	// Age 年龄
	Age []enum.AudienceAge `json:"age,omitempty"`
	// AndroidOsv 最低安卓版本，当app_type为"APP_ANDROID"选填,其余情况不填,
	AndroidOsv string `json:"android_osv,omitempty"`
	// IosOsv 最低IOS版本，当app_type为"APP_IOS"选填,其余情况不填
	IosOsv string `json:"ios_osv,omitempty"`
	// Ac 网络类型
	Ac []string `json:"ac,omitempty"`
	// Carrier 运营商
	Carrier []enum.Carrier `json:"carrier,omitempty"`
	// DeviceBrand 手机品牌
	DeviceBrand []string `json:"device_brand,omitempty"`
	// ArticleCategory 文章分类
	ArticleCategory []string `json:"article_category,omitempty"`
	// ActivateType 用户首次激活时间, 详见【附录-用户首次激活时间】
	// 允许值:"WITH_IN_A_MONTH","ONE_MONTH_2_THREE_MONTH","THREE_MONTH_EAILIER"
	ActivateType enum.ActivateType `json:"activate_type,omitempty"`
	// Platform 平台，当下载方式包含下载链接时，平台类型需与选择的下载链接类型对应，当下载方式不包含下载方式的时候，平台可多选。为保证投放效果,平台类型定向PC与移动端互斥
	Platform []enum.AudiencePlatform `json:"platform,omitempty"`
	// District 地域;取值: "CITY"省市, "COUNTY"区县, "BUSINESS_DISTRICT"商圈,"NONE"不限，省市传法："city": [12],"district": "CITY",区县的传法："city": [130102],"district": "COUNTY";暂不支持"海外"
	District enum.District `json:"district,omitempty"`
	// City 地域定向省市或者区县列表(当传递省份ID时,旗下市县ID可省略不传),当district为"CITY"或"COUNTY"时有值
	City []uint64 `json:"city,omitempty"`
	// BusinessIDs 商圈ID数组，district为"BUSINESS_DISTRICT"时有值
	BusinessIDs []uint64 `json:"business_ids,omitempty"`
	// LocationType 位置类型;取值：CURRENT正在该地区的用户，HOME居住在该地区的用户，TRAVEL到该地区旅行的用户，ALL该地区内的所有用户;当city和district有值时返回值
	LocationType enum.LocationType `json:"location_type,omitempty"`
	// AdTags （老版行为兴趣）兴趣分类,如果传"空数组"表示不限，如果"数组传0"表示系统推荐,如果按兴趣类型传表示自定义
	AdTags []uint64 `json:"ad_tags,omitempty"`
	// InterestTags （老版行为兴趣）兴趣关键词, 传入具体的词id，非兴趣词包id，可以通过词包相关接口或者兴趣关键词word2id接口获取词id，一个计划下最多创建1000个关键词。
	InterestTags []uint64 `json:"interest_tags,omitempty"`
	// AppBehaviorTarget （老版行为兴趣）APP行为; 取值：NONE不限，CATEGORY按分类，APP按APP
	AppBehaviorTarget string `json:"app_behavior_target,omitempty"`
	// AppCategory 老版行为兴趣）APP行为定向——按分类
	AppCategory []uint64 `json:"app_category,omitempty"`
	// AppIDs （老版行为兴趣）APP行为定向——按APP（请注意如果投放的是"应用下载-IOS"不支持设置APP行为定向，请勿传值。）可通过【工具-查询工具-查询应用信息】获取。当app_behavior_target为APP时有值
	AppIDs []uint64 `json:"app_ids,omitempty"`
	// SuperiorPopularityType 媒体定向;
	SuperiorPopularityType enum.SuperiorPopularityType `json:"superior_popularity_type,omitempty"`
	// FlowPackage 定向逻辑
	FlowPackage []uint64 `json:"flow_package,omitempty"`
	// ExcludeFlowPackage 排除定向逻辑
	ExcludeFlowPackage []uint64 `json:"exclude_flow_package,omitempty"`
	// IncludeCustomActions 包含人群包((DPA推广目的特有,格式举例[{"days": 7, "code": 1001},]， dpa_local_audience为1时有值; day可选范围:1, 7, 14, 28, 60, 90, 120, 180。
	IncludeCustomActions []interface{} `json:"include_custom_actions,omitempty"`
	// ExcludeCustomActions 排除人群包((DPA推广目的特有,格式举例{"days": 7, "code": 1002},]，day可选范围: 1, 7, 14, 28, 60, 90, 120, 180。
	ExcludeCustomActions []interface{} `json:"exclude_custom_actions,omitempty"`
}

// Encode implement GetRequest interface
func (r EstimateAudienceRequest) Encode() string {
	values, _ := query.Values(r)
	return values.Encode()
}

// EstimateAudienceResponse 查询受众预估结果 API Response
type EstimateAudienceResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *EstimateAudienceResponseData `json:"data,omitempty"`
}

// EstimateAudienceResponseData 受众预估结果
type EstimateAudienceResponseData struct {
	// TouTiao 今日头条预估用户覆盖量结果
	TouTiao EstimateAudienceResult `json:"toutiao,omitempty"`
	// Aweme 抖音视频预估用户覆盖量结果
	Aweme EstimateAudienceResult `json:"aweme,omitempty"`
	// VideoApp 西瓜视频预估用户覆盖量结果
	VideoApp EstimateAudienceResult `json:"video_app,omitempty"`
	// Hotsoon 火山视频预估用户覆盖量结果
	Hotsoon EstimateAudienceResult `json:"hotsoon,omitempty"`
}

// EstimateAudienceResult
type EstimateAudienceResult struct {
	// Num 覆盖量
	Num int64 `json:"num,omitempty"`
}
