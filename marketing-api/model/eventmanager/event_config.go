package eventmanager

import "github.com/bububa/oceanengine/marketing-api/enum"

// EventConfig 事件详情
type EventConfig struct {
	// EventID 事件ID
	EventID uint64 `json:"event_id,omitempty"`
	// EventType 事件类型
	EventType string `json:"event_type,omitempty"`
	// EventCnName 事件中文名称
	EventCnName string `json:"event_cn_name,omitempty"`
	// Description 事件描述信息
	Description string `json:"description,omitempty"`
	// AttributionConfiguration 属性配置
	AttributionConfiguration *AttributionConfiguration `json:"attribution_configuration,omitempty"`
	// DebugingStatus 激活免联调状态，枚举值：Active 已激活、Inactive 未激活
	DebugingStatus string `json:"debuging_status,omitempty"`
	// StatisticalType 统计方式，枚举值：ONLY_ONE 仅一次
	StatisticalType string `json:"statistical_type,omitempty"`
	// TrackTypes 事件回传方式列表，枚举值：JSSDK JS埋码 、EXTERNAL_API API回传 、XPATH XPath圈选
	TrackTypes []enum.EventTrackType `json:"track_types,omitempty"`
	// CreateTime 事件创建时间
	CreateTime string `json:"create_time,omitempty"`
	// Properties 事件的附加属性
	Properties []EventConfigProperty `json:"properties,omitempty"`
}

// EventConfigProperty 事件的附加属性
type EventConfigProperty struct {
	// Field 附加属性英文名称
	Field string `json:"field,omitempty"`
	// FieldName 附加属性中文名称
	FieldName string `json:"field_name,omitempty"`
	// VariableType 附加属性值类型
	VariableType string `json:"variable_type,omitempty"`
	// EnumValue 附加属性枚举值
	EnumValue interface{} `json:"enum_value,omitempty"`
	// Unit 附加属性单位
	Unit string `json:"unit,omitempty"`
	// Description 附加属性描述
	Description string `json:"description,omitempty"`
}

// AttributionConfiguration  属性配置
type AttributionConfiguration struct {
	// AttributionWindow 归因窗口
	AttributionWindow int64 `json:"attribution_window,omitempty"`
}
