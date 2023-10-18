package file

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// RebateMaterialSearchRequest 获取低效素材List API Request
type RebateMaterialSearchRequest struct {
	// AgentID 代理商账户ID
	AgentID uint64 `json:"agent_id,omitempty"`
	// Period 年月
	Period string `json:"period,omitempty"`
	// Filtering 过滤条件
	Filtering *RebateMaterialSearchFilter `json:"filtering,omitempty"`
	// Page 页码
	Page int `json:"page,omitempty"`
	// PageSize 页面大小, 默认10, 最大值500
	PageSize int `json:"page_size,omitempty"`
}

type RebateMaterialSearchFilter struct {
	// AdvertierID 广告主ID
	AdvertierID uint64 `json:"advertier_id,omitempty"`
	// CustomerName  客户名称: 需转为英文括号
	CustomerName string `json:"customer_name,omitempty"`
	// MaterialTag 素材标签筛选: 若传入多个标签，则会取交集
	// INEFFICIENT_MATERIAL 低效素材
	// AD_HIGH_QUALITY 广告高质量
	// FIRST_PUBLISH_MATERIAL 首发素材
	MaterialTag []enum.MaterialProperty `json:"material_tag,omitempty"`
	// MaterialIsEffective 当月最新一天素材是否在投, 允许值:
	// 1: 是, 2: 否
	MaterialIsEffective int `json:"material_is_effective,omitempty"`
	// IsValidVideoMaterial [政策粒度] 是否有效素材, 允许值:
	// 1: 是, 2: 否 (默认值 2)
	IsValidVideoMaterial int `json:"is_valid_video_material,omitempty"`
	// PolicyCostMin [政策力度] 小号范围下限, 格式限制: 正数
	PolicyCostMin int64 `json:"policy_cost_min,omitempty"`
	// PolicyCostMax [政策力度] 大号范围上限, 格式限制: 正数
	PolicyCostMax int64 `json:"policy_cost_max,omitempty"`
	// OperatorTag 运营标签
	// 1: 自运营
	// 2: 走量
	// 3: 收量
	OperatorTag int `json:"operator_tag,omitempty"`
	// RebateCalcPolicyType 政策类型
	// 1: 综代政策
	// 2: 优待政策
	RebateCalcPolicyType int `json:"rebate_calc_policy_type,omitempty"`
	// RebateCalcSettlementStatsType 结算行业统计类型
	// 1: 行业类目
	// 2: 引流电商
	// 4: 任务激励
	// 5: 微信加粉
	RebateCalcSettlementStatsType int `json:"rebate_calc_settlement_stats_type,omitempty"`
	// RebateCalcFirstIndustryName 一级结算行业
	RebateCalcFirstIndustryName string `json:"rebate_calc_first_industry_name,omitempty"`
	// RebateCalcSecondIndustryName 二级结算行业
	RebateCalcSecondIndustryName string `json:"rebate_calc_second_industry_name,omitempty"`
	// MaterialCreateStartDate 素材创建日期范围开始时间, 格式: yyyy-MM-dd
	MaterialCreateStartDate string `json:"material_create_start_date,omitempty"`
	// MaterialCreateEndDate 素材创建日期范围结束时间, 格式: yyyy-MM-dd
	MaterialCreateEndDate string `json:"material_create_end_date,omitempty"`
}

// Encode implement GetRequest interface
func (r RebateMaterialSearchRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("agent_id", strconv.FormatUint(r.AgentID, 10))
	values.Set("period", r.Period)
	if r.Filtering != nil {
		filtering, _ := json.Marshal(r.Filtering)
		values.Set("filtering", string(filtering))
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

// RebateMaterialSearchResponse 获取低效素材List API Response
type RebateMaterialSearchResponse struct {
	model.BaseResponse
	Data *RebateMaterialSearchResult `json:"data,omitempty"`
}

type RebateMaterialSearchResult struct {
	// PageInfo 翻页信息
	PageInfo  *model.PageInfo  `json:"page_info,omitempty"`
	Materials []RebateMaterial `json:"materials,omitempty"`
}

// RebateMaterial 低效素材
type RebateMaterial struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// OperatorTag 运营标签
	OperatorTag int `json:"operator_tag,omitempty"`
	// RebateCalcPolicyType 政策类型
	RebateCalcPolicyType int `json:"rebate_calc_policy_type,omitempty"`
	// RebateCalcSettlementStatsType 结算行业统计类型
	RebateCalcSettlementStatsType int `json:"rebate_calc_settlement_stats_type,omitempty"`
	// RebateCalcFirstIndustryName 一级结算行业
	RebateCalcFirstIndustryName string `json:"rebate_calc_first_industry_name,omitempty"`
	// RebateCalcSecondIndustryName 二级结算行业
	RebateCalcSecondIndustryName string `json:"rebate_calc_second_industry_name,omitempty"`
	// CustomerName 客户名称
	CustomerName string `json:"customer_name,omitempty"`
	// OptimizerName 代理商优化师
	OptimizerName string `json:"optimizer_name,omitempty"`
	// AgentSaleName 代理商销售
	AgentSaleName string `json:"agent_sale_name,omitempty"`
	// SecondAdAgentID 二级代理商账户ID
	SecondAdAgentID uint64 `json:"second_ad_agent_id,omitempty"`
	// MaterialInfo 素材信息
	MaterialInfo *RebateMaterialInfo `json:"material_info,omitempty"`
	// UpdateDate 数据更新日期
	UpdateDate string `json:"update_date,omitempty"`
}

// RebateMaterialInfo 素材信息
type RebateMaterialInfo struct {
	// MaterialID 素材ID
	MaterialID uint64 `json:"material_id,omitempty"`
	// MaterialIsEffective 素材是否有效
	MaterialIsEffective int `json:"material_is_effective,omitempty"`
	// IsValidVideoMaterial [政策粒度] 是否有效素材
	IsValidVideoMaterial int `json:"is_valid_video_material,omitempty"`
	// PolicyCost [政策力度] 消耗
	PolicyCost float64 `json:"policy_cost,omitempty"`
	// Cost 消耗
	Cost float64 `json:"cost,omitempty"`
	// MaterialCreateTime 素材创建时间
	MaterialCreateTime string `json:"material_create_time,omitempty"`
	// MaterialTag 素材属性标签
	MaterialTag []enum.MaterialProperty `json:"material_tag,omitempty"`
	// AdIDs 计划ID数组
	AdIDs []uint64 `json:"ad_ids,omitempty"`
	// CdpPromotionIDs 广告ID数组
	CdpPromotionIDs []uint64 `json:"cdp_promotion_ids,omitempty"`
}
