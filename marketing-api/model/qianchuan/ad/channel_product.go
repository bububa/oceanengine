package ad

import "github.com/bububa/oceanengine/marketing-api/enum/qianchuan"

// ChannelProduct 渠道品信息
type ChannelProduct struct {
	// ProductID 商品id，需要与product_ids保持一致
	ProductID uint64 `json:"product_id,omitempty"`
	// ChannelID 渠道ID
	ChannelID uint64 `json:"channel_id,omitempty"`
	// ChannelType 渠道类型
	// 达人自播 STAR_SELL
	// 商家自卖 SHOP_SELL
	ChannelType qianchuan.ProductChannelType `json:"channel_type,omitempty"`
}
