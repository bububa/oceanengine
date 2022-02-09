package creativecomponent

import (
	"encoding/json"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// PostRequest 创建组件 API Request
type PostRequest struct {
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ComponentInfo 这里不详细拆分此结构体字段，因为这里的字段会根据不同条件而变化
	ComponentInfo json.RawMessage `json:"component_info,omitempty"`
}

func (r PostRequest) Encode() []byte {
	ret, _ := json.Marshal(r)

	return ret
}

type PostResponse struct {
	model.BaseResponse
	Data *PostResponseData `json:"data,omitempty"`
}

type PostResponseData struct {
	ComponentId  uint64 `json:"component_id,omitempty"`
	CreateTime   string `json:"create_time,omitempty"`
	AdvertiserId uint64 `json:"advertiser_id,omitempty"`
	ModifyTime   string `json:"modify_time,omitempty"`
	Status       string `json:"status,omitempty"`
}
