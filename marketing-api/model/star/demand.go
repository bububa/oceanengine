package star

import "github.com/bububa/oceanengine/marketing-api/enum"

// Demand 星图客户任务
type Demand struct {
	// DemandID 任务id
	DemandID uint64 `json:"demand_id,omitempty"`
	// DemandName 任务名称
	DemandName string `json:"name,omitempty"`
	// TaskCategory 所创建的星图任务类型。枚举值详见【附录-枚举值-星图任务类型】
	TaskCategory enum.StarTaskCategory `json:"task_category,omitempty"`
	// ComponentType 组件类型
	ComponentType enum.StarComponentType `json:"component_type,omitempty"`
	// UniversalSettlementType 结算方式
	UniversalSettlementType enum.UniversalSettlementType `json:"universal_settlement_type,omitempty"`
	// CreateTime 任务创建时间，格式：%Y-%m-%d %H:%M:%S
	CreateTime string `json:"create_time,omitempty"`
}
