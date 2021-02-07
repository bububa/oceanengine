package enum

// 分组条件
type StatGroupBy string

const (
	STAT_GROUP_BY_FIELD_ID               StatGroupBy = "STAT_GROUP_BY_FIELD_ID"               // ID 类型-按照 ID 分组
	STAT_GROUP_BY_ADVERTISER_ID          StatGroupBy = "STAT_GROUP_BY_ADVERTISER_ID"          // ID 类型-按照广告主分组
	STAT_GROUP_BY_CAMPAIGN_ID            StatGroupBy = "STAT_GROUP_BY_CAMPAIGN_ID"            // ID 类型-按照广告组分组
	STAT_GROUP_BY_AD_ID                  StatGroupBy = "STAT_GROUP_BY_AD_ID"                  // ID 类型-按照广告计划分组
	STAT_GROUP_BY_CREATIVE_ID            StatGroupBy = "STAT_GROUP_BY_CREATIVE_ID"            // ID 类型-按照创意分组
	STAT_GROUP_BY_FIELD_STAT_TIME        StatGroupBy = "STAT_GROUP_BY_FIELD_STAT_TIME"        // 时间类型-按照时间分组
	STAT_GROUP_BY_TIME_MONTH             StatGroupBy = "STAT_GROUP_BY_TIME_MONTH"             // 时间类型-按照月分组
	STAT_GROUP_BY_TIME_WEEK              StatGroupBy = "STAT_GROUP_BY_TIME_WEEK"              // 时间类型-按照周分组
	STAT_GROUP_BY_TIME_DAY               StatGroupBy = "STAT_GROUP_BY_TIME_DAY"               // 时间类型-按照天分组
	STAT_GROUP_BY_TIME_HOUR              StatGroupBy = "STAT_GROUP_BY_TIME_HOUR"              // 时间类型-按照小时分组
	STAT_GROUP_BY_PRICING                StatGroupBy = "STAT_GROUP_BY_PRICING"                // 细分基础类型-按照出价类型分组
	STAT_GROUP_BY_IMAGE_MODE             StatGroupBy = "STAT_GROUP_BY_IMAGE_MODE"             // 细分基础类型-按照素材类型分组
	STAT_GROUP_BY_INVENTORY              StatGroupBy = "STAT_GROUP_BY_INVENTORY"              // 细分基础类型-按照投放位置分组
	STAT_GROUP_BY_CAMPAIGN_TYPE          StatGroupBy = "STAT_GROUP_BY_CAMPAIGN_TYPE"          // 细分基础类型-按照广告组类型分组
	STAT_GROUP_BY_CREATIVE_MATERIAL_MODE StatGroupBy = "STAT_GROUP_BY_CREATIVE_MATERIAL_MODE" // 细分基础类型-按照创意类型分组
	STAT_GROUP_BY_EXTERNAL_ACTION        StatGroupBy = "STAT_GROUP_BY_EXTERNAL_ACTION"        // 细分基础类型-按照转化类型分组
	STAT_GROUP_BY_LANDING_TYPE           StatGroupBy = "STAT_GROUP_BY_LANDING_TYPE"           // 细分基础类型-按照推广类型分组
	STAT_GROUP_BY_PRICING_CATEGORY       StatGroupBy = "STAT_GROUP_BY_PRICING_CATEGORY"       // 细分基础类型-按照广告类型分组
	STAT_GROUP_BY_PROVINCE_NAME          StatGroupBy = "STAT_GROUP_BY_PROVINCE_NAME"          // 细分受众类型-按照省份分组
	STAT_GROUP_BY_CITY_NAME              StatGroupBy = "STAT_GROUP_BY_CITY_NAME"              // 细分受众类型-按照城市分组
	STAT_GROUP_BY_GENDER                 StatGroupBy = "STAT_GROUP_BY_GENDER"                 // 细分受众类型-按照性别分组
	STAT_GROUP_BY_AGE                    StatGroupBy = "STAT_GROUP_BY_AGE"                    // 细分受众类型-按照年龄分组
	STAT_GROUP_BY_PLATFORM               StatGroupBy = "STAT_GROUP_BY_PLATFORM"               // 细分受众类型-按照平台分组
	STAT_GROUP_BY_AC                     StatGroupBy = "STAT_GROUP_BY_AC"                     // 细分受众类型-按照网络类型分组
	STAT_GROUP_BY_AD_TAG                 StatGroupBy = "STAT_GROUP_BY_AD_TAG"                 // 细分受众类型-按照兴趣分类分组
	STAT_GROUP_BY_INTEREST_TAG           StatGroupBy = "STAT_GROUP_BY_INTEREST_TAG"           // 细分受众类型-按照兴趣标签分组
	STAT_GROUP_BY_BIDWORD_ID             StatGroupBy = "STAT_GROUP_BY_BIDWORD_ID"             // 关键词词类型-按照关键词分组
	STAT_GROUP_BY_QUERY                  StatGroupBy = "STAT_GROUP_BY_QUERY"                  // 关键词词类型-按照搜索词分组
	STAT_GROUP_BY_PRODUCT_ID             StatGroupBy = "STAT_GROUP_BY_PRODUCT_ID"             // DPA 类型-按照商品 ID 分组
	STAT_GROUP_BY_PRODUCT_PLATFORM_ID    StatGroupBy = "STAT_GROUP_BY_PRODUCT_PLATFORM_ID"    // DPA 类型-按照商品库 ID 分组
	STAT_GROUP_BY_MATERIAL_ID            StatGroupBy = "STAT_GROUP_BY_MATERIAL_ID"            // 素材类型-按照素材 ID 分组
	STAT_GROUP_BY_PLAYABLE_ID            StatGroupBy = "STAT_GROUP_BY_PLAYABLE_ID"            // 素材类型-按照试玩素材 ID 分组
)
