package agent

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
)

type AdvertiserUpdateRequest struct {
	AdvertiserID uint64 `json:"advertiser_id,omitempty"` // 广告主ID
	Name         string `json:"name,omitempty"`          // 账户名称
	Contacter    string `json:"contacter,omitempty"`     // 联系人
	Phonenumber  string `json:"phonenumber,omitempty"`   // 手机号码
	Telephone    string `json:"telephone,omitempty"`     // 固定电话
}

func (r AdvertiserUpdateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

type AdvertiserUpdateResponse struct {
	model.BaseResponse
	Data *AdvertiserUpdateResponseData `json:"data,omitempty"`
}

type AdvertiserUpdateResponseData struct {
	AdvertiserID uint64 `json:"advertiser_id,omitempty"` // 广告主ID
	NeedAudit    int    `json:"need_audit,omitempty"`    // 此次修改是否触发进入待审状态(1表示进入待审状态,0表示不进入待审状态)
}
