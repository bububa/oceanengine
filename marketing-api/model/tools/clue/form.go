package clue

// 表单类型
type FormType string

const (
	// NORMAL_FORM 普通表单
	NORMAL_FORM FormType = "NORMAL_FORM"
	// ADVANCED_CREATIVE_FORM 附加创意表单
	ADVANCED_CREATIVE_FORM FormType = "ADVANCED_CREATIVE_FORM"
)

// Form 表单信息
type Form struct {
	// AdvID 广告主ID
	AdvID uint64 `json:"adv_id,omitempty"`
	// InstanceID 表单ID
	InstanceID uint64 `json:"instance_id,omitempty"`
	// Name 表单名称
	Name string `json:"name,omitempty"`
	// Version 版本信息
	Version int `json:"version,omitempty"`
	// IsDel 是否已删除，1（删除）0（未删除）
	IsDel int `json:"is_del,omitempty"`
	// SubType 线索表单类型，返回值：
	// NORMAL_FORM（普通表单）
	// ADVANCED_CREATIVE_FORM（附加创意表单）
	SubType FormType `json:"sub_type,omitempty"`
	// EnableLayer 是否未分层表单，1（是）、0（否）
	EnableLayer int `json:"enable_layer,omitempty"`
	// FormType 表单类型，返回值：
	// NORMAL_FORM（普通表单）
	// ADVANCED_CREATIVE_FORM（附加创意表单）
	FormType FormType `json:"form_type,omitempty"`
	// ContainPhone 表单内是否包含电话元素; 1（是）、0（否）
	ContainPhone int `json:"contain_phone,omitempty"`
	// LightingPageURL 轻落地页URL，仅附加创意表单中存在
	LigntingPageURL string `json:"lighting_page_url,omitempty"`
	// CreateTime 创建时间，格式：Y-m-d H:M:S
	CreateTime string `json:"create_time,omitempty"`
}
