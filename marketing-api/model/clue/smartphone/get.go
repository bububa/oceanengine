package smartphone

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取智能电话列表 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// InstanceIDs 智能电话ID列表; 若指定instance_ids进行过滤，其他的过滤条件将不生效
	InstanceIDs []string `json:"instance_ids,omitempty"`
	// StartTime 起始时间，格式：%Y-%m-%d，默认7天前，即不指定起止时间获取最近7天数据
	StartTime string `json:"start_time,omitempty"`
	// EndTime 截止时间，格式：%Y-%m-%d，默认当天，即不指定起止时间获取最近7天数据
	EndTime string `json:"end_time,omitempty"`
	// IsDel 删除标记，默认获取未删除的数据; 可选值：0，1（0表示未删除，1表示已删除）
	IsDel int `json:"is_del,omitempty"`
	// Page 页码，默认值为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值为10，最大不超过50
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement GetRequest interface
func (r GetRequest) Encode() string {
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
	if r.IsDel > 0 {
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

// GetResponse 获取智能电话列表 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *GetResponseData `json:"data,omitempty"`
}

// GetResponseData json返回值
type GetResponseData struct {
	// Key advertiser_id的加密字段
	Key string `json:"key,omitempty"`
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 智能电话列表
	List []SmartPhone `json:"list,omitempty"`
}
