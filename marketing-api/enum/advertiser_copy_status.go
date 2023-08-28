package enum

// AdvertiserCopyStatus 复制结果状态码
type AdvertiserCopyStatus int

const (
	// AdvertiserCopySuccess 全部成功
	AdvertiserCopySuccess AdvertiserCopyStatus = 1
	// AdvertiserCopyPartSuccess 部分成功
	AdvertiserCopyPartSuccess AdvertiserCopyStatus = 2
	// AdvertiserCopyFailed 全部失败
	AdvertiserCopyFailed AdvertiserCopyStatus = 3
)
