package creative

// AuditRecord 审核详细内容
type AuditRecord struct {
	// Desc 审核内容，即审核的内容类型，如 视频，图片，标题 等
	Desc string `json:"desc,omitempty"`
	// Content 拒绝内容（文字类型）
	Content string `json:"content,omitempty"`
	// ImageID 拒绝内容id（图片类型）
	ImageID uint64 `json:"image_id,omitempty"`
	// VideoID 拒绝内容id（视频类型）
	VideoID uint64 `json:"video_id,omitempty"`
	// AuditPlatform 审核来源类型，返回值：AD 广告审核、CONTENT 内容审核
	AuditPlatform string `json:"audit_platform,omitempty"`
	// RejectReason 拒绝原因，可能会有多条
	RejectReason []string `json:"reject_reason,omitempty"`
	// Suggestion 审核建议，可能会有多条
	Suggestion []string `json:"suggestion,omitempty"`
}
