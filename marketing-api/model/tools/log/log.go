package log

// Log 日志详情
type Log struct {
	// ContentTitle 操作内容
	ContentTitle string `json:"content_title,omitempty"`
	// ObjectType 操作对象类型
	ObjectType string `json:"object_type,omitempty"`
	// ObjectID 操作对象ID
	ObjectID uint64 `json:"object_id,omitempty"`
	// ObjectName 操作对象名称
	ObjectName string `json:"object_name,omitempty"`
	// CreateTime 操作时间
	CreateTime string `json:"create_time,omitempty"`
	// ContentLog 操作前后内容
	ContentLog []string `json:"content_log,omitempty"`
	// Operator 操作人
	Operator string `json:"operator,omitempty"`
	// OptIP 操作IP
	OptIP string `json:"opt_ip,omitempty"`
}
