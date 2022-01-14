package audience

import "encoding/json"

// Metrics 指标
type Metrics struct {
	// StatCost 消耗
	StatCost json.Number `json:"stat_cost,omitempty"`
	// ShowCnt 展示数
	ShowCnt json.Number `json:"show_cnt,omitempty"`
	// CpmPlatform 平均千次展示成本
	CpmPlatform json.Number `json:"cpm_platform,omitempty"`
	// ClickCnt 点击数
	ClickCnt json.Number `json:"click_cnt,omitempty"`
	// Ctr 点击率
	Ctr json.Number `json:"ctr,omitempty"`
	// CpcPlatform 平均点击单价
	CpcPlatform json.Number `json:"cpc_platform,omitempty"`
	// ConvertCnt 转化数
	ConvertCnt json.Number `json:"convert_cnt,omitempty"`
	// ConversionCost 转化成本
	ConversionCost json.Number `json:"conversion_cost,omitempty"`
	// ConversionRate 转化率
	ConversionRate json.Number `json:"conversion_rate,omitempty"`
	// DeepConvertCnt 深度转化数
	DeepConvertCnt json.Number `json:"deep_convert_cnt,omitempty"`
	// DeepConvertCost 深度转化成本
	DeepConvertCost json.Number `json:"deep_convert_cost,omitempty"`
	// DeepConvertRate 深度转化率
	DeepConvertRate json.Number `json:"deep_convert_rate,omitempty"`
	// TotalPlay 播放数
	TotalPlay json.Number `json:"total_play,omitempty"`
	// ValidPlay 有效播放数
	ValidPlay json.Number `json:"valid_play,omitempty"`
	// ValidPlayRate 有效播放率
	ValidPlayRate json.Number `json:"valid_play_rate,omitempty"`
	// ValidPlayCost 有效播放成本
	ValidPlayCost json.Number `json:"valid_play_cost,omitempty"`
	// DyFollow 新增关注数
	DyFollow json.Number `json:"dy_follow,omitempty"`
	// DyLike 点赞数
	DyLike json.Number `json:"dy_like,omitempty"`
	// DyComment 评论提交数
	DyComment json.Number `json:"dy_comment,omitempty"`
	// DyShare 分享数
	DyShare json.Number `json:"dy_share,omitempty"`
}

// MetricsDict 查询指标详细数据
type MetricsDict struct {
	// Cost 总消耗(单位元,精确到分)
	Cost float64 `json:"cost,omitempty"`
	// Show 展示数
	Show int64 `json:"show,omitempty"`
	// Click 点击数
	Click int64 `json:"click,omitempty"`
	// DownloadFinish 下载完成数
	DownloadFinish int64 `json:"download_finish,omitempty"`
	// Convert 转化数
	Convert int64 `json:"convert,omitempty"`
}
