package creativecomponent

import (
	"encoding/json"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// GetRequest 查询组件列表 API Request
type GetRequest struct {
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	Page         uint64 `json:"page,omitempty"`
	PageSize     uint64 `json:"page_size,omitempty"`
	Filtering    struct {
		ComponentID    uint64   `json:"component_id,omitempty"`
		ComponentName  string   `json:"component_name,omitempty"`
		ComponentTypes []string `json:"component_types,omitempty"`
		Status         []string `json:"status,omitempty"`
	} `json:"filtering,omitempty"`
}

func (r GetRequest) GetJsonBody() []byte {
	ret, _ := json.Marshal(r)

	return ret
}

type GetResponse struct {
	model.BaseResponse
	Data *GetResponseData `json:"data,omitempty"`
}

type GetResponseData struct {
	List []struct {
		ComponentID   uint64          `json:"component_id,omitempty"`
		ComponentName string          `json:"component_name,omitempty"`
		ComponentType string          `json:"component_type,omitempty"`
		ComponentData json.RawMessage `json:"component_data,omitempty"`
		CreateTime    string          `json:"create_time,omitempty"`
		Status        string          `json:"status,omitempty"`
	} `json:"list,omitempty"`
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
