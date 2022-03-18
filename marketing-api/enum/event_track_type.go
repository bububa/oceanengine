package enum

// EventTrackType 事件回传方式
type EventTrackType string

const (
	// EventTrackType_JSSDK JS埋码
	EventTrackType_JSSDK EventTrackType = "JSSDK"
	// EventTrackType_EXTERNAL_API API回传
	EventTrackType_EXTERNAL_API EventTrackType = "EXTERNAL_API"
	// EventTrackType_XPATH XPath圈选
	EventTrackType_XPATH EventTrackType = "XPATH"
)
