package qianchuan

// AdStatusForSearch 广告计划查询状态
type AdStatusForSearch string

const (
	// AdStatusForSearch_DELETED 已删除
	AdStatusForSearch_DELETED AdStatusForSearch = "DELETED"
	// AdStatusForSearch_AUDIT 新建审核中
	AdStatusForSearch_AUDIT AdStatusForSearch = "AUDIT"
	// AdStatusForSearch_TIME_DONE 已完成
	AdStatusForSearch_TIME_DONE AdStatusForSearch = "TIME_DONE"
	// AdStatusForSearch_DISABLE 已暂停
	AdStatusForSearch_DISABLE AdStatusForSearch = "DISABLE"
	// AdStatusForSearch_TIME_NO_REACH 未到达投放时间
	AdStatusForSearch_TIME_NO_REACH AdStatusForSearch = "TIME_NO_REACH"
	// AdStatusForSearch_OFFLINE_BALANCE 账户余额不足
	AdStatusForSearch_OFFLINE_BALANCE AdStatusForSearch = "OFFLINE_BALANCE"
	// AdStatusForSearch_OFFLINE_BUDGET 广告预算不足
	AdStatusForSearch_OFFLINE_BUDGET AdStatusForSearch = "OFFLINE_BUDGET"
	// AdStatusForSearch_DELIVERY_OK 投放中
	AdStatusForSearch_DELIVERY_OK AdStatusForSearch = "DELIVERY_OK"
	// AdStatusForSearch_NO_SCHEDULE 不在投放时段
	AdStatusForSearch_NO_SCHEDULE AdStatusForSearch = "NO_SCHEDULE"
	// AdStatusForSearch_REAUDIT 修改审核中
	AdStatusForSearch_REAUDIT AdStatusForSearch = "REAUDIT"
	// AdStatusForSearch_OFFLINE_AUDIT 审核不通过
	AdStatusForSearch_OFFLINE_AUDIT AdStatusForSearch = "OFFLINE_AUDIT"
	// AdStatusForSearch_ENABLE 已启用
	AdStatusForSearch_ENABLE AdStatusForSearch = "ENABLE"
	// AdStatusForSearch_EXTERNAL_URL_DISABLE 落地页暂不可用
	AdStatusForSearch_EXTERNAL_URL_DISABLE AdStatusForSearch = "EXTERNAL_URL_DISABLE"
	// AdStatusForSearch_LIVE_ROOM_OFF 关联直播间未开播
	AdStatusForSearch_LIVE_ROOM_OFF AdStatusForSearch = "LIVE_ROOM_OFF"
	// AdStatusForSearch_FROZEN 已终止
	AdStatusForSearch_FROZEN AdStatusForSearch = "FROZEN"
	// AdStatusForSearch_SYSTEM_DISABLE 系统暂停
	AdStatusForSearch_SYSTEM_DISABLE AdStatusForSearch = "SYSTEM_DISABLE"
	// AdStatusForSearch_ALL_INCLUDE_DELETED 全部（包含已删除）
	AdStatusForSearch_ALL_INCLUDE_DELETED AdStatusForSearch = "ALL_INCLUDE_DELETED"
	// AdStatusForSearch_QUOTA_DISABLE 在投计划配额超限
	AdStatusForSearch_QUOTA_DISABLE AdStatusForSearch = "QUOTA_DISABLE"
)
