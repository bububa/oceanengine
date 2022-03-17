package form

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// DeleteRequest 删除表单 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// InstanceID 表单实例ID
	InstanceID uint64 `json:"instance_id,omitempty"`
}

// Encode implement PostRequest interface
func (r DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// DeleteResponse 删除表单 API Response
type DeleteResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// Success 删除结果，true表示成功，false表示失败
		Success bool `json:"success,omitempty"`
	} `json:"data,omitempty"`
}
