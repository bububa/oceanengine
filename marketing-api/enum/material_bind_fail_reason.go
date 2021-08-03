package enum

// MaterialBindFailReason 推送失败原因
type MaterialBindFailReason string

const (
	// VIDEO_BINDING_EXISTED  视频已存在
	VIDEO_BINDING_EXISTED MaterialBindFailReason = "VIDEO_BINDING_EXISTED"
	// IMAGE_BINDING_EXISTED 图片已存在
	IMAGE_BINDING_EXISTED MaterialBindFailReason = "IMAGE_BINDING_EXISTED"
)
