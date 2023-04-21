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
)
