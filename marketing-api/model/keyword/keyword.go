package keyword

import "github.com/bububa/oceanengine/marketing-api/enum"

// Keyword 关键词
type Keyword struct {
	// MatchType 匹配类型; 枚举值: "PHRASE"（短语匹配）, "EXTENSIVE"（广泛匹配）,"PRECISION"（精准匹配）
	MatchType enum.KeywordMatchType `json:"match_type"`
	// Word 关键词
	Word string `json:"word,omitempty"`
	// Bid 出价。与bid_type联动，当bit_type="CUSTOM"此字段才有效
	Bid float64 `json:"bid,omitempty"`
	// BidType 出价类型
	BidType enum.KeywordBidType `json:"bid_type,omitempty"`
	// IsPause 是否停用; 枚举值：0（启用），1（停用）
	IsPause *int `json:"is_pause,omitempty"`
	// KeywordID 关键词id，唯一标志计划下的关键词。可根据该id进行搜索、更新、删除
	KeywordID uint64 `json:"keyword_id,omitempty"`
	// WordID 词ID，不同计划下如果关键词的内容相同，词id会相同
	WordID uint64 `json:"word_id,omitempty"`
	// Msv 月搜索量
	Msg int64 `json:"msv,omitempty"`
	// Status  审核状态:
	// AUDITING 审核中
	// AUDIT_ACCEPTED 审核通过
	// AUDIT_REJECTED 审核拒绝
	// DELETED 已删除
	Status enum.KeywordAuditStatus `json:"status,omitempty"`
	// ErrorReason 关键词添加失败原因
	ErrorReason string `json:"error_reason,omitempty"`
}
