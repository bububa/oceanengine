package qianchuan

// KeywordStatus 关键词状态
type KeywordStatus string

const (
	// KeywordStatus_CONFIRM 审核通过且可代入
	KeywordStatus_CONFIRM KeywordStatus = "CONFIRM"
	// KeywordStatus_REJECT 审核拒绝
	KeywordStatus_REJECT KeywordStatus = "REJECT"
	// KeywordStatus_AUDIT 新建审核中
	KeywordStatus_AUDIT KeywordStatus = "AUDIT"
	// KeywordStatus_ENABLE
	KeywordStatus_ENABLE KeywordStatus = "ENABLE"
	// KeywordStatus_DELETE 已删除
	KeywordStatus_DELETE KeywordStatus = "DELETE"
	// KeywordStatus_PAUSED 词暂停
	KeywordStatus_PAUSED KeywordStatus = "PAUSED"
)
