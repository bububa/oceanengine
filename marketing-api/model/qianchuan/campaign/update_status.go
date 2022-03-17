package campaign

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateStatusRequest 广告组更新状态 API Request
type UpdateStatusRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignIDs 广告组ID，不超过100个，且广告组ID属于广告主ID否则会报错；
	CampaignIDs []uint64 `json:"campaign_ids,omitempty"`
	// 操作类型，允许值: "ENABLE"：启用, "DELETE"：删除, "DISABLE"：暂停；对于删除的广告组不可进行任何操作。
	OptStatus string `json:"opt_status,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateStatusRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

type UpdateStatusResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *UpdateStatusResponseData `json:"data,omitempty"`
}

// UpdateStatusResponseData json返回值
type UpdateStatusResponseData struct {
	// Success 更新成功的广告组ID列表
	Success []uint64 `json:"success,omitempty"`
	// Errors 更新失败的广告组列表
	Errors []UpdateError `json:"errors,omitempty"`
}

// UpdateError 更新失败错误
type UpdateError struct {
	// CampaignID 广告组ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// ErrorMessage 错误信息
	ErrorMessage string `json:"error_message,omitempty"`
}
