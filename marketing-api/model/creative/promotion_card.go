package creative

// PromotionCard 商品推广卡片
type PromotionCard struct {
	// EnableStorePack 是否使用门店包，true为使用，false为不使用，推广目的非门店推广时会忽略该字段。若选择使用，则卡片标题为最近门店名称
	EnableStorePack bool `json:"enable_store_pack,omitempty"`
	// ProductSellingPoints 商品卖点，对应广告投放平台的推荐卖点
	ProductSellingPoints []string `json:"product_selling_points,omitempty"`
	// ProductDescription 商品描述
	ProductDescription string `json:"product_description,omitempty"`
	// CallToAction 行动号召
	CallToAction string `json:"call_to_action,omitempty"`
	// EnablePersonalAction 是否使用智能优选，true为使用，false为不使用
	EnablePersonalAction bool `json:"enable_personal_action,omitempty"`
	// ProductImageID 商品图片ID，对应广告投放平台的推广卡片主图
	ProductImageID string `json:"product_image_id,omitempty"`
}
