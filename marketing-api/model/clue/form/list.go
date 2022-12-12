package form

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ListRequest 获取表单列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// InstanceIDs 表单实例ID，可从【表单创建】【表单列表获取】接口获取; 若指定instance_ids进行过滤，其他的过滤条件将不生效
	InstanceIDs []uint64 `json:"instance_ids,omitempty"`
	// StartTime 起始时间，格式：%Y-%m-%d，默认7天前，即不指定起止时间获取最近7天数据
	StartTime string `json:"start_time,omitempty"`
	// EndTime 截止时间，格式：%Y-%m-%d，默认当天，即不指定起止时间获取最近7天数据
	EndTime string `json:"end_time,omitempty"`
	// IsDel 删除标记，默认获取未删除的数据; 可选值：0，1（0表示未删除，1表示已删除）
	IsDel int `json:"is_del,omitempty"`
	// Page 起始页，默认值为1
	Page int `json:"page,omitempty"`
	// PageSize 页大小，默认值为10，最大不超过50
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r ListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if len(r.InstanceIDs) > 0 {
		ids, _ := json.Marshal(r.InstanceIDs)
		values.Set("instance_ids", string(ids))
	}
	if r.StartTime != "" {
		values.Set("start_time", r.StartTime)
	}
	if r.EndTime != "" {
		values.Set("end_time", r.EndTime)
	}
	if r.IsDel == 1 {
		values.Set("is_del", strconv.Itoa(r.IsDel))
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

// ListResponse 获取表单列表 API Response
type ListResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *ListResponseData `json:"data,omitempty"`
}

// ListResponseData json返回值
type ListResponseData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 表单列表
	List []Form `json:"list,omitempty"`
}
