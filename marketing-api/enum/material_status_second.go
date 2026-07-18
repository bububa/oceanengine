package enum

// MaterialStatusSecond 素材二级状态
// 该枚举为视频/图片/图文/试玩/直玩等素材类型 status_second 字段的并集，具体可用值随素材类型不同而不同
type MaterialStatusSecond string

const (
	// MaterialStatusSecond_AUDIT 未投放-新建审核中
	MaterialStatusSecond_AUDIT MaterialStatusSecond = "MATERIAL_STATUS_AUDIT"
	// MaterialStatusSecond_AWEME_ACCOUNT_OFFLINE 抖音账号不可投放
	MaterialStatusSecond_AWEME_ACCOUNT_OFFLINE MaterialStatusSecond = "MATERIAL_STATUS_AWEME_ACCOUNT_OFFLINE"
	// MaterialStatusSecond_DELETE 素材已删除
	MaterialStatusSecond_DELETE MaterialStatusSecond = "MATERIAL_STATUS_DELETE"
	// MaterialStatusSecond_DISABLE 手动暂停/已暂停
	MaterialStatusSecond_DISABLE MaterialStatusSecond = "MATERIAL_STATUS_DISABLE"
	// MaterialStatusSecond_NO_SCHEDULE 项目不在投放时段
	MaterialStatusSecond_NO_SCHEDULE MaterialStatusSecond = "MATERIAL_STATUS_NO_SCHEDULE"
	// MaterialStatusSecond_OFFLINE_AUDIT 审核不通过
	MaterialStatusSecond_OFFLINE_AUDIT MaterialStatusSecond = "MATERIAL_STATUS_OFFLINE_AUDIT"
	// MaterialStatusSecond_OFFLINE_BALANCE 账户余额不足
	MaterialStatusSecond_OFFLINE_BALANCE MaterialStatusSecond = "MATERIAL_STATUS_OFFLINE_BALANCE"
	// MaterialStatusSecond_OFFLINE_BUDGET 账户超出预算
	MaterialStatusSecond_OFFLINE_BUDGET MaterialStatusSecond = "MATERIAL_STATUS_OFFLINE_BUDGET"
	// MaterialStatusSecond_PRODUCT_OFFLINE 商品不可投放
	MaterialStatusSecond_PRODUCT_OFFLINE MaterialStatusSecond = "MATERIAL_STATUS_PRODUCT_OFFLINE"
	// MaterialStatusSecond_PROJECT_DISABLE 项目已暂停
	MaterialStatusSecond_PROJECT_DISABLE MaterialStatusSecond = "MATERIAL_STATUS_PROJECT_DISABLE"
	// MaterialStatusSecond_PROJECT_OFFLINE_BUDGET 项目超出预算
	MaterialStatusSecond_PROJECT_OFFLINE_BUDGET MaterialStatusSecond = "MATERIAL_STATUS_PROJECT_OFFLINE_BUDGET"
	// MaterialStatusSecond_TIME_NO_REACH 项目未达投放时间
	MaterialStatusSecond_TIME_NO_REACH MaterialStatusSecond = "MATERIAL_STATUS_TIME_NO_REACH"
	// MaterialStatusSecond_FROZEN 素材已终止
	MaterialStatusSecond_FROZEN MaterialStatusSecond = "MATERIAL_STATUS_FROZEN"
	// MaterialStatusSecond_PROCEDURAL_DISABLE 系统暂停
	MaterialStatusSecond_PROCEDURAL_DISABLE MaterialStatusSecond = "MATERIAL_STATUS_PROCEDURAL_DISABLE"
	// MaterialStatusSecond_PROJECT_STATUS_DELETE 项目已删除
	MaterialStatusSecond_PROJECT_STATUS_DELETE MaterialStatusSecond = "PROJECT_STATUS_DELETE"
	// MaterialStatusSecond_PROJECT_STATUS_FROZEN 项目已终止
	MaterialStatusSecond_PROJECT_STATUS_FROZEN MaterialStatusSecond = "PROJECT_STATUS_FROZEN"
	// MaterialStatusSecond_PROJECT_STATUS_REAUDIT 未投放-修改审核中
	MaterialStatusSecond_PROJECT_STATUS_REAUDIT MaterialStatusSecond = "PROJECT_STATUS_REAUDIT"
	// MaterialStatusSecond_PROJECT_STATUS_AUDIT 项目新建审核中
	MaterialStatusSecond_PROJECT_STATUS_AUDIT MaterialStatusSecond = "PROJECT_STATUS_AUDIT"
	// MaterialStatusSecond_PROJECT_AUDIT_REJECTED 项目审核不通过
	MaterialStatusSecond_PROJECT_AUDIT_REJECTED MaterialStatusSecond = "PROJECT_AUDIT_REJECTED"
	// MaterialStatusSecond_PROJECT_REACHED_DELIVERY_DURATION 项目已达投放时长
	MaterialStatusSecond_PROJECT_REACHED_DELIVERY_DURATION MaterialStatusSecond = "PROJECT_REACHED_DELIVERY_DURATION"
	// MaterialStatusSecond_SHARED_WALLET_EXCEEDS_BUDGET 共享钱包超出预算
	MaterialStatusSecond_SHARED_WALLET_EXCEEDS_BUDGET MaterialStatusSecond = "SHARED_WALLET_EXCEEDS_BUDGET"
	// MaterialStatusSecond_ACCOUNT_AVATAR_NOT_AVAILABLE_FOR_DELIVERY 账户头像不可投放
	MaterialStatusSecond_ACCOUNT_AVATAR_NOT_AVAILABLE_FOR_DELIVERY MaterialStatusSecond = "ACCOUNT_AVATAR_NOT_AVAILABLE_FOR_DELIVERY"
	// MaterialStatusSecond_APP_NOT_AVAILABLE_FOR_DELIVERY 应用不可投放
	MaterialStatusSecond_APP_NOT_AVAILABLE_FOR_DELIVERY MaterialStatusSecond = "APP_NOT_AVAILABLE_FOR_DELIVERY"
	// MaterialStatusSecond_DOUYIN_ITEM_NOT_AVAILABLE_FOR_DELIVERY 抖音素材不可投放
	MaterialStatusSecond_DOUYIN_ITEM_NOT_AVAILABLE_FOR_DELIVERY MaterialStatusSecond = "DOUYIN_ITEM_NOT_AVAILABLE_FOR_DELIVERY"
	// MaterialStatusSecond_GUIDE_VIDEO_NOT_EXIST 引导视频不存在
	MaterialStatusSecond_GUIDE_VIDEO_NOT_EXIST MaterialStatusSecond = "GUIDE_VIDEO_NOT_EXIST"
	// MaterialStatusSecond_LACK_BASIC_MATERIAL_STATUS 项目素材待补充
	MaterialStatusSecond_LACK_BASIC_MATERIAL_STATUS MaterialStatusSecond = "LACK_BASIC_MATERIAL_STATUS"
	// MaterialStatusSecond_LIVE_ROOM_NOT_AVAILABLE_FOR_DELIVERY 直播间不可投/不可投放
	MaterialStatusSecond_LIVE_ROOM_NOT_AVAILABLE_FOR_DELIVERY MaterialStatusSecond = "LIVE_ROOM_NOT_AVAILABLE_FOR_DELIVERY"
	// 图文素材二级状态
	MaterialStatusSecond_ACCOUNT_EXCEEDS_BUDGET         MaterialStatusSecond = "ACCOUNT_EXCEEDS_BUDGET"
	MaterialStatusSecond_ALL_KEYWORDS_DELETED           MaterialStatusSecond = "ALL_KEYWORDS_DELETED"
	MaterialStatusSecond_ALL_MATERIALS_DELETED          MaterialStatusSecond = "ALL_MATERIALS_DELETED"
	MaterialStatusSecond_AUDIT_REJECTED                 MaterialStatusSecond = "AUDIT_REJECTED"
	MaterialStatusSecond_DOUYIN_ACCOUNT_NOT_AVAILABLE   MaterialStatusSecond = "DOUYIN_ACCOUNT_NOT_AVAILABLE_FOR_DELIVERY"
	MaterialStatusSecond_GOODS_NOT_AVAILABLE            MaterialStatusSecond = "GOODS_NOT_AVAILABLE_FOR_DELIVERY"
	MaterialStatusSecond_LIVE_ROOM_NOT_STARTED          MaterialStatusSecond = "LIVE_ROOM_NOT_STARTED"
	MaterialStatusSecond_MANUAL_PAUSE                   MaterialStatusSecond = "MANUAL_PAUSE"
	MaterialStatusSecond_MODIFY_AUDITING                MaterialStatusSecond = "MODIFY_AUDITING"
	MaterialStatusSecond_NEW_CREATED_AUDITING           MaterialStatusSecond = "NEW_CREATED_AUDITING"
	MaterialStatusSecond_PAUSED                         MaterialStatusSecond = "PAUSED"
	MaterialStatusSecond_PROJECT_AUDIENCE_SETTING_CONFLICT MaterialStatusSecond = "PROJECT_AUDIENCE_SETTING_CONFLICT"
	MaterialStatusSecond_PROJECT_BALANCE_INSUFFICIENT   MaterialStatusSecond = "PROJECT_BALANCE_INSUFFICIENT"
	MaterialStatusSecond_PROJECT_COMPLETED             MaterialStatusSecond = "PROJECT_COMPLETED"
	MaterialStatusSecond_PROJECT_DELETED               MaterialStatusSecond = "PROJECT_DELETED"
	MaterialStatusSecond_PROJECT_EXCEEDS_BUDGET         MaterialStatusSecond = "PROJECT_EXCEEDS_BUDGET"
	MaterialStatusSecond_PROJECT_MODIFY_AUDITING        MaterialStatusSecond = "PROJECT_MODIFY_AUDITING"
	MaterialStatusSecond_PROJECT_NEW_CREATED_AUDITING  MaterialStatusSecond = "PROJECT_NEW_CREATED_AUDITING"
	MaterialStatusSecond_PROJECT_NOT_IN_DELIVERY_PERIOD MaterialStatusSecond = "PROJECT_NOT_IN_DELIVERY_PERIOD"
	MaterialStatusSecond_PROJECT_NOT_REACH_DELIVERY_TIME MaterialStatusSecond = "PROJECT_NOT_REACH_DELIVERY_TIME"
	MaterialStatusSecond_PROJECT_PAUSED                MaterialStatusSecond = "PROJECT_PAUSED"
	MaterialStatusSecond_PROJECT_TERMINATED            MaterialStatusSecond = "PROJECT_TERMINATED"
	MaterialStatusSecond_REACHED_REGULATION_BUDGET_UPPER_LIMIT MaterialStatusSecond = "REACHED_REGULATION_BUDGET_UPPER_LIMIT"
	MaterialStatusSecond_TASK_BID_EXCEEDS_PROJECT_BUDGET MaterialStatusSecond = "TASK_BID_EXCEEDS_PROJECT_BUDGET"
	MaterialStatusSecond_TASK_DELETED                  MaterialStatusSecond = "TASK_DELETED"
	MaterialStatusSecond_TASK_REACHED_REGULATION_DURATION MaterialStatusSecond = "TASK_REACHED_REGULATION_DURATION"
	MaterialStatusSecond_TASK_TERMINATED               MaterialStatusSecond = "TASK_TERMINATED"
)
