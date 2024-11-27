package star

import (
	"strconv"

	"github.com/bububa/oceanengine/marketing-api/model"
	"github.com/bububa/oceanengine/marketing-api/util"
)

// ReportOrderOverviewGetRequest 获取订单投后分析报表 API Request
type ReportOrderOverviewGetRequest struct {
	// StarID 星图id，星图客户授权后，通过“获取已授权账户”接口，查询到账号角色为”6-星图账号“的账户id，即为星图id
	StarID uint64 `json:"star_id,omitempty"`
	// OrderID 订单id，可通过“获取星图客户任务订单列表”获取
	OrderID uint64 `json:"order_id,omitempty"`
}

// Encode implement GetRequest interface
func (r ReportOrderOverviewGetRequest) Encode() string {
	values := util.GetUrlValues()
	values.Set("star_id", strconv.FormatUint(r.StarID, 10))
	values.Set("order_id", strconv.FormatUint(r.OrderID, 10))
	ret := values.Encode()
	util.PutUrlValues(values)
	return ret
}

// ReportOrderOverviewGetResponse 获取订单投后分析报表 API Response
type ReportOrderOverviewGetResponse struct {
	// Data json返回值
	Data *ReportOrderOverviewGetResult `json:"data,omitempty"`
	model.BaseResponse
}

// ReportOrderOverviewGetResult json返回值
type ReportOrderOverviewGetResult struct {
	// Comment 舆情表现
	Comment *CommentReport `json:"comment,omitempty"`
	// Convert 转化表现
	Convert *ConvertReport `json:"convert,omitempty"`
	// CostEffectiveness 性价比表现
	CostEffectiveness *CostEffectiveness `json:"cost_effectiveness,omitempty"`
	// CreativeReport 创意表现
	Creative *CreativeReport `json:"creative,omitempty"`
	// Spread 传播表现
	Spread *SpreadReport `json:"spread,omitempty"`
	// UpdateTime 数据更新时间，格式%Y-%m-%d %H:%M:%S
	UpdateTime string `json:"update_time,omitempty"`
}

// CommentReport 舆情表现
type CommentReport struct {
	// HighFrequencyWords 热词top10
	HighFrequencyWords []string `json:"high_frequency_words,omitempty"`
	// NegRate 负向评论率（neg_rate/100）%
	NegRate int `json:"neg_rate,omitempty"`
	// NeuRate 中立评论率（neu_rate/100）%
	NeuRate int `json:"neu_rate,omitempty"`
	// PosRate 正向评论率（pos_rate/100）%
	PosRate int `json:"pos_rate,omitempty"`
}

// ConvertReport 转化表现
type ConvertReport struct {
	// Click 组件点击量
	Click int64 `json:"click,omitempty"`
	// Ctr 组件点击率 (ctr/100)%
	Ctr int `json:"ctr,omitempty"`
	// Show 组件展示量
	Show int64 `json:"show,omitempty"`
}

// CostEffectiveness 性价比表现
type CostEffectiveness struct {
	// Cpm 千次播放成本(分)
	Cpm int64 `json:"cpm,omitempty"`
	// Play 播放次数
	Play int64 `json:"play,omitempty"`
	// Price 订单金额(分)
	Price int64 `json:"price,omitempty"`
}

// CreativeReport 创意表现
type CreativeReport struct {
	// FinishRate 完播率 (finish_rate/100)%
	FinishRate int `json:"finish_rate,omitempty"`
	// FiveSPlayRate 有效播放率（播放5s以上记为有效播放）(five_s_play_rate/100)%
	FiveSPlayRate int `json:"five_s_play_rate,omitempty"`
	// PlayRate 平均播放率（=用户观看该任务视频的平均观看时长/视频总时长）(play_rate/100)%
	PlayRate int `json:"play_rate,omitempty"`
}

// SpreadReport 传播表现
type SpreadReport struct {
	// Comment 评论量
	Comment int64 `json:"comment,omitempty"`
	// Like 点赞量
	Like int64 `json:"like,omitempty"`
	// Play 播放量
	Play int64 `json:"play,omitempty"`
	// Share 分享量
	Share int64 `json:"share,omitempty"`
}
