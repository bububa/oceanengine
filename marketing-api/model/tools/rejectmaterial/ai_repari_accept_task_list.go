package rejectmaterial

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AIRepairAcceptTaskListRequest 获取采纳素材修复建议任务结果 API Request
type AIRepairAcceptTaskListRequest struct {
	// AdvertiserID 广告账户id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AIRepairIDs 已创建采纳任务的修复建议id
	AIRepairIDs []uint64 `json:"ai_repair_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r AIRepairAcceptTaskListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AIRepairAcceptTaskListResponse 获取采纳素材修复建议任务结果 API Response
type AIRepairAcceptTaskListResponse struct {
	model.BaseResponse
	Data struct {
		// ResultList 采纳结果
		ResultList []AIRepairAcceptTask `json:"result_list,omitempty"`
	}
}
