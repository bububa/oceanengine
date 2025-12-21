package hotmaterialderive

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 查询爆款裂变任务详情 API Request
type GetRequest struct {
	// AdvertiserID 广告主账户 ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TaskIDs 任务 ID，数量限制 1 - 50 个
	TaskIDs []uint64 `json:"task_ids,omitempty"`
}

// Encode implements GetRequest interface
func (r GetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("task_ids", string(util.JSONMarshal(r.TaskIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// GetResponse 查询爆款裂变任务详情 API Response
type GetResponse struct {
	model.BaseResponse
	Data struct {
		// TaskDetails 任务列表
		TaskDetails []Task `json:"task_details,omitempty"`
	} `json:"data,omitempty"`
}
