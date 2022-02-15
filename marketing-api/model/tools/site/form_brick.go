package site

// FormType 表单元素类型
type FormType string

const (
	// FormType_NAME 姓名
	FormType_NAME FormType = "NAME"
	// FormType_TELEPHONE 电话
	FormType_TELEPHONE FormType = "TELEPHONE"
	// FormType_EMAIL 邮箱
	FormType_EMAIL FormType = "EMAIL"
	// FormType_NUMBER 数值
	FormType_NUMBER FormType = "NUMBER"
	// FormType_SEX 性别
	FormType_SEX FormType = "SEX"
	// FormType_DATE 日期
	FormType_DATE FormType = "DATE"
	// FormType_CITY 城市
	FormType_CITY FormType = "CITY"
	// FormType_TEXT 文本
	FormType_TEXT FormType = "TEXT"
	// FormType_TEXTAREA 文本域
	FormType_TEXTAREA FormType = "TEXTAREA"
)

// FormBrick 表单组件
type FormBrick struct {
	BaseBrick
	// FormData 表单数据
	FormData *FormData `json:"form_data,omitempty"`
}

// Type implement Brick interface
func (b FormBrick) Type() BrickType {
	return XrForm
}

// FormData 表单数据
type FormData struct {
	// InstanceID 已有表单，如果传入instance_id，以此为准，其他表单数据不再读取; 当传表单数据时，submit_text和name为空时必填
	InstanceID uint64 `json:"instance_id,omitempty"`
	// SubmitText 提交文案;当传表单数据时，无instance_id时， 必填
	SubmitText string `json:"submit_text,omitempty"`
	// Name 表单名称;当传表单数据时，无instance_id时， 必填
	Name string `json:"name,omitempty"`
	// Elements 表单元素
	Elements []FormElement `json:"elements,omitempty"`
	// SuccessLink 用户提交成功跳转链接
	SuccessLink *Link `json:"success_link,omitempty"`
	// FailureLink 提交失败跳转链接
	FailureLink *Link `json:"failure_link,omitempty"`
	// Linkable 表单提交后是否跳转; 允许值：0（不跳转），1（跳转）
	Linkable int `json:"linkable"`
	// SuccessTip 用户提交成功提示语表示用户填写表单成功后以小窗方式弹出的提示信息，比如告诉用户24小时后商家会主动联系用户，长度不超过128，默认值“表单提交成功”
	SuccessTip string `json:"success_tip,omitempty"`
}

// FormElement 表单元素
type FormElement struct {
	// Label 元素标签
	Label string `json:"label,omitempty"`
	// AllowEmpty 是否允许为空 允许值：0（必填），1（可为空）
	AllowEmpty int `json:"allow_empty,omitempty"`
	// Type 表单元素类型
	// 允许值：
	// NAME（姓名）
	// TELEPHONE（电话）
	// EMAIL（邮箱）
	// NUMBER（数值）
	// SEX（性别）
	// DATE（日期）
	// CITY（城市）
	// TEXT（文本）
	// TEXTAREA（文本域)
	// 当allow_empty为0时必填
	Type FormType `json:"type,omitempty"`
	// IsUnique 是否可重复; 允许值：0（唯一），1（可重复）
	IsUnique int `json:"is_unique"`
}
