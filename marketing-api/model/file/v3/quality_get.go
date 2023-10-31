package v3

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QualityGetRequest 连山投前分析结果查询 API Request
type QualityGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// MaterialIDs 预审ID
	MaterialIDs []uint64 `json:"material_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r QualityGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("material_ids", string(util.JSONMarshal(r.MaterialIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// QualityGetResponse 连山投前分析结果查询 API Response
type QualityGetResponse struct {
	model.BaseResponse
	Data struct {
		// List
		List []MaterialQuality `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

type MaterialQuality struct {
	// MaterialID 预审ID
	MaterialID uint64 `json:"material_id,omitempty"`
	// Status  可选值:
	// AUDITING 审核中
	// AUDIT_ACCEPTED 预审通过
	// AUDIT_FAILED 预审失败
	// AUDIT_REJECT 预审拒绝
	// AUDIT_SUBMIT 发起预审
	// AUDIT_TIMEOUT 预审超时
	Status enum.MaterialQualityStatus `json:"status,omitempty"`
	// Errmsg 预审拒绝视频此处会展示拒审原因
	Errmsg string `json:"errmsg,omitempty"`
	// IsFirstPublishMaterial 是否是首发素材
	IsFirstPublishMaterial bool `json:"is_first_publish_material,omitempty"`
	// IsSimilarQueueMaterial 是否同质化素材风险-排队投放素材
	IsSimilarQueueMaterial bool `json:"is_similar_queue_material,omitempty"`
	// IsSimilarExpectedQueueMaterial 是否同质化素材风险-未投放预计排队素材
	IsSimilarExpectedQueueMaterial bool `json:"is_similar_expected_queue_material,omitempty"`
}
