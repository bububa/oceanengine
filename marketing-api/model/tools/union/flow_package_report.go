package union

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// FlowPackageReportRequest 查看rit数据 API Request
type FlowPackageReportRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// Filtering 过滤字段
	Filtering *FlowPackageReportFilter `json:"filtering,omitempty"`
	// OrderField 排序字段，所有的统计指标均可参与排序
	// 默认按rit排序
	OrderField string `json:"order_field,omitempty"`
	// OrderType 排序方式
	// 枚举值：ASC（升序），DESC（降序）
	// 默认：DESC
	OrderType enum.OrderType `json:"order_type,omitempty"`
	// Page 页数
	// 默认值: 1
	Page int `json:"page,omitempty"`
	// PageSize 页面大小
	// 默认值: 10，最大值100
	PageSize int `json:"page_size,omitempty"`
}

// FlowPackageReportFilter 过滤字段
type FlowPackageReportFilter struct {
	// AdIDs 广告 id 列表，最大数量限制：1000
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// CampaignIDs 广告组列表，最大数量限制：1000
	CampaignIDs []uint64 `json:"campaign_ids,omitempty"`
	// Rits 广告位列表，最大数量限制：1000
	Rits []uint64 `json:"rits,omitempty"`
	// StartTime 开始时间，格式为"yyyy-mm-dd"
	// 限制范围100天内
	// 默认7天前（不包括当天），即不指定起止时间获取最近7天数据
	StartTime string `json:"start_time,omitempty"`
	// EndTime 结束时间，格式为"yyyy-mm-dd"
	// 默认昨天，即不指定起止时间获取最近7天数据
	EndTime string `json:"end_time,omitempty"`
}

// Encode implement GetRequest interface
func (r FlowPackageReportRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("advertiser_id", strconv.FormatUint(r.AdvertiserID, 10))
	if r.Filtering != nil {
		bs, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(bs))
	}
	if r.OrderField != "" {
		values.Set("order_field", r.OrderField)
	}
	if r.OrderType != "" {
		values.Set("order_type", string(r.OrderType))
	}
	if r.Page > 1 {
		values.Set("page", strconv.Itoa(r.Page))
	}
	if r.PageSize > 10 {
		values.Set("page_size", strconv.Itoa(r.PageSize))
	}
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// FlowPackageReportResponse 查看rit数据 API Response
type FlowPackageReportResponse struct {
	model.BaseResponse
	// Data json 返回值`
	Data *FlowPackageReportData `json:"data,omitempty"`
}

// FlowPackageReportData json 返回值
type FlowPackageReportData struct {
	// PageInfo 分页信息
	PageInfo *model.PageInfo `json:"page_info,omitempty"`
	// List 数据列表
	List []FlowPackageReport `json:"list,omitempty"`
}

// FlowPackageReport .
type FlowPackageReport struct {
	// Rit 广告位
	Rit uint64 `json:"rit,omitempty"`
	// ShowCnt 展示指标
	ShowCnt int64 `json:"sbow_cnt,omitempty"`
	// ClickCnt 点击指标
	ClickCnt int64 `json:"click_cnt,omitempty"`
	// ActivePayRate 付费率
	ActivePayRate float64 `json:"active_pay_rate,omitemty"`
	// ActiveRegisterRate 注册率
	ActiveRegisterRate float64 `json:"active_register_rate,omitempty"`
	// Cost 消耗，单位： 元/10w
	Cost float64 `json:"cost,omitempty"`
	// Form 表单提交数
	Form int64 `json:"form,omitempty"`
	// ConversionCost 转化成本
	ConversionCost float64 `json:"conversion_cost,omitempty"`
	// ActiveCost 激活成本
	ActiveCost float64 `json:"active_cost,omitempty"`
	// GameAdditionCost 关键行为成本
	GameAdditionCost float64 `json:"game_addition_cost,omitempty"`
	// NextDayOpenCost 次留成本
	NextDayOpenCost float64 `json:"next_day_open_cost,omitempty"`
	// StatCost 消耗，单位：元
	StatCost float64 `json:"stat_cost,omitempty"`
	// ActivePayCost 付费成本
	ActivePayCost float64 `json:"active_pay_cost,omitempty"`
	// NextDayOpenRate 次留率
	NextDayOpenRate float64 `json:"next_day_open_rate,omitempty"`
	// ActivePay 付费数
	ActivePay int64 `json:"active_pay,omitempty"`
	// GameAdditionRate 关键行为率
	GameAdditionRate float64 `json:"game_addition_rate,omitempty"`
	// GameAddition 关键行为数
	GameAddition int64 `json:"game_addition,omitempty"`
	// Active 激活数
	Active int64 `json:"active,omitempty"`
	// ActiveRegisgter 注册数
	ActiveRegister int64 `json:"active_register,omitempty"`
	// ConvertCnt 转化量
	ConvertCnt int64 `json:"convert_cnt,omitempty"`
	// ActiveRegisterCost 注册成本`
	ActiveRegisterCost float64 `json:"active_register_cost,omitempty"`
	// NextDayOpen 次留数
	NextDayOpen int64 `json:"next_day_open,omitempty"`
	// FormCost 表单
	FormCost float64 `json:"form_cost,omitempty"`
	// LoanCompletion 完件数
	LoanCompletion int64 `json:"loan_completion,omitempty"`
	// PreLoanCredit 预授信数
	PreLoanCredit int64 `json:"pre_loan_credit,omitempty"`
	// LoanCredit 授信数
	LoanCredit int64 `json:"loan_credit,omitempty"`
	// LoanCompletionCost 完件成本
	LoanCompletionCost float64 `json:"loan_completion_cost,omitempty"`
	// PreLoanCreditCost 预授信成本
	PreLoanCreditCost float64 `json:"pre_loan_credit_cost,omitempty"`
	// LoanCreditCost 授信成本
	LoanCreditCost float64 `json:"loan_credit_cost,omitempty"`
}
