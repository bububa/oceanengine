package local

// ExternalAction 优化目标
type ExternalAction string

const (
	// ExternalAction_NATIVE_ACTION 用户互动
	ExternalAction_NATIVE_ACTION ExternalAction = "NATIVE_ACTION"
	// ExternalAction_FOLLOW_ACTION 粉丝增长
	ExternalAction_FOLLOW_ACTION ExternalAction = "FOLLOW_ACTION"
	// ExternalAction_SHOW 展示量
	ExternalAction_SHOW ExternalAction = "SHOW"
	// ExternalAction_OTO_PAY 团购购买
	ExternalAction_OTO_PAY ExternalAction = "OTO_PAY"
	// ExternalAction_POI_RECOMMEND 门店引流
	ExternalAction_POI_RECOMMEND ExternalAction = "POI_RECOMMEND"
	// ExternalAction_LIVE_OTO_CLICK 商品点击
	ExternalAction_LIVE_OTO_CLICK ExternalAction = "LIVE_OTO_CLICK"
	// ExternalAction_LIVE_OTO_GROUP_BUYING 直播间团购购买
	ExternalAction_LIVE_OTO_GROUP_BUYING ExternalAction = "LIVE_OTO_GROUP_BUYING"
	// ExternalAction_LIVE_ENGAGE 直播加热
	ExternalAction_LIVE_ENGAGE ExternalAction = "LIVE_ENGAGE"
	// ExternalAction_LIVE_ENGAGEMENT 直播加热
	ExternalAction_LIVE_ENGAGEMENT ExternalAction = "LIVE_ENGAGEMENT"
	// ExternalAction_FOLLOWER_COUNT 粉丝增长
	ExternalAction_FOLLOWER_COUNT ExternalAction = "FOLLOWER_COUNT"
	// ExternalAction_LIVE_ENTER_ACTION 直播间观看
	ExternalAction_LIVE_ENTER_ACTION ExternalAction = "LIVE_ENTER_ACTION"
	// ExternalAction_LIVE_STAY_TIME 直播间停留
	ExternalAction_LIVE_STAY_TIME ExternalAction = "LIVE_STAY_TIME"
)
