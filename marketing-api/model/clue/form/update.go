package form

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateRequest 更新表单 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// InstanceID 表单实例ID
	InstanceID uint64 `json:"instance_id,omitempty"`
	// Name 表单名称，不超过20个字符
	Name string `json:"name,omitempty"`
	// SubmitText 提交文案说明，如“提交表单”，文案长度不超过8个字符
	SubmitText string `json:"submit_text,omitempty"`
	// EnableLayer 是否开启分层，默认不开启; 枚举值：0，1（0表示不启用，1表示启用）
	EnableLayer int `json:"enable_layer,omitempty"`
	// LayerSubmitText 分层提交文案说明，分层信息在Elements中，当不开启分层时，所有分层信息无效。文案长度不超过8个字符
	LayerSubmitText string `json:"layer_submit_text,omitempty"`
	// ExtendInfo 扩展属性
	ExtendInfo *ExtendInfo `json:"extend_info,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// UpdateResponse 更新表单 API Response
type UpdateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data struct {
		// InstanceID 创建的表单实例id
		InstanceID uint64 `json:"instance_id,omitempty"`
	} `json:"data,omitempty"`
}
