package security

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// OpenMaterialAuditRequest 广告素材预审 API Request
type OpenMaterialAuditRequest struct {
	// AccountID 业务线对应广告主ID
	AccountID uint64 `json:"account_id,omitempty"`
	// BusinessType 业务线，允许值：
	// AD 巨量广告业务
	// QIANCHUAN 千川业务
	// LOCAL_LIFE 巨量本地推业务
	BusinessType string `json:"business_type,omitempty"`
	// Type 素材类型，允许值：
	// VIDEO 视频
	Type string `json:"type,omitempty"`
	// Data 素材内容
	// 当type = VIDEO 视频时，填写视频对应video_id
	Data string `json:"data,omitempty"`
	// MsgType 操作内容，允许值：
	// SEND 送审
	// RECALL 取消送审
	MsgType string `json:"msg_type,omitempty"`
}

// Encode implements PostRequest interface
func (r OpenMaterialAuditRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// OpenMaterialAuditResponse 广告素材预审 API Response
type OpenMaterialAuditResponse struct {
	model.BaseResponse
	Data *OpenMaterialAuditResult `json:"data,omitempty"`
}

type OpenMaterialAuditResult struct {
	// Result 送审是否成功
	Result bool `json:"result,omitempty"`
	// TaskID 链路task_id，每次送审生成一个唯一值
	TaskID uint64 `json:"task_id,omitempty"`
	// ObjectID 审核对象id
	ObjectID uint64 `json:"object_id,omitempty"`
}
