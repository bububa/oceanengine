package enum

// ComponentType 组件类型
type ComponentType string

const (
	// ComponentType_CHOICE_MAGNET 选择磁贴
	ComponentType_CHOICE_MAGNET ComponentType = "CHOICE_MAGNET"
	// ComponentType_VOTE_MAGNET 投票磁贴
	ComponentType_VOTE_MAGNET ComponentType = "VOTE_MAGNET"
	// ComponentType_IMAGE_MAGNET 图片磁贴
	ComponentType_IMAGE_MAGNET ComponentType = "IMAGE_MAGNET"
	// ComponentType_COMMERCE_MAGNET 电商磁贴
	ComponentType_COMMERCE_MAGNET ComponentType = "COMMERCE_MAGNET"
	// ComponentType_COUPON_MAGNET 优惠券磁贴
	ComponentType_COUPON_MAGNET ComponentType = "COUPON_MAGNET"
	// ComponentType_GAME_PACK 游戏礼包码
	ComponentType_GAME_PACK ComponentType = "GAME_PACK"
	// ComponentType_PROMOTION_CARD 推广卡片
	ComponentType_PROMOTION_CARD ComponentType = "PROMOTION_CARD"
)
