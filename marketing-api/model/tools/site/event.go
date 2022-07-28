package site

// EventType 事件行为类型
type EventType string

const (
	// EventType_DownloadEvent DownloadEvent事件行为
	EventType_DownloadEvent EventType = "DownloadEvent"
	// EventType_LinkEvent LinkEvent事件行为描述
	EventType_LinkEvent EventType = "LinkEvent"
	// EventType_TelephoneEvent TelephoneEvent事件行为描述
	EventType_TelephoneEvent EventType = "TelephoneEvent"
)

// Event 事件行为信息
type Event interface {
	Type() EventType
}

type BaseEvent struct {
	Name string `json:"name,omitempty"`
}

func (e BaseEvent) Type() EventType {
	return EventType(e.Name)
}
