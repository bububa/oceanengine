package diagnosis

// Suggest 诊断建议
type Suggest struct {
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// SceneList 计划对应的场景列表
	SceneList []SuggestScene `json:"scene_list,omitempty"`
}

// SuggestScene 计划对应的场景
type SuggestScene struct {
	// Scene 场景名称，允许值：CLEAN清理低质计划场景、POTENTIAL获取潜力计划场景
	Scene string `json:"scene,omitempty"`
	// Suggestions 建议列表
	Suggestions []Suggestion `json:"suggestions,omitempty"`
}

// Suggestion 建议
type Suggestion struct {
	// Msg 该场景下所有建议的详细描述
	Msg string `json:"msg,omitempty"`
	// Name 建议名称
	Name string `json:"name,omitempty"`
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
}
