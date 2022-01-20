package enum

// UniversalSettlementType 星图任务结算方式
type UniversalSettlementType string

const (
	// UniversalSettlementType_ALL 不限
	UniversalSettlementType_ALL UniversalSettlementType = "ALL"
	// UniversalSettlementType_CPM 按播放量计费
	UniversalSettlementType_CPM UniversalSettlementType = "CPM"
	// UniversalSettlementType_FIXED_PRICE 一口价
	UniversalSettlementType_FIXED_PRICE UniversalSettlementType = "FIXED_PRICE"
	// UniversalSettlementType_EXCHANGE 资源置换
	UniversalSettlementType_EXCHANGE UniversalSettlementType = "EXCHANGE"
	// UniversalSettlementType_CPA 按转化付费
	UniversalSettlementType_CPA UniversalSettlementType = "CPA"
	// UniversalSettlementType_RANK 按视频等级结算
	UniversalSettlementType_RANK UniversalSettlementType = "RANK"
	// UniversalSettlementType_EFFECT 按效果结算
	UniversalSettlementType_EFFECT UniversalSettlementType = "EFFECT"
	// UniversalSettlementType_GIFT 礼品奖励
	UniversalSettlementType_GIFT UniversalSettlementType = "GIFT"
	// UniversalSettlementType_MONEY_SHARE 现金奖励
	UniversalSettlementType_MONEY_SHARE UniversalSettlementType = "MONEY_SHARE"
	// UniversalSettlementType_FLOW_SHARE 流量奖励
	UniversalSettlementType_FLOW_SHARE UniversalSettlementType = "FLOW_SHARE"
	// UniversalSettlementType_STAR_SUPPORT 应援奖励
	UniversalSettlementType_STAR_SUPPORT UniversalSettlementType = "STAR_SUPPORT"
	// UniversalSettlementType_DOU_PLUS DOU+奖励
	UniversalSettlementType_DOU_PLUS UniversalSettlementType = "DOU_PLUS"
	// UniversalSettlementType_CUSTOMIZE 自定义结算
	UniversalSettlementType_CUSTOMIZE UniversalSettlementType = "CUSTOMIZE"
)
