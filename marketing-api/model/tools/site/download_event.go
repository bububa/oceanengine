package site

// DownloadEvent DownloadEvent事件行为描述
// 特殊说明：基于应用下载安全要求，在创建和修改站点接口中，如果希望添加应用下载链接，有两种情况：①DownloadEvent类型的按钮组件可以填入安卓/IOS下载链接；②LinkEvent类型的按钮组件的URL支持落地页链接和IOS下载链接，不支持安卓下载链接，仅DownloadEvent事件下的安卓URL参数下可填写安卓下载链接；综上，除了这两种情况外其他建站组件中的URL参数都不允许填入应用下载链接，如果填入将会创建失败。
type DownloadEvent struct {
	BaseEvent
	// IOSLink ios链接信息
	IOSLink *Link `json:"ios_link,omitempty"`
	// AndroidLink 安卓链接信息
	AndroidLink *Link `json:"android_link,omitempty"`
}

// Type implement Event interface
func (e DownloadEvent) Type() EventType {
	return EventType_DownloadEvent
}
