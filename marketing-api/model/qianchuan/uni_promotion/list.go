package unipromotion

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/enum/qianchuan"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ListRequest 获取全域推广列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// StartTime 开始时间，格式 2021-04-05 00:00:00
	StartTime string `json:"start_time,omitempty"`
	// EndTime 结束时间，格式 2021-04-06 00:00:00
	EndTime string `json:"end_time,omitempty"`
	// MarketingGoal 按营销目标过滤，允许值
	// LIVE_PROM_GOODS：直播带货
	MarketingGoal enum.MarketingGoal `json:"marketing_goal,omitempty"`
	// Fields 需要查询的消耗指标，见返回参数
	Fields []string `json:"fields,omitempty"`
	// NeedCompensateInfo 是否获取成本保障信息，默认不获取
	// false：不获取
	// true：获取
	NeedCompensateInfo bool `json:"need_compensate_info,omitempty"`
	// OrderType 排序方式，允许值：
	// ASC 升序（默认）
	// DESC 降序
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// OrderField  排序字段，默认create_time，同时支持根据消耗指标排序
	// create_time
	// stat_cost
	// total_cost_per_pay_order_for_roi2
	// total_pay_order_count_for_roi2
	// total_pay_order_gmv_for_roi2
	// total_prepay_and_pay_order_roi2
	// total_prepay_order_count_for_roi2
	OrderField string `json:"order_field,omitempty"`
	// Page 页码，默认值：1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，允许值：10, 20, 50, 100, 200，默认值：10
	PageSize int `json:"page_size,omitempty"`
	// Filtering 过滤
	Filtering *ListFilter `json:"filtering,omitempty"`
}

// ListFilter 过滤
type ListFilter struct {
	// SmartBidType 投放方式 可选值:
	// SMART_BID_CONSERVATIVE 放量投放
	// SMART_BID_CUSTOM 控成本投放
	//  默认值: SMART_BID_CUSTOM
	SmartBidType enum.SmartBidType `json:"smart_bid_type,omitempty"`
	// SearchKeyword 搜索关键词，与search_keyword_type搭配使用
	// 注意：仅当marketing_goal=VIDEO_PROM_GOODS支持
	SearchKeyword string `json:"search_keyword,omitempty"`
	// Status 计划状态 可选值:
	// ALL 所有，不包括已删除
	// ALL_INCLUDE_DELETED 所有,包含已删除
	// AUDIT 新建审核中
	// DELETED 已删除
	// DELIVERY_OK 投放中
	// DISABLE 已暂停
	// EXTERNAL_URL_DISABLE 落地页暂不可用
	// FROZEN 已终止
	// LIVE_ROOM_OFF 关联直播间未开播
	// NO_SCHEDULE 不在投放时段
	// OFFLINE_AUDIT 审核不通过
	// OFFLINE_BALANCE 账户余额不足
	// OFFLINE_BUDGET 营销超出预算
	// QUOTA_DISABLE quota暂停
	// REAUDIT 修改审核中
	// SYSTEM_DISABLE 系统暂停
	// TIME_DONE 已完成
	// TIME_NO_REACH 未达到投放时间
	// ADVERTISER_OFFLINE_BUDGET账户预算不足
	// 注意：仅当marketing_goal=VIDEO_PROM_GOODS支持
	Status qianchuan.AdStatus `json:"status,omitempty"`
	// HavingCost 消耗情况 可选值:
	// ALL 全部
	// YES 有消耗
	// NO 无消耗
	// 注意：仅当marketing_goal=VIDEO_PROM_GOODS支持
	HavingCost string `json:"having_cost,omitempty"`
	// CreateStartDate 计划创建开始日期，时间格式 YYYY-MM-DD
	// 注意：仅当marketing_goal=VIDEO_PROM_GOODS支持
	CreateStartDate string `json:"create_start_date,omitempty"`
	// CreateEndDate 计划创建结束日期，时间格式 YYYY-MM-DD
	// 注意：仅当marketing_goal=VIDEO_PROM_GOODS支持
	CreateEndDate string `json:"create_end_date,omitempty"`
}

// Encode implement GetRequest interface
func (r ListRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("start_time", r.StartTime)
	values.Set("end_time", r.EndTime)
	values.Set("marketing_goal", string(r.MarketingGoal))
	values.Set("fields", string(util.JSONMarshal(r.Fields)))
	if r.NeedCompensateInfo {
		values.Set("need_compensate_info", "true")
	}
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	if r.Page > 0 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 0 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ListResponse 获取全域推广列表 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResult `json:"data,omitempty"`
}

type ListResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// AdList 全域推广列表
	AdList []Ad `json:"ad_list,omitempty"`
}
