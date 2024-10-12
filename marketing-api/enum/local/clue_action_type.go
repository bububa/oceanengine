package local

// ClueActionType 互动场景
type ClueActionType string

const (
	// ClueActionType_SHORT_VIDEO 短视频
	ClueActionType_SHORT_VIDEO ClueActionType = "SHORT_VIDEO"
	// ClueActionType_LIVE_VIDEO 直播
	ClueActionType_LIVE_VIDEO ClueActionType = "LIVE_VIDEO"
	// ClueActionType_HOME_PAGE 企业主页
	ClueActionType_HOME_PAGE ClueActionType = "HOME_PAGE"
	// ClueActionType_IM_MESSAGE 消息列表
	ClueActionType_IM_MESSAGE ClueActionType = "IM_MESSAGE"
	// ClueActionType_GROUPON_ORDER 团购tab
	ClueActionType_GROUPON_ORDER ClueActionType = "GROUPON_ORDER"
	// ClueActionType_ALIEN_CARD 异形卡
	ClueActionType_ALIEN_CARD ClueActionType = "ALIEN_CARD"
	// ClueActionType_OTHERS 其他
	ClueActionType_OTHERS ClueActionType = "OTHERS"
)
