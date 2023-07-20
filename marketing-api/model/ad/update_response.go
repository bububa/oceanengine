package ad

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// UpdateResponse 更新计划API Response
type UpdateResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *UpdateResponseData `json:"data,omitempty"`
}

// UpdateResponseData json返回值
type UpdateResponseData struct {
	// AdID 广告计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// NeedAudit 此次修改是否触发进入待审状态(1表示进入待审状态,0表示不进入待审状态)
	NeedAudit int `json:"need_audit,omitempty"`
	// AdIDs 广告计划ID集合
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// Errors 更新失败的广告计划列表
	Errors []UpdateError `json:"errors,omitempty"`
}

// UpdateError 更新失败的广告计划
type UpdateError struct {
	// AdID 广告计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// ErrorMessage 错误信息
	ErrorMessage string `json:"error_message"`
}

// Error implement error interface
func (r UpdateError) Error() string {
	return util.StringsJoin("广告计划ID:", strconv.FormatUint(r.AdID, 10), ", 错误信息", r.ErrorMessage)
}
