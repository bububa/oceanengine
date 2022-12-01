package enum

// DownloadType 下载方式
type DownloadType string

const (
	// DownloadType_DOWNLOAD_URL 直接下载
	DownloadType_DOWNLOAD_URL DownloadType = "DOWNLOAD_URL"
	// DownloadType_EXTERNAL_URL 落地页下载
	DownloadType_EXTERNAL_URL DownloadType = "EXTERNAL_URL"
	// DownloadType_QUICK_APP_URL 快应用+下载链接
	DownloadType_QUICK_APP_URL DownloadType = "QUICK_APP_URL"
)
