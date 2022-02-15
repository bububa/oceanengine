package smartphone

import "github.com/bububa/oceanengine/marketing-api/enum"

// SmartPhone 智能电话
type SmartPhone struct {
	// CallDisplay 是否显示真实手机号
	// 1：显示真实手机号
	// 0：显示虚拟号（默认）
	CallDisplay int `json:"call_display,omitempty"`
	// CreateTime 表单创建时间，格式：%Y-%m-%d %H:%M:%S
	CreateTime string `json:"create_time,omitempty"`
	// InstanceID 生成的智能电话ID
	InstanceID uint64 `json:"instance_id,omitempty"`
	// Name 智能电话名称，长度不超过20
	Name string `json:"name,omitempty"`
	// NoDisturbTimes 免打扰开启时间段(最多5组，不填或者空数组代表不开启免打扰)
	// 所选时间内，用户通过广告页所拨电话将由语音机器人应答，商家将不会收到来电
	// 非工作时段所有来电及录音可在飞鱼查看，并支持回拨
	// 触发非工作模式同样记转化
	NoDisturbTimes []TimeRange `json:"no_disturb_times,omitempty"`
	// PhoneNumber 电话号码
	// 支持固话(如01012345678)、手机号
	// 暂不支持400号码绑定
	PhoneNumber string `json:"phone_number,omitempty"`
	// ValidateType 对用户手机号的验证类型，可选值:
	// CLUE_PRIORITY：优先线索量。对系统判定异常电话进行语音验证，未通过验证的用户线索仍然获取
	// VALIDITY_PRIORITY： 优先有效性。对系统判定异常电话进行语音验证，只获取通过验证的用户线索
	// NONE_VERIFICATION：不进行验证。无论什么类型用户，均不进行语音验证
	ValidateType enum.ClueValidateType `json:"validate_type,omitempty"`
}
