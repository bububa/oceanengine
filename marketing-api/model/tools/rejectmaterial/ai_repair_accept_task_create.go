package rejectmaterial

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AIRepairAcceptTaskCreateRequest 创建采纳「拒审素材修复建议」任务 API Request
type AIRepairAcceptTaskCreateRequest struct {
	// AdvertiserID 广告账户ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AIRepairIDs 需要采纳的修复建议id列表，一次最多50个
	// 可从「获取拒审素材修复建议」接口获取
	AIRepairIDs []uint64 `json:"ai_repair_ids,omitempty"`
}

// Encode implements PostRequest interface
func (r AIRepairAcceptTaskCreateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AIRepairAcceptTaskCreateResponse 创建采纳「拒审素材修复建议」任务 API Response
type AIRepairAcceptTaskCreateResponse struct {
	model.BaseResponse
	Data *AIRepairAcceptTaskCreateResult `json:"data,omitempty"`
}

type AIRepairAcceptTaskCreateResult struct {
	// AdoptAIRepairIDs 创建采纳任务成功的修复id列表
	AdoptAIRepairIDs []uint64 `json:"adopt_ai_repair_ids,omitempty"`
	// AIRepairAcceptTaskCreateErrors 创建采纳任务失败详情
	Errors []AIRepairAcceptTaskCreateError `json:"errors,omitempty"`
}

// AIRepairAcceptTaskCreateError 创建采纳任务失败详情
type AIRepairAcceptTaskCreateError struct {
	// ErrorMessage 创建采纳任务失败的原因
	ErrorMessage string `json:"error_message,omitempty"`
	// AIRepairID 创建采纳任务失败的修复id
	AIRepairID uint64 `json:"ai_repair_id,omitempty"`
}

func (r AIRepairAcceptTaskCreateError) Error() string {
	return util.StringsJoin("ai_repair_id:", strconv.FormatUint(r.AIRepairID, 10), "message: ", r.ErrorMessage)
}
