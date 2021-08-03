package audience

// Metrics 指标
type Metrics struct {
	// StatCost 消耗
	StatCost float64 `json:"stat_cost,omitempty"`
	// ShowCnt 展示数
	ShowCnt int64 `json:"show_cnt,omitempty"`
	// CpmPlatform 平均千次展示成本
	CpmPlatform float64 `json:"cpm_platform,omitempty"`
	// ClickCnt 点击数
	ClickCnt int64 `json:"click_cnt,omitempty"`
	// Ctr 点击率
	Ctr float64 `json:"ctr,omitempty"`
	// CpcPlatform 平均点击单价
	CpcPlatform float64 `json:"cpc_platform,omitempty"`
	// ConvertCnt 转化数
	ConvertCnt int64 `json:"convert_cnt,omitempty"`
	// ConversionCost 转化成本
	ConversionCost float64 `json:"conversion_cost,omitempty"`
	// ConversionRate 转化率
	ConversionRate float64 `json:"conversion_rate,omitempty"`
	// DeepConvertCnt 深度转化数
	DeepConvertCnt int64 `json:"deep_convert_cnt,omitempty"`
	// DeepConvertCost 深度转化成本
	DeepConvertCost float64 `json:"deep_convert_cost,omitempty"`
	// DeepConvertRate 深度转化率
	DeepConvertRate float64 `json:"deep_convert_rate,omitempty"`
	// TotalPlay 播放数
	TotalPlay int64 `json:"total_play,omitempty"`
	// ValidPlay 有效播放数
	ValidPlay int64 `json:"valid_play,omitempty"`
	// ValidPlayRate 有效播放率
	ValidPlayRate float64 `json:"valid_play_rate,omitempty"`
	// ValidPlayCost 有效播放成本
	ValidPlayCost float64 `json:"valid_play_cost,omitempty"`
	// DyFollow 新增关注数
	DyFollow int64 `json:"dy_follow,omitempty"`
	// DyLike 点赞数
	DyLike int64 `json:"dy_like,omitempty"`
	// DyComment 评论提交数
	DyComment int64 `json:"dy_comment,omitempty"`
	// DyShare 分享数
	DyShare int64 `json:"dy_share,omitempty"`
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
