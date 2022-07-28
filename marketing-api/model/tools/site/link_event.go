package site

// LinkEvent LinkEvent事件行为描述
type LinkEvent struct {
	BaseEvent
	// Link 链接信息（支持落地页链接和IOS下载链接，不支持安卓下载链接，仅DownloadEvent事件下的安卓URL参数下可填写安卓下载链接）
	Link *Link `json:"link,omitempty"`
}

// Type implement Event interface
func (e LinkEvent) Type() EventType {
	return EventType_LinkEvent
}
