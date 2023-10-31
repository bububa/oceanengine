package ad

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/model/qianchuan/creative"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// RejectReasonRequest 获取计划审核建议 API Request
type RejectReasonRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// AdIDs 查询审核意见的的计划id
	AdIDs []uint64 `json:"ad_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r RejectReasonRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ids, _ := json.Marshal(r.AdIDs)
	values.Set("ad_ids", string(ids))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// RejectReasonResponse 获取计划审核建议  API Response
type RejectReasonResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		List []RejectReasonList `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// RejectReasonList 审核详细信息
type RejectReasonList struct {
	// AdID 广告计划id
	AdID uint64 `json:"ad_id,omitempty"`
	// AuditRecords 审核详细内容
	AuditRecords []creative.AuditRecord `json:"audit_records,omitempty"`
}
