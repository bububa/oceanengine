package ad

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
)

type UpdateStatusRequest struct {
	AdvertiserID uint64   `json:"advertiser_id,omitempty"` // 广告主ID
	AdIDs        []uint64 `json:"ad_ids,omitempty"`        // 计划ID集合，限制1～100. 广告计划id需属于广告主
	OptStatus    string   `json:"opt_status,omitempty"`    // 操作, "enable"表示启用, "delete"表示删除, "disable"表示暂停；允许值: "enable", "delete", "disable",; 对于删除的广告组不可进行任何操作，否则会报错哈
}

func (r UpdateStatusRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}

type UpdateStatusResponse struct {
	model.BaseResponse
	Data *UpdateStatusResponseData `json:"data,omitempty"`
}

type UpdateStatusResponseData struct {
	AdIDs []uint64 `json:"ad_ids,omitempty"`
}
