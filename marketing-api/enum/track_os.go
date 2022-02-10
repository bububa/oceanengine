package enum

// TrackOS  转化追踪客户端系统
type TrackOS = int

const (
	// Track_ANDROID Android
	Track_ANDROID TrackOS = 0
	// Track_IOS iOS
	Track_IOS TrackOS = 1
	// Track_UNKNOWN 未知设备系统
	Track_UNKNOWN TrackOS = 3
)
