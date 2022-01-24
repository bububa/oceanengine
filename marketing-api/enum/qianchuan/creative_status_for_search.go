package qianchuan

// CreativeStatusForSearch 广告创意查询状态
type CreativeStatusForSearch string

const (
	// CreativeStatusForSearch_DELETED 已删除
	CreativeStatusForSearch_DELETED CreativeStatusForSearch = "DELETED"
	// CreativeStatusForSearch_AUDIT 新建审核中
	CreativeStatusForSearch_AUDIT CreativeStatusForSearch = "AUDIT"
	// CreativeStatusForSearch_REAUDIT 修改审核中
	CreativeStatusForSearch_REAUDIT CreativeStatusForSearch = "REAUDIT"
	// CreativeStatusForSearch_TIME_DONE 已完成
	CreativeStatusForSearch_TIME_DONE CreativeStatusForSearch = "TIME_DONE"
	// CreativeStatusForSearch_DISABLE 已暂停
	CreativeStatusForSearch_DISABLE CreativeStatusForSearch = "DISABLE"
	// CreativeStatusForSearch_AD_DISABLE 广告计划已暂停
	CreativeStatusForSearch_AD_DISABLE CreativeStatusForSearch = "AD_DISABLE"
	// CreativeStatusForSearch_TIME_NO_REACH 未到达投放时间
	CreativeStatusForSearch_TIME_NO_REACH CreativeStatusForSearch = "TIME_NO_REACH"
	// CreativeStatusForSearch_AD_OFFLINE_BUDGET 广告预算不足
	CreativeStatusForSearch_AD_OFFLINE_BUDGET CreativeStatusForSearch = "AD_OFFLINE_BUDGET"
	// CreativeStatusForSearch_OFFLINE_BALANCE 账户余额不足
	CreativeStatusForSearch_OFFLINE_BALANCE CreativeStatusForSearch = "OFFLINE_BALANCE"
	// CreativeStatusForSearch_DELIVERY_OK 投放中
	CreativeStatusForSearch_DELIVERY_OK CreativeStatusForSearch = "DELIVERY_OK"
	// CreativeStatusForSearch_NO_SCHEDULE 不在投放时段
	CreativeStatusForSearch_NO_SCHEDULE CreativeStatusForSearch = "NO_SCHEDULE"
	// CreativeStatusForSearch_OFFLINE_AUDIT 审核不通过
	CreativeStatusForSearch_OFFLINE_AUDIT CreativeStatusForSearch = "OFFLINE_AUDIT"
	// CreativeStatusForSearch_EXTERNAL_URL_DISABLE 落地页不可用
	CreativeStatusForSearch_EXTERNAL_URL_DISABLE CreativeStatusForSearch = "EXTERNAL_URL_DISABLE"
	// CreativeStatusForSearch_LIVE_ROOM_OFF 关联直播间未开播
	CreativeStatusForSearch_LIVE_ROOM_OFF CreativeStatusForSearch = "LIVE_ROOM_OFF"
	// CreativeStatusForSearch_FROZEN 已终止
	CreativeStatusForSearch_FROZEN CreativeStatusForSearch = "FROZEN"
	// CreativeStatusForSearch_ALL_INCLUDE_DELETED 全部（包含已删除）
	CreativeStatusForSearch_ALL_INCLUDE_DELETED CreativeStatusForSearch = "ALL_INCLUDE_DELETED"
)
