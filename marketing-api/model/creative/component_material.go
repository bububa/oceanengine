package creative

// ComponentMaterial 组件信息
type ComponentMaterial struct {
	// ComponentID 组件id，通过【查询组件列表】接口获取
	ComponentID uint64 `json:"component_id,omitempty"`
}
