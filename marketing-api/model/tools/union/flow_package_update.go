package union

import (
	"encoding/json"

	"github.com/bububa/oceanengine/marketing-api/model"
)

// FlowPackageUpdateRequest 修改穿山甲流量包 API Request
type FlowPackageUpdateRequest struct {
	// AdvertiserID 获取穿山甲流量包
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Name 流量包名称，限制：[1-20] 个字符
	Name string `json:"name,omitempty"`
	// FlowPackageID 流量包ID，从【获取穿山甲流量包】获取
	FlowPackageID uint64 `json:"flow_package_id,omitempty"`
	// Rit 穿山甲广告位，限制：[1-100]
	Rit []uint64 `json:"rit,omitempty"`
}

// Encode implement PostRequest interface
func (r FlowPackageUpdateRequest) Encode() []byte {
	bs, _ := json.Marshal(r)
	return bs
}

// FlowPackageUpdateResponse 修改穿山甲流量包 API Response
type FlowPackageUpdateResponse struct {
	model.BaseResponse
	// Data json 返回值`
	Data struct {
		// FlowPackageID 流量包ID``
		FlowPackageID uint64 `json:"flow_package_id,omitempty"`
	} `json:"data,omitempty"`
}
