package agent

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AdvertiserUpdateRequest 修改广告主 API Request
type AdvertiserUpdateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Name 账户名称
	Name string `json:"name,omitempty"`
	// Contacter 联系人
	Contacter string `json:"contacter,omitempty"`
	// Phonenumber 手机号码
	Phonenumber string `json:"phonenumber,omitempty"`
	// Telephone 固定电话
	Telephone string `json:"telephone,omitempty"`
	// Note 备注
	Note string `json:"note,omitempty"`
}

// Encode implement PostRequest interface
func (r AdvertiserUpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AdvertiserUpdateResponse 修改广告主 API Response
type AdvertiserUpdateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *AdvertiserUpdateResponseData `json:"data,omitempty"`
}

// AdvertiserUpdateResponseData json返回值
type AdvertiserUpdateResponseData struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// NeedAudit 此次修改是否触发进入待审状态(1表示进入待审状态,0表示不进入待审状态)
	NeedAudit int `json:"need_audit,omitempty"`
}
