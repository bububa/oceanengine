package coupon

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// GetRequest 获取卡券列表 API Request
type GetRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// ActivityIDs 青鸟营销活动ID列表，当设置时，会忽略其他搜索条件
	ActivityIDs []uint64 `json:"activity_ids,omitempty"`
	// ActivityTypes 卡券类型，可选值:
	// DIRECT_NEED_PHONE：直接发券，收集手机号，need_phone须为true
	// DIRECT_NOT_NEED_PHONE ：直接发券，不收集手机号，need_phone须为false 并且 BindFormId会被忽略
	ActivityTypes []ActivityType `json:"activity_types,omitempty"`
	// StartTime 起始时间，格式：%Y-%m-%d，默认7天前，即不指定起止时间获取最近7天数据
	StartTime string `json:"start_time,omitempty"`
	// EndTime 截止时间，格式：%Y-%m-%d，默认当天，即不指定起止时间获取最近7天数据
	EndTime string `json:"end_time,omitempty"`
	// IsDel 删除标记，默认获取未删除的数据
	// 可选值：0，1（0表示未删除，1表示已删除）
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
	if len(r.ActivityIDs) > 0 {
		ids, _ := json.Marshal(r.ActivityIDs)
		values.Set("activity_ids", string(ids))
	}
	if len(r.ActivityTypes) > 0 {
		types, _ := json.Marshal(r.ActivityTypes)
		values.Set("activity_types", string(types))
	}
	if r.StartTime != "" {
		values.Set("start_time", r.StartTime)
	}
	if r.EndTime != "" {
		values.Set("end_time", r.EndTime)
	}
	if r.IsDel == 1 {
		values.Set("is_del", "1")
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

// GetResponse 获取卡券列表 API Response
type GetResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *GetResponseData `json:"data,omitempty"`
}

// GetResponseData json返回值
type GetResponseData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 卡券列表
	List []DetailResponseData `json:"list,omitempty"`
}
