package star

import "github.com/bububa/oceanengine/marketing-api/enum"

// Demand 星图客户任务
type Demand struct {
	// ID 任务id
	ID uint64 `json:"demand_id,omitempty"`
	// Name 任务名称
	Name string `json:"name,omitempty"`
	// ComponentType 组件类型
	ComponentType enum.StarComponentType `json:"component_type,omitempty"`
	// UniversalSettlementType 结算方式
	UniversalSettlementType enum.UniversalSettlementType `json:"universal_settlement_type,omitempty"`
	// CreateTime 任务创建时间，格式：%Y-%m-%d %H:%M:%S
	CreateTime string `json:"create_time,omitempty"`
}
