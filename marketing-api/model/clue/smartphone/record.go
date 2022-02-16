package smartphone

import (
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util/query"
)

// RecordRequest 查询智能电话拨打记录 API Request
type RecordRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// InstanceIDs 智能电话ID，不超过100个
	InstanceIDs []uint64 `json:"instance_ids,omitempty"`
	// AdIDs 广告计划ID，不超过100个
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// ClueIDs 线索ID，不超过100个
	ClueIDs []uint64 `json:"clue_ids,omitempty"`
	// SiteIDs 建站落地页ID，不超过100个
	SiteIDs []uint64 `json:"site_ids,omitempty"`
	// StartTime 起始时间，格式：%Y-%m-%d，默认7天前，即不指定起止时间获取最近7天数据
	StartTime string `json:"start_time,omitempty"`
	// EndTime 截止时间，格式：%Y-%m-%d，默认当天，即不指定起止时间获取最近7天数据
	EndTime string `json:"end_time,omitempty"`
	// Page 页码，默认值为1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值10，不超过50
	PageSize int `json:"page_size,omitempty"`
}

func (r RecordRequest) Encode() string {
	values, err := query.Values(r)
	if err != nil {
		return ""
	}
	return values.Encode()
}

// RecordResponse 查询智能电话拨打记录 API Response
type RecordResponse struct {
	model.BaseResponse
	// Data json返回值
	Data *RecordResponseData `json:"data,omitempty"`
}

// RecordResponseData json返回值
type RecordResponseData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 智能电话列表
	List []Record `json:"list,omitempty"`
}

// Record 智能电话拨打记录
type Record struct {
	// InstanceID 智能电话ID
	InstanceID uint64 `json:"instance_id,omitempty"`
	// AdID 广告计划ID
	AdID uint64 `json:"ad_id,omitempty"`
	// SiteID 建站落地页ID
	SiteID uint64 `json:"site_id,omitempty"`
	// ClueID 线索ID
	ClueID uint64 `json:"clue_id,omitempty"`
	// CallDuration 呼叫时长，秒单位
	CallDuration int64 `json:"call_duration,omitempty"`
	// RealDuration 通话时长，秒单位
	RealDuration int64 `json:"real_duration,omitempty"`
	// VirtualNumber 呼叫的虚拟号码
	VirtualNumber string `json:"virtual_number,omitempty"`
	// StartTIme 拨打时间，格式：%Y-%m-%d %H:%M:%S
	StartTime string `json:"start_time,omitempty"`
	// CreateTime 创建时间，格式：%Y-%m-%d %H:%M:%S
	CreateTime string `json:"create_time,omitempty"`
}
