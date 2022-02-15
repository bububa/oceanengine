package clue

// SmartPhone 智能电话
type SmartPhone struct {
	// InstanceID 智能电话组件ID
	InstanceID uint64 `json:"instance_id,omitempty"`
	// Name 智能电话名称
	Name string `json:"name,omitempty"`
	// CallDisplay 是否为真实可接通的电话：0：表示不向商家显示用户真实手机号，而是虚拟号1：表示商家看到用户的真实号码
	CallDisplay int `json:"call_display,omitempty"`
	// PhoneID 电话id
	PhoneID uint64 `json:"phone_id,omitempty"`
	// PhoneNumber 电话号码，支持固话(如01012345678)、手机号；暂不支持400号码
	PhoneNumber string `json:"phone_number,omitempty"`
	// CreateTime 创建时间，格式：Y-m-d H:M:S
	CreateTime string `json:"create_time,omitempty"`
}
