package form

import "github.com/bububa/oceanengine/marketing-api/enum"

// FormType 表单类型
type FormType string

const (
	// NORMAL_FORM 普通落地页
	NORMAL_FORM FormType = "NORMAL_FORM"
	// ADVANCED_CREATIVE_FORM 附加创意
	ADVANCED_CREATIVE_FORM FormType = "ADVANCED_CREATIVE_FORM"
	// NATIVE_FORM 原生表单
	NATIVE_FORM FormType = "NATIVE_FORM"
)

// Form 表单详细
type Form struct {
	// InstanceID 表单ID
	InstanceID uint64 `json:"instance_id,omitempty"`
	// Name 表单名称，不超过20个字符
	Name string `json:"name,omitempty"`
	// FormType 表单类型，目前不支持设置，仅支持查询时返回; 可选值：NORMAL_FORM（普通落地页）、ADVANCED_CREATIVE_FORM（附加创意）、NATIVE_FORM（原生表单）
	FormType FormType `json:"form_type,omitempty"`
	// ValidateType 线索验证类型，主要指C端是否进行短信验证，表单中含有电话类型时生效; 可选值：CLUE_PRIORITY（线索量优先）、VALIDITY_PRIORITY（有效性优先）、NONE_VERIFICATION（全不出）、AUTO_VERIFICATION（智能验证，历史逻辑，效果等同于全不出）、ALL_VERIFICATION（全出），详情见【附录-青鸟表单线索验证类型】
	ValidateType enum.ClueValidateType `json:"validate_type,omitempty"`
	// SubmitText 提交文案说明，如“提交表单”，文案长度不超过8个字符
	SubmitText string `json:"submit_text,omitempty"`
	// EnableLayer 是否开启分层; 枚举值：0，1（0表示不启用，1表示启用）
	EnableLayer int `json:"enable_layer,omitempty"`
	// LayerSubmitText 分层提交文案说明，分层信息在Elements中，当不开启分层时，所有分层信息无效。文案长度不超过8个字符
	LayerSubmitText string `json:"layer_submit_text,omitempty"`
	// Elements 表单元素
	Elements []FormElement `json:"elements,omitempty"`
	// ExtendInfo 扩展属性
	ExtendInfo *ExtendInfo `json:"extend_info,omitempty"`
	// CreateTime 表单创建时间，格式：%Y-%m-%d %H:%M:%S
	CreateTime string `json:"create_time,omitempty"`
	// UpdateTime 表单上一次修改时间，格式：%Y-%m-%d %H:%M:%S
	UpdateTime string `json:"update_time,omitempty"`
	// IsDel 是否已删除; 可选值：0，1（0表示未删除，1表示已删除）
	IsDel int `json:"is_del,omitempty"`
}
