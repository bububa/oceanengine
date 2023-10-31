package qianchuan

// OrderFlowSource 订单来源
type OrderFlowSource string

const (
	// OrderFlowSource_ACTIVITY 直播-活动页面
	OrderFlowSource_ACTIVITY OrderFlowSource = "ACTIVITY"
	// OrderFlowSource_DOUYIN_SHOPPING_CENTER 直播-抖音商城
	OrderFlowSource_DOUYIN_SHOPPING_CENTER OrderFlowSource = "DOUYIN_SHOPPING_CENTER"
	// OrderFlowSource_GENERAL_SEARCH 直播-搜索
	OrderFlowSource_GENERAL_SEARCH OrderFlowSource = "GENERAL_SEARCH"
	// OrderFlowSource_GUESS_YOU_LIKE 短视频-猜你喜欢
	OrderFlowSource_GUESS_YOU_LIKE OrderFlowSource = "GUESS_YOU_LIKE"
	// OrderFlowSource_HOMEPAGE_FOLLOW 直播-关注来源
	OrderFlowSource_HOMEPAGE_FOLLOW OrderFlowSource = "HOMEPAGE_FOLLOW"
	// OrderFlowSource_LIVE 直播
	OrderFlowSource_LIVE OrderFlowSource = "LIVE"
	// OrderFlowSource_LIVE_OTHER 直播-其他
	OrderFlowSource_LIVE_OTHER OrderFlowSource = "LIVE_OTHER"
	// OrderFlowSource_OTHERS_HOMEPAGE 直播-个人主页
	OrderFlowSource_OTHERS_HOMEPAGE OrderFlowSource = "OTHERS_HOMEPAGE"
	// OrderFlowSource_OTHER_PROFILE 短视频-达人个人主页
	OrderFlowSource_OTHER_PROFILE OrderFlowSource = "OTHER_PROFILE"
	// OrderFlowSource_PRODUCT_CARD 商品卡
	OrderFlowSource_PRODUCT_CARD OrderFlowSource = "PRODUCT_CARD"
	// OrderFlowSource_PRODUCT_CARD_GENERAL_SEARCH 商户卡-搜索
	OrderFlowSource_PRODUCT_CARD_GENERAL_SEARCH OrderFlowSource = "PRODUCT_CARD_GENERAL_SEARCH"
	// OrderFlowSource_PRODUCT_CARD_OTHER 商户卡-购物车/收藏
	OrderFlowSource_PRODUCT_CARD_OTHER OrderFlowSource = "PRODUCT_CARD_OTHER"
	// OrderFlowSource_QIANCHUAN_PROMOTE 千川推广
	OrderFlowSource_QIANCHUAN_PROMOTE OrderFlowSource = "QIANCHUAN_PROMOTE"
	// OrderFlowSource_RECOMMEND_LIVE 直播-自然推荐
	OrderFlowSource_RECOMMEND_LIVE OrderFlowSource = "RECOMMEND_LIVE"
	// OrderFlowSource_RECOMMEND_PRODUCT 商户卡-商城/商品推荐
	OrderFlowSource_RECOMMEND_PRODUCT OrderFlowSource = "RECOMMEND_PRODUCT"
	// OrderFlowSource_RECOMMEND_VIDEO 短视频-短视频推荐
	OrderFlowSource_RECOMMEND_VIDEO OrderFlowSource = "RECOMMEND_VIDEO"
	// OrderFlowSource_SHOP_WINDOW 商户卡-橱窗/店铺
	OrderFlowSource_SHOP_WINDOW OrderFlowSource = "SHOP_WINDOW"
	// OrderFlowSource_UNKNOWN 未知
	OrderFlowSource_UNKNOWN OrderFlowSource = "UNKNOWN"
	// OrderFlowSource_VIDEO 短视频
	OrderFlowSource_VIDEO OrderFlowSource = "VIDEO"
	// OrderFlowSource_VIDEO_ACTIVITY 短视频-活动页
	OrderFlowSource_VIDEO_ACTIVITY OrderFlowSource = "VIDEO_ACTIVITY"
	// OrderFlowSource_VIDEO_GENERAL_SEARCH 短视频-搜索
	OrderFlowSource_VIDEO_GENERAL_SEARCH OrderFlowSource = "VIDEO_GENERAL_SEARCH"
	// OrderFlowSource_VIDEO_HOMEPAGE_FOLLOW 短视频-关注来源
	OrderFlowSource_VIDEO_HOMEPAGE_FOLLOW OrderFlowSource = "VIDEO_HOMEPAGE_FOLLOW"
	// OrderFlowSource_VIDEO_OTHER 短视频-其他
	OrderFlowSource_VIDEO_OTHER OrderFlowSource = "VIDEO_OTHER"
	// OrderFlowSource_VIDEO_TO_LIVE 直播-短视频引流
	OrderFlowSource_VIDEO_TO_LIVE OrderFlowSource = "VIDEO_TO_LIVE"
)
