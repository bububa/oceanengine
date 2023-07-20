package campaign

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateResponse 广告组更新 API Response
type UpdateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *UpdateResponseData `json:"data,omitempty"`
}

// UpdateResponseData json返回值
type UpdateResponseData struct {
	// CampaignID 广告组id
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CampaignIDs 广告组id
	CampaignIDs []uint64 `json:"campaign_ids,omitempty"`
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

// Error implement error interface
func (r UpdateError) Error() string {
	return util.StringsJoin("广告组ID: ", strconv.FormatUint(r.CampaignID, 10), ", 错误信息", r.ErrorMessage)
}
