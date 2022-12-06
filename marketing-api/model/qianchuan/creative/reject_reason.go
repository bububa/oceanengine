package creative

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// RejectReasonRequest 获取创意审核建议 API Request
type RejectReasonRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CreativeIDs 查询审核意见的的创意id
	CreativeIDs []uint64 `json:"creative_ids,omitempty"`
}

// Encode implement GetRequest interface
func (r RejectReasonRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	ids, _ := json.Marshal(r.CreativeIDs)
	values.Set("creative_ids", string(ids))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// RejectReasonResponse 获取创意审核建议  API Response
type RejectReasonResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		List []RejectReasonList `json:"list,omitempty"`
	} `json:"data,omitempty"`
}

// RejectReasonList 审核详细信息
type RejectReasonList struct {
	// 广告创意id
	CreativeID uint64 `json:"creative_id,omitempty"`
	// AuditRecords 审核详细内容
	AuditRecords []AuditRecord `json:"audit_records,omitempty"`
}
