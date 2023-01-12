package v3

// Dimension 维度
type Dimension struct {
	// Field 维度字段
	Field string `json:"field,omitempty"`
	// Name 维度名称
	Name string `json:"name,omitempty"`
	// Description 维度描述
	Description string `json:"description,omitempty"`
	// Sortable 是否支持排序
	Sortable bool `json:"sort_able,omtiempty"`
	// Filterable 是否支持可筛选
	Filterable bool `json:"filter_able,omitempty"`
	// FilterConfig 过滤条件
	FilterConfig *FilterConfig `json:"filter_config,omitempty"`
}

// FilterConfig 过滤条件
type FilterConfig struct {
	// Type 字段类型。
	// 1 -固定枚举值
	// 2 - 固定输入值
	// 3 -数值类型
	Type int `json:"type,omitempty"`
	// Operator 处理方式。
	// 1-等于
	// 2 -小于
	// 3 -小于等于
	// 4 -大于
	// 5 -大于等于
	// 6 -不等于
	// 7-包含
	// 8 -不包含
	// 9 -范围查询
	// 10 -多个值匹配包含
	// 11 -多个值匹配都要包含
	Operator int `json:"operator,omitempty"`
	// ValueLimit 过滤字段传入数量上限
	ValueLimit int64 `json:"value_limit,omitempty"`
	// RangeValue 维度指标过滤枚举值列表
	RangeValue []RangeValue `json:"range_value,omitempty"`
	// ExclusionDims 互斥的维度列表
	ExclusionDims []string `json:"exclusion_dims,omitempty"`
	// ExclusionMetrics 互斥的指标列表
	ExclusionMetrics []string `json:"exclusion_metrics,omitempty"`
}

// RangeValue 维度指标过滤枚举值列表
type RangeValue struct {
	// Label 维度过滤字段名称
	Label string `json:"label,omitempty"`
	// Value 维度过滤字段具体值
	Value string `json:"value,omitempty"`
}
