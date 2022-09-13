package union

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// FlowPackageDeleteRequest 删除穿山甲流量包 API Request
type FlowPackageDeleteRequest struct {
	// AdvertiserID 获取穿山甲流量包
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// FlowPackageID 流量包ID，从【获取穿山甲流量包】获取
	FlowPackageID uint64 `json:"flow_package_id,omitempty"`
}

// Encode implement PostRequest interface
func (r FlowPackageDeleteRequest) Encode() []byte {
	bs, _ := json.Marshal(r)
	return bs
}

// FlowPackageDeleteResponse 删除`j穿山甲流量包 API Response
type FlowPackageDeleteResponse struct {
	model.BaseResponse
	// Data json 返回值`
	Data struct {
		// FlowPackageID 流量包ID``
		FlowPackageID uint64 `json:"flow_package_id,omitempty"`
	} `json:"data,omitempty"`
}
