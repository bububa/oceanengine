package quickappmanagement

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// QuickAppGetRequest 查询快应用信息 API Request
type QuickAppGetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// QuickAppIDs 快应用ids，单次上限100
	QuickAppIDs []uint64 `json:"quick_app_ids,omitempty"`
	// SearchKey 搜索关键字，快应用名称或者包名的模糊搜索，可以为空，可以传中文
	SearchKey string `json:"search_key,omitempty"`
	// UpdateTime 按快应用的更新时间查询，需要输入时间范围
	UpdateTime *UpdateTime `json:"update_time,omitempty"`
	// Status 应用状态；允许值：
	// AUDIT_DOING:审核中
	// AUDIT_SEND_REJECTED：送审失败
	// AUDIT_REJECTED:审核失败
	// AUDIT_ACCEPTED:审核成功
	// REMOVED：已下架
	Status enum.QuickAppStatus `json:"status,omitempty"`
	// Page 页码，默认值为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值为10，最大不超过100
	PageSize int `json:"page_size,omitempty"`
}

// UpdateTime 按快应用的更新时间查询，需要输入时间范围
type UpdateTime struct {
	// StartTime 更新起始时间，默认为空。格式为i64类型的时间戳（单位：秒）
	StartTime int64 `json:"start_time,omitempty"`
	// EndTime 更新结束时间，默认为当天。格式为i64类型的时间戳（单位：秒）
	EndTime int64 `json:"end_time,omitempty"`
}

func (r QuickAppGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if len(r.QuickAppIDs) > 0 {
		values.Set("quick_app_ids", string(util.JSONMarshal(r.QuickAppIDs)))
	}
	if r.Status != "" {
		values.Set("status", string(r.Status))
	}
	if r.UpdateTime != nil {
		values.Set("update_time", string(util.JSONMarshal(r.UpdateTime)))
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// QuickAppGetResponse 查询快应用信息 API Response
type QuickAppGetResponse struct {
	model.BaseResponse
	Data *QuickAppGetResult `json:"data,omitempty"`
}

type QuickAppGetResult struct {
	// QuickAppInfo
	QuickAppInfo []QuickApp `json:"quick_app_info,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
}
