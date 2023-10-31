package ad

import "github.com/bububa/oceanengine/marketing-api/enum/qianchuan"

// RecommendKeyword 推荐关键词
type RecommendKeyword struct {
	// Keyword 关键词
	Keyword string `json:"keyword,omitempty"`
	// SuggestReason 推荐理由类型，枚举值:
	// CLICK: 高点买词
	// DARKHORSE: 黑马词
	// LOW_COST: 低成本词
	// POTENTIAL: 蓝海词
	// PREFERENCE: 同行买词
	SuggestReason qianchuan.KeywordSuggestReason `json:"suggest_reason,omitempty"`
	// SearchSum 月搜索量
	SearchSum int64 `json:"search_sum,omitempty"`
	// Degree 竞争程度，枚举值：
	// 1、2、3、4、5（对应pc页面上点亮的格数）
	Degree int `json:"degree,omitempty"`
	// CacheID 本次查询的唯一标识ID，有效期12h。
	// 可用于后续请求，用以保证查询到数据的完整性。
	CacheID string `json:"cache_id,omitempty"`
}
