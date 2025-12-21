package enum

// GrayKey 白名单能力key
type GrayKey string

const (
	// GrayKey_DEFAULT_2_0_WHITELIST 白名单id为692，默认进入巨量广告升级版
	GrayKey_DEFAULT_2_0_WHITELIST GrayKey = "default_2_0_whitelist"

	// GrayKey_DEFAULT_2_0_BLACKLIST 白名单id为694，默认进入巨量广告
	GrayKey_DEFAULT_2_0_BLACKLIST GrayKey = "default_2_0_blacklist"

	// GrayKey_BAN_CREATE_WHITELIST_1_0 白名单id为688，禁止巨量广告新建计划，可创建巨量广告升级版
	GrayKey_BAN_CREATE_WHITELIST_1_0 GrayKey = "ban_create_whitelist_1_0"

	// GrayKey_BAN_CRATE_BLACKLIST_1_0 白名单id为689，允许巨量广告新建计划
	GrayKey_BAN_CRATE_BLACKLIST_1_0 GrayKey = "ban_create_blacklist_1_0"

	// GrayKey_BAN_CRATE_WHITELIST_1_0_EXCEPT_LIVE 白名单id为761，禁止巨量广告新建直播链路以外的计划，允许新建直播链路广告计划，可创建巨量广告升级版
	GrayKey_BAN_CRATE_WHITELIST_1_0_EXCEPT_LIVE GrayKey = "ban_create_whitelist_1_0_except_live"

	// GrayKey_BAN_CREATE_BLACKLIST_1_0_EXCEPT_LIVE 白名单id为763，允许巨量广告新建直播链路以外的计划，禁止新建直播链路广告计划，可创建巨量广告升级版
	GrayKey_BAN_CREATE_BLACKLIST_1_0_EXCEPT_LIVE GrayKey = "ban_create_blacklist_1_0_except_live"

	// GrayKey_BAN_UPDATE_WHITELIST_1_0 白名单id为690，禁止巨量广告编辑计划
	GrayKey_BAN_UPDATE_WHITELIST_1_0 GrayKey = "ban_update_whitelist_1_0"

	// GrayKey_BAN_UPDATE_BLACKLIST_1_0 白名单id为691，允许巨量广告编辑计划
	GrayKey_BAN_UPDATE_BLACKLIST_1_0 GrayKey = "ban_update_blacklist_1_0"

	// GrayKey_BAN_UPDATE_WHITELIST_1_0_EXCEPT_LIVE 白名单id为765，禁止巨量广告编辑直播链路以外的计划，允许编辑直播链路广告计划
	GrayKey_BAN_UPDATE_WHITELIST_1_0_EXCEPT_LIVE GrayKey = "ban_update_whitelist_1_0_except_live"

	// GrayKey_BAN_UPDATE_BLACKLIST_1_0_EXCEPT_LIVE 白名单id为766，允许巨量广告编辑直播链路以外的计划，禁止编辑直播链路广告计划
	GrayKey_BAN_UPDATE_BLACKLIST_1_0_EXCEPT_LIVE GrayKey = "ban_update_blacklist_1_0_except_live"

	// GrayKey_XIAOSHUO 白名单id为570，允许巨量广告创建landing_type为线索的项目
	GrayKey_XIAOSHUO GrayKey = "xiaoshuo"

	// GrayKey_UBA 白名单id为506，游戏粒度自动化投放（UBA）
	GrayKey_UBA GrayKey = "uba"

	// GrayKey_UBA_INTERNET_SERVICE 白名单id为506，游戏粒度自动化投放（UBA）
	GrayKey_UBA_INTERNET_SERVICE GrayKey = "uba_internet_service"

	// GrayKey_MUCH_DEEPLINK 白名单id为995，电商多直达链接能力白名单
	GrayKey_MUCH_DEEPLINK GrayKey = "much_deeplink"

	// GrayKey_MOREAPP 多平台优选投放
	GrayKey_MOREAPP GrayKey = "moreapp"

	// GrayKey_TONGXIN_YUNYINGSHANG_HAOMA_DINGXIANG 白名单id为1686，支持通信行业运营商号段设置地域定向
	GrayKey_TONGXIN_YUNYINGSHANG_HAOMA_DINGXIANG GrayKey = "tongxin_yunyingshang_haoma_dingxiang"

	// GrayKey_MallLongBid 商城广告支持长效出价
	// 对应字段external_action、deep_external_action和deep_bid_type
	// 商品购买-7日总成交：AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7DAYS
	// 商品支付ROI-7日总支付ROI：external_action=AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY_7DAYS且deep_external_action=AD_CONVERT_TYPE_LIVE_PAY_ROI且deep_bid_type=MIN
	GrayKey_MallLongBid GrayKey = "mallLongBid"

	// GrayKey_SearchProductNewAccelerate 推商品-日常销售-搜索-托管链路新品加速能力
	GrayKey_SearchProductNewAccelerate GrayKey = "search_product_new_accelerate"

	// GrayKey_CommRoi “推直播间-日常销售-通投-自定义-官方/自运营-控成本”场景下的佣金ROI
	GrayKey_CommRoi GrayKey = "comm_roi"
)
