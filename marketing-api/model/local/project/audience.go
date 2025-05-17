package project

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/local"
)

// Audience 定向设置
type Audience struct {
	// District 地域类型，允许值：
	// ALL 不限
	// LOCAL 自定义/商圈
	// POI 门店附近
	// 当delivery_content_type=POI&& 门店数量>2000家 时，不支持POI_AROUND门店附近
	// REGION 按行政区域划分
	// 使用行政区域示例：{"district":"REGION": "city":[31], "region_versio":"1.0.0"}
	District enum.District `json:"district,omitempty"`
	// Region 行政区域信息
	// 当district=REGION 行政区域划分时有效且必填
	Region *Region `json:"region,omitempty"`
	// CustomArea 自定义区域设置
	// 当district=LOCAL时必填
	CustomArea *CustomArea `json:"custom_area,omitempty"`
	// PoiAround 门店附近定向设置
	// 当district=POI门店附近时有效且必填
	PoiAround *PoiAround `json:"poi_around,omitempty"`
	// Age 年龄，允许值：
	// 年龄不可同时传入 5 分段枚举和 10 分段枚举
	// 5分段（age_5_part_set）
	// AGE_BETWEEN_18_19
	// AGE_BETWEEN_20_23
	// AGE_BETWEEN_24_30
	// AGE_BETWEEN_31_35
	// AGE_BETWEEN_36_40
	// AGE_BETWEEN_41_45
	// AGE_BETWEEN_46_50
	// AGE_BETWEEN_51_55
	// AGE_BETWEEN_56_59
	// AGE_ABOVE_60
	// 10分段（age_10_part_set）
	// AGE_BETWEEN_18_23
	// AGE_BETWEEN_24_30
	// AGE_BETWEEN_31_40
	// AGE_BETWEEN_41_49
	// AGE_ABOVE_50
	Age enum.AudienceAge `json:"age,omitempty"`
	// Gender 性别，不传默认不限，允许值：
	// FEMALE 女
	// MALE 男
	// NONE 不限
	Gender local.Gender `json:"gender,omitempty"`
	// RetargetingTags 定向人群包ID，可通过【查询本地推创编可用人群包】接口获取
	// 数量上限：200
	RetargetingTags []uint64 `json:"retargeting_tags,omitempty"`
	// RetargetingTagsExclude 排除人群包ID，可通过【查询本地推创编可用人群包】接口获取
	// 数量上限：200
	RetargetingTagsExclude []uint64 `json:"retargeting_tags_exclude,omitempty"`
	// HideIfConverted 过滤已转化的维度，可以避免该广告再次投放给已转化过的用户，不传默认不过滤，允许值：
	// NOT_EXCLUDE 不过滤
	// ADVERTISER 广告主账户
	// CC 组织账户
	// CUSTOMER 公司账户
	// PROJECT 项目
	// PROMOTION 广告
	// 当external_action=SHOW 展示量时，不支持该字段
	HideIfConverted enum.HideIfConverted `json:"hide_if_converted,omitempty"`
	// ConvertedTimeDuration 过滤已转化投放时间，允许值：
	// SEVEN_DAY 七天（默认值）
	// ONE_MONTH 1个月
	// THREE_MONTH 3个月
	// SIX_MONTH 6个月
	// TWELVE_MONTH 12个月
	// TODAY 当天
	// 仅当hide_if_converted设置为CUSTOMER或ORGANIZATION，该字段有效，不传默认7天，其余场景该字段无效。
	ConvertedTimeDuration enum.ConvertedTimeDuration `json:"converted_time_duration,omitempty"`
}

// Region 行政区域信息
type Region struct {
	// City 地域定向省市或者区县列表(当传递省份ID时，旗下市县ID可省略不传)，通过【获取行政信息】接口获取
	City []uint64 `json:"city,omitempty"`
	// CityDivide 城市划分类型，允许值：
	// BY_LEVEL 按发展等级划分
	// BY_LOCATION 按地理划分（默认值）
	CityDivide local.CityDivide `json:"city_divide,omitempty"`
	// LocationType 区域内人群定向设置，允许值：
	// ALL 该地区的所有用户
	// CURRENT 正在该地区的用户（默认值）
	// HOME 居住在该地区的用户
	// TRAVEL 到该地区旅行的用户
	LocationType local.LocationType `json:"location_type,omitempty"`
	// RegionVer 行政区域版本号，通过【获取行政信息】接口获取
	RegionVer string `json:"region_ver,omitempty"`
}

// CustomArea 自定义区域设置
type CustomArea struct {
	// Geolocation 地理位置设置，数量上限：1000
	Geolocation []Geolocation `json:"geolocation,omitempty"`
}

// Geolocation 地理位置信息
type Geolocation struct {
	// Name 地点名称
	Name string `json:"name,omitempty"`
	// AreaRadius 半径
	AreaRadius int64 `json:"area_radius,omitempty"`
	// Long 经度
	Long float64 `json:"long,omitempty"`
	// Lat 纬度
	Lat float64 `json:"lat,omitempty"`
}

// PoiAround 门店附近定向设置
type PoiAround struct {
	// PoiIDs 定向门店id列表，可通过【获取门店列表】接口查询
	// 数量上限：2000
	// 「短视频推门店」无需传入，自动定位至推广的门店附近
	// 填写说明：
	// marketing_goal=LIVE直播 && district =POI_AROUND门店附近，该字段有效且必填
	// 门店id要求：传入门店id需为本广告账户下的门店ID
	// marketing_goal=VIDEO_AND_IMAGE短视频/图文，且delivery_content_type=PRODUCT为商品，且district =POI_AROUND时，该字段有效。
	// 门店id要求：传入门店id必须为推广的商品适用门店
	// 所推广的商品适用门店数大于2000时必填，且最多传入2000个；
	// 所推广的商品适用门店数不超过2000时，不传默认推商品适用门店附近
	// marketing_goal=VIDEO_AND_IMAGE短视频/图文 && delivery_content_type=POI门店时，该字段无效，无需传入。选择门店附近定向，则默认投所推广的门店附近（仅推广门店数不超过2000家时可选择门店附近定向）
	PoiIDs []uint64 `json:"poi_ids,omitempty"`
	// PoiAroundRadius 半径，门店附近仅支持固定半径的设置，允许值：
	// KM_6 门店附近6km
	// KM_8 门店附近8km
	// KM_10 门店附近10km（默认值）
	// KM_12 门店附近12km
	// KM_15 门店附近15km
	// KM_20 门店附近20km
	PoiAroundRadius local.PoiAroundRadius `json:"poi_around_radius,omitempty"`
}
