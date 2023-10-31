package enum

// DeepBidType 深度优化方式
type DeepBidType string

const (
	// DeepBidType_DEEP_BID_DEFAULT 不启用，无深度优化
	DeepBidType_DEEP_BID_DEFAULT DeepBidType = "DEEP_BID_DEFAULT"
	// DeepBidType_DEEP_BID_PACING 自动优化（手动出价方式下）
	DeepBidType_DEEP_BID_PACING DeepBidType = "DEEP_BID_PACING"
	// DeepBidType_DEEP_BID_MIN 自定义双出价（手动出价方式下）
	DeepBidType_DEEP_BID_MIN DeepBidType = "DEEP_BID_MIN"
	// DeepBidType_SMARTBID 自动优化（自动出价方式下）
	DeepBidType_SMARTBID DeepBidType = "SMARTBID"
	// DeepBidType_AUTO_MIN_SECOND_STAGE 自定义双出价（自动出价方式下）
	DeepBidType_AUTO_MIN_SECOND_STAGE DeepBidType = "AUTO_MIN_SECOND_STAGE"
	// DeepBidType_ROI_COEFFICIENT ROI系数
	DeepBidType_ROI_COEFFICIENT DeepBidType = "ROI_COEFFICIENT"
	// DeepBidType_ROI_PACING ROI系数——自动优化
	DeepBidType_ROI_PACING DeepBidType = "ROI_PACING"
	// DeepBidType_MIN_SECOND_STAGE 两阶段优化
	DeepBidType_MIN_SECOND_STAGE DeepBidType = "MIN_SECOND_STAGE"
	// DeepBidType_PACING_SECOND_STAGE 动态两阶段
	DeepBidType_PACING_SECOND_STAGE DeepBidType = "PACING_SECOND_STAGE"
	// DeepBidType_BID_PER_ACTION 每次付费出价
	DeepBidType_BID_PER_ACTION DeepBidType = "BID_PER_ACTION"
	// DeepBidType_SOCIAL_ROIROI 三出价, 对于每次付费的转化，深度优化类型需要设置为BID_PER_ACTION（每次付费出价）
	DeepBidType_SOCIAL_ROIROI DeepBidType = "SOCIAL_ROIROI"
	// DeepBidType_FORM_BID 优选表单出价（landing_type=link&&external_action=表单提交/多转化&&deep_external_action为空时，支持优选表单出价/不启用）
	DeepBidType_FORM_BID DeepBidType = "FORM_BID"
	// DeepBidType_PHONE_CONNECT_BID 电话接通出价
	// (landing_type=link&&external_action=表单提交，deep_external_action=电话接通时，deep_bid_type仅支持PHONE_CONNECT_BID）
	DeepBidType_PHONE_CONNECT_BID DeepBidType = "PHONE_CONNECT_BID"
	// DeepBidType_ROI_DIRECT_MAIL ROI直投（landing_type=shop/app/DPA&&external_action=app内下单、app内付费&&deep_external_action为空时，支持ROI直投/不启用）
	DeepBidType_ROI_DIRECT_MAIL DeepBidType = "ROI_DIRECT_MAIL"
)
