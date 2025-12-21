package rejectmaterial

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AIRepairCrossAccountGetRequest 根据mid查询同主体账户下修复建议列表 API Request
type AIRepairCrossAccountGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MaterialID 素材Id
	MaterialID uint64 `json:"material_id,omitempty"`
}

// Encode implements GetRequest interface
func (r AIRepairCrossAccountGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("material_id", strconv.FormatUint(r.MaterialID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AIRepairCrossAccountGetResponse 根据mid查询同主体账户下修复建议列表 API Response
type AIRepairCrossAccountGetResponse struct {
	model.BaseResponse
	Data *AIRepairCrossAccountGetResult
}

type AIRepairCrossAccountGetResult struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MaterialID 素材Id
	MaterialID uint64 `json:"material_id,omitempty"`
	// AIRepairID 拒审素材在当前账户下的修复id，可能有多个
	AIRepairID []uint64 `json:"ai_repair_id,omitempty"`
	// RejectMaterialRelatedRepairTasks 原账户&拒审素材相关的同主体账户及修复建议
	RejectMaterialRelatedRepairTasks []RejectMaterialRelatedRepairTask `json:"reject_material_related_repair_tasks,omitempty"`
}

// RejectMaterialRelatedRepairTask 原账户&拒审素材相关的同主体账户及修复建议
type RejectMaterialRelatedRepairTask struct {
	// RelatedAdvertiserID 关联的同主体广告主id
	RelatedAdvertiserID uint64 `json:"related_advertiser_id,omitempty"`
	// RelatedAIRepairIDs 修复建议id列表，获取后可用于采纳接口
	RelatedAIRepairIDs []uint64 `json:"related_ai_repair_ids,omitempty"`
}
