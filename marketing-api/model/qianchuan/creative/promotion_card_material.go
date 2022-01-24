package creative

// PromotionCardMaterial 推广卡片信息
type PromotionCardMaterial struct {
	// ID 素材唯一标识
	ID uint64 `json:"id,omitempty"`
	// ComponentID 组件唯一标识
	ComponentID uint64 `json:"component_id,omitempty"`
	// Title 推广卡片标题
	Title string `json:"title,omitempty"`
	// SellingPoints 推广卡片卖点列表
	SellingPoints []string `json:"selling_points,omitempty"`
	// ImageID 推广卡片配图ID
	ImageID string `json:"image_id,omitempty"`
	// ActionButton 推广卡片行动号召按钮文案
	ActionButton string `json:"action_button,omitempty"`
	// ButtonSmartOptimization 智能优选行动号召按钮文案开关
	ButtonSmartOptimization *int `json:"button_smart_optimization,omitempty"`
}
