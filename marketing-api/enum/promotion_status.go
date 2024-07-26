package enum

// PromotionStatus 广告状态
type PromotionStatus string

const (
	// PromotionStatus_NOT_DELETED NOT_DELETED 不限
	PromotionStatus_NOT_DELETED PromotionStatus = "NOT_DELETED"
	// PromotionStatus_ALL ALL 不限（包含已删除）
	PromotionStatus_ALL PromotionStatus = "ALL"
	// PromotionStatus_OK OK 投放中
	PromotionStatus_OK PromotionStatus = "OK"
	// PromotionStatus_DELETED DELETED 已删除
	PromotionStatus_DELETED PromotionStatus = "DELETED"
	// PromotionStatus_PROJECT_OFFLINE_BUDGET 项目超出预算
	PromotionStatus_PROJECT_OFFLINE_BUDGET PromotionStatus = "PROJECT_OFFLINE_BUDGET"
	// PromotionStatus_PROJECT_PREOFFLINE_BUDGET 项目接近预算
	PromotionStatus_PROJECT_PREOFFLINE_BUDGET PromotionStatus = "PROJECT_REOFFLINE_BUDGET"
	// PromotionStatus_TIME_NO_REACH 未到达投放时间
	PromotionStatus_TIME_NO_REACH PromotionStatus = "TIME_NO_REACH"
	// PromotionStatus_TIME_DONE 已完成
	PromotionStatus_TIME_DONE PromotionStatus = "TIME_DONE"
	// PromotionStatus_NO_SCHEDULE 不在投放时段
	PromotionStatus_NO_SCHEDULE PromotionStatus = "NO_SCHEDULE"
	// PromotionStatus_AUDIT 新建审核中
	PromotionStatus_AUDIT PromotionStatus = "AUDIT"
	// PromotionStatus_REAUDIT 修改审核中
	PromotionStatus_REAUDIT PromotionStatus = "REAUDIT"
	// PromotionStatus_FROZEN 已终止
	PromotionStatus_FROZEN PromotionStatus = "FROZEN"
	// PromotionStatus_AUDIT_DENY 审核不通过
	PromotionStatus_AUDIT_DENY PromotionStatus = "AUDIT_DENY"
	// PromotionStatus_OFFLINE_BUDGET 广告超出预算
	PromotionStatus_OFFLINE_BUDGET PromotionStatus = "OFFLINE_BUDGET"
	// PromotionStatus_OFFLINE_BALANCE 账户余额不足
	PromotionStatus_OFFLINE_BALANCE PromotionStatus = "OFFLINE_BALANCE"
	// PromotionStatus_PREOFFLINE_BUDGET 广告接近预算
	PromotionStatus_PREOFFLINE_BUDGET PromotionStatus = "PREOFFLINE_BUDGET"
	// PromotionStatus_BALANCE_OFFLINE_BUDGET 账户超出预算
	PromotionStatus_BALANCE_OFFLINE_BUDGET PromotionStatus = "BALANCE_OFFLINE_BUDGET"
	// PromotionStatus_PROMOTION_OFFLINE_BUDGET 广告超出预算
	PromotionStatus_PROMOTION_OFFLINE_BUDGET PromotionStatus = "PROMOTION_OFFLINE_BUDGET"
	// PromotionStatus_DISABLED 已暂停
	PromotionStatus_DISABLED PromotionStatus = "DISABLED"
	// PromotionStatus_PROJECT_DISABLED 已被项目暂停
	PromotionStatus_PROJECT_DISABLED PromotionStatus = "PROJECT_DISABLED"
	// PromotionStatus_LIVE_ROOM_OFF 关联直播间不可投
	PromotionStatus_LIVE_ROOM_OFF PromotionStatus = "LIVE_ROOM_OFF"
	// PromotionStatus_PRODUCT_OFFLINE 关联商品不可投
	PromotionStatus_PRODUCT_OFFLINE PromotionStatus = "PRODUCT_OFFLINE"
	// PromotionStatus_AWEME_ACCOUNT_DISABLED 关联抖音账号不可投
	PromotionStatus_AWEME_ACCOUNT_DISABLED PromotionStatus = "AWEME_ACCOUNT_DISABLED"
	// PromotionStatus_AWEME_ANCHOR_DISABLED 锚点不可投
	PromotionStatus_AWEME_ANCHOR_DISABLED PromotionStatus = "AWEME_ANCHOR_DISABLED"
	// PromotionStatus_AWEME_ACCOUNT_OPTIMIZABLE 关联抖音号可优化
	PromotionStatus_AWEME_ACCOUNT_OPTIMIZABLE PromotionStatus = "AWEME_ACCOUNT_OPTIMIZABLE"
	// PromotionStatus_AWEME_VIDEO_OPTIMIZABLE 关联抖音视频可优化
	PromotionStatus_AWEME_VIDEO_OPTIMIZABLE PromotionStatus = "AWEME_VIDEO_OPTIMIZABLE"
	// PromotionStatus_DISABLE_BY_QUOTA 已暂停（配额达限）
	PromotionStatus_DISABLE_BY_QUOTA PromotionStatus = "DISABLE_BY_QUOTA"
	// PromotionStatus_CREATE 新建
	PromotionStatus_CREATE PromotionStatus = "CREATE"
	// PromotionStatus_ADVERTISER_OFFLINE_BUDGET 账号超出预算
	PromotionStatus_ADVERTISER_OFFLINE_BUDGET PromotionStatus = "ADVERTISER_OFFLINE_BUDGET "
	// PromotionStatus_ADVERTISER_PREOFFLINE_BUDGET 账号接近预算
	PromotionStatus_ADVERTISER_PREOFFLINE_BUDGET PromotionStatus = "ADVERTISER_PREOFFLINE_BUDGET "
)

// PromotionStatusFirst 广告一级状态过滤
type PromotionStatusFirst string

const (
	// PROMOTION_STATUS_ENABLE 投放中
	PROMOTION_STATUS_ENABLE PromotionStatusFirst = "PROMOTION_STATUS_ENABLE"
	// PROMOTION_STATUS_DISABLE 未投放
	PROMOTION_STATUS_DISABLE PromotionStatusFirst = "PROMOTION_STATUS_DISABLE"
	// PROMOTION_STATUS_FROZEN 已终止
	PROMOTION_STATUS_FROZEN PromotionStatusFirst = "PROMOTION_STATUS_FROZEN"
	// PROMOTION_STATUS_DONE 已完成
	PROMOTION_STATUS_DONE PromotionStatusFirst = "PROMOTION_STATUS_DONE"
	// PROMOTION_STATUS_DELETED 已删除
	PROMOTION_STATUS_DELETED PromotionStatusFirst = "PROMOTION_STATUS_DELETED"
	// PROMOTION_STATUS_ALL 不限包含已删除
	PROMOTION_STATUS_ALL PromotionStatusFirst = "PROMOTION_STATUS_ALL"
	// PROMOTION_STATUS_NOT_DELETE 不限（不包含已删除）
	PROMOTION_STATUS_NOT_DELETE PromotionStatusFirst = "PROMOTION_STATUS_NOT_DELETE"
)
