package audience

type Metrics struct {
	StatCost        float64 `json:"stat_cost,omitempty"`         // 消耗
	ShowCnt         int64   `json:"show_cnt,omitempty"`          // 展示数
	CpmPlatform     float64 `json:"cpm_platform,omitempty"`      // 平均千次展示成本
	ClickCnt        int64   `json:"click_cnt,omitempty"`         // 点击数
	Ctr             float64 `json:"ctr,omitempty"`               // 点击率
	CpcPlatform     float64 `json:"cpc_platform,omitempty"`      // 平均点击单价
	ConvertCnt      int64   `json:"convert_cnt,omitempty"`       // 转化数
	ConversionCost  float64 `json:"conversion_cost,omitempty"`   // 转化成本
	ConversionRate  float64 `json:"conversion_rate,omitempty"`   // 转化率
	DeepConvertCnt  int64   `json:"deep_convert_cnt,omitempty"`  // 深度转化数
	DeepConvertCost float64 `json:"deep_convert_cost,omitempty"` // 深度转化成本
	DeepConvertRate float64 `json:"deep_convert_rate,omitempty"` // 深度转化率
	TotalPlay       int64   `json:"total_play,omitempty"`        // 播放数
	ValidPlay       int64   `json:"valid_play,omitempty"`        // 有效播放数
	ValidPlayRate   float64 `json:"valid_play_rate,omitempty"`   // 有效播放率
	ValidPlayCost   float64 `json:"valid_play_cost,omitempty"`   // 有效播放成本
	DyFollow        int64   `json:"dy_follow,omitempty"`         // 新增关注数
	DyLike          int64   `json:"dy_like,omitempty"`           // 点赞数
	DyComment       int64   `json:"dy_comment,omitempty"`        // 评论提交数
	DyShare         int64   `json:"dy_share,omitempty"`          // 分享数
}

type MetricsDict struct {
	Cost           float64 `json:"cost,omitempty"`            // 总消耗(单位元,精确到分)
	Show           int64   `json:"show,omitempty"`            // 展示数
	Click          int64   `json:"click,omitempty"`           // 点击数
	DownloadFinish int64   `json:"download_finish,omitempty"` // 下载完成数
	Convert        int64   `json:"convert,omitempty"`         // 转化数
}
