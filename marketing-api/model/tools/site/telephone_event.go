package site

// TelephoneEvent TelephoneEvent事件行为描述
type TelephoneEvent struct {
	BaseEvent
	// SmartPhone 智能电话信息
	SmartPhone *SmartPhone `json:"smart_phone,omitempty"`
}

// Type implement Event interface
func (e TelephoneEvent) Type() EventType {
	return EventType_TelephoneEvent
}

// SmartPhone 智能电话信息
type SmartPhone struct {
	// InstanceID 表单ID，如果传入instance_id，以此为准，其他智能电话数据不再读取
	InstanceID uint64 `json:"instance_id,omitempty"`
	// Name 新建智能电话的名称，无instance_id时， 必填
	Name string `json:"name,omitempty"`
	// PhoneNumber 电话号码，无instance_id时， 必填
	PhoneNumber string `json:"phone_number,omitempty"`
	// CallDisplay 是否展示真实号码，默认现实中间虚拟号码。
	// 默认值：0（中间虚拟号码）
	// 允许值：0（中间虚拟号码）、 1（用户真实号码）
	// 【注】显示用户真实号码可能会被运营商拦截，发生拨打异常情况时可选择来电展示中间虚拟号码。
	CallDisplay int `json:"call_display,omitempty"`
}
