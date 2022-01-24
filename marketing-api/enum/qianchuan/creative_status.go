package qianchuan

// CreativeStatus 广告创意投放状态
type CreativeStatus string

const (
	// CreativeStatus_DELIVERY_OK 投放中
	CreativeStatus_DELIVERY_OK CreativeStatus = "DELIVERY_OK"
	// CreativeStatus_AUDIT 新建审核中
	CreativeStatus_AUDIT CreativeStatus = "AUDIT"
	// CreativeStatus_REAUDIT 修改审核中
	CreativeStatus_REAUDIT CreativeStatus = "REAUDIT"
	// CreativeStatus_DELETE 已删除
	CreativeStatus_DELETE CreativeStatus = "DELETE"
	// CreativeStatus_DISABLE 已暂停
	CreativeStatus_DISABLE CreativeStatus = "DISABLE"
	// CreativeStatus_AD_DISABLE 广告计划已暂停
	CreativeStatus_AD_DISABLE CreativeStatus = "AD_DISABLE"
	// CreativeStatus_TIME_NO_REACH 未到达投放时间
	CreativeStatus_TIME_NO_REACH CreativeStatus = "TIME_NO_REACH"
	// CreativeStatus_TIME_DONE 已完成
	CreativeStatus_TIME_DONE CreativeStatus = "TIME_DONE"
	// CreativeStatus_NO_SCHEDULE 不在投放时段
	CreativeStatus_NO_SCHEDULE CreativeStatus = "NO_SCHEDULE"
	// CreativeStatus_CREATE 计划新建
	CreativeStatus_CREATE CreativeStatus = "CREATE"
	// CreativeStatus_FROZEN 已终止
	CreativeStatus_FROZEN CreativeStatus = "FROZEN"
	// CreativeStatus_OFFLINE_AUDIT 审核不通过
	CreativeStatus_OFFLINE_AUDIT CreativeStatus = "OFFLINE_AUDIT"
	// CreativeStatus_ADVERTISER_PRE_OFFLINE_BUDGET 广告计划超出预算
	CreativeStatus_AD_OFFLINE_BUDGET CreativeStatus = "AD_OFFLINE_BUDGET"
	// CreativeStatus_OFFLINE_BALANCE 账户余额不足
	CreativeStatus_OFFLINE_BALANCE CreativeStatus = "OFFLINE_BALANCE"
	// CreativeStatus_AD_PRE_OFFLINE_BUDGET 广告计划接近预算
	CreativeStatus_AD_PRE_OFFLINE_BUDGET CreativeStatus = "AD_PRE_OFFLINE_BUDGET"
	// CreativeStatus_PRE_ONLINE 预上线
	CreativeStatus_PRE_ONLINE CreativeStatus = "PRE_ONLINE"
	// CreativeStatus_AD_REAUDIT 广告计划修改审核中
	CreativeStatus_AD_REAUDIT CreativeStatus = "AD_REAUDIT"
	// CreativeStatus_AD_AUDIT 广告计划新建审核中
	CreativeStatus_AD_AUDIT CreativeStatus = "AD_AUDIT"
	// CreativeStatus_AD_OFFLINE_AUDIT 广告计划审核不通过
	CreativeStatus_AD_OFFLINE_AUDIT CreativeStatus = "AD_OFFLINE_AUDIT"
	// CreativeStatus_ERROR 数据错误
	CreativeStatus_ERROR CreativeStatus = "ERROR"
	// CreativeStatus_AD_EXTERNAL_URL_DISABLE 计划落地页不可用
	CreativeStatus_AD_EXTERNAL_URL_DISABLE CreativeStatus = "AD_EXTERNAL_URL_DISABLE"
	// CreativeStatus_LIVE_CAN_NOT_LAUNCH 直播间状态不可投
	CreativeStatus_LIVE_CAN_NOT_LAUNCH CreativeStatus = "LIVE_CAN_NOT_LAUNCH"
	// CreativeStatus_AWEME_ITEM_DISABLED 关联抖音视频不可投
	CreativeStatus_AWEME_ITEM_DISABLED CreativeStatus = "AWEME_ITEM_DISABLED"
	// CreativeStatus_CAMPAIGN_DISABLE 已被广告组暂停
	CreativeStatus_CAMPAIGN_DISABLE CreativeStatus = "CAMPAIGN_DISABLE"
	// CreativeStatus_CAMPAIGN_OFFLINE_BUDGET 广告组超出预算
	CreativeStatus_CAMPAIGN_OFFLINE_BUDGET CreativeStatus = "CAMPAIGN_OFFLINE_BUDGET"
	// CreativeStatus_CAMPAIGN_PREOFFLINE_BUDGET 广告组接近预算
	CreativeStatus_CAMPAIGN_PREOFFLINE_BUDGET CreativeStatus = "CAMPAIGN_PREOFFLINE_BUDGET"
)
