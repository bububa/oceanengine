package clue

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// FormGetRequest 建站工具——查询已有表单列表 API Request
type FormGetRequest struct {
	// AdvertiserID 广告主id，范围：1 <= advertiser_id <= MAX_INT64
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页码，默认为1，范围：page >= 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认10条每页，范围：page_size >= 1
	PageSize int `json:"page_size,omitempty"`
	// StartTIme 起始时间，格式：%Y-%m-%d，默认三天前
	StartTime string `json:"start_time,omitempty"`
	// EndTime 截止时间，格式：%Y-%m-%d，默认今天
	EndTime string `json:"end_time,omitempty"`
	// InstanceID 表单ID
	InstanceID uint64 `json:"instance_id,omitempty"`
	// IsDel 是否删除，允许值：1（未删除）2（已删除）默认未删除
	IsDel int `json:"is_del,omitempty"`
	// FormType 表单类型，允许值：
	// NORMAL_FORM（普通表单）
	// ADVANCED_CREATIVE_FORM（附加创意表单）
	FormType FormType `json:"form_type,omitempty"`
}

// Encode implement GetRequest interface
func (r FormGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	if r.StartTime != "" {
		values.Set("start_time", r.StartTime)
	}
	if r.EndTime != "" {
		values.Set("end_time", r.EndTime)
	}
	if r.InstanceID > 0 {
		values.Set("instance_id", strconv.FormatUint(r.InstanceID, 10))
	}
	if r.IsDel == 1 || r.IsDel == 2 {
		values.Set("is_del", strconv.Itoa(r.IsDel))
	}
	if r.FormType != "" {
		values.Set("form_type", string(r.FormType))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// FormGetResponse 建站工具——查询已有表单列表 API Response
type FormGetResponse struct {
	model.BaseResponse
	// Data json 返回值
	Data *FormGetResponseData `json:"data,omitempty"`
}

// FormGetResponseData json 返回值
type FormGetResponseData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 表单列表
	List []Form `json:"list,omitempty"`
}
