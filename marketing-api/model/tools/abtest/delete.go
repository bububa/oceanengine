package abtest

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DeleteRequest 删除实验 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TestIDs 实验ID列表，目前仅支持一个实验ID
	TestIDs []uint64 `json:"test_ids,omitempty"`
}

func (r DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

type DeleteResponse struct {
	model.BaseResponse
	Data *DeleteResult `json:"data,omitempty"`
}

type DeleteResult struct {
	// Success 删除成功的实验列表
	Success []uint64 `json:"success,omitempty"`
	// Errors 删除失败的实验列表
	Errors []Error `json:"errors,omitempty"`
}
