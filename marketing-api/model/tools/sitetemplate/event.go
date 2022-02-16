package sitetemplate

// EventType 按钮事件类型
type EventType string

const (
	// EventType_APPOINT_EVENT 预约事件行为
	EventType_APPOINT_EVENT EventType = "APPOINT_EVENT"
	// EventType_DOWNLOAD_EVENT 下载事件行为
	EventType_DOWNLOAD_EVENT EventType = "DOWNLOAD_EVENT"
	// EventType_LINK_EVENT 链接事件行为
	EventType_LINK_EVENT EventType = "LINK_EVENT"
	// EventType_TELEPHONE_EVENT 电话事件行为
	EventType_TELEPHONE_EVENT EventType = "TELEPHONE_EVENT"
)

// DownloadEvent downloadEvent事件行为描述
type DownloadEvent struct {
	// IOSLink ios链接信息
	IOSLink *Link `json:"ios_link,omitempty"`
	// AndroidLink 安卓链接信息
	AndroidLink *Link `json:"android_link,omitempty"`
}

// LinkEvent LinkEvent事件行为描述
type LinkEvent struct {
	// Link 链接信息
	Link *Link `json:"link,omitempty"`
}

// PhoneEvent phoneEvent事件行为描述
type PhoneEvent struct {
	// InstanceID 智能电话ID，当phone_event不为空时，有返回值。用户可以通过【获取智能电话列表】接口或【青鸟线索通平台】获取智能电话ID
	InstanceID uint64 `json:"instance_id,omitempty"`
}

// AppointEvent appointEvent事件行为描述
type AppointEvent struct {
	// GameID 游戏站ID
	GameID string `json:"game_id,omitempty"`
	// Link 链接信息
	Link *Link `json:"link,omitempty"`
}
