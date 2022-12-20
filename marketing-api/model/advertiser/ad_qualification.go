package advertiser

// AdQualification 投放资质
type AdQualification struct {
	// QualificationID 资质id
	QualificationID uint64 `json:"qualification_id,omitempty"`
	// QualificationType 投放资质类型
	// 1：肖像、商标或其他授权
	// 2: 专利证明
	// 3: 商标证
	// 4: 广告内容中的数据证明
	// 5: 质检报告
	QualificationType int `json:"qualification_type,omitempty"`
	// PicturePreviewURL 投放资质预览url
	PicturePreviewURL string `json:"picture_preview_url,omitempty"`
	// Status 状态
	// 0: 待提交
	// 1: 待审核
	// 2: 审核中
	// 3: 审核通过
	// 4: 审核拒绝
	Status int `json:"status,omitempty"`
	// RejectReason 审核原因
	RejectReason string `json:"reject_reason,omitempty"`
	// AuditTime 审核时间，形如 yyyy-MM-dd HH:mm:ss
	AuditTime string `json:"audit_time,omitempty"`
	// ImageList 资质图片列表
	ImageList []Image `json:"image_list,omitempty"`
}

// Image 图片信息
type Image struct {
	// AttachmentID 附件id
	AttachmentID uint64 `json:"attachment_id,omitempty"`
	// ImageURL 图片预览地址
	ImageURL string `json:"image_url,omitempty"`
}
