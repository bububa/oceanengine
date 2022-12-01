package enum

// DownloadMode 优先从系统应用商店下载（下载模式）
type DownloadMode string

const (
	// DownloadMode_APP_STORE_DELIVERY 优先商店下载
	DownloadMode_APP_STORE_DELIVERY DownloadMode = "APP_STORE_DELIVERY"
	// DownloadMode_DEFAULT 默认下载
	DownloadMode_DEFAULT DownloadMode = "DEFAULT"
)
