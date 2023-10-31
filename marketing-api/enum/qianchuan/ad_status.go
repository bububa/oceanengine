package qianchuan

// AdStatus 广告计划投放状态
type AdStatus string

const (
	// AdStatus_DELIVERY_OK 投放中
	AdStatus_DELIVERY_OK AdStatus = "DELIVERY_OK"
	// AdStatus_AUDIT 新建审核中
	AdStatus_AUDIT AdStatus = "AUDIT"
	// AdStatus_REAUDIT 修改审核中
	AdStatus_REAUDIT AdStatus = "REAUDIT"
	// AdStatus_DELETE 已删除
	AdStatus_DELETE AdStatus = "DELETE"
	// AdStatus_DISABLE 已暂停
	AdStatus_DISABLE AdStatus = "DISABLE"
	// AdStatus_DRAFT 草稿
	AdStatus_DRAFT AdStatus = "DRAFT"
	// AdStatus_TIME_NO_REACH 未到达投放时间
	AdStatus_TIME_NO_REACH AdStatus = "TIME_NO_REACH"
	// AdStatus_TIME_DONE 已完成
	AdStatus_TIME_DONE AdStatus = "TIME_DONE"
	// AdStatus_NO_SCHEDULE 不在投放时段
	AdStatus_NO_SCHEDULE AdStatus = "NO_SCHEDULE"
	// AdStatus_CREATE 计划新建
	AdStatus_CREATE AdStatus = "CREATE"
	// AdStatus_OFFLINE_AUDIT 审核不通过
	AdStatus_OFFLINE_AUDIT AdStatus = "OFFLINE_AUDIT"
	// AdStatus_OFFLINE_BUDGET 广告预算不足（已超出预算）
	AdStatus_OFFLINE_BUDGET AdStatus = "OFFLINE_BUDGET"
	// AdStatus_OFFLINE_BALANCE 账户余额不足
	AdStatus_OFFLINE_BALANCE AdStatus = "OFFLINE_BALANCE"
	// AdStatus_PRE_OFFLINE_BUDGET 广告预算不足（即将超出预算）
	AdStatus_PRE_OFFLINE_BUDGET AdStatus = "PRE_OFFLINE_BUDGET"
	// AdStatus_PRE_ONLINE 预上线
	AdStatus_PRE_ONLINE AdStatus = "PRE_ONLINE"
	// AdStatus_FROZEN 已终止
	AdStatus_FROZEN AdStatus = "FROZEN"
	// AdStatus_ERROR 数据错误
	AdStatus_ERROR AdStatus = "ERROR"
	// AdStatus_AUDIT_STATUS_ERROR 异常，请联系审核人员
	AdStatus_AUDIT_STATUS_ERROR AdStatus = "AUDIT_STATUS_ERROR"
	// AdStatus_ADVERTISER_OFFLINE_BUDGET 账户超出预算
	AdStatus_ADVERTISER_OFFLINE_BUDGET AdStatus = "ADVERTISER_OFFLINE_BUDGET"
	// AdStatus_ADVERTISER_PRE_OFFLINE_BUDGET 账户接近预算
	AdStatus_ADVERTISER_PRE_OFFLINE_BUDGET AdStatus = "ADVERTISER_PRE_OFFLINE_BUDGET"
	// AdStatus_EXTERNAL_URL_DISABLE 落地页暂不可用
	AdStatus_EXTERNAL_URL_DISABLE AdStatus = "EXTERNAL_URL_DISABLE"
	// AdStatus_LIVE_ROOM_OFF 关联直播间未开播
	AdStatus_LIVE_ROOM_OFF AdStatus = "LIVE_ROOM_OFF"
	// AdStatus_CAMPAIGN_DISABLE 已被广告组暂停
	AdStatus_CAMPAIGN_DISABLE AdStatus = "CAMPAIGN_DISABLE"
	// AdStatus_CAMPAIGN_OFFLINE_BUDGET 广告组超出预算
	AdStatus_CAMPAIGN_OFFLINE_BUDGET AdStatus = "CAMPAIGN_OFFLINE_BUDGET"
	// AdStatus_CAMPAIGN_PREOFFLINE_BUDGET 广告组接近预算
	AdStatus_CAMPAIGN_PREOFFLINE_BUDGET AdStatus = "CAMPAIGN_PREOFFLINE_BUDGET"
	// AdStatus_SYSTEM_DISABLE 系统暂停
	AdStatus_SYSTEM_DISABLE AdStatus = "SYSTEM_DISABLE"
	// AdStatus_QUOTA_DISABLE 在投计划配额超限
	AdStatus_QUOTA_DISABLE AdStatus = "QUOTA_DISABLE"
)
