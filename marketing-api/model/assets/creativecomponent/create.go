package creativecomponent

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// CreateRequest 创建组件 API Request
type CreateRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ComponentInfo 组件信息
	ComponentInfo *ComponentInfo `json:"component_info,omitempty"`
}

// Encode implement Post Request interface
func (r CreateRequest) Encode() []byte {
	ret, _ := json.Marshal(r)

	return ret
}

// CreateResponse 创建组件 API Response
type CreateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *CreateResponseData `json:"data,omitempty"`
}

// CreateResponseData json返回值
type CreateResponseData struct {
	AdvertiserID uint64               `json:"advertiser_id,omitempty"`
	ComponentID  uint64               `json:"component_id,omitempty"`
	CreateTime   string               `json:"create_time,omitempty"`
	ModifyTime   string               `json:"modify_time,omitempty"`
	Status       enum.ComponentStatus `json:"status,omitempty"`
}
