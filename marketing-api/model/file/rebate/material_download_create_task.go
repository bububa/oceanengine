package rebate

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// MaterialDownloadCreateTaskRequest 创建下载任务 API Request
type MaterialDownloadCreateTaskRequest struct {
	// AgentID 代理商帐户Id
	AgentID uint64 `json:"agent_id,omitempty"`
	// Period 业绩消耗年月 年月 格式：yyyy-MM
	// 默认值：当前月份（如2024-04 ）
	// 最早允许传入时间：2024-04
	Period string `json:"period,omitempty"`
	// MaterialPolicyType 素材政策类型 可选值:
	// VALID_ORIGINAL_MATERIAL_RATE 有效首投占比
	// VALID_FIRST_EFFECTIVE_MATERIAL 有效首发素材
	// VALID_HIGH_QUALITY_MATERIAL 有效优质素材
	MaterialPolicyType enum.MaterialPolicyType `json:"material_policy_type,omitempty"`
	// Filtering 过滤条件
	Filtering *MaterialDownloadCreateTaskFilter `json:"filtering,omitempty"`
}

// MaterialDownloadCreateTaskFilter
type MaterialDownloadCreateTaskFilter struct {
	// FirstEffectivePeroid 素材首投所在月 格式：yyyy-MM
	// 默认值：当前月份（如2024-04 ）
	// 最早允许传入时间：2024-04
	FirstEffectivePeroid string `json:"first_effective_period,omitempty"`
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CustomerName 客户名称
	// 如果名称涉及括号，使用英文括号
	CustomerName string `json:"customer_name,omitempty"`
	// OperatorTag 运营标签 可选值:
	// SHOULIANG 收量
	// ZOULIANG 走量
	// ZIYUNYING 自运营
	OperatorTag enum.RebateOperatorTag `json:"operator_tag,omitempty"`
	// RebateCalcPolicyType 政策类型 可选值:
	// NORMAL_POLICY 综代政策
	// EXCLUSIVE_POLICY 优代政策
	// CAR_POLICY 汽车政策
	RebateCalcPolicyType enum.RebateCalcPolicyType `json:"rebate_calc_policy_type,omitempty"`
	// RebateCalcSettlementStatsType 结算行业统计类型 可选值:
	// NORMAL_INDUSTRY 行业类目
	// E_COMMERCE 引流电商
	// TASK_INCENTIVES 任务激励
	// WECHAT_INCREASE微信加粉
	RebateCalcSettlementStatsType enum.RebateCalcSettlementStatsType `json:"rebate_calc_settlement_stats_type,omitempty"`
	// RebateCalcFirstIndustryName 一级结算行业
	RebateCalcFirstIndustryName string `json:"rebate_calc_first_industry_name,omitempty"`
	// RebateCalcSecondIndustryName 二级结算行业
	RebateCalcSecondIndustryName string `json:"rebate_calc_second_industry_name,omitempty"`
	// MaterialTag 素材标签筛选项（如传入多个标签，取交集）
	// HIGH_QUALITY_MATERIAL（优质素材）
	// LOW_QUALITY_MATERIAL（低质素材）
	// FIRST_EFFECTIVE_MATERIAL（首发素材）
	MaterialTag []enum.RebateMaterialTag `json:"material_tag,omitempty"`
	// IsLiveRebateType 是否直播返点类型（仅支持有效优质或有效首发查询时筛选，入参有效首投，该筛选需要置空） 可选值:
	// NO 否
	// YES 是
	IsLiveRebateType enum.YesNo `json:"is_live_rebate_type,omitempty"`
	// IsAccumulation 是否累量 可选值:
	// NO 否
	// YES 是
	IsAccumulation enum.YesNo `json:"is_accumulation,omitempty"`
	// IsValidMaterial 【政策粒度】是否有效素材 可选值:
	// NO 否
	// YES 是
	IsValidMaterial enum.YesNo `json:"is_valid_material,omitempty"`
	// PolicyCostMin 【政策粒度】累量消耗范围区间下限
	// - 仅支持录入整数,单位：元
	PolicyCostMin int64 `json:"policy_cost_min,omitempty"`
	// PolicyCostMax 【政策粒度】累量消耗范围区间上限
	PolicyCostMax int64 `json:"policy_cost_max,omitempty"`
	// MaterialIsEffective 数据更新日期前一日是否在投 可选值:
	// NO 否
	// YES 是
	MaterialIsEffective enum.YesNo `json:"material_is_effective,omitempty"`
	// MaterialFirstEffectiveStartDate 素材首投日期范围开始日期
	// 格式：yyyy-MM-dd
	MaterialFirstEffectiveStartDate string `json:"material_first_effective_start_date,omitempty"`
	// MaterialFirstEffectiveEndDate 素材首投日期范围结束日期
	// 格式：yyyy-MM-dd
	MaterialFirstEffectiveEndDate string `json:"material_first_effective_end_date,omitempty"`
	// ThisPeriodCumDayNum 本期已累量天数
	ThisPeriodCumDayNum int `json:"this_period_cum_day_num,omitempty"`
}

// Encode implements PostRequest interface
func (r MaterialDownloadCreateTaskRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// MaterialDownloadCreateTaskResponse 创建下载任务 API Response
type MaterialDownloadCreateTaskResponse struct {
	model.BaseResponse
	Data struct {
		// QueryID 下载任务对应的查询ID
		QueryID string `json:"query_id,omitempty"`
	} `json:"data,omitempty"`
}
