package promotion

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
)

// AutoGenerateConfig 配置
type AutoGenerateConfig struct {
	// ConfigID 配置ID
	ConfigID uint64 `json:"config_id,omitempty"`
	// Templates 模板列表(仅当version=Template时有值)
	Templates []ConfigTemplate `json:"templates,omitempty"`
	// UpdateTime 上次修改时间戳
	UpdateTime int64 `json:"update_time,omitempty"`
	// Version 版本：Template版本(老版本)，Strategy版本（新版本） 可选值:
	// Strategy: strategy版本
	// Template: template版本
	Version enum.AutoGenerateConfigVersion `json:"version,omitempty"`
	// StrategyData 策略配置详情列表(仅当version=Strategy时有值)
	StrategyData []StrategyData `json:"strategy_data,omitempty"`
}

// ConfigTemplate 配置模版
type ConfigTemplate struct {
	// TemplateID 模板ID
	TemplateID uint64 `json:"template_id,omitempty"`
	// TemplateType 模板类型 可选值:
	// HORIZONTAL_TO_VERTICAL_DYNAMIC_BACK_VIDEO: 横转竖动效背景视频
	// VERTICAL_TO_HORIZONTAL_DYNAMIC_BACK_VIDEO: 竖转横动效背景视频
	// VERTICAL_TO_HORIZONTAL_STATIC_BACK_VIDEO: 竖转横静态背景视频
	TemplateType enum.AutoGenerateConfigTemplateType `json:"template_type,omitempty"`
	// TemplateTextScheme 模板填充的文本内容
	TemplateTextScheme []TemplateScheme `json:"template_text_scheme,omitempty"`
	// TemplateImgScheme 模板填充的图片内容
	TemplateImgScheme []TemplateScheme `json:"template_img_scheme,omitempty"`
}

// TemplateScheme 模板填充的内容
type TemplateScheme struct {
	// Key 填充内容的键
	Key string `json:"key,omitempty"`
	// Value 填充内容的值
	Value string `json:"value,omitempty"`
}

// StrategyData 策略数据(列表中strategy_id需唯一, 即同一个策略（strategy）的策略配置项列表(strategy_state)，需放到同一个对象内，不可分开传递)
type StrategyData struct {
	// StrategyID 策略id
	StrategyID uint64 `json:"strategy_id,omitempty"`
	// StrategyType 策略类型 可选值:
	// Horizontal2Horizontal: 横转横
	// Horizontal2Vertical: 横转竖
	// Vertical2Horizontal: 竖转横
	// Vertical2Vertical: 竖转竖
	StrategyType enum.StrategyType `json:"strategy_type,omitempty"`
	// StrategyName 策略名称
	StrategyName string `json:"strategy_name,omitempty"`
	// Info 策略基本信息
	Info *StrategyInfo `json:"info,omitempty"`
	// StrategyState 策略配置项列表
	StrategyState []StrategyState `json:"strategy_state,omitempty"`
}

// StrategyInfo 策略基本信息
type StrategyInfo struct {
	// CoverURL 样例视频封面url
	CoverURL string `json:"cover_url,omitempty"`
	// VideoURL 样例视频url
	VideoURL string `json:"video_url,omitempty"`
	// Vid 样例视频vid
	Vid string `json:"vid,omitempty"`
}

// StrategyState 策略配置项列表
type StrategyState struct {
	// StateKey 配置项key
	StateKey string `json:"state_key,omitempty"`
	// StateType 配置项类型 可选值:
	// Image: 图片
	// Text: 文案
	StateType enum.StrategyStateType `json:"state_type,omitempty"`
	// StateValue 配置项值
	StateValue string `json:"state_value,omitempty"`
	// Limit 配置项限制条件
	Limit *StrategyStateLimit `json:"limite,omitempty"`
}

// StrategyStateLimit 配置项限制条件
type StrategyStateLimit struct {
	// TextMaxLength 文案最大长度(仅对state_type=Text有效)
	TextMaxLength int `json:"text_maxlength,omitempty"`
	// TextMinLength 文案最小长度(仅对state_type=Text有效)
	TextMinLength int `json:"text_min_length,omitempty"`
	// ImgWidth 图片宽度(仅对state_type=Image有效)
	ImgWidth int `json:"img_width,omitempty"`
	// ImgHeight 图片高度(仅对state_type=Image有效)
	ImgHeight int `json:"img_height,omitempty"`
}
