package diagnosis

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// AdSuggestion 诊断建议
type AdSuggestion struct {
	// AdID 计划id
	AdID model.Uint64 `json:"ad_id,omitempty"`
	// SceneList 计划对应的场景列表
	SceneList []SuggestScene `json:"scene_list,omitempty"`
}

// SuggestScene 计划对应的场景
type SuggestScene struct {
	// Scene 场景名称，允许值：CLEAN清理低质计划场景、POTENTIAL获取潜力计划场景, ZOMBIE 活跃度低广告（暂为白名单功能，如需获取请联系运营或销售）
	Scene string `json:"scene,omitempty"`
	// Suggestions 建议列表
	Suggestions []Suggestion `json:"suggestions,omitempty"`
}

// Suggestion 建议
type Suggestion struct {
	// Conclusion 诊断结论，当scene=POTENTIAL时，该字段为空
	Conclusion string `json:"conculsion,omitempty"`
	// Msg 该场景下所有建议的详细描述
	Msg string `json:"msg,omitempty"`
	// Name 建议名称
	Name string `json:"name,omitempty"`
	// SceneType 场景下的细分类型，允许值:
	// 当 scene=CLEAN时，值为学习期失败率高，空耗，活跃度低，创意重复无法跑量，低效素材无法跑量
	// 当 scene=POTENTIAL时，值为潜力计划
	// 当 scene=PROBLEM时，值为起量困难，掉量，衰减，优化潜力
	SceneType string `json:"scene_type,omitempty"`
	// ToolType 工具类型，允许值：ACTION操作类建议（可直接采纳）、TEXT文案类建议
	ToolType string `json:"tool_type,omitempty"`
	// Tools 工具列表
	Tools []Tool `json:"tools,omitempty"`
}

// Tool 工具
type Tool struct {
	// Tool 工具名称
	Tool string `json:"tool,omitempty"`
	// Params 工具参数列表
	Params []Param `json:"params,omitempty"`
}

// Param 工具参数
type Param struct {
	// ParamName 工具参数名称
	ParamName string `json:"param_name,omitempty"`
	// ParamValue 工具参数值
	ParamValue ParamValue `json:"param_value,omitempty"`
}

// ParamValue 工具参数值
type ParamValue struct {
	// StringParam 字符类型参数
	StringParam string `json:"string_param,omitempty"`
	// BoolParam 布尔类型参数
	BoolParam string `json:"bool_param,omitempty"`
	// ListParam 列表类型参数
	ListParam []string `json:"list_param,omitempty"`
	// ObjectListParam 对象列表类型参数，详细信息见下方【工具参数名称&工具参数值 对应说明】
	ObjectListParam []ObjectParam `json:"object_list_param,omitempty"`
}

// ObjectType 对象类型
type ObjectType int

const (
	// ObjectType_TITLE 标题对象
	ObjectType_TITLE ObjectType = iota
	// ObjectType_IMAGE 图片素材对象
	ObjectType_IMAGE
	// ObjectType_VIDEO 视频素材对象
	ObjectType_VIDEO
)

// ObjectParam 对象列表类型参数
type ObjectParam interface {
	Type() ObjectType
}

// TitleMaterial 标题对象
type TitleMaterial struct {
	// Title 标题文案
	Title string `json:"title,omitempty"`
	// WordIDs 标题文案中如果有词包时，传入词包ID，多个词包ID用英文逗号分隔
	WordIDs []uint64 `json:"word_ids,omitempty"`
}

// Type implement ObjectType
func (m TitleMaterial) Type() ObjectType {
	return ObjectType_TITLE
}

// ImageMaterial 图片素材对象
type ImageMaterial struct {
	// ImageMode 素材类型枚举值
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
	// ImageID 图片素材URI
	ImageID string `json:"image_id,omitempty"`
}

// Type implement ObjectType
func (m ImageMaterial) Type() ObjectType {
	return ObjectType_IMAGE
}

// VideoMaterial 视频素材对象
type VideoMaterial struct {
	// ImageMode 素材类型枚举值
	ImageMode enum.ImageMode `json:"image_mode,omitempty"`
	// VideoID 视频vid
	VideoID string `json:"video_id,omitempty"`
	// CoverImageID 封面图片URI
	CoverImageID string `json:"cover_image_id,omitempty"`
}
