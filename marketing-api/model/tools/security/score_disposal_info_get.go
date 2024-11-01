package security

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ScoreDisposalInfoGetRequest 查看积分处置详情 API Request
type ScoreDisposalInfoGetRequest struct {
	// Filtering 过滤条件
	Filtering *ScoreDisposalInfoGetFilter `json:"filtering,omitempty"`
	// BusinessLine 业务线，允许值：
	// AD 巨量广告
	BusinessLine string `json:"business_line,omitempty"`
	// OrderField 排序字段 ，允许值：
	// disposal_start_time 处置开始时间
	// disposal_end_time 处置结束时间
	OrderField string `json:"order_field,omitempty"`
	// OrderType 排序方式 可选值:
	// ASC 升序，默认
	// DESC 降序
	OrderType string `json:"order_type,omitempty"`
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Page 页数，默认值：1，page范围为[1,99999]
	Page int `json:"page,omitempty"`
	// PageSize 页面大小，默认值：10，page_size范围为[1,100]
	PageSize int `json:"page_size,omitempty"`
}

type ScoreDisposalInfoGetFilter struct {
	// DisposalAction 处置行为，允许值：
	// ACCOUNTCLEAR 账号封停+主体限制新开
	//  默认值: ACCOUNTCLEAR
	DisposalAction string `json:"disposal_action,omitempty"`
	// DisposalStatus 处置状态，允许值：
	// DISPOSAL 已处置
	//  默认值: DISPOSAL
	DisposalStatus string `json:"disposal_status,omitempty"`
}

func (r ScoreDisposalInfoGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	values.Set("page", strconv.Itoa(r.Page))
	values.Set("page_size", strconv.Itoa(r.PageSize))
	if r.BusinessLine != "" {
		values.Set("business_line", r.BusinessLine)
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	if r.OrderType != "" {
		values.Set("order_type", r.OrderType)
	}
	if r.Filtering != nil {
		values.Set("filtering", string(util.JSONMarshal(r.Filtering)))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ScoreDisposalInfoGetResponse 查看积分处置详情 API Response
type ScoreDisposalInfoGetResponse struct {
	Data *ScoreDisposalInfoGetResult `json:"data,omitempty"`
	model.BaseResponse
}

type ScoreDisposalInfoGetResult struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// DisposalInfoList 处置信息
	DisposalInfoList []DisposalInfo `json:"disposal_info_list,omitempty"`
}

// DisposalInfo 处置信息
type DisposalInfo struct {
	// DisposalAction 处置行为
	// ACCOUNTCLEAR 账户封停+主体限制新建账户
	DisposalAction string `json:"disposal_action,omitempty"`
	// DisposalTerm 处置时长
	DisposalTerm string `json:"disposal_term,omitempty"`
	// DisposalStartTime 处置开始时间
	DisposalStartTime string `json:"disposal_start_time,omitempty"`
	// DisposalEndTime 处置结束时间
	DisposalEndTime string `json:"disposal_end_time,omitempty"`
	// DisposalStatus 处置状态
	// DISPOSAL 已处置
	DisposalStatus string `json:"disposal_status,omitempty"`
	// OuterText 对外文案
	OuterText string `json:"outer_text,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// DisposalScore 落罚分值
	DisposalScore int64 `json:"disposal_score,omitempty"`
}
