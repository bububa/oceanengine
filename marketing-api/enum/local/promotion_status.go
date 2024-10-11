package local

// PromotionStatusFirst  广告一级状态
type PromotionStatusFirst string

const (
	// PROMOTION_STATUS_ALL 不限（包含已删除）
	PROMOTION_STATUS_ALL PromotionStatusFirst = "PROMOTION_STATUS_ALL"
	// PROMOTION_STATUS_DELETED 已删除:
	PROMOTION_STATUS_DELETED PromotionStatusFirst = "PROMOTION_STATUS_DELETED"
	// PROMOTION_STATUS_DISABLE 未投放
	PROMOTION_STATUS_DISABLE PromotionStatusFirst = "PROMOTION_STATUS_DISABLE"
	// PROMOTION_STATUS_DONE 已完成
	PROMOTION_STATUS_DONE PromotionStatusFirst = "PROMOTION_STATUS_DONE"
	// PROMOTION_STATUS_ENABLE 投放中
	PROMOTION_STATUS_ENABLE PromotionStatusFirst = "PROMOTION_STATUS_ENABLE"
	// PROMOTION_STATUS_FROZEN 已终止
	PROMOTION_STATUS_FROZEN PromotionStatusFirst = "PROMOTION_STATUS_FROZEN"
	// PROMOTION_STATUS_NOT_DELETE 不限（不包含已删除）
	PROMOTION_STATUS_NOT_DELETE PromotionStatusFirst = "PROMOTION_STATUS_NOT_DELETE"
)

// PromotionStatusSecond 广告二级状态
type PromotionStatusSecond string

const (
	// PromotionStatusSecond_AUDIT_DENY 审核不通过
	PromotionStatusSecond_AUDIT_DENY PromotionStatusSecond = "AUDIT_DENY"
	// PromotionStatusSecond_AUDIT 新建审核中
	PromotionStatusSecond_AUDIT PromotionStatusSecond = "AUDIT"
	// PromotionStatusSecond_REAUDIT 修改审核中
	PromotionStatusSecond_REAUDIT PromotionStatusSecond = "REAUDIT"
	// PromotionStatusSecond_DISABLED 已暂停
	PromotionStatusSecond_DISABLED PromotionStatusSecond = "DISABLED"
	// PromotionStatusSecond_DISABLE_BY_QUOTA 配额达限
	PromotionStatusSecond_DISABLE_BY_QUOTA PromotionStatusSecond = "DISABLE_BY_QUOTA"
	// PromotionStatusSecond_PROJECT_DISABLED 项目已被暂停
	PromotionStatusSecond_PROJECT_DISABLED PromotionStatusSecond = "PROJECT_DISABLED"
	// PromotionStatusSecond_NO_SCHEDULE 不在投放时段
	PromotionStatusSecond_NO_SCHEDULE PromotionStatusSecond = "NO_SCHEDULE"
	// PromotionStatusSecond_TIME_NO_REACH 未到达投放时间
	PromotionStatusSecond_TIME_NO_REACH PromotionStatusSecond = "TIME_NO_REACH"
	// PromotionStatusSecond_OFFLINE_BALANCE 账户余额不足
	PromotionStatusSecond_OFFLINE_BALANCE PromotionStatusSecond = "OFFLINE_BALANCE"
	// PromotionStatusSecond_BALANCE_OFFLINE_BUDGET 账户超出预算
	PromotionStatusSecond_BALANCE_OFFLINE_BUDGET PromotionStatusSecond = "BALANCE_OFFLINE_BUDGET"
	// PromotionStatusSecond_PROJECT_OFFLINE_BUDGET 项目超出预算
	PromotionStatusSecond_PROJECT_OFFLINE_BUDGET PromotionStatusSecond = "PROJECT_OFFLINE_BUDGET"
	// PromotionStatusSecond_PROMOTION_OFFLINE_BUDGET 广告超出预算
	PromotionStatusSecond_PROMOTION_OFFLINE_BUDGET PromotionStatusSecond = "PROMOTION_OFFLINE_BUDGET"
	// PromotionStatusSecond_LIVE_ROOM_OFF 直播间不可投放
	PromotionStatusSecond_LIVE_ROOM_OFF PromotionStatusSecond = "LIVE_ROOM_OFF"
	// PromotionStatusSecond_AWEME_ACCOUNT_DISABLED 抖音账号不可投放
	PromotionStatusSecond_AWEME_ACCOUNT_DISABLED PromotionStatusSecond = "AWEME_ACCOUNT_DISABLED"
	// PromotionStatusSecond_PRODUCT_OR_POI_OFFLINE 商品/门店不可投
	PromotionStatusSecond_PRODUCT_OR_POI_OFFLINE PromotionStatusSecond = "PRODUCT_OR_POI_OFFLINE"
)
