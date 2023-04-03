package abtest

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// StopRequest 关停实验 API Request
type StopRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TestIDs 实验ID列表，目前仅支持一个实验ID
	TestIDs []uint64 `json:"test_ids,omitempty"`
}

func (r StopRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

type StopResponse struct {
	model.BaseResponse
	Data *StopResult `json:"data,omitempty"`
}

type StopResult struct {
	// Success 关停成功的实验列表
	Success []uint64 `json:"success,omitempty"`
	// Errors 关停失败的实验列表，目前实验关停失败会直接报错。
	Errors []Error `json:"errors,omitempty"`
}
