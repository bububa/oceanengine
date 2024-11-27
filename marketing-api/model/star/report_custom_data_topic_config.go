package star

import "github.com/bububa/oceanengine/marketing-api/model"

// ReportCustomDataTopicConfigResponse 获取投后数据主题累计数据 API Response
type ReportCustomDataTopicConfigResponse struct {
	Data *ReportCustomDataTopicConfigResult `json:"data,omitempty"`
	model.BaseResponse
}

type ReportCustomDataTopicConfigResult struct {
	// Data 对应请求的数据主题数组
	Data []DataTopicConfig `json:"data,omitempty"`
	// ItemID 请求的交付物Id
	ItemID uint64 `json:"item_id,omitempty"`
	// DemandID 请求的任务Id
	DemandID uint64 `json:"demand_id,omitempty"`
}
