package subscribe

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// AccountsListRequest 查询订阅 Adv API Request
type AccountsListRequest struct {
	// AppID 开放平台调用MAPI/订阅RDS时所使用的APPID
	AppID uint64 `json:"app_id,omitempty"`
	// SubscribeTaskID 创建成功的RDS订阅任务所生成的订阅ID
	SubscribeTaskID uint64 `json:"subscribe_task_id,omitempty"`
	// Events 子事件扩展字段，可不填
	Events []string `json:"events,omitempty"`
	// CoreUserID 授权用户ID，获取方式https://open.oceanengine.com/labels/7/docs/1696710507039756
	CoreUserID uint64 `json:"core_user_id,omitempty"`
	// AdvertiserIDs 授权广告主账号ID
	AdvertiserIDs []uint64 `json:"advertiser_ids,omitempty"`
	// Statuses 可选值:
	// OK 推送中
	// PENDING 新增状态
	// UNAUTHORIZED 无权限
	// UNKNOWN 未知
	Statuses []enum.SubscribeAdvStatus `json:"statuses,omitempty"`
	// Cursor 游标，默认值 0
	Cursor int `json:"cursor,omitempty"`
	// Count 页大小，默认 100， 最大值 500
	Count int `json:"count,omitempty"`
}

// Encode implement GetRequest interface
func (r AccountsListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("app_id", strconv.FormatUint(r.AppID, 10))
	values.Set("subscribe_task_id", strconv.FormatUint(r.SubscribeTaskID, 10))
	if len(r.Events) > 0 {
		values.Set("events", string(util.JSONMarshal(r.Events)))
	}
	if r.CoreUserID > 0 {
		values.Set("core_user_id", strconv.FormatUint(r.CoreUserID, 10))
	}
	if len(r.AdvertiserIDs) > 0 {
		values.Set("advertiser_ids", string(util.JSONMarshal(r.AdvertiserIDs)))
	}
	if len(r.Statuses) > 0 {
		values.Set("statuses", string(util.JSONMarshal(r.Statuses)))
	}
	if r.Cursor > 0 {
		values.Set("cursor", strconv.Itoa(r.Cursor))
	}
	if r.Count > 0 {
		values.Set("count", strconv.Itoa(r.Count))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// AccountsListResponse 查询订阅 Adv API Response
type AccountsListResponse struct {
	model.BaseResponse
	Data *AccountsListResult `json:"data,omitempty"`
}

type AccountsListResult struct {
	// NextCursor 游标值
	NextCursor int `json:"next_cursor,omitempty"`
	// Count 页大小值
	Count int `json:"count,omitempty"`
	// Advertisers
	Advertisers []SubscribeAccount `json:"advertisers,omitempty"`
}

type SubscribeAccount struct {
	Event string `json:"event,omitempty"`
	// CoreUserID 授权用户ID
	CoreUserID uint64 `json:"core_user_id,omitempty"`
	// AdvertiserID 授权广告主账号ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Status 可选值:
	// OK 推送中
	// PENDING 新增状态
	// UNAUTHORIZED 无权限
	// UNKNOWN 未知
	Status enum.SubscribeAdvStatus `json:"status,omitempty"`
	Reason string                  `json:"reason,omitempty"`
}
