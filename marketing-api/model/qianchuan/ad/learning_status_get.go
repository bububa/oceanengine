package ad

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// LearningStatusGetRequest 获取计划学习期状态 API Request
type LearningStatusGetRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdIDs 计划id列表，每次最多传入50个
	AdIDs []uint64 `json:"ad_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r LearningStatusGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("ad_ids", string(util.JSONMarshal(r.AdIDs)))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// LearningStatusGetResponse 获取计划学习期状态 API Response
type LearningStatusGetResponse struct {
	model.BaseResponse
	Data struct {
		// List 计划列表
		List []LearningStatus `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// LearingStatus 计划学习期状态
type LearningStatus struct {
	// AdID 计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// Status 学习期状态，允许值：
	// LEARNING（学习期）
	// LEARNED（学习期结束）
	// LEARN_FAILED（学习期失败)
	// DEFAULT：无学习期状态
	// 具体可以参考此文档的说明：关于学习期
	Status qianchuan.LearningStatus `json:"status,omitempty"`
}
