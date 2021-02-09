package enum

type DeepBidType string

const (
	DeepBidType_DEEP_BID_DEFAULT      DeepBidType = "DEEP_BID_DEFAULT"      //不启用，无深度优化
	DeepBidType_DEEP_BID_PACING       DeepBidType = "DEEP_BID_PACING"       // 自动优化（手动出价方式下）
	DeepBidType_DEEP_BID_MIN          DeepBidType = "DEEP_BID_MIN"          // 自定义双出价（手动出价方式下）
	DeepBidType_SMARTBID              DeepBidType = "SMARTBID"              // 自动优化（自动出价方式下）
	DeepBidType_AUTO_MIN_SECOND_STAGE DeepBidType = "AUTO_MIN_SECOND_STAGE" // 自定义双出价（自动出价方式下）
	DeepBidType_ROI_COEFFICIENT       DeepBidType = "ROI_COEFFICIENT"       // ROI系数
	DeepBidType_ROI_PACING            DeepBidType = "ROI_PACING"            // ROI系数——自动优化
	DeepBidType_MIN_SECOND_STAGE      DeepBidType = "MIN_SECOND_STAGE"      // 两阶段优化
	DeepBidType_PACING_SECOND_STAGE   DeepBidType = "PACING_SECOND_STAGE"   // 动态两阶段
	DeepBidType_BID_PER_ACTION        DeepBidType = "BID_PER_ACTION"        // 每次付费出价
)
