package coupon

// FormElementType 表单元素类型
type FormElementType string

const (
	// FormElementType_NAME 姓名
	FormElementType_NAME FormElementType = "NAME"
	// FormElementType_TELEPHONE 电话
	FormElementType_TELEPHONE FormElementType = "TELEPHONE"
	// FormElementType_EMAIL 邮箱
	FormElementType_EMAIL FormElementType = "EMAIL"
	// FormElementType_NUMBER 数值
	FormElementType_NUMBER FormElementType = "NUMBER"
	// FormElementType_SEX 性别
	FormElementType_SEX FormElementType = "SEX"
	// FormElementType_DATE 日期
	FormElementType_DATE FormElementType = "DATE"
	// FormElementType_CITY 城市
	FormElementType_CITY FormElementType = "CITY"
	// FormElementType_TEXT 文本
	FormElementType_TEXT FormElementType = "TEXT"
	// FormElementType_TEXTAREA 文本域
	FormElementType_TEXTAREA FormElementType = "TEXTAREA"
	// FormElementType_SELECT 下拉框
	FormElementType_SELECT FormElementType = "SELECT"
	// FormElementType_RADIO 单选框
	FormElementType_RADIO FormElementType = "RADIO"
	// FormElementType_CHECKBOX 多选框
	FormElementType_CHECKBOX FormElementType = "CHECKBOX"
	// FormElementType_CALCULATOR 计算器
	FormElementType_CALCULATOR FormElementType = "CALCULATOR"
)

// FormElement 表单元素
type FormElement struct {
	// ElementID 元素ID
	ElementID uint64 `json:"element_id,omitempty"`
	// Label 元素标签名称，如“姓名“等，文案长度不超过8个字符
	Label string `json:"label,omitempty"`
	// ElementType 元素类型;可选值：NAME（姓名）、TELEPHONE（电话）、EMAIL（邮箱）、NUMBER（数值）、SEX（性别）、DATE（日期）、CITY（城市）、TEXT（文本）、TEXTAREA（文本域）、SELECT（下拉框）、RADIO（单选框）、CHECKBOX（多选框）、CALCULATOR（计算器）；暂不支持SELECT、RADIO、CHECKBOX、CALCULATOR的创建
	ElementType FormElementType `json:"element_type,omitempty"`
	// Sequence 元素排序，即元素流式展示分布的先后顺序，默认按list顺序
	Sequence int `json:"sequence,omitempty"`
	// DefaultValue 元素默认值，目前已开放的元素类中，目前仅支持element_type=NUMBER，即数值类型时填写，表示默认值
	DefaultValue float64 `json:"default_value,omitempty"`
	// Layer 是否分层元素，和enable_layer，layer_submit_text结合用; 可选值：0，1（0表示不是，1表示是）
	Layer int `json:"layer,omitempty"`
	// AllowEmpty 是否允许为空，默认允许; 可选值：0，1（0表示不允许，1表示允许）
	AllowEmpty int `json:"allow_empty,omitempty"`
	// IsUnique 同一表单是否限制不可重复提交，默认可重复; 可选值：0，1（0表示可重复，1表示唯一）
	IsUnique int `json:"is_unique,omitempty"`
}
