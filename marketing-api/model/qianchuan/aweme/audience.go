package aweme

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// Audience 定向设置
type Audience struct {
	// AudienceMode 想吸引的观众类型，可选值:
	// ATUO: 系统智能推荐
	// CUSTOM: 自定义
	// FANS: 达人相似粉丝
	// LIVEFANS: 达人相似观众
	AudienceMode enum.AudienceMode `json:"audience_mode,omitempty"`
	// Gender 性别，可选值:
	// GENDER_FEMALE: 女性
	// GENDER_MALE: 男性
	// 若不传值，则默认为不限
	// 注意：仅想吸引的观众类型为“自定义”时，有效
	Gender enum.AudienceGender `json:"gender,omitempty"`
	// Age 年龄，允许值：若不传值，则默认为不限
	// AGE_BETWEEN_18_23：18-23岁、
	// AGE_BETWEEN_24_30：24-30岁、
	// AGE_BETWEEN_31_40：31-40岁、
	// AGE_BETWEEN_41_50：41-50岁
	// 注意：仅想吸引的观众类型为“自定义”时，有效
	Age enum.AudienceAge `json:"age,omitempty"`
	// ExcludeLimitedRegion 排除限运地区，允许值：
	// 0 ：否，默认值；1 ：是
	// 注意：
	// 1、当营销目标为VIDEO_PROM_GOODS时，audience_mode为“ATUO”&“FANS”时，有效
	// 2、当营销目标为LIVE_PROM_GOODS时，audience_mode为“ATUO”&“LIVEFANS”时，有效
	ExcludeLimitedRegion int `json:"exclude_limited_region,omitempty"`
	// District 地域定向类型，配合 city 字段使用，允许值：
	// 若不传值，则默认为不限
	// CITY ：省市
	// COUNTY ：区县
	// 注意：仅营销目标为VIDEO_PROM_GOODS且想吸引的观众类型为“自定义”时，有效
	District enum.District `json:"district,omitempty"`
	// City 具体定向的城市列表，当 district 为COUNTY、CITY为必填。
	// 省市的传法："city" : [12], "district" : "CITY"，不能超过200个省
	// 区县的传法："city" : [130102], "district" : "COUNTY"，不能超过100和区县
	// 注意：仅营销目标为VIDEO_PROM_GOODS且想吸引的观众类型为“自定义”时，有效
	City []uint64 `json:"city,omitempty"`
	// InterestCategories 兴趣类目词，可以通过【随心推投放-获取随心推兴趣标签】接口获取
	// 注意：仅营销目标为VIDEO_PROM_GOODS且想吸引的观众类型为“自定义”时，有效
	InterestCategories []uint64 `json:"interest_categories,omitempty"`
	// AuthorIDs 抖音达人ID列表，可通过【工具-抖音达人-查询抖音类目下的推荐达人】接口获取
	// 注意：限制20个
	// 仅想吸引的观众类型为“达人相似观众”和“达人相似粉丝”时，必填
	AuthorIDs []uint64 `json:"author_ids,omitempty"`
	// Behaviors 抖音用户行为类型 ，允许值：
	// 1、当吸引的观众类型为“达人相似观众”且营销目标为LIVE_PROM_GOODS时，允许值
	// ALL全部，默认
	// FOLLOWED_USER 关注过此人
	// LIVE_WATCH 观看过直播
	// GOODS_SHARE 下单过商品
	// 2、当吸引的观众类型为“达人相似粉丝”且营销目标为VIDEO_PROM_GOODS时，允许值
	// ALL全部，默认
	// FOLLOWED_USER 关注过达人
	// LIKED_USER 点赞过视频
	// GOODS_SHARE 下单过商品
	Behaviors []enum.Behavior `json:"behaviors,omitempty"`
	// AwemeFanAccounts 抖音达人ID列表
	AwemeFanAccounts model.Uint64s `json:"aweme_fan_accounts,omitempty"`
	// AwemeFanBehaviors 抖音用户行为类型
	AwemeFanBehaviors []enum.Behavior `json:"aweme_fan_behaviors,omitempty"`
}
