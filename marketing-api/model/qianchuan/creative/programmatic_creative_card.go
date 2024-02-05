package creative

// ProgrammaticCreativeCard 程序化创意推广卡片信息
type ProgrammaticCreativeCard struct {
	// PromotionCardID 推广卡片ID
	PromotionCardID uint64 `json:"promotion_card_id,omitempty"`
	// ComponentID 组件ID
	ComponentID uint64 `json:"component_id,omitempty"`
	// Title 推广卡片标题
	Title string `json:"promotion_card_title,omitempty"`
	// SellingPoints 推广卡片卖点列表
	SellingPoints []string `json:"promotion_card_selling_points,omitempty"`
	// ImageID 推广卡片配图ID
	ImageID string `json:"promotion_card_image_id,omitempty"`
	// ActionButton 推广卡片行动号召按钮文案
	ActionButton string `json:"promotion_card_action_button,omitempty"`
	// ButtonSmartOptimization 智能优选行动号召按钮文案开关
	ButtonSmartOptimization int `json:"promotion_card_button_smart_optimization,omitempty"`
}
