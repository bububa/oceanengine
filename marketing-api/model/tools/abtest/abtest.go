package abtest

import (
	"github.com/bububa/oceanengine/marketing-api/enum"
	"github.com/bububa/oceanengine/marketing-api/model"
)

// AbTest 实验信息
type AbTest struct {
	// TestID 实验ID
	TestID uint64 `json:"test_id,omitempty"`
	// TestName 实验名称
	TestName string `json:"test_name,omitempty"`
	// TestStartTime 实验开始时间
	TestStartTime string `json:"test_start_time,omitempty"`
	// TestEndTime 实验结束时间
	TestEndTime string `json:"test_end_time,omitempty"`
	// Status 实验状态，"CREATED"：排期中，"PROCESSING"：进行中，"FINISH"：结束，"FAILED"：不满足实验条件。
	Status enum.AbTestStatus `json:"status,omitempty"`
	// Conclusion 实验结论，"NOT_START"：未开始，"OBVIOUS"：有明显结论，"INSUFFICIENT"：数据量不足；"FAILED"：不满足实验条件，"TMP_OBVIOUS"：有初步结论。
	Conclusion enum.AbTestConclusion `json:"conclusion,omitempty"`
	// TestType 实验类型，"CAMPAIGN"：广告组，"AD"：广告计划。
	TestType enum.AbTestType `json:"test_type,omitempty"`
	// CreateTime 实验创建时间，格式："2020-12-25 07:12:08"
	CreateTime string `json:"create_time,omitempty"`
	// ReportUpdateTime 实验报告更新时间，格式："2020-12-25 07:12:08"
	ReportUpdateTime string `json:"report_update_time,omitempty"`
	// ReportUpdateDate 实验报告更新日期，格式："2020-12-25 07:12:08"
	ReportUpdateDate string `json:"report_update_date,omitempty"`
	// TestVersion 实验版本，"V1"：V1版本，"V2"：V2版本。
	TestVersion string `json:"test_version,omitempty"`
	// ObjectList 实验对象ID列表
	ObjectList []AbTestObject `json:"object_list,omitempty"`
	// ReportDate 报告日期，格式："20201225"
	ReportDate string `json:"report_date,omitempty"`
	// ReportTime 报告时间，格式："2020-12-25 07:12:08"
	ReportTime string `json:"report_time,omitempty"`
	// ConclusionDetail 结论详情
	ConclusionDetail *ConclusionDetail `json:"conclusion_detail,omitempty"`
	// TestReports 实验对象详细报告
	TestReports []Report `json:"test_reports,omitempty"`
}

// ConclusionDetail 结论详情
type ConclusionDetail struct {
	// Winner 获胜对象ID列表
	Winner []uint64 `json:"winner,omitempty"`
	// Common 普通对象ID列表
	Common []uint64 `json:"common,omitempty"`
	// Failed 失败对象ID列表
	Failed []uint64 `json:"failed,omitempty"`
}

// Report 实验对象详细报告
type Report struct {
	// ObjectID 实验对象ID
	ObjectID uint64 `json:"object_id,omitempty"`
	// Metrics 该实验对象详细数据
	Metrics *Metrics `json:"metrics,omitempty"`
	// DailyStats 实验对象分日统计数据
	DailyStats []DailyStats `json:"daily_stats,omitempty"`
}

// Metrics 该实验对象详细数据
type Metrics struct {
	// StatCost 消耗
	StatCost model.Float64 `json:"stat_cost,omitempty"`
	// ShowCnt 展示数
	ShowCnt model.Int64 `json:"show_cnt,omitempty"`
	// ClickCnt 点击数
	ClickCnt model.Int64 `json:"click_cnt,omitempty"`
	// Ctr 点击率
	Ctr model.Float64 `json:"ctr,omitempty"`
	// CpcPlatform 平均点击单价
	CpcPlatform model.Float64 `json:"cpc_platform,omitempty"`
	// ConvertCnt 转化数
	ConvertCnt model.Int64 `json:"convert_cnt,omitempty"`
	// ConversionCost 转化成本
	ConversionCost model.Float64 `json:"conversion_cost,omitempty"`
	// ConversionRate 转化率
	ConversionRate model.Float64 `json:"conversion_rate,omitempty"`
	// CpmPlatform 平均前次展现费用
	CpmPlatform model.Float64 `json:"cpm_platform,omitempty"`
	// CtrProb ctr获胜概率
	CtrProb model.Float64 `json:"ctr_prob,omitempty"`
	// CvrProb cvr获胜概率
	CvrProb model.Float64 `json:"cvr_prob,omitempty"`
	// ConversionCostProb 转化成本获胜概率
	ConversionCostProb model.Float64 `json:"conversion_cost_prob,omitempty"`
	// CompositeProb 综合指标获胜概率
	CompositeProb model.Float64 `json:"composite_prob,omitempty"`
	// CtrVariation ctr波动范围
	CtrVariation *Variation `json:"ctr_variation,omitempty"`
	// CvrVariation cvr波动范围
	CvrVariation *Variation `json:"cvr_variation,omitempty"`
	// ConversionCostVariation 转化成本波动范围
	ConversionCostVariation *Variation `json:"conversion_cost_variation,omitempty"`
}

// Variation 波动范围
type Variation struct {
	// VariationMin 波动范围下限
	VariationMin model.Float64 `json:"variation_min,omitempty"`
	// VariationMax 波动范围上限
	VariationMax model.Float64 `json:"variation_max,omitempty"`
}

// DailyStats 实验对象分日统计数据
type DailyStats struct {
	// ObjectID 实验对象ID
	ObjectID uint64 `json:"object_id,omitempty"`
	// Date 数据统计日期
	Date string `json:"date,omitempty"`
	// Metrics 详细数据
	Metrics *Metrics `json:"metrics,omitempty"`
}
