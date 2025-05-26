package rejectmaterial

import "github.com/bububa/oceanengine/marketing-api/enum"

type AIRepairAcceptTask struct {
	// AIRepairID 已创建采纳任务的修复建议id
	AIRepairID uint64 `json:"ai_repair_id,omitempty"`
	// AIRepairStatus 采纳状态，枚举值如下
	// SUCCESS 采纳成功
	// FAILED  采纳失败
	// PROCESSING 采纳中
	// NO_TASK  采纳任务不存在，可能为未创建采纳任务/已过期/广告已删除/广告下素材已删除
	AIRepairStatus enum.AIRepairStatus `json:"ai_repair_status,omitempty"`
	// AIRepairErrorMessage 采纳失败的原因，仅当ai_repair_status=FAILED 时返回
	AIRepairErrorMessage string `json:"ai_repair_error_message,omitempty"`
}
