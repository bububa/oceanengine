package creative

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// ComponentMaterial 组件信息
type ComponentMaterial struct {
	// ComponentID 组件id，通过【查询组件列表】接口获取
	ComponentID model.Uint64 `json:"component_id,omitempty"`
	// ComponentType 组件类型
	ComponentType enum.ComponentType `json:"component_type,omitempty"`
	// MaterialID
	MaterialID model.Uint64 `json:"material_id,omitempty"`
}
